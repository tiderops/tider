package endpoint

import (
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
)

type WorkloadEndpoint struct {
	podUseCase        usecase.PodUseCase
	deploymentUseCase usecase.DeploymentUseCase
	resourceUseCase   usecase.ResourceUseCase
}

func NewWorkloadEndpoint(podUseCase usecase.PodUseCase, deploymentUseCase usecase.DeploymentUseCase) *WorkloadEndpoint {
	return &WorkloadEndpoint{podUseCase: podUseCase, deploymentUseCase: deploymentUseCase}
}

func (we *WorkloadEndpoint) GetPods(clusterCtx string) ([]model.PodDto, error) {
	return we.podUseCase.GetPods(clusterCtx)
}

func (we *WorkloadEndpoint) GetPod(name string, namespace string) (model.PodDto, error) {
	return we.podUseCase.GetPod(name, namespace)
}

func (we *WorkloadEndpoint) UpdatePod(name string, namespace string, dto model.PodDto) error {
	return we.podUseCase.UpdatePod(name, namespace, dto)
}

func (we *WorkloadEndpoint) RestartPod(name string, namespace string) error {
	return we.podUseCase.RestartPod(name, namespace)
}

func (we *WorkloadEndpoint) GetDeployments(clusterCtx string) ([]model.DeploymentDto, error) {
	return we.deploymentUseCase.GetDeployments(clusterCtx)
}

func (we *WorkloadEndpoint) GetDeployment(name string, namespace string) (model.DeploymentDto, error) {
	return we.deploymentUseCase.GetDeployment(name, namespace)
}

func (we *WorkloadEndpoint) UpdateDeployment(name string, namespace string, dto model.DeploymentDto) error {
	return we.deploymentUseCase.UpdateDeployment(name, namespace, dto)
}

func (we *WorkloadEndpoint) DeleteDeployment(name string, namespace string) error {
	return we.deploymentUseCase.DeleteDeployment(name, namespace)
}

func (we *WorkloadEndpoint) ResourceTuning(namespace string) {
	we.resourceUseCase.Invoke(namespace)
}

func (we *WorkloadEndpoint) TroubleshootPod(name string, namespace string) {
	we.podUseCase.TroubleshootPod(name, namespace)
}

func (we *WorkloadEndpoint) TroubleshootDeployment(name string, namespace string) {
	we.deploymentUseCase.TroubleshootDeployment(name, namespace)
}
