package service

import (
	"Kubexplorer/backend/model"
	"Kubexplorer/backend/repository"
)

type MetricV2Service interface {
	EnableMetrics(dto model.MetricEnableDto)
	CollectMetrics()
	GetMetricsStatus(namespace string)
}

type metricV2Service struct {
	repository repository.MetricRepository
}

func NewMetricV2Service(repository repository.MetricRepository) MetricV2Service {
	return &metricV2Service{repository: repository}
}

func (m metricV2Service) EnableMetrics(dto model.MetricEnableDto) {

	if dto.EnablePods == true {
		m.repository.UpdatePodMetricsState(dto.EnablePods, dto.Namespace)
	}

	if dto.EnableNode == true {
		m.repository.UpdateNodeMetricsState(dto.EnableNode)
	}
}

func (m metricV2Service) CollectMetrics() {
	//TODO implement me
	panic("implement me")
	
}

func (m metricV2Service) GetMetricsStatus(namespace string) {
	podStatus, nodeStatus := m.repository.SearchMetricStatus(namespace)

	// gorutine for pods
	if podStatus != true {
		timer := 15

		m.repository.PersistPodMetrics(model.PodMetric{})
	}

	// gorutine for node
	if nodeStatus != true {
		timer := 15

		m.repository.PersistNodeMetrics(model.NodeMetric{})
	}

}
