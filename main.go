package main

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/middleware"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"os"
	"path/filepath"
)

//go:embed all:frontend/dist
var assets embed.FS

const (
	PROGRAM_NAME = "Kubexplorer"
	WIDTH        = 1024
	HEIGHT       = 768
)

func getPath() string {
	home, _ := os.UserHomeDir()
	kubeConfigPath := filepath.Join(home, ".kube/config")

	return kubeConfigPath
}

func main() {
	//clusterManager := kubeclient.NewClusterManager()
	//client, _ := clusterManager.GetClient("dev", "/.kube/config")

	//podEp := bootstrap2.Execute()
	//parameterEp := bootstrap2.ParameterEp()
	//environmentEp := bootstrap2.EnvironmentEp()
	//serviceEp := bootstrap2.ServiceEp()
	//nodeEp := bootstrap2.NodeEp()
	//storageEp := bootstrap2.StorageEp()
	//metricEp := bootstrap2.MetricEp()

	//pc := di.SetupWorkloadContainer()
	//podEp := pc.MustResolve("IWorkloadEndpoint").(endpoint.IWorkloadEndpoint)
	//
	//pr := di.SetupParameterContainer()
	//parameterEp := pr.MustResolve("IParameterEndpoint").(endpoint.IParameterEndpoint)
	//
	//e := di.SetupEnvironmentContainer()
	//environmentEp := e.MustResolve("IEnvironmentEndpoint").(endpoint.IEnvironmentEndpoint)
	//
	//s := di.SetupNetworkContainer()
	//serviceEp := s.MustResolve("INetworkEndpoint").(endpoint.INetworkEndpoint)
	//
	//n := di.SetupGeneralContainer()
	//nodeEp := n.MustResolve("IGeneralEndpoint").(endpoint.IGeneralEndpoint)
	//
	//stg := di.SetupStorageContainer()
	//storageEp := stg.MustResolve("IStorageEndpoint").(endpoint.IStorageEndpoint)
	//
	//mtc := di.SetupMetricsContainer2()
	//metricEp := mtc.MustResolve("IMetricEndpoint").(endpoint.IMetricEndpoint)
	//
	//client := kubeclient.NewPod()
	//client2 := kubeclient.NewDeployment()
	//usecase := usecase2.NewPodUseCase(client)
	//usecase2 := usecase2.NewDeploymentUseCase(client2)
	//ep := endpoint.NewWorkloadEndpoint(usecase, usecase2)

	// KubeClient - INSTANCIA DE MANAGER
	conf := kubeclient.NewClusterManager()
	client, _ := conf.GetClient("minikube", getPath())
	//dynamicClient, _ := conf.GetDynamicClient("minikube", getPath()))
	//conf.GetValue("minikube")

	// Create an instance of the app structure
	app := middleware.NewApp()

	// Build
	//bw := middleware.BuildWorkload(client)
	//bs := middleware.BuildStorage(client)

	// Instance
	//workloadMiddleware := middleware.NewWorkloadMiddleware(nil)
	//parameterMiddleware := middleware.NewParameterMiddleware(nil)
	//environmentMiddleware := middleware.NewEnvironmentMiddleware(nil)
	//networkMiddleware := middleware.NewNetworkMiddleware(nil)
	//generalMiddleware := middleware.NewGeneralMiddleware(nil)
	//storageMiddleware := middleware.NewStorageMiddleware(nil)
	//metricMiddleware := middleware.NewMetricMiddleware(nil)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  PROGRAM_NAME,
		Width:  WIDTH,
		Height: HEIGHT,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		OnShutdown:       app.Shutdown,
		Bind: []interface{}{
			app,
			middleware.BuildEnvironment(client),
			middleware.BuildGeneral(client),
			middleware.BuildNetwork(client),
			middleware.BuildParameters(client),
			middleware.BuildStorage(client),
			middleware.BuildWorkload(client),
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
