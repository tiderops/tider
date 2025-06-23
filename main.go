package main

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/middleware"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

const (
	PROGRAM_NAME = "Kubexplorer"
	WIDTH        = 1024
	HEIGHT       = 768
)

func main() {
	// Create an instance of the app structure
	app := middleware.NewApp()

	manager := kubeclient.GlobalClusterManager

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
			middleware.BuildEnvironment(),
			middleware.BuildParameters(),
			middleware.BuildGeneral(manager),
			middleware.BuildNetwork(manager),
			middleware.BuildStorage(manager),
			middleware.BuildWorkload(manager),
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
