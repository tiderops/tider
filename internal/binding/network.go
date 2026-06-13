package binding

import (
	"Kubexplorer/internal/apperr"
	"Kubexplorer/internal/k8s"
	"Kubexplorer/internal/model"
	"Kubexplorer/internal/usecase"
)

type Network struct {
	app     *App
	service usecase.ServiceUseCase
	ingress usecase.IngressUseCase
}

func BuildNetwork(app *App, manager *k8s.ClusterManager) *Network {
	return &Network{
		app:     app,
		service: usecase.NewServiceUseCase(k8s.NewServiceClient(manager)),
		ingress: usecase.NewIngressUseCase(k8s.NewIngressClient(manager)),
	}
}

func (n *Network) GetServices(clusterCtx string) ([]model.ServiceDto, error) {
	ctx, cancel := n.app.requestContext()
	defer cancel()

	services, err := n.service.GetServices(ctx, clusterCtx)
	return services, apperr.Normalize(err)
}

func (n *Network) GetService(name string, namespace string, clusterCtx string) (model.ServiceDto, error) {
	ctx, cancel := n.app.requestContext()
	defer cancel()

	svc, err := n.service.GetService(ctx, model.ResourceRef{Cluster: clusterCtx, Namespace: namespace, Name: name})
	return svc, apperr.Normalize(err)
}

func (n *Network) UpdateService(name string, namespace string, dto model.ServiceUpdate, clusterCtx string) error {
	ctx, cancel := n.app.requestContext()
	defer cancel()

	return apperr.Normalize(n.service.UpdateService(ctx, model.ResourceRef{Cluster: clusterCtx, Namespace: namespace, Name: name}, dto))
}

func (n *Network) DeleteService(name string, namespace string, clusterCtx string) error {
	ctx, cancel := n.app.requestContext()
	defer cancel()

	return apperr.Normalize(n.service.DeleteService(ctx, model.ResourceRef{Cluster: clusterCtx, Namespace: namespace, Name: name}))
}

func (n *Network) GetIngresses(clusterCtx string) ([]model.IngressDto, error) {
	ctx, cancel := n.app.requestContext()
	defer cancel()

	ingresses, err := n.ingress.GetIngresses(ctx, clusterCtx)
	return ingresses, apperr.Normalize(err)
}

func (n *Network) GetIngress(name string, namespace string, clusterCtx string) (model.IngressDto, error) {
	ctx, cancel := n.app.requestContext()
	defer cancel()

	ingress, err := n.ingress.GetIngress(ctx, model.ResourceRef{Cluster: clusterCtx, Namespace: namespace, Name: name})
	return ingress, apperr.Normalize(err)
}

func (n *Network) UpdateIngress(name string, namespace string, dto model.IngressDto, clusterCtx string) error {
	ctx, cancel := n.app.requestContext()
	defer cancel()

	return apperr.Normalize(n.ingress.UpdateIngress(ctx, model.ResourceRef{Cluster: clusterCtx, Namespace: namespace, Name: name}, dto))
}

func (n *Network) DeleteIngress(name string, namespace string, clusterCtx string) error {
	ctx, cancel := n.app.requestContext()
	defer cancel()

	return apperr.Normalize(n.ingress.DeleteIngress(ctx, model.ResourceRef{Cluster: clusterCtx, Namespace: namespace, Name: name}))
}

func (n *Network) ExportManifest(name string, namespace string, clusterCtx string, object string) ([]byte, error) {
	ctx, cancel := n.app.requestContext()
	defer cancel()

	var data []byte
	var err error
	if object == "service" {
		data, err = n.service.ExportManifest(ctx, model.ResourceRef{Cluster: clusterCtx, Namespace: namespace, Name: name})
	} else {
		data, err = n.ingress.ExportManifest(ctx, model.ResourceRef{Cluster: clusterCtx, Namespace: namespace, Name: name})
	}
	return data, apperr.Normalize(err)
}
