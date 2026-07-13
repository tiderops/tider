package k8s

import (
	"Kubexplorer/internal/model"
	"context"
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetPodsEmptyCluster(t *testing.T) {
	client := NewPod(&fakeResolver{client: fake.NewSimpleClientset()})

	pods, err := client.GetPods(context.Background(), "test")
	if err != nil {
		t.Fatalf("GetPods() error: %v", err)
	}
	if len(pods) != 0 {
		t.Errorf("GetPods() on empty cluster returned %d pods", len(pods))
	}
}

func TestGetPodsMapsAllContainers(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:              "web-1",
			Namespace:         "default",
			Labels:            map[string]string{"app": "web"},
			CreationTimestamp: metav1.NewTime(time.Now().Add(-2 * time.Hour)),
		},
		Spec: corev1.PodSpec{
			NodeName: "node-1",
			Containers: []corev1.Container{
				{
					Name:  "web",
					Image: "nginx:1.27",
					Resources: corev1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("500m"),
							corev1.ResourceMemory: resource.MustParse("128M"),
						},
					},
				},
				{Name: "sidecar", Image: "envoy:v1.30"},
			},
		},
		Status: corev1.PodStatus{Phase: corev1.PodRunning},
	}

	client := NewPod(&fakeResolver{client: fake.NewSimpleClientset(pod)})

	pods, err := client.GetPods(context.Background(), "test")
	if err != nil {
		t.Fatalf("GetPods() error: %v", err)
	}
	if len(pods) != 1 {
		t.Fatalf("GetPods() returned %d pods, want 1", len(pods))
	}

	got := pods[0]
	if len(got.Containers) != 2 {
		t.Fatalf("pod has %d containers in DTO, want 2", len(got.Containers))
	}
	if got.Containers[0].Name != "web" || got.Containers[1].Name != "sidecar" {
		t.Errorf("container names = %s, %s; want web, sidecar", got.Containers[0].Name, got.Containers[1].Name)
	}
	if got.Containers[0].Limit.Cpu != 500 {
		t.Errorf("container CPU limit = %d, want 500", got.Containers[0].Limit.Cpu)
	}
	if got.Age == "" || got.CreatedAt == 0 {
		t.Errorf("age not mapped: Age=%q CreatedAt=%d", got.Age, got.CreatedAt)
	}
	if got.Node != "node-1" || got.Status != "Running" {
		t.Errorf("node/status = %s/%s, want node-1/Running", got.Node, got.Status)
	}
}

func TestGetPodNotFound(t *testing.T) {
	client := NewPod(&fakeResolver{client: fake.NewSimpleClientset()})

	_, err := client.GetPod(context.Background(), model.ResourceRef{Cluster: "test", Namespace: "default", Name: "missing"})
	if err == nil {
		t.Fatal("GetPod() expected error for missing pod")
	}
}

func TestUpdatePodWithoutLabels(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "bare", Namespace: "default"},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{{Name: "main", Image: "old:1"}},
		},
	}
	clientset := fake.NewSimpleClientset(pod)
	client := NewPod(&fakeResolver{client: clientset})

	ref := model.ResourceRef{Cluster: "test", Namespace: "default", Name: "bare"}
	dto := model.PodUpdate{App: "bare-app", Container: model.ContainerUpdate{Image: "new:2"}}

	if err := client.UpdatePod(context.Background(), ref, dto); err != nil {
		t.Fatalf("UpdatePod() error: %v", err)
	}

	updated, _ := clientset.CoreV1().Pods("default").Get(context.Background(), "bare", metav1.GetOptions{})
	if updated.Labels["app"] != "bare-app" {
		t.Errorf("label app = %q, want bare-app", updated.Labels["app"])
	}
	if updated.Spec.Containers[0].Image != "new:2" {
		t.Errorf("image = %q, want new:2", updated.Spec.Containers[0].Image)
	}
}
