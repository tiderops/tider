package service

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
)

type ResourceService interface {
	ResourceTuning(namespace string, clusterCtx string)
}

type resourceService struct {
	deployment kubeclient.DeploymentClient
	metric     kubeclient.MetricClient
}

func NewResourceService(client kubeclient.DeploymentClient, metric kubeclient.MetricClient) ResourceService {
	return &resourceService{deployment: client, metric: metric}
}

// TODO: improve resource tuning logic
func (r *resourceService) ResourceTuning(namespace string, clusterCtx string) {
	metrics, err := r.metric.GetPodMetricsV2(namespace, clusterCtx)
	if err != nil {
		fmt.Errorf("Error getting metrics: %v", err)
	}

	for _, metric := range metrics.Items {
		name := metric.Labels["app"]
		fmt.Println("name:: ", name)

		d, err := r.deployment.GetDeployment(name, namespace, clusterCtx)
		if err != nil {
			fmt.Errorf("Error getting pods: %v", err)
		}

		cpu := tuningCpu(d, metric.Containers[0].Usage.Cpu().MilliValue())
		memory := tuningMemory(d, metric.Containers[0].Usage.Memory().ScaledValue(resource.Mega))

		update := model.DeploymentUpdate{
			Container: model.ContainerUpdate{
				Resource: model.ResourceUpdate{
					LCpu:    cpu,
					LMemory: memory,
				},
			},
		}

		err = r.deployment.UpdateDeployment(name, namespace, update, clusterCtx)

		if err != nil {
			fmt.Errorf("Error updating deployment: %v", err)
		}
	}
}

func tuningCpu(pod model.DeploymentDto, currentCpu int64) int64 {
	baseCpu := pod.Container.Limit.Cpu
	newCpu := currentCpu

	newCpu = increaseCPU(baseCpu, currentCpu)
	newCpu = decreaseCPU(baseCpu, currentCpu)
	fmt.Println("baseCpu: ", baseCpu)
	fmt.Println("currentCpu: ", currentCpu)
	fmt.Println("newCpu: ", newCpu)

	pod.Container.Limit.Cpu = newCpu

	return newCpu
}

func tuningMemory(pod model.DeploymentDto, currentMemory int64) int64 {
	baseMemory := pod.Container.Limit.Memory
	newMemory := currentMemory

	newMemory = increaseMemory(baseMemory, currentMemory)
	newMemory = decreaseMemory(baseMemory, currentMemory)
	fmt.Println("baseMemory: ", baseMemory)
	fmt.Println("currentMemory: ", currentMemory)
	fmt.Println("newMemory: ", newMemory)

	fmt.Println("---------------")

	pod.Container.Limit.Memory = newMemory

	return newMemory
}

func increaseCPU(baseCpu int64, currentCpu int64) int64 {
	ninetyThreshold := (baseCpu * 90) / 100
	eightyThreshold := (baseCpu * 80) / 100

	var newCpu = currentCpu

	if currentCpu > ninetyThreshold {
		newCpu = baseCpu + ((baseCpu * 50) / 100)
	}

	if currentCpu >= eightyThreshold && currentCpu < ninetyThreshold {
		newCpu = baseCpu + ((baseCpu * 30) / 100)
	}

	return newCpu
}

func decreaseCPU(baseCpu int64, currentCpu int64) int64 {
	threshold := baseCpu - ((baseCpu * 70) / 100)

	var newCpu = currentCpu

	if currentCpu < threshold {
		newCpu = baseCpu - ((baseCpu * 50) / 100)
	}

	return newCpu
}

func increaseMemory(baseMemory int64, currentMemory int64) int64 {
	ninetyThreshold := (baseMemory * 90) / 100
	eightyThreshold := (baseMemory * 80) / 100

	var newCpu = currentMemory

	if currentMemory > ninetyThreshold {
		newCpu = baseMemory + ((baseMemory * 50) / 100)
	}

	if currentMemory >= eightyThreshold && currentMemory < ninetyThreshold {
		newCpu = baseMemory + ((baseMemory * 30) / 100)
	}

	return newCpu
}

func decreaseMemory(baseMemory int64, currentMemory int64) int64 {
	threshold := baseMemory - ((baseMemory * 70) / 100)

	var newCpu = currentMemory

	if currentMemory < threshold {
		newCpu = baseMemory - ((baseMemory * 50) / 100)
	}

	return newCpu
}
