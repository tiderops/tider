package main

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/service"
	"fmt"
)

func main() {
	fmt.Println("Start Process")

	clusters := kubeclient.NewCluster()
	clusters.ListAvailableClusters()

	manager := kubeclient.GlobalClusterManager
	//d := kubeclient.NewDeployment(manager)
	//s := kubeclient.NewServiceClient(manager)
	//i := kubeclient.NewIngressClient(manager)
	p := kubeclient.NewDeployment(manager)
	m := kubeclient.NewMetric(manager)
	rs := service.NewResourceService(p, m)

	rs.ResourceTuning("west", "minikube")

	//d.ExportManifest("java-person", "west", "minikube")
	//fmt.Println("-------------")
	//s.ExportManifest("java-person-service-private", "default", "minikube")
	//fmt.Println("-------------")
	//i.ExportManifest("example-ingress", "default", "minikube")
	//fmt.Println("-------------")
	//x, _ := d.GetDeployment("java-person", "west", "minikube")
	//fmt.Println(x.Age)

	fmt.Println("-------------")
	fmt.Println("Finish Process")
}
