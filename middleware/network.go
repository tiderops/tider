package middleware

import (
	"Kubexplorer/backend/endpoint"
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
)

type NetworkMiddleware struct {
	endpoint endpoint.NetworkEndpoint
}

func NewNetworkMiddleware(endpoint *endpoint.NetworkEndpoint) *NetworkMiddleware {
	return &NetworkMiddleware{endpoint: *endpoint}
}

func (n *NetworkMiddleware) GetServices(clusterCtx string) ([]model.ServiceDto, error) {
	return n.endpoint.GetServices(clusterCtx)
}

func (n *NetworkMiddleware) GetService(name string, namespace string, clusterCtx string) (model.ServiceDto, error) {
	return n.endpoint.GetService(name, namespace, clusterCtx)
}

func (n *NetworkMiddleware) UpdateService(name string, namespace string, dto model.ServiceUpdate, clusterCtx string) error {
	return n.endpoint.UpdateService(name, namespace, dto, clusterCtx)
}

func (n *NetworkMiddleware) DeleteService(name string, namespace string, clusterCtx string) error {
	return n.endpoint.DeleteService(name, namespace, clusterCtx)
}

func (n *NetworkMiddleware) GetIngresses(clusterCtx string) ([]model.IngressDto, error) {
	return n.endpoint.GetIngresses(clusterCtx)
}

func (n *NetworkMiddleware) GetIngress(name string, namespace string, clusterCtx string) (model.IngressDto, error) {
	return n.endpoint.GetIngress(name, namespace, clusterCtx)
}

func (n *NetworkMiddleware) UpdateIngress(name string, namespace string, dto model.IngressDto, clusterCtx string) error {
	return n.endpoint.UpdateIngress(name, namespace, dto, clusterCtx)
}

func (n *NetworkMiddleware) DeleteIngress(name string, namespace string, clusterCtx string) error {
	return n.endpoint.DeleteIngress(name, namespace, clusterCtx)
}

func (n *NetworkMiddleware) ExportManifest(name string, namespace string, clusterCtx string, object string) ([]byte, error) {
	return n.endpoint.ExportManifest(name, namespace, clusterCtx, object)
}

func BuildNetwork(manager kubeclient.ClusterResolver) *NetworkMiddleware {
	serviceClient := kubeclient.NewServiceClient(manager)
	ingressClient := kubeclient.NewIngressClient(manager)

	serviceUseCase := usecase.NewServiceUseCase(serviceClient)
	ingressUseCase := usecase.NewIngressUseCase(ingressClient)

	networkEndpoint := endpoint.NewNetworkEndpoint(serviceUseCase, ingressUseCase)

	return NewNetworkMiddleware(networkEndpoint)
}
