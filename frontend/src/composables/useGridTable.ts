import { ref } from 'vue'
import { database, model } from '../../wailsjs/go/models'
import { fetchGetDeployments, fetchGetPods } from '../services/workload.service'
import { fetchHeaderParams } from '../services/layout.service'
import HeadParamsDto = database.HeadParamsDto
import { fetchGetNamespaces, fetchGetNodes } from '../services/general.service'
import NamespaceDto = model.NamespaceDto
import { fetchGetIngresses, fetchGetServices } from '../services/network.service'
import ServiceDto = model.ServiceDto
import IngressDto = model.IngressDto
import { fetchGetPersistentVolumes, fetchGetPersistentVolumesClaim } from '../services/storage.service'
import PersistentVolumeDto = model.PersistentVolumeDto
import PodDto = model.PodDto
import DeploymentDto = model.DeploymentDto
import { IDeployment, IPod } from '../types/workload.type'
import { IIngress, IService } from '../types/network.type'
import { IPersistentVolume, IPersistentVolumeClaim } from '../types/storage.type'
import PersistentVolumeClaimDto = model.PersistentVolumeClaimDto
import { INamespace, INode } from '../types/general.type'
import NodeDtoV2 = model.NodeDtoV2

interface GridResponse {
	fetchData: () => Promise<any>
	content?: {
		head: any
		body: any
	}
}

export function useGridTable(clusterCtx: string, k8sObject: string): GridResponse {
	switch (k8sObject) {
		case 'node': {
			console.log('Get Nodes')
			return gridBodyNodes(clusterCtx, k8sObject)
		}
		case 'namespace': {
			console.log('Get Namespaces')
			return gridBodyNamespaces(clusterCtx, k8sObject)
		}
		case 'service': {
			console.log('Get Services')
			return gridBodyServices(clusterCtx, k8sObject)
		}
		case 'ingress': {
			console.log('Get Ingresses')
			return gridBodyIngresses(clusterCtx, k8sObject)
		}
		case 'persistentVolume': {
			console.log('Get Persistent Volume')
			return gridBodyPersistentVolumes(clusterCtx, k8sObject)
		}
		case 'persistentVolumeClaim': {
			console.log('Get Persistent Volume Claim')
			return gridBodyPersistentVolumesClaim(clusterCtx, k8sObject)
		}
		case 'deployment': {
			console.log('Get Deployments')
			return gridBodyDeployments(clusterCtx, k8sObject)
		}
		case 'pod': {
			console.log('Get Pods')
			return gridBodyPods(clusterCtx, k8sObject)
		}
		default: {
			console.log('K8s object not found')
			return {
				fetchData: async () => {},
				content: {
					head: null,
					body: null,
				},
			}
		}
	}
}

export function gridBodyPods(clusterCtx: string, k8sObject: string) {
	const head = ref<Array<any>>([])
	const body = ref<IPod[]>([])

	const fetchData = async () => {
		try {
			const pods = ref<PodDto[]>([])
			pods.value = await fetchGetPods(clusterCtx)
			console.log('PODS: ', pods.value)

			body.value = pods.value.map((p: PodDto) => ({
				name: p.Name,
				namespace: p.Namespace,
				replicas: p.Replicas,
				cpu: `${p.Container.Limit.Cpu}/${p.Container.Request.Cpu}`,
				memory: `${p.Container.Limit.Memory}/${p.Container.Request.Memory}`,
				storage: `${p.Container.Limit.Storage}/${p.Container.Request.Storage}`,
				age: p.Age,
				status: p.Status,
			}))

			head.value = await gridHead(k8sObject)
		} catch (error) {
			console.log('Error fetching pod data: ', error)
			throw error
		}
	}

	const response: GridResponse = {
		fetchData: fetchData,
		content: {
			head: head,
			body: body,
		},
	}

	return response
}

export function gridBodyDeployments(clusterCtx: string, k8sObject: string) {
	const head = ref<Array<any>>([])
	const body = ref<IDeployment[]>([])

	const fetchData = async () => {
		try {
			const deployment = ref<DeploymentDto[]>([])
			deployment.value = await fetchGetDeployments(clusterCtx)
			console.log('DEPLOYMENTS', deployment.value)

			body.value = deployment.value.map((d: DeploymentDto) => ({
				name: d.Name,
				namespace: d.Namespace,
				status: d.Status,
				age: d.Age,
			}))

			head.value = await gridHead(k8sObject)
		} catch (error) {
			console.log('Error fetching deployment data: ', error)
			throw error
		}
	}

	const response: GridResponse = {
		fetchData: fetchData,
		content: {
			head: head,
			body: body,
		},
	}

	return response
}

export function gridBodyServices(clusterCtx: string, k8sObject: string) {
	const head = ref<Array<any>>([])
	const body = ref<IService[]>([])

	const fetchData = async () => {
		try {
			const service = ref<ServiceDto[]>([])
			service.value = await fetchGetServices(clusterCtx)
			console.log('SERVICES', service.value)

			body.value = service.value.map((s: ServiceDto) => ({
				name: s.Name,
				namespace: s.Namespace,
				labels: s.Labels,
				status: s.Status,
				age: s.CreationTimestamp,
				spec: s.Spec,
			}))

			head.value = await gridHead(k8sObject)
		} catch (error) {
			console.log('Error fetching service data: ', error)
			throw error
		}
	}

	const response: GridResponse = {
		fetchData: fetchData,
		content: {
			head: head,
			body: body,
		},
	}

	return response
}

export function gridBodyIngresses(clusterCtx: string, k8sObject: string) {
	const head = ref<HeadParamsDto[]>([])
	const body = ref<IIngress[]>([])

	const fetchData = async () => {
		try {
			const ingress = ref<IngressDto[]>([])
			ingress.value = await fetchGetIngresses(clusterCtx)
			console.log('INGRESSES: ', ingress.value)

			body.value = ingress.value.map((i: IngressDto) => ({
				name: i.Name,
				namespace: i.Namespace,
				rules: i.Rules,
				host: i.Rules.map((r) => r.Host),
				path: i.Rules.map((r) => r.Path),
				age: i.Creation,
				label: i.Labels,
			}))

			head.value = await gridHead(k8sObject)
		} catch (error) {
			console.log('Error fetching ingress data: ', error)
			throw error
		}
	}

	const response: GridResponse = {
		fetchData: fetchData,
		content: {
			head: head,
			body: body,
		},
	}

	return response
}

export function gridBodyPersistentVolumes(clusterCtx: string, k8sObject: string) {
	const head = ref<any[]>([])
	const body = ref<IPersistentVolume[]>([])

	const fetchData = async () => {
		try {
			const persistentVolume = ref<PersistentVolumeDto[]>([])
			persistentVolume.value = await fetchGetPersistentVolumes(clusterCtx)
			console.log('PVS: ', persistentVolume.value)

			body.value = persistentVolume.value.map((p: PersistentVolumeDto) => ({
				name: p.Name,
				namespace: p.Namespace,
				age: p.CreationTimestamp,
				label: p.Labels,
				spec: p.VolumeSpec,
			}))

			head.value = await gridHead(k8sObject)
		} catch (error) {
			console.log('Error fetching pv data: ', error)
			throw error
		}
	}

	const response: GridResponse = {
		fetchData: fetchData,
		content: {
			head: head,
			body: body,
		},
	}

	return response
}

export function gridBodyPersistentVolumesClaim(clusterCtx: string, k8sObject: string) {
	const head = ref<any[]>([])
	const body = ref<IPersistentVolumeClaim[]>([])

	const fetchData = async () => {
		try {
			const persistentVolumeClaim = ref<PersistentVolumeClaimDto[]>([])
			persistentVolumeClaim.value = await fetchGetPersistentVolumesClaim(clusterCtx)
			console.log('PVCS: ', persistentVolumeClaim.value)

			body.value = persistentVolumeClaim.value.map((p: PersistentVolumeClaimDto) => ({
				name: p.Name,
				namespace: p.Namespace,
				age: p.CreationTimestamp,
				label: p.Labels,
				spec: p.VolumeClaimSpec,
			}))

			head.value = await gridHead(k8sObject)
		} catch (error) {
			console.log('Error fetching pvc data: ', error)
			throw error
		}
	}

	const response: GridResponse = {
		fetchData: fetchData,
		content: {
			head: head,
			body: body,
		},
	}

	return response
}

export function gridBodyNamespaces(clusterCtx: string, k8sObject: string) {
	const head = ref<Array<any>>([])
	const body = ref<INamespace[]>([])

	const fetchData = async () => {
		try {
			const namespaces = ref<NamespaceDto[]>([])
			namespaces.value = await fetchGetNamespaces(clusterCtx)
			console.log('NAMESPACES: ', namespaces.value)

			body.value = namespaces.value.map((n: NamespaceDto) => ({
				name: n.Name,
				version: n.Version,
				label: n.Labels,
				age: n.CreationTime,
				status: n.Status,
			}))

			head.value = await gridHead(k8sObject)
		} catch (error) {
			console.log('Error fetching namespace data: ', error)
			throw error
		}
	}

	const response: GridResponse = {
		fetchData: fetchData,
		content: {
			head: head,
			body: body,
		},
	}

	return response
}

export function gridBodyNodes(clusterCtx: string, k8sObject: string) {
	const head = ref<any[]>([])
	const body = ref<INode[]>([])

	const fetchData = async () => {
		try {
			const node = ref<NodeDtoV2[]>([])
			node.value = await fetchGetNodes(clusterCtx)
			console.log('NODE: ', node.value)

			body.value = node.value.map((n: NodeDtoV2) => ({
				name: n.Name,
				namespace: n.Namespace,
				memory: n.Resource.Memory,
				cpu: n.Resource.Cpu,
				storage: n.Resource.Storage,
				ephemeralStorage: n.Resource.StorageEphemeral,
				version: n.Version,
				age: n.CreationTimestamp,
				labels: n.Labels,
			}))

			head.value = await gridHead(k8sObject)
		} catch (error) {
			console.log('Error fetching node data: ', error)
			throw error
		}
	}

	const response: GridResponse = {
		fetchData: fetchData,
		content: {
			head: head,
			body: body,
		},
	}

	return response
}

async function gridHead(k8sObject: string): ref<any[]> {
	const head = ref<any[]>([])
	try {
		const header = ref<any[]>([])
		header.value = await fetchHeaderParams(k8sObject)

		head.value = header.value.map((h: any) => ({
			title: h.Title,
			key: h.Key,
			align: h.Align,
			sortable: h.Sortable,
		}))
	} catch (error) {
		console.log('Error fetching node data: ', error)
		throw error
	}
	return head
}
