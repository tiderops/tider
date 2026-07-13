import {
	AutoTroubleshoot,
	DeleteDeployment,
	GetDeployment,
	GetDeployments,
	GetPod,
	GetPods,
	ResourceTuning,
	RestartPod,
	UpdateDeployment,
	UpdatePod,
} from '../../wailsjs/go/binding/Workload'
import { model } from '../../wailsjs/go/models'

export const fetchGetPods = async (clusterCtx: string) => GetPods(clusterCtx)
export const fetchGetPod = async (name: string, namespace: string, clusterCtx: string) => GetPod(name, namespace, clusterCtx)
export const fetchUpdatePod = async (name: string, namespace: string, dto: model.PodUpdate, clusterCtx: string) => UpdatePod(name, namespace, dto, clusterCtx)
export const fetchRestartPod = async (name: string, namespace: string, clusterCtx: string) => RestartPod(name, namespace, clusterCtx)

export const fetchGetDeployments = async (clusterCtx: string) => GetDeployments(clusterCtx)
export const fetchGetDeployment = async (name: string, namespace: string, clusterCtx: string) => GetDeployment(name, namespace, clusterCtx)
export const fetchUpdateDeployment = async (name: string, namespace: string, dto: model.DeploymentUpdate, clusterCtx: string) => UpdateDeployment(name, namespace, dto, clusterCtx)
export const fetchDeleteDeployment = async (name: string, namespace: string, clusterCtx: string) => DeleteDeployment(name, namespace, clusterCtx)

export const fetchResourceTuning = async (namespace: string, clusterCtx: string) => ResourceTuning(namespace, clusterCtx)
export const fetchAutoTroubleshoot = async (name: string, namespace: string, clusterCtx: string, resource: string) => AutoTroubleshoot(name, namespace, clusterCtx, resource)
