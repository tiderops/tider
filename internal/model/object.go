package model

import core "k8s.io/api/core/v1"
import app "k8s.io/api/apps/v1"

type ObjectMapDto struct {
	ClusterInfo            ClusterInfo
	Nodes                  []Node
	Namespaces             []Namespace
	Pods                   []Pod
	Deployments            []Deployment
	ReplicaSets            []ReplicaSet
	Services               []Service
	Ingresses              []Ingress
	PersistentVolumeClaims []PersistentVolumeClaim
	PersistentVolumes      []PersistentVolume
	ConfigMaps             []ConfigMap
	Secrets                []Secret
	Jobs                   []Job
	CronJobs               []CronJob
}

type Pod struct {
	Name          string
	Namespace     string
	Labels        map[string]string // Labels to match with services
	NodeName      string
	OwnerKind     string // ReplicaSet, Job, CronJob, etc.
	OwnerName     string
	Deployment    string // optional, resolved from ReplicaSet
	PVCNames      []string
	ConfigMapRefs []string
	SecretRefs    []string
	ServiceRefs   []string
}

type Deployment struct {
	Name      string
	Namespace string
	PodNames  []string
}

type Service struct {
	Name      string
	Namespace string
	Selector  map[string]string
	PodNames  []string // matched by selector
}

type Ingress struct {
	Name        string
	Namespace   string
	ServiceRefs []string // backend services
}

type PersistentVolumeClaim struct {
	Name       string
	Namespace  string
	VolumeName string
	UsedByPods []string
}

type PersistentVolume struct {
	Name      string
	ClaimName string // PVC name
	Namespace string // useful to resolve claim
}

type Job struct {
	Name      string
	Namespace string
	PodNames  []string
	OwnerKind string // usually "CronJob"
	OwnerName string
}

type CronJob struct {
	Name      string
	Namespace string
	JobNames  []string
}

type Namespace struct {
	Name         string
	PodNames     []string
	ServiceNames []string
}

type Node struct {
	Name     string
	PodNames []string
}

type ConfigMap struct {
	Name       string
	Namespace  string
	UsedByPods []string
}

type Secret struct {
	Name       string
	Namespace  string
	UsedByPods []string
}

type ReplicaSet struct {
	Name       string
	Namespace  string
	Deployment string
	PodNames   []string
}

// Mappers

func MapPodsToDto(pods *core.PodList) []Pod {
	var result []Pod

	for _, p := range pods.Items {
		pod := Pod{
			Name:      p.Name,
			Namespace: p.Namespace,
			NodeName:  p.Spec.NodeName,
		}

		for _, ownerRef := range p.OwnerReferences {
			if ownerRef.Controller != nil && *ownerRef.Controller {
				pod.OwnerKind = ownerRef.Kind
				pod.OwnerName = ownerRef.Name
				break
			}
		}

		for _, volume := range p.Spec.Volumes {
			if volume.PersistentVolumeClaim != nil {
				pod.PVCNames = append(pod.PVCNames, volume.PersistentVolumeClaim.ClaimName)
			}

			if volume.Secret != nil {
				pod.SecretRefs = append(pod.SecretRefs, volume.Secret.SecretName)
			}

			if volume.ConfigMap != nil {
				pod.ConfigMapRefs = append(pod.ConfigMapRefs, volume.ConfigMap.Name)
			}
		}

		for _, container := range p.Spec.Containers {
			for _, envFrom := range container.EnvFrom {
				if envFrom.SecretRef != nil {
					pod.SecretRefs = append(pod.SecretRefs, envFrom.SecretRef.Name)
				}
				if envFrom.ConfigMapRef != nil {
					pod.ConfigMapRefs = append(pod.ConfigMapRefs, envFrom.ConfigMapRef.Name)
				}
			}

			for _, env := range container.Env {
				if env.ValueFrom != nil {
					if env.ValueFrom.ConfigMapKeyRef != nil {
						pod.ConfigMapRefs = append(pod.ConfigMapRefs, env.ValueFrom.ConfigMapKeyRef.Name)
					}
				}
				if env.ValueFrom.SecretKeyRef != nil {
					pod.SecretRefs = append(pod.SecretRefs, env.ValueFrom.SecretKeyRef.Name)
				}
			}

		}
		result = append(result, pod)
	}

	return result
}

func MapReplicaSetsToDto(replicaSets *app.ReplicaSetList) []ReplicaSet {

	return nil
}

func MapDeploymentsToDto(replicaSets *app.DeploymentList) []Deployment {

	return nil
}

func MapServicesToDto(replicaSets *core.ServiceList) []Service {

	return nil
}

func MapNodesToDto(replicaSets *core.NodeList) []Node {

	return nil
}

func MapPVCsToDto(replicaSets *core.PersistentVolumeClaimList) []PersistentVolumeClaim {

	return nil
}

func MapConfigMapsToDto(replicaSets *core.ConfigMapList) []ConfigMap {

	return nil
}

func MapSecretsToDto(replicaSets *core.SecretList) []Secret {

	return nil
}

func MapNamespacesToDto(replicaSets *core.NamespaceList) []Namespace {

	return nil
}
