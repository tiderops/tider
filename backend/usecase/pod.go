package usecase

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/service"
)

type PodUseCase interface {
	GetPods(clusterCtx string) ([]model.PodDto, error)
	GetPod(name string, namespace string, clusterCtx string) (model.PodDto, error)
	UpdatePod(name string, namespace string, dto model.PodRequest, clusterCtx string) error
	RestartPod(name string, namespace string, clusterCtx string) error
	TroubleshootPod(name string, namespace string, clusterCtx string)
}

type podUseCase struct {
	client  kubeclient.PodClient
	service service.DiagnosticService
}

func NewPodUseCase(client kubeclient.PodClient, service service.DiagnosticService) PodUseCase {
	return &podUseCase{client: client, service: service}
}

func (p *podUseCase) GetPods(clusterCtx string) ([]model.PodDto, error) {
	return p.client.GetPods(clusterCtx)
}

func (p *podUseCase) GetPod(name string, namespace string, clusterCtx string) (model.PodDto, error) {
	return p.client.GetPod(name, namespace, clusterCtx)
}

func (p *podUseCase) UpdatePod(name string, namespace string, dto model.PodRequest, clusterCtx string) error {
	return p.client.UpdatePod(name, namespace, dto, clusterCtx)
}

func (p *podUseCase) RestartPod(name string, namespace string, clusterCtx string) error {
	return p.client.DeletePod(name, namespace, clusterCtx)
}

func (p *podUseCase) TroubleshootPod(name string, namespace string, clusterCtx string) {
	p.service.Analyse(name, namespace, service.Pod)
}
