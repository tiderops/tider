package binding

import (
	"Kubexplorer/internal/apperr"
	"Kubexplorer/internal/k8s"
	"Kubexplorer/internal/model"
	"Kubexplorer/internal/usecase"
)

type General struct {
	app       *App
	node      usecase.NodeUseCase
	namespace usecase.NamespaceUseCase
}

func BuildGeneral(app *App, manager *k8s.ClusterManager) *General {
	return &General{
		app:       app,
		node:      usecase.NewNodeUseCase(k8s.NewNode(manager)),
		namespace: usecase.NewNamespaceUseCase(k8s.NewNamespaceClient(manager)),
	}
}

func (g *General) GetNodes(clusterCtx string) ([]model.NodeDto, error) {
	ctx, cancel := g.app.requestContext()
	defer cancel()

	nodes, err := g.node.GetNodes(ctx, clusterCtx)
	return nodes, apperr.Normalize(err)
}

func (g *General) GetNode(name string, clusterCtx string) (model.NodeDto, error) {
	ctx, cancel := g.app.requestContext()
	defer cancel()

	node, err := g.node.GetNode(ctx, model.ResourceRef{Cluster: clusterCtx, Name: name})
	return node, apperr.Normalize(err)
}

func (g *General) GetNamespaces(clusterCtx string) ([]model.NamespaceDto, error) {
	ctx, cancel := g.app.requestContext()
	defer cancel()

	namespaces, err := g.namespace.GetNamespaces(ctx, clusterCtx)
	return namespaces, apperr.Normalize(err)
}

func (g *General) GetNamespace(name string, clusterCtx string) (model.NamespaceDto, error) {
	ctx, cancel := g.app.requestContext()
	defer cancel()

	namespace, err := g.namespace.GetNamespace(ctx, model.ResourceRef{Cluster: clusterCtx, Name: name})
	return namespace, apperr.Normalize(err)
}

func (g *General) UpdateNamespace(name string, dto model.NamespaceDto, clusterCtx string) error {
	ctx, cancel := g.app.requestContext()
	defer cancel()

	return apperr.Normalize(g.namespace.UpdateNamespace(ctx, model.ResourceRef{Cluster: clusterCtx, Name: name}, dto))
}

func (g *General) DeleteNamespace(name string, clusterCtx string) error {
	ctx, cancel := g.app.requestContext()
	defer cancel()

	return apperr.Normalize(g.namespace.DeleteNamespace(ctx, model.ResourceRef{Cluster: clusterCtx, Name: name}))
}

func (g *General) ExportNamespaceObjects(namespace string, directory string, clusterCtx string) error {
	ctx, cancel := g.app.requestContext()
	defer cancel()

	return apperr.Normalize(g.namespace.ExportNamespaceObjects(ctx, clusterCtx, namespace, directory))
}
