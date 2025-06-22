package service

import (
	"Kubexplorer/backend/kubeclient"
	"Kubexplorer/backend/model"
	"fmt"
	"strconv"
)

type ResourceService interface {
	ResourceTuning(namespace string) []model.PodDto
}

type resourceService struct {
	pod    kubeclient.PodClient
	metric kubeclient.MetricClient
}

func NewResourceService(client kubeclient.PodClient) ResourceService {
	return &resourceService{pod: client}
}

func (r *resourceService) ResourceTuning(namespace string) []model.PodDto {
	metrics := r.metric.GetPodMetricsV2(namespace)

	var result []model.PodDto

	for _, metric := range metrics.Items {
		pod, err := r.pod.GetPod(metric.Name, metric.Namespace, "")
		if err != nil {
			fmt.Errorf("Error getting pods: %v", err)
		}
		tuningCpu(pod, metric.Containers[0].Usage.Cpu().Value())
		tuningMemory(pod, metric.Containers[0].Usage.Memory().Value())

		result = append(result, pod)
	}

	return result
}

func tuningCpu(pod model.PodDto, currentCpu int64) model.PodDto {
	baseCpu, err := strconv.ParseInt(pod.Container.Limit.Cpu, 10, 64)
	if err != nil {
		fmt.Errorf("Error parsing cpu limit: %v", err)
	}

	newCpu := currentCpu

	newCpu = scaleUpCPU(baseCpu, currentCpu)
	newCpu = scaleDownCPU(baseCpu, currentCpu)

	pod.Container.Limit.Cpu = strconv.FormatInt(newCpu, 10)

	return pod
}

func tuningMemory(pod model.PodDto, currentMemory int64) model.PodDto {
	baseMemory, err := strconv.ParseInt(pod.Container.Limit.Memory, 10, 64)
	if err != nil {
		fmt.Errorf("Error parsing memory limit: %v", err)
	}

	newMemory := currentMemory

	newMemory = scaleUpMemory(baseMemory, currentMemory)
	newMemory = scaleDownMemory(baseMemory, currentMemory)

	pod.Container.Limit.Memory = strconv.FormatInt(newMemory, 10)

	return pod
}

func scaleUpCPU(baseCpu int64, currentCpu int64) int64 {
	ninetyThreshold := (baseCpu * 90) / 100
	eightyThreshold := (baseCpu * 80) / 100

	var newCpu = currentCpu

	if currentCpu > ninetyThreshold {
		newCpu = baseCpu + ((baseCpu * 50) / 100)
	}

	if currentCpu > eightyThreshold && currentCpu < ninetyThreshold {
		newCpu = baseCpu + ((baseCpu * 30) / 100)
	}

	return newCpu
}

func scaleDownCPU(baseCpu int64, currentCpu int64) int64 {
	threshold := currentCpu - ((currentCpu * 70) / 100)

	var newCpu = currentCpu

	if currentCpu < threshold {
		newCpu = baseCpu - ((baseCpu * 30) / 100)
	}

	return newCpu
}

func scaleUpMemory(baseMemory int64, currentMemory int64) int64 {
	ninetyThreshold := (baseMemory * 90) / 100
	eightyThreshold := (baseMemory * 80) / 100

	var newCpu = currentMemory

	if currentMemory > ninetyThreshold {
		newCpu = baseMemory + ((baseMemory * 50) / 100)
	}

	if currentMemory > eightyThreshold && currentMemory < ninetyThreshold {
		newCpu = baseMemory + ((baseMemory * 30) / 100)
	}

	return newCpu
}

func scaleDownMemory(baseMemory int64, currentMemory int64) int64 {
	threshold := currentMemory - ((currentMemory * 70) / 100)

	var newCpu = currentMemory

	if currentMemory < threshold {
		newCpu = baseMemory - ((baseMemory * 30) / 100)
	}

	return newCpu
}
