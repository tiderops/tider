package middleware

import (
	"Kubexplorer/backend/endpoint"
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/service"
	"Kubexplorer/backend/usecase"
	"k8s.io/client-go/kubernetes"
)

type WorkloadMiddleware struct {
	endpoint endpoint.WorkloadEndpoint
}

func NewWorkloadMiddleware(endpoint *endpoint.WorkloadEndpoint) *WorkloadMiddleware {
	return &WorkloadMiddleware{endpoint: *endpoint}
}

func (w *WorkloadMiddleware) GetPods() []model.PodDto {
	x, _ := w.endpoint.GetPods()
	return x
}

func (w *WorkloadMiddleware) GetPod(name string, namespace string) (model.PodDto, error) {
	return w.endpoint.GetPod(name, namespace)
}

func (w *WorkloadMiddleware) UpdatePod(name string, namespace string, dto model.PodDto) error {
	return w.endpoint.UpdatePod(name, namespace, dto)
}

func (w *WorkloadMiddleware) RestartPod(name string, namespace string) error {
	return w.endpoint.RestartPod(name, namespace)
}

func (w *WorkloadMiddleware) GetDeployments() []model.DeploymentDto {
	x, _ := w.endpoint.GetDeployments()
	return x
}

func (w *WorkloadMiddleware) GetDeployment(name string, namespace string) (model.DeploymentDto, error) {
	return w.endpoint.GetDeployment(name, namespace)
}

func (w *WorkloadMiddleware) UpdateDeployment(name string, namespace string, dto model.DeploymentDto) error {
	return w.endpoint.UpdateDeployment(name, namespace, dto)
}

func (w *WorkloadMiddleware) DeleteDeployment(name string, namespace string) error {
	return w.endpoint.DeleteDeployment(name, namespace)
}

func (w *WorkloadMiddleware) ResourceTuning(namespace string) {
	w.endpoint.ResourceTuning(namespace)
}

func (w *WorkloadMiddleware) TroubleshootPod(name string, namespace string) {
	w.endpoint.TroubleshootPod(name, namespace)
}

func (w *WorkloadMiddleware) TroubleshootDeployment(name string, namespace string) {
	w.endpoint.TroubleshootDeployment(name, namespace)
}

func BuildWorkload(client kubernetes.Interface) *WorkloadMiddleware {
	deploymentClient := kubeclient.NewDeployment(client)
	podClient := kubeclient.NewPod(client)
	diagnosticService := service.NewDiagnosticService(client)

	deploymentUseCase := usecase.NewDeploymentUseCase(deploymentClient, diagnosticService)
	podUseCase := usecase.NewPodUseCase(podClient, diagnosticService)

	workloadEndpoint := endpoint.NewWorkloadEndpoint(podUseCase, deploymentUseCase)

	return NewWorkloadMiddleware(workloadEndpoint)
}
