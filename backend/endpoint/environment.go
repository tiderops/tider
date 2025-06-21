package endpoint

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/usecase"
	"fmt"
)

type EnvironmentEndpoint struct {
	useCase usecase.EnvironmentUseCase
}

func NewEnvironmentEndpoint(useCase usecase.EnvironmentUseCase) *EnvironmentEndpoint {
	return &EnvironmentEndpoint{useCase: useCase}
}

func (ee *EnvironmentEndpoint) GetClusters() ([]model.ClusterInfo, error) {
	return ee.useCase.ListAvailableClusters()
}

func (ee *EnvironmentEndpoint) GetCurrentEnvironment(env string, name string) (model.EnvironmentDto, error) {
	return ee.useCase.GetCurrentEnvironment(env, name)
}

func (ee *EnvironmentEndpoint) GetObjectsView() (model.ObjectMapDto, error) {
	return ee.useCase.GetObjectsView()
}

func (ee *EnvironmentEndpoint) CreateClusterManager(clusters []Cluster) {
	for _, c := range clusters {
		conf := kubeclient.NewClusterManager()
		_, err := conf.GetClient(c.name, c.path)
		if err != nil {
			return
		}
		fmt.Println("Clusters already registered")
	}
}

type Cluster struct {
	name string
	path string
}
