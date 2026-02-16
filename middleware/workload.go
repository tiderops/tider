package middleware

import (
	"Kubexplorer/backend/endpoint"
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/service"
	"Kubexplorer/backend/usecase"
	"fmt"
)

type WorkloadMiddleware struct {
	endpoint endpoint.WorkloadEndpoint
}

func NewWorkloadMiddleware(endpoint *endpoint.WorkloadEndpoint) *WorkloadMiddleware {
	return &WorkloadMiddleware{endpoint: *endpoint}
}

func (w *WorkloadMiddleware) GetPods(clusterCtx string) []model.PodDto {
	x, _ := w.endpoint.GetPods(clusterCtx)
	return x
}

func (w *WorkloadMiddleware) GetPod(name string, namespace string, clusterCtx string) (model.PodDto, error) {
	return w.endpoint.GetPod(name, namespace, clusterCtx)
}

func (w *WorkloadMiddleware) UpdatePod(name string, namespace string, dto model.PodUpdate, clusterCtx string) error {
	return w.endpoint.UpdatePod(name, namespace, dto, clusterCtx)
}

func (w *WorkloadMiddleware) RestartPod(name string, namespace string, clusterCtx string) error {
	return w.endpoint.RestartPod(name, namespace, clusterCtx)
}

func (w *WorkloadMiddleware) GetDeployments(clusterCtx string) []model.DeploymentDto {
	fmt.Println("clusterCtx-GETDEP", clusterCtx)

	x, _ := w.endpoint.GetDeployments(clusterCtx)
	return x
}

func (w *WorkloadMiddleware) GetDeployment(name string, namespace string, clusterCtx string) (model.DeploymentDto, error) {
	return w.endpoint.GetDeployment(name, namespace, clusterCtx)
}

func (w *WorkloadMiddleware) UpdateDeployment(name string, namespace string, dto model.DeploymentUpdate, clusterCtx string) error {
	return w.endpoint.UpdateDeployment(name, namespace, dto, clusterCtx)
}

func (w *WorkloadMiddleware) DeleteDeployment(name string, namespace string, clusterCtx string) error {
	return w.endpoint.DeleteDeployment(name, namespace, clusterCtx)
}

func (w *WorkloadMiddleware) ResourceTuning(namespace string, clusterCtx string) {
	w.endpoint.ResourceTuning(namespace, clusterCtx)
}

func (w *WorkloadMiddleware) AutoTroubleshoot(name string, namespace string, clusterCtx string, resource string) {
	w.endpoint.AutoTroubleshoot(name, namespace, clusterCtx, resource)
}

func (w *WorkloadMiddleware) ExportManifest(name string, namespace string, clusterCtx string) ([]byte, error) {
	return w.endpoint.ExportManifest(name, namespace, clusterCtx)
}

func BuildWorkload(manager kubeclient.ClusterResolver) *WorkloadMiddleware {
	deploymentClient := kubeclient.NewDeployment(manager)
	podClient := kubeclient.NewPod(manager)
	jobClient := kubeclient.NewJob(manager)

	diagnosticService := service.NewDiagnosticService(podClient, deploymentClient, jobClient)

	deploymentUseCase := usecase.NewDeploymentUseCase(deploymentClient, diagnosticService)
	podUseCase := usecase.NewPodUseCase(podClient, diagnosticService)

	workloadEndpoint := endpoint.NewWorkloadEndpoint(podUseCase, deploymentUseCase)

	return NewWorkloadMiddleware(workloadEndpoint)
}
