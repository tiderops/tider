package usecase

import (
	"context"
	"testing"

	"Kubexplorer/internal/model"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

func TestTuneValue(t *testing.T) {
	tests := []struct {
		name    string
		limit   int64
		current int64
		want    int64
	}{
		{name: "usage above 90% grows limit by 50%", limit: 1000, current: 950, want: 1500},
		{name: "usage at 85% grows limit by 30%", limit: 1000, current: 850, want: 1300},
		{name: "usage at exactly 80% grows limit by 30%", limit: 1000, current: 800, want: 1300},
		{name: "usage below 30% shrinks limit by 50%", limit: 1000, current: 100, want: 500},
		{name: "usage in normal range keeps limit", limit: 1000, current: 500, want: 1000},
		{name: "usage just below 80% keeps limit", limit: 1000, current: 799, want: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tuneValue(tt.limit, tt.current); got != tt.want {
				t.Errorf("tuneValue(%d, %d) = %d, want %d", tt.limit, tt.current, got, tt.want)
			}
		})
	}
}

type fakeMetricClient struct {
	metrics *v1beta1.PodMetricsList
}

func (f *fakeMetricClient) GetPodMetrics(ctx context.Context, clusterCtx string, namespace string) (*v1beta1.PodMetricsList, error) {
	return f.metrics, nil
}

func TestResourceTuningReturnsRecommendationsWithoutApplying(t *testing.T) {
	metrics := &v1beta1.PodMetricsList{
		Items: []v1beta1.PodMetrics{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "web-7d9f",
					Labels: map[string]string{"app": "web"},
				},
				Containers: []v1beta1.ContainerMetrics{
					{
						Name: "web",
						Usage: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("950m"),
							corev1.ResourceMemory: resource.MustParse("100M"),
						},
					},
				},
			},
		},
	}
	deployment := &fakeDeploymentClient{
		depDto: model.DeploymentDto{
			Name:      "web",
			Namespace: "default",
			Containers: []model.Container{
				{Name: "web", Limit: model.Resource{Cpu: 1000, Memory: 200}},
			},
		},
	}

	uc := NewResourceUseCase(deployment, &fakeMetricClient{metrics: metrics})

	recs, err := uc.ResourceTuning(context.Background(), "my-cluster", "default")
	if err != nil {
		t.Fatalf("ResourceTuning() error: %v", err)
	}
	if len(recs) != 1 {
		t.Fatalf("ResourceTuning() returned %d recommendations, want 1", len(recs))
	}

	rec := recs[0]
	if rec.Deployment != "web" || rec.Container != "web" {
		t.Errorf("recommendation identifies %s/%s, want web/web", rec.Deployment, rec.Container)
	}
	// CPU at 95% of the 1000m limit -> suggest 1500m; memory at 50% -> unchanged.
	if rec.SuggestedLimit.Cpu != 1500 {
		t.Errorf("suggested CPU = %d, want 1500", rec.SuggestedLimit.Cpu)
	}
	if rec.SuggestedLimit.Memory != 200 {
		t.Errorf("suggested memory = %d, want 200 (unchanged)", rec.SuggestedLimit.Memory)
	}
}
