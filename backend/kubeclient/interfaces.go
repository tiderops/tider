package kubeclient

import (
	"Kubexplorer/backend/model"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

type ClusterClient interface {
	ListAvailableClusters() ([]model.ClusterInfo, error)
	GetCurrentCluster() (model.EnvironmentDto, error)
	GetNode(name string) (model.NodeDtoV2, error)
	GetNodes() ([]model.NodeDtoV2, error)
}

type MetricClient interface {
	GetPodMetrics(namespace string, chMetricDto <-chan []model.PodMetricDto) []model.PodMetricDto
	GetPodMetricsV2(namespace string) *v1beta1.PodMetricsList
}

type PodClient interface {
	GetPodsMock() ([]model.PodDto, error)
	GetPods(clusterCtx string) ([]model.PodDto, error)
	GetPod(name string, namespace string, clusterCtx string) (model.PodDto, error)
	UpdatePod(name string, namespace string, dto model.PodDto, clusterCtx string) error
	DeletePod(name string, namespace string, clusterCtx string) error
}

type DeploymentClient interface {
	GetDeployments(clusterCtx string) ([]model.DeploymentDto, error)
	GetDeploymentsMock() ([]model.DeploymentDto, error)
	GetDeployment(name string, namespace string, clusterCtx string) (model.DeploymentDto, error)
	UpdateDeployment(name string, namespace string, dto model.DeploymentDto, clusterCtx string) error
	DeleteDeployment(name string, namespace string, clusterCtx string) error
}

type StorageClient interface {
	GetPersistentVolumes(clusterCtx string) ([]model.PersistentVolumeDto, error)
	GetPersistentVolume(name string, clusterCtx string) (model.PersistentVolumeDto, error)
	UpdatePersistentVolume(name string, dto model.PersistentVolumeDto, clusterCtx string) error
	DeletePersistentVolume(name string, clusterCtx string) error
	GetPersistentVolumesClaim(clusterCtx string) ([]model.PersistentVolumeClaimDto, error)
	GetPersistentVolumeClaim(name string, namespace string, clusterCtx string) (model.PersistentVolumeClaimDto, error)
	UpdatePersistentVolumeClaim(name string, namespace string, dto model.PersistentVolumeClaimDto, clusterCtx string) error
	DeletePersistentVolumeClaim(name string, namespace string, clusterCtx string) error
}

type NamespaceClient interface {
	GetNamespaces() ([]model.NamespaceDto, error)
	GetNamespace(name string) (model.NamespaceDto, error)
	UpdateNamespace(name string, dto model.NamespaceDto) error
	DeleteNamespace(name string) error
}

type ServiceClient interface {
	GetServices(clusterCtx string) ([]model.ServiceDto, error)
	GetService(name string, namespace string, clusterCtx string) (model.ServiceDto, error)
	UpdateService(name string, namespace string, dto model.ServiceDto, clusterCtx string) error
	DeleteService(name string, namespace string, clusterCtx string) error
}

type IngressClient interface {
	GetIngresses(clusterCtx string) ([]model.IngressDto, error)
	GetIngress(name string, namespace string, clusterCtx string) (model.IngressDto, error)
	UpdateIngress(name string, namespace string, dto model.IngressDto, clusterCtx string) error
	DeleteIngress(name string, namespace string, clusterCtx string) error
}
