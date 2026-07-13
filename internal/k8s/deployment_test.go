package k8s

import (
	"Kubexplorer/internal/model"
	"context"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

// New deployments have no status conditions yet; this used to panic on Conditions[0]
func TestGetDeploymentsWithoutConditions(t *testing.T) {
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: "fresh", Namespace: "default"},
	}
	client := NewDeployment(&fakeResolver{client: fake.NewSimpleClientset(dep)})

	deployments, err := client.GetDeployments(context.Background(), "test")
	if err != nil {
		t.Fatalf("GetDeployments() error: %v", err)
	}
	if len(deployments) != 1 {
		t.Fatalf("GetDeployments() returned %d, want 1", len(deployments))
	}
	if deployments[0].Status != "Unknown" {
		t.Errorf("status = %q, want Unknown for a deployment without conditions", deployments[0].Status)
	}
}

func TestGetDeploymentNotFound(t *testing.T) {
	client := NewDeployment(&fakeResolver{client: fake.NewSimpleClientset()})

	_, err := client.GetDeployment(context.Background(), model.ResourceRef{Cluster: "test", Namespace: "default", Name: "missing"})
	if err == nil {
		t.Fatal("GetDeployment() expected error for missing deployment")
	}
}

func TestGetDeploymentsUnregisteredCluster(t *testing.T) {
	client := NewDeployment(&fakeResolver{})

	if _, err := client.GetDeployments(context.Background(), "nope"); err == nil {
		t.Fatal("GetDeployments() expected error for unregistered cluster")
	}
}
