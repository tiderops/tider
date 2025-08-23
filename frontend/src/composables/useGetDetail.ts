import { fetchGetDeployment, fetchGetPod } from '../services/workload.service'
import { ref } from 'vue'
import { fetchGetService } from '../services/network.service'
import { IDeploymentDetail, IPodDetail } from '../types/workload.type'
import { IServiceDetail } from '../types/network.type'
import { model } from '../../wailsjs/go/models'
import PodDto = model.PodDto
import DeploymentDto = model.DeploymentDto
import ServiceDto = model.ServiceDto

export function useGetDetail(name: string, namespace: string, clusterCtx: string, k8sObject: string) {
	const k8sComponents = new Map<string, Function>()
	k8sComponents.set('pod', getPod(name, namespace, clusterCtx))
	k8sComponents.set('deployment', getDeployment(name, namespace, clusterCtx))
	k8sComponents.set('service', getService(name, namespace, clusterCtx))

	// return getPod(name, namespace, clusterCtx)
	console.log('OBJECT:: ', k8sObject)
	return k8sComponents.get(k8sObject)
}

export async function getPod(name: string, namespace: string, clusterCtx: string) {
	const body = ref<IPodDetail>([])

	try {
		const pods = ref<PodDto>([])
		pods.value = await fetchGetPod(name, namespace, clusterCtx)
		console.log('PODS: ', pods.value)
		const p = pods.value

		body.value = {
			name: p.Name,
			namespace: p.Namespace,
			image: p.Container.Image,
			pullPolicy: p.Container.PullPolicy,
			port: p.Container.Port,
			cpu: `${p.Container.Limit.Cpu}/${p.Container.Request.Cpu}`,
			memory: `${p.Container.Limit.Memory}/${p.Container.Request.Memory}`,
			storage: `${p.Container.Limit.Storage}/${p.Container.Request.Storage}`,
			age: p.Age,
			status: p.Status,
			editable: p.Editable,
		}
	} catch (error) {
		console.log('Error fetching pod data: ', error)
		throw error
	}

	return body
}

export async function getDeployment(name: string, namespace: string, clusterCtx: string) {
	const body = ref<IDeploymentDetail>()

	try {
		const deployment = ref<DeploymentDto>([])
		deployment.value = await fetchGetDeployment(name, namespace, clusterCtx)
		const d = deployment.value

		body.value = {
			name: d.Name,
			namespace: d.Namespace,
			replicas: d.Replicas,
			age: d.Age,
			status: d.Status,
		}
	} catch (error) {
		console.log('Error fetching pod data: ', error)
		throw error
	}

	return body
}

export async function getService(name: string, namespace: string, clusterCtx: string) {
	const body = ref<IServiceDetail[]>()

	try {
		const service = ref<ServiceDto>([])
		service.value = await fetchGetService(name, namespace, clusterCtx)
		const s = service.value

		body.value = {
			name: s.Name,
			namespace: s.Namespace,
			age: s.Age,
			status: s.Status,
		}
	} catch (error) {
		console.log('Error fetching pod data: ', error)
		throw error
	}

	return body
}
