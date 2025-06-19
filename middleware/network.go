package middleware

import (
	"Kubexplorer/backend/endpoint"
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
	"k8s.io/client-go/kubernetes"
)

type NetworkMiddleware struct {
	endpoint endpoint.NetworkEndpoint
}

func NewNetworkMiddleware(endpoint *endpoint.NetworkEndpoint) *NetworkMiddleware {
	return &NetworkMiddleware{endpoint: *endpoint}
}

func (n *NetworkMiddleware) GetServices(namespace string) ([]model.ServiceDto, error) {
	return n.endpoint.GetServices(namespace)
}

func (n *NetworkMiddleware) GetService(name string, namespace string) (model.ServiceDto, error) {
	return n.endpoint.GetService(name, namespace)
}

func (n *NetworkMiddleware) UpdateService(name string, namespace string, dto model.ServiceDto) error {
	return n.endpoint.UpdateService(name, namespace, dto)
}

func (n *NetworkMiddleware) DeleteService(name string, namespace string) error {
	return n.endpoint.DeleteService(name, namespace)
}

func (n *NetworkMiddleware) GetIngresses(namespace string) ([]model.IngressDto, error) {
	return n.endpoint.GetIngresses(namespace)
}

func (n *NetworkMiddleware) GetIngress(name string, namespace string) (model.IngressDto, error) {
	return n.endpoint.GetIngress(name, namespace)
}

func (n *NetworkMiddleware) UpdateIngress(name string, namespace string, dto model.IngressDto) error {
	return n.endpoint.UpdateIngress(name, namespace, dto)
}

func (n *NetworkMiddleware) DeleteIngress(name string, namespace string) error {
	return n.endpoint.DeleteIngress(name, namespace)
}

func BuildNetwork(client kubernetes.Interface) *NetworkMiddleware {
	serviceClient := kubeclient.NewServiceClient(client)
	ingressClient := kubeclient.NewIngressClient(client)

	serviceUseCase := usecase.NewServiceUseCase(serviceClient)
	ingressUseCase := usecase.NewIngressUseCase(ingressClient)

	networkEndpoint := endpoint.NewNetworkEndpoint(serviceUseCase, ingressUseCase)

	return NewNetworkMiddleware(networkEndpoint)
}
