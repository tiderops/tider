package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/service"
)

type NamespaceUseCase interface {
	GetNamespaces(clusterCtx string) ([]model.NamespaceDto, error)
	GetNamespace(name string, clusterCtx string) (model.NamespaceDto, error)
	UpdateNamespace(name string, dto model.NamespaceDto, clusterCtx string) error
	DeleteNamespace(name string, clusterCtx string) error
	ExportNamespaceObjects(namespace string, directory string, clusterCtx string) error
}

type namespaceUseCase struct {
	client  kubeclient.NamespaceClient
	service service.NamespaceService
}

func NewNamespaceUseCase(client kubeclient.NamespaceClient) NamespaceUseCase {
	return &namespaceUseCase{client: client}
}

func (n *namespaceUseCase) GetNamespaces(clusterCtx string) ([]model.NamespaceDto, error) {
	return n.client.GetNamespaces(clusterCtx)
}
func (n *namespaceUseCase) GetNamespace(name string, clusterCtx string) (model.NamespaceDto, error) {
	return n.client.GetNamespace(name, clusterCtx)
}

func (n *namespaceUseCase) UpdateNamespace(name string, dto model.NamespaceDto, clusterCtx string) error {
	return n.client.UpdateNamespace(name, dto, clusterCtx)
}

func (n *namespaceUseCase) DeleteNamespace(name string, clusterCtx string) error {
	return n.client.DeleteNamespace(name, clusterCtx)
}

func (n *namespaceUseCase) ExportNamespaceObjects(namespace string, directory string, clusterCtx string) error {
	return n.service.ExportObjects(namespace, directory, clusterCtx)
}
