import { fetchGetDeployment, fetchGetPod } from '../services/workload.service'
import { ref } from 'vue'
import { fetchGetService } from '../services/network.service'
import { IDeploymentDetail, IPod, IPodDetail } from '../types/workload.type'
import { IService, IServiceDetail } from '../types/network.type'
import { model } from '../../wailsjs/go/models'
import PodDto = model.PodDto

export function useGetDetail(name: string, namespace: string, clusterCtx: string, k8sObject: string) {
	const k8sComponents = new Map<string, Function>()
	// k8sComponents.set('pod', getPod(name, namespace, clusterCtx))
	// k8sComponents.set('deployment', getDeployment(name, namespace, clusterCtx))
	// k8sComponents.set('service', getService(name, namespace, clusterCtx))

	// const res = k8sComponents.get(k8sObject)

	// if (res == undefined) {
	// 	return {
	// 		fetchData: null,
	// 		content: {
	// 			body: null,
	// 		},
	// 	}
	// } else {
	// 	return res
	// }

	return getPod(name, namespace, clusterCtx)
}

export async function getPod(name: string, namespace: string, clusterCtx: string) {
	const body = ref<IPodDetail>([])

	// const fetchData = async () => {
	try {
		const pods = ref<PodDto>([])
		pods.value = await fetchGetPod(name, namespace, clusterCtx)
		console.log('PODS: ', pods.value)

		// body.value = pods.value((p: PodDto) => ({
		// 	name: p.Name,
		// 	namespace: p.Namespace,
		// 	image: p.Container.Image,
		// 	pullPolicy: p.Container.PullPolicy,
		// 	port: p.Container.Port,
		// 	cpu: `${p.Container.Limit.Cpu}/${p.Container.Request.Cpu}`,
		// 	memory: `${p.Container.Limit.Memory}/${p.Container.Request.Memory}`,
		// 	storage: `${p.Container.Limit.Storage}/${p.Container.Request.Storage}`,
		// 	age: p.Age,
		// 	status: p.Status,
		// 	editable: p.Editable
		// }))
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
	// }

	// const response: Response = {
	// 	fetchData: null,
	// 	content: {
	// 		body: body,
	// 	},
	// }

	return body
}

export function getDeployment(name: string, namespace: string, clusterCtx: string) {
	const body = ref<IDeploymentDetail>()

	const fetchData = async () => {
		try {
			// const deployment = ref<>()
		} catch (error) {
			console.log('Error fetching pod data: ', error)
			throw error
		}
	}

	return body
}

export function getService(name: string, namespace: string, clusterCtx: string) {
	const body = ref<IServiceDetail>()

	const fetchData = async () => {
		try {
		} catch (error) {
			console.log('Error fetching pod data: ', error)
			throw error
		}
	}

	return body
}
