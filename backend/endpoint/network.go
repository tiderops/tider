package endpoint

import (
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
)

type NetworkEndpoint struct {
	serviceUseCase usecase.ServiceUseCase
	ingressUseCase usecase.IngressUseCase
}

func NewNetworkEndpoint(serviceUseCase usecase.ServiceUseCase, ingressUseCase usecase.IngressUseCase) *NetworkEndpoint {
	return &NetworkEndpoint{serviceUseCase: serviceUseCase, ingressUseCase: ingressUseCase}
}

func (ne *NetworkEndpoint) GetServices(clusterCtx string) ([]model.ServiceDto, error) {
	return ne.serviceUseCase.GetServices(clusterCtx)
}

func (ne *NetworkEndpoint) GetService(name string, namespace string, clusterCtx string) (model.ServiceDto, error) {
	return ne.serviceUseCase.GetService(name, namespace, clusterCtx)
}

func (ne *NetworkEndpoint) UpdateService(name string, namespace string, dto model.ServiceUpdate, clusterCtx string) error {
	return ne.serviceUseCase.UpdateService(name, namespace, dto, clusterCtx)
}

func (ne *NetworkEndpoint) DeleteService(name string, namespace string, clusterCtx string) error {
	return ne.serviceUseCase.DeleteService(name, namespace, clusterCtx)
}

func (ne *NetworkEndpoint) GetIngresses(clusterCtx string) ([]model.IngressDto, error) {
	return ne.ingressUseCase.GetIngresses(clusterCtx)
}

func (ne *NetworkEndpoint) GetIngress(name string, namespace string, clusterCtx string) (model.IngressDto, error) {
	return ne.ingressUseCase.GetIngress(name, namespace, clusterCtx)
}

func (ne *NetworkEndpoint) UpdateIngress(name string, namespace string, dto model.IngressDto, clusterCtx string) error {
	return ne.ingressUseCase.UpdateIngress(name, namespace, dto, clusterCtx)
}

func (ne *NetworkEndpoint) DeleteIngress(name string, namespace string, clusterCtx string) error {
	return ne.ingressUseCase.DeleteIngress(name, namespace, clusterCtx)
}
