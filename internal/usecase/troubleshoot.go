package usecase

import (
	"Kubexplorer/internal/model"
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

type SourceObject string

const (
	Pod        SourceObject = "POD"
	Job        SourceObject = "JOB"
	Deployment SourceObject = "DEPLOYMENT"
)

type WellKnownPodError string
type WellKnownDeploymentError string
type WellKnownJobError string

const (
	// Pod-level
	CrashLoopBackOff           WellKnownPodError = "CrashLoopBackOff"
	OOMKilled                  WellKnownPodError = "OOMKilled"
	ImagePullBackOff           WellKnownPodError = "ImagePullBackOff"
	ErrImagePull               WellKnownPodError = "ErrImagePull"
	CreateContainerConfigError WellKnownPodError = "CreateContainerConfigError"
	ContainerCannotRun         WellKnownPodError = "ContainerCannotRun"
	Unschedulable              WellKnownPodError = "Unschedulable"
	Evicted                    WellKnownPodError = "Evicted"
	NodeLost                   WellKnownPodError = "NodeLost"
	Completed                  WellKnownPodError = "Completed"
	Pending                    WellKnownPodError = "Pending"
	Terminating                WellKnownPodError = "Terminating"

	// Deployment-level
	UnavailableReplicas        WellKnownDeploymentError = "UnavailableReplicas"
	MinimumReplicasUnavailable WellKnownDeploymentError = "MinimumReplicasUnavailable"
	ProgressDeadlineExceeded   WellKnownDeploymentError = "ProgressDeadlineExceeded"

	// Job-level
	DeadlineExceeded     WellKnownJobError = "DeadlineExceeded"
	BackoffLimitExceeded WellKnownJobError = "BackoffLimitExceeded"
)

var PodErrorMessages = map[WellKnownPodError]string{
	CrashLoopBackOff:           "The container keeps crashing on startup. Check container logs with `kubectl logs` and verify entrypoint, configs, or dependencies.",
	OOMKilled:                  "The container was killed due to exceeding memory limits. Increase memory limits/requests or optimize application memory usage.",
	ImagePullBackOff:           "Kubernetes cannot pull the container image. Verify image name, tag, registry credentials, or network access.",
	ErrImagePull:               "The image could not be pulled. Check that the image exists and that the registry is accessible.",
	CreateContainerConfigError: "Invalid container configuration. Review environment variables, volume mounts, and container spec.",
	ContainerCannotRun:         "The container failed to start. Check entrypoint, permissions, or binary compatibility.",
	Unschedulable:              "The pod cannot be scheduled. Verify resource requests, node selectors, taints, or affinity rules.",
	Evicted:                    "The pod was evicted due to resource pressure. Reduce requests/limits or add more cluster resources.",
	NodeLost:                   "The node running this pod is unreachable. Check node health and networking.",
	Completed:                  "The pod has successfully finished execution (Job/Pod complete). No action needed unless it was expected to keep running.",
	Pending:                    "The pod is stuck in Pending. Check scheduler logs, resource availability, or PVC binding.",
	Terminating:                "The pod is stuck in Terminating. Check for finalizers, stuck volumes, or force delete with `kubectl delete pod --force --grace-period=0`.",
}

var DeploymentErrorMessages = map[WellKnownDeploymentError]string{
	UnavailableReplicas:        "Not enough replicas are available. Check pod errors, resource limits, or scheduling constraints.",
	MinimumReplicasUnavailable: "Minimum replicas not met. Scale your cluster or adjust replica settings.",
	ProgressDeadlineExceeded:   "Deployment rollout is stuck. Check pod logs, events, and ensure readiness/liveness probes are correct.",
}

var JobErrorMessages = map[WellKnownJobError]string{
	DeadlineExceeded:     "The Job exceeded its active deadline. Increase `.spec.activeDeadlineSeconds` or optimize the job workload.",
	BackoffLimitExceeded: "The Job retried too many times and failed. Investigate pod logs and fix underlying issues.",
}

// TroubleshootUseCase diagnoses well-known failure states of pods,
// deployments, and jobs and returns a meaning plus a recommendation.
type TroubleshootUseCase interface {
	Analyse(ctx context.Context, ref model.ResourceRef, resource string) model.Troubleshoot
}

type troubleshootUseCase struct {
	pod        PodClient
	deployment DeploymentClient
	job        JobClient
}

func NewTroubleshootUseCase(pod PodClient, deployment DeploymentClient, job JobClient) TroubleshootUseCase {
	return &troubleshootUseCase{pod: pod, deployment: deployment, job: job}
}

func (d *troubleshootUseCase) Analyse(ctx context.Context, ref model.ResourceRef, resource string) model.Troubleshoot {
	switch resource {
	case string(Pod):
		pod, err := d.pod.GetPodObject(ctx, ref)
		if err != nil {
			return model.Troubleshoot{Meaning: fmt.Sprintf("Error retrieving Pod: %v", err)}
		}
		return CheckPodErrors(*pod)

	case string(Deployment):
		dep, err := d.deployment.GetDeploymentObject(ctx, ref)
		if err != nil {
			return model.Troubleshoot{Meaning: fmt.Sprintf("Error retrieving Deployment: %v", err)}
		}
		return CheckDeploymentErrors(*dep)

	case string(Job):
		job, err := d.job.GetJob(ctx, ref)
		if err != nil {
			return model.Troubleshoot{Meaning: fmt.Sprintf("Error retrieving Job: %v", err)}
		}
		return CheckJobErrors(*job)

	default:
		return model.Troubleshoot{Meaning: "Unsupported object type"}
	}
}

func CheckPodErrors(pod corev1.Pod) model.Troubleshoot {
	// Check phase-level
	switch pod.Status.Phase {
	case corev1.PodPending:
		if msg, ok := PodErrorMessages[Pending]; ok {
			return model.Troubleshoot{
				Meaning:        string(Pending),
				Recommendation: msg,
			}
		}
	case corev1.PodSucceeded:
		if msg, ok := PodErrorMessages[Completed]; ok {
			return model.Troubleshoot{
				Meaning:        string(Completed),
				Recommendation: msg,
			}
		}
	case corev1.PodFailed:
		if msg, ok := PodErrorMessages[Terminating]; ok {
			return model.Troubleshoot{
				Meaning:        string(Terminating),
				Recommendation: msg,
			}
		}
	}

	// Check conditions
	for _, cond := range pod.Status.Conditions {
		if cond.Type == corev1.PodScheduled && cond.Status == corev1.ConditionFalse {
			if msg, ok := PodErrorMessages[Unschedulable]; ok {
				return model.Troubleshoot{
					Meaning:        string(Unschedulable),
					Recommendation: msg,
				}
			}
		}
		if cond.Reason == "Evicted" {
			if msg, ok := PodErrorMessages[Evicted]; ok {
				return model.Troubleshoot{
					Meaning:        string(Evicted),
					Recommendation: msg,
				}
			}
		}
		if cond.Reason == "NodeLost" {
			if msg, ok := PodErrorMessages[NodeLost]; ok {
				return model.Troubleshoot{
					Meaning:        string(NodeLost),
					Recommendation: msg,
				}
			}
		}
	}

	// Check container states
	for _, cs := range pod.Status.ContainerStatuses {
		if cs.State.Waiting != nil {
			if msg, ok := PodErrorMessages[WellKnownPodError(cs.State.Waiting.Reason)]; ok {
				return model.Troubleshoot{
					Meaning:        cs.State.Waiting.Reason,
					Recommendation: msg,
				}
			}
		}
		if cs.State.Terminated != nil {
			if msg, ok := PodErrorMessages[WellKnownPodError(cs.State.Terminated.Reason)]; ok {
				return model.Troubleshoot{
					Meaning:        cs.State.Terminated.Reason,
					Recommendation: msg,
				}
			}
		}
		if cs.LastTerminationState.Terminated != nil {
			if msg, ok := PodErrorMessages[WellKnownPodError(cs.LastTerminationState.Terminated.Reason)]; ok {
				return model.Troubleshoot{
					Meaning:        cs.LastTerminationState.Terminated.Reason,
					Recommendation: msg,
				}
			}
		}
	}

	// Fallback: if pod.Status.Reason is set
	if pod.Status.Reason != "" {
		if msg, ok := PodErrorMessages[WellKnownPodError(pod.Status.Reason)]; ok {
			return model.Troubleshoot{
				Meaning:        pod.Status.Reason,
				Recommendation: msg,
			}
		}
	}

	return model.Troubleshoot{Meaning: "No known pod errors detected."}
}

func CheckDeploymentErrors(dep appsv1.Deployment) model.Troubleshoot {
	desired := int32(1)
	if dep.Spec.Replicas != nil {
		desired = *dep.Spec.Replicas
	}
	available := dep.Status.AvailableReplicas
	ready := dep.Status.ReadyReplicas

	if available < desired {
		if msg, ok := DeploymentErrorMessages[UnavailableReplicas]; ok {
			return model.Troubleshoot{
				Meaning:        string(UnavailableReplicas),
				Recommendation: msg,
			}
		}
	}
	if ready < desired {
		if msg, ok := DeploymentErrorMessages[MinimumReplicasUnavailable]; ok {
			return model.Troubleshoot{
				Meaning:        string(MinimumReplicasUnavailable),
				Recommendation: msg,
			}
		}
	}

	for _, cond := range dep.Status.Conditions {
		if cond.Type == appsv1.DeploymentProgressing && cond.Reason == "ProgressDeadlineExceeded" {
			if msg, ok := DeploymentErrorMessages[ProgressDeadlineExceeded]; ok {
				return model.Troubleshoot{
					Meaning:        string(ProgressDeadlineExceeded),
					Recommendation: msg,
				}
			}
		}
	}

	return model.Troubleshoot{Meaning: "No known deployment errors detected."}
}

func CheckJobErrors(job batchv1.Job) model.Troubleshoot {
	for _, cond := range job.Status.Conditions {
		if cond.Type == batchv1.JobFailed && cond.Status == corev1.ConditionTrue {
			switch cond.Reason {
			case "BackoffLimitExceeded":
				if msg, ok := JobErrorMessages[BackoffLimitExceeded]; ok {
					return model.Troubleshoot{
						Meaning:        string(BackoffLimitExceeded),
						Recommendation: msg,
					}
				}
			case "DeadlineExceeded":
				if msg, ok := JobErrorMessages[DeadlineExceeded]; ok {
					return model.Troubleshoot{
						Meaning:        string(DeadlineExceeded),
						Recommendation: msg,
					}
				}
			default:
				return model.Troubleshoot{Meaning: cond.Message}
			}
		}
	}

	return model.Troubleshoot{Meaning: "No known job errors detected."}
}
