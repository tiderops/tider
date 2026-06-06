package repository

import "Kubexplorer/backend/model"

type MetricRepository interface {
	UpdatePodMetricsState(flag bool, namespace string)
	UpdateNodeMetricsState(flag bool)
	SearchMetricStatus(namespace string) (podsFlag bool, nodesFlag bool)
	PersistPodMetrics(metric model.PodMetric)
	PersistNodeMetrics(metric model.NodeMetric)
}

type metricRepository struct {
	// database Conn
}

func NewMetricRepository() MetricRepository {
	return &metricRepository{}
}

func (m metricRepository) UpdatePodMetricsState(flag bool, namespace string) {
	//TODO implement me
	panic("implement me")
}

func (m metricRepository) UpdateNodeMetricsState(flag bool) {
	//TODO implement me
	panic("implement me")
}

func (m metricRepository) SearchMetricStatus(namespace string) (podsFlag bool, nodesFlag bool) {

	panic("implement me")

	return true, true
}

func (m metricRepository) PersistPodMetrics(metric model.PodMetric) {
	//TODO implement me
	panic("implement me")
}

func (m metricRepository) PersistNodeMetrics(metric model.NodeMetric) {
	//TODO implement me
	panic("implement me")
}
