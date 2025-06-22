package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/service"
)

type DeploymentUseCase interface {
	GetDeployments(clusterCtx string) ([]model.DeploymentDto, error)
	GetDeployment(name string, namespace string, clusterCtx string) (model.DeploymentDto, error)
	UpdateDeployment(name string, namespace string, dto model.DeploymentDto, clusterCtx string) error
	DeleteDeployment(name string, namespace string, clusterCtx string) error
	TroubleshootDeployment(name string, namespace string, clusterCtx string)
}

type deploymentUseCase struct {
	client  kubeclient.DeploymentClient
	service service.DiagnosticService
}

func NewDeploymentUseCase(client kubeclient.DeploymentClient, service service.DiagnosticService) DeploymentUseCase {
	return &deploymentUseCase{client: client, service: service}
}

func (d *deploymentUseCase) GetDeployments(clusterCtx string) ([]model.DeploymentDto, error) {
	return d.client.GetDeployments(clusterCtx)
}

func (d *deploymentUseCase) GetDeployment(name string, namespace string, clusterCtx string) (model.DeploymentDto, error) {
	return d.client.GetDeployment(name, namespace, clusterCtx)
}

func (d *deploymentUseCase) UpdateDeployment(name string, namespace string, dto model.DeploymentDto, clusterCtx string) error {
	return d.client.UpdateDeployment(name, namespace, dto, clusterCtx)
}

func (d *deploymentUseCase) DeleteDeployment(name string, namespace string, clusterCtx string) error {
	return d.client.DeleteDeployment(name, namespace, clusterCtx)
}

func (d *deploymentUseCase) TroubleshootDeployment(name string, namespace string, clusterCtx string) {
	d.service.Analyse(name, namespace, service.Deployment)
}
