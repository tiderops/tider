package model

type MetricEnableDto struct {
	EnablePods bool
	Namespace  string
	EnableNode bool
}

type PodMetric struct {
}

type NodeMetric struct {
}
