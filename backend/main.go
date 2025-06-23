package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getPath() string {
	home, _ := os.UserHomeDir()
	kubeConfigPath := filepath.Join(home, ".kube/config")
	return kubeConfigPath
}

func main() {
	fmt.Println("Start Process")

	// Build ClientSet
	//conf := kubeclient.NewClusterManager()
	//client, _ := conf.GetClient("minikube", getPath())

	// Backend test
	//pClient := kubeclient.NewPod(client)
	//dClient := kubeclient.NewDeployment(client)
	//svc := service.NewDiagnosticService(client)

	//ucPoc := usecase.NewPodUseCase(pClient, svc)
	//ucDep := usecase.NewDeploymentUseCase(dClient, svc)

	//workload := endpoint.NewWorkloadEndpoint(ucPoc, ucDep)

	//workload.TroubleshootPod("java-person-7ddd44cfb-zkfzp", "west")
	//workload.TroubleshootDeployment("java-course", "west")
	fmt.Println("-------------")

	fmt.Println("Finish Process")
}
