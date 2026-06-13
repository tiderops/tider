package binding

import (
	"Kubexplorer/internal/apperr"
	"Kubexplorer/internal/k8s"
	"Kubexplorer/internal/model"
	"Kubexplorer/internal/usecase"
)

type Environment struct {
	app *App
	env usecase.EnvironmentUseCase
}

func BuildEnvironment(app *App, manager *k8s.ClusterManager) *Environment {
	return &Environment{
		app: app,
		env: usecase.NewEnvironmentUseCase(k8s.NewCluster(manager)),
	}
}

func (e *Environment) GetClusters() ([]model.ClusterInfo, error) {
	ctx, cancel := e.app.requestContext()
	defer cancel()

	clusters, err := e.env.ListAvailableClusters(ctx)
	return clusters, apperr.Normalize(err)
}

func (e *Environment) GetCurrentEnvironment(env string, name string) (model.EnvironmentDto, error) {
	ctx, cancel := e.app.requestContext()
	defer cancel()

	environment, err := e.env.GetCurrentEnvironment(ctx, env, name)
	return environment, apperr.Normalize(err)
}

func (e *Environment) GetObjectsView(clusterCtx string) (model.ObjectMapDto, error) {
	ctx, cancel := e.app.requestContext()
	defer cancel()

	objects, err := e.env.GetObjectsView(ctx, clusterCtx)
	return objects, apperr.Normalize(err)
}
