package kubeclient

import (
	"Kubexplorer/backend/model"
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

type metric struct {
	manager ClusterResolver
}

func NewMetric(manager ClusterResolver) MetricClient {
	return &metric{manager: manager}
}
func (m *metric) GetPodMetricsV2(namespace string, clusterCtx string) (*v1beta1.PodMetricsList, error) {
	client, err := m.manager.ResolveClusterMetric(clusterCtx)
	if err != nil {
		return nil, fmt.Errorf("kubeclient: error resolving cluster context: %v", err)
	}

	metrics, err := client.MetricsV1beta1().PodMetricses(namespace).List(context.Background(), metav1.ListOptions{})

	return metrics, nil
}

func (m *metric) GetPodMetrics(namespace string, chMetricDto <-chan []model.PodMetricDto) []model.PodMetricDto {
	//inClusterConfig, err := rest.InClusterConfig()
	//if err != nil {
	//	fmt.Printf("Error creating in-cluster config: %v", err)
	//}
	//
	//metricsClient, err := metricsv.NewForConfig(inClusterConfig)
	//if err != nil {
	//	fmt.Print(err)
	//}

	return pollMetrics(nil, namespace)
}

func pollMetrics(metricsClient metricsv.Interface, namespace string) []model.PodMetricDto {
	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error to get pod metrics")
	}

	var metrics []model.PodMetricDto

	for _, podMetric := range podMetrics.Items {
		for _, container := range podMetric.Containers {
			metrics = append(metrics, model.PodMetricDto{
				Name:      podMetric.Name,
				Namespace: podMetric.Namespace,
				Consume: model.Resource{
					Cpu:              container.Usage.Cpu().MilliValue(),
					Memory:           container.Usage.Memory().ScaledValue(resource.Mega),
					Storage:          container.Usage.Storage().ScaledValue(resource.Mega),
					StorageEphemeral: container.Usage.StorageEphemeral().ScaledValue(resource.Mega),
				},
			})
		}
	}

	return metrics
}
