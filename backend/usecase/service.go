package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
)

type ServiceUseCase interface {
	GetServices(clusterCtx string) ([]model.ServiceDto, error)
	GetService(name string, namespace string, clusterCtx string) (model.ServiceDto, error)
	UpdateService(name string, namespace string, dto model.ServiceUpdate, clusterCtx string) error
	DeleteService(name string, namespace string, clusterCtx string) error
	ExportManifest(name string, namespace string, clusterCtx string) ([]byte, error)
}

type serviceUseCase struct {
	client kubeclient.ServiceClient
}

func NewServiceUseCase(client kubeclient.ServiceClient) ServiceUseCase {
	return &serviceUseCase{client: client}
}

func (s *serviceUseCase) GetServices(clusterCtx string) ([]model.ServiceDto, error) {
	return s.client.GetServices(clusterCtx)
}

func (s *serviceUseCase) GetService(name string, namespace string, clusterCtx string) (model.ServiceDto, error) {
	return s.client.GetService(name, namespace, clusterCtx)
}

func (s *serviceUseCase) UpdateService(name string, namespace string, dto model.ServiceUpdate, clusterCtx string) error {
	return s.client.UpdateService(name, namespace, dto, clusterCtx)
}

func (s *serviceUseCase) DeleteService(name string, namespace string, clusterCtx string) error {
	return s.client.DeleteService(name, namespace, clusterCtx)
}

func (s *serviceUseCase) ExportManifest(name string, namespace string, clusterCtx string) ([]byte, error) {
	return s.client.ExportManifest(name, namespace, clusterCtx)
}
