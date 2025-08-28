package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
)

type IngressUseCase interface {
	GetIngresses(clusterCtx string) ([]model.IngressDto, error)
	GetIngress(name string, namespace string, clusterCtx string) (model.IngressDto, error)
	UpdateIngress(name string, namespace string, dto model.IngressDto, clusterCtx string) error
	DeleteIngress(name string, namespace string, clusterCtx string) error
	ExportManifest(name string, namespace string, clusterCtx string) ([]byte, error)
}

type ingressUseCase struct {
	client kubeclient.IngressClient
}

func NewIngressUseCase(client kubeclient.IngressClient) IngressUseCase {
	return &ingressUseCase{client: client}
}

func (i *ingressUseCase) GetIngresses(clusterCtx string) ([]model.IngressDto, error) {
	return i.client.GetIngresses(clusterCtx)
}

func (i *ingressUseCase) GetIngress(name string, namespace string, clusterCtx string) (model.IngressDto, error) {
	return i.client.GetIngress(name, namespace, clusterCtx)
}

func (i *ingressUseCase) UpdateIngress(name string, namespace string, dto model.IngressDto, clusterCtx string) error {
	return i.client.UpdateIngress(name, namespace, dto, clusterCtx)
}

func (i *ingressUseCase) DeleteIngress(name string, namespace string, clusterCtx string) error {
	return i.client.DeleteIngress(name, namespace, clusterCtx)
}

func (i *ingressUseCase) ExportManifest(name string, namespace string, clusterCtx string) ([]byte, error) {
	return i.client.ExportManifest(name, namespace, clusterCtx)
}
