package usecase

import (
	"Kubexplorer/internal/model"
	"context"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

type fakePodClient struct {
	gotRef model.ResourceRef
	pod    *corev1.Pod
}

func (f *fakePodClient) GetPods(ctx context.Context, clusterCtx string) ([]model.PodDto, error) {
	return nil, nil
}
func (f *fakePodClient) GetPod(ctx context.Context, ref model.ResourceRef) (model.PodDto, error) {
	return model.PodDto{}, nil
}
func (f *fakePodClient) GetPodObject(ctx context.Context, ref model.ResourceRef) (*corev1.Pod, error) {
	f.gotRef = ref
	return f.pod, nil
}
func (f *fakePodClient) UpdatePod(ctx context.Context, ref model.ResourceRef, dto model.PodUpdate) error {
	return nil
}
func (f *fakePodClient) DeletePod(ctx context.Context, ref model.ResourceRef) error {
	return nil
}

type fakeDeploymentClient struct {
	dep    *appsv1.Deployment
	depDto model.DeploymentDto
}

func (f *fakeDeploymentClient) GetDeployments(ctx context.Context, clusterCtx string) ([]model.DeploymentDto, error) {
	return nil, nil
}
func (f *fakeDeploymentClient) GetDeployment(ctx context.Context, ref model.ResourceRef) (model.DeploymentDto, error) {
	return f.depDto, nil
}
func (f *fakeDeploymentClient) GetDeploymentObject(ctx context.Context, ref model.ResourceRef) (*appsv1.Deployment, error) {
	return f.dep, nil
}
func (f *fakeDeploymentClient) UpdateDeployment(ctx context.Context, ref model.ResourceRef, dto model.DeploymentUpdate) error {
	return nil
}
func (f *fakeDeploymentClient) DeleteDeployment(ctx context.Context, ref model.ResourceRef) error {
	return nil
}
func (f *fakeDeploymentClient) ExportManifest(ctx context.Context, ref model.ResourceRef) ([]byte, error) {
	return nil, nil
}

type fakeJobClient struct {
	job *batchv1.Job
}

func (f *fakeJobClient) GetJob(ctx context.Context, ref model.ResourceRef) (*batchv1.Job, error) {
	return f.job, nil
}

func TestAnalysePassesRef(t *testing.T) {
	podClient := &fakePodClient{pod: &corev1.Pod{Status: corev1.PodStatus{Phase: corev1.PodPending}}}
	svc := NewTroubleshootUseCase(podClient, &fakeDeploymentClient{}, &fakeJobClient{})

	ref := model.ResourceRef{Cluster: "my-cluster", Namespace: "my-namespace", Name: "my-pod"}
	got := svc.Analyse(context.Background(), ref, string(Pod))

	if podClient.gotRef != ref {
		t.Errorf("Analyse() passed ref %+v, want %+v", podClient.gotRef, ref)
	}
	if got.Meaning != string(Pending) {
		t.Errorf("Analyse() meaning = %q, want %q", got.Meaning, string(Pending))
	}
}

func TestAnalyseUnsupportedResource(t *testing.T) {
	svc := NewTroubleshootUseCase(&fakePodClient{}, &fakeDeploymentClient{}, &fakeJobClient{})

	got := svc.Analyse(context.Background(), model.ResourceRef{}, "CONFIGMAP")

	if got.Meaning != "Unsupported object type" {
		t.Errorf("Analyse() meaning = %q, want %q", got.Meaning, "Unsupported object type")
	}
}

func TestAnalyseDeploymentAndJob(t *testing.T) {
	replicas := int32(2)
	dep := &appsv1.Deployment{
		Spec:   appsv1.DeploymentSpec{Replicas: &replicas},
		Status: appsv1.DeploymentStatus{AvailableReplicas: 0},
	}
	job := &batchv1.Job{
		Status: batchv1.JobStatus{
			Conditions: []batchv1.JobCondition{
				{Type: batchv1.JobFailed, Status: corev1.ConditionTrue, Reason: "BackoffLimitExceeded"},
			},
		},
	}
	svc := NewTroubleshootUseCase(&fakePodClient{}, &fakeDeploymentClient{dep: dep}, &fakeJobClient{job: job})

	ref := model.ResourceRef{Cluster: "c", Namespace: "ns", Name: "x"}
	if got := svc.Analyse(context.Background(), ref, string(Deployment)); got.Meaning != string(UnavailableReplicas) {
		t.Errorf("Analyse(deployment) meaning = %q, want %q", got.Meaning, string(UnavailableReplicas))
	}
	if got := svc.Analyse(context.Background(), ref, string(Job)); got.Meaning != string(BackoffLimitExceeded) {
		t.Errorf("Analyse(job) meaning = %q, want %q", got.Meaning, string(BackoffLimitExceeded))
	}
}
