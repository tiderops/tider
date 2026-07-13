package usecase

import (
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

func TestCheckPodErrors(t *testing.T) {
	tests := []struct {
		name        string
		pod         corev1.Pod
		wantMeaning string
	}{
		{
			name: "pending pod",
			pod: corev1.Pod{
				Status: corev1.PodStatus{Phase: corev1.PodPending},
			},
			wantMeaning: string(Pending),
		},
		{
			name: "succeeded pod reports completed",
			pod: corev1.Pod{
				Status: corev1.PodStatus{Phase: corev1.PodSucceeded},
			},
			wantMeaning: string(Completed),
		},
		{
			name: "failed pod reports terminating",
			pod: corev1.Pod{
				Status: corev1.PodStatus{Phase: corev1.PodFailed},
			},
			wantMeaning: string(Terminating),
		},
		{
			name: "unschedulable condition",
			pod: corev1.Pod{
				Status: corev1.PodStatus{
					Phase: corev1.PodRunning,
					Conditions: []corev1.PodCondition{
						{Type: corev1.PodScheduled, Status: corev1.ConditionFalse},
					},
				},
			},
			wantMeaning: string(Unschedulable),
		},
		{
			name: "crashloopbackoff waiting container",
			pod: corev1.Pod{
				Status: corev1.PodStatus{
					Phase: corev1.PodRunning,
					ContainerStatuses: []corev1.ContainerStatus{
						{State: corev1.ContainerState{
							Waiting: &corev1.ContainerStateWaiting{Reason: string(CrashLoopBackOff)},
						}},
					},
				},
			},
			wantMeaning: string(CrashLoopBackOff),
		},
		{
			name: "imagepullbackoff waiting container",
			pod: corev1.Pod{
				Status: corev1.PodStatus{
					Phase: corev1.PodRunning,
					ContainerStatuses: []corev1.ContainerStatus{
						{State: corev1.ContainerState{
							Waiting: &corev1.ContainerStateWaiting{Reason: string(ImagePullBackOff)},
						}},
					},
				},
			},
			wantMeaning: string(ImagePullBackOff),
		},
		{
			name: "oomkilled terminated container",
			pod: corev1.Pod{
				Status: corev1.PodStatus{
					Phase: corev1.PodRunning,
					ContainerStatuses: []corev1.ContainerStatus{
						{State: corev1.ContainerState{
							Terminated: &corev1.ContainerStateTerminated{Reason: string(OOMKilled)},
						}},
					},
				},
			},
			wantMeaning: string(OOMKilled),
		},
		{
			name: "healthy running pod",
			pod: corev1.Pod{
				Status: corev1.PodStatus{
					Phase: corev1.PodRunning,
					ContainerStatuses: []corev1.ContainerStatus{
						{State: corev1.ContainerState{Running: &corev1.ContainerStateRunning{}}},
					},
				},
			},
			wantMeaning: "No known pod errors detected.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckPodErrors(tt.pod)
			if got.Meaning != tt.wantMeaning {
				t.Errorf("CheckPodErrors() meaning = %q, want %q", got.Meaning, tt.wantMeaning)
			}
			if tt.wantMeaning != "No known pod errors detected." && got.Recommendation == "" {
				t.Errorf("CheckPodErrors() returned no recommendation for %q", tt.wantMeaning)
			}
		})
	}
}

func TestCheckDeploymentErrors(t *testing.T) {
	replicas := func(n int32) *int32 { return &n }

	tests := []struct {
		name        string
		dep         appsv1.Deployment
		wantMeaning string
	}{
		{
			name: "unavailable replicas",
			dep: appsv1.Deployment{
				Spec:   appsv1.DeploymentSpec{Replicas: replicas(3)},
				Status: appsv1.DeploymentStatus{AvailableReplicas: 1, ReadyReplicas: 3},
			},
			wantMeaning: string(UnavailableReplicas),
		},
		{
			name: "ready replicas below desired",
			dep: appsv1.Deployment{
				Spec:   appsv1.DeploymentSpec{Replicas: replicas(2)},
				Status: appsv1.DeploymentStatus{AvailableReplicas: 2, ReadyReplicas: 1},
			},
			wantMeaning: string(MinimumReplicasUnavailable),
		},
		{
			name: "progress deadline exceeded",
			dep: appsv1.Deployment{
				Spec: appsv1.DeploymentSpec{Replicas: replicas(1)},
				Status: appsv1.DeploymentStatus{
					AvailableReplicas: 1,
					ReadyReplicas:     1,
					Conditions: []appsv1.DeploymentCondition{
						{Type: appsv1.DeploymentProgressing, Reason: "ProgressDeadlineExceeded"},
					},
				},
			},
			wantMeaning: string(ProgressDeadlineExceeded),
		},
		{
			name: "healthy deployment",
			dep: appsv1.Deployment{
				Spec:   appsv1.DeploymentSpec{Replicas: replicas(2)},
				Status: appsv1.DeploymentStatus{AvailableReplicas: 2, ReadyReplicas: 2},
			},
			wantMeaning: "No known deployment errors detected.",
		},
		{
			name: "nil replicas defaults to one",
			dep: appsv1.Deployment{
				Status: appsv1.DeploymentStatus{AvailableReplicas: 0, ReadyReplicas: 0},
			},
			wantMeaning: string(UnavailableReplicas),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckDeploymentErrors(tt.dep)
			if got.Meaning != tt.wantMeaning {
				t.Errorf("CheckDeploymentErrors() meaning = %q, want %q", got.Meaning, tt.wantMeaning)
			}
		})
	}
}

func TestCheckJobErrors(t *testing.T) {
	tests := []struct {
		name        string
		job         batchv1.Job
		wantMeaning string
	}{
		{
			name: "backoff limit exceeded",
			job: batchv1.Job{
				Status: batchv1.JobStatus{
					Conditions: []batchv1.JobCondition{
						{Type: batchv1.JobFailed, Status: corev1.ConditionTrue, Reason: "BackoffLimitExceeded"},
					},
				},
			},
			wantMeaning: string(BackoffLimitExceeded),
		},
		{
			name: "deadline exceeded",
			job: batchv1.Job{
				Status: batchv1.JobStatus{
					Conditions: []batchv1.JobCondition{
						{Type: batchv1.JobFailed, Status: corev1.ConditionTrue, Reason: "DeadlineExceeded"},
					},
				},
			},
			wantMeaning: string(DeadlineExceeded),
		},
		{
			name: "unknown failure reason falls back to condition message",
			job: batchv1.Job{
				Status: batchv1.JobStatus{
					Conditions: []batchv1.JobCondition{
						{Type: batchv1.JobFailed, Status: corev1.ConditionTrue, Reason: "Other", Message: "boom"},
					},
				},
			},
			wantMeaning: "boom",
		},
		{
			name:        "healthy job",
			job:         batchv1.Job{},
			wantMeaning: "No known job errors detected.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckJobErrors(tt.job)
			if got.Meaning != tt.wantMeaning {
				t.Errorf("CheckJobErrors() meaning = %q, want %q", got.Meaning, tt.wantMeaning)
			}
		})
	}
}
