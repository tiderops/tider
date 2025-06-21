import { onMounted, ref } from 'vue'
import { database, model } from '../../wailsjs/go/models'
import { fetchGetDeployments, fetchGetPods } from '../services/workload.service'
import { fetchHeaderParams } from '../services/layout.service'
import HeadParamsDto = database.HeadParamsDto
import { fetchGetNamespace, fetchGetNodes } from '../services/general.service'
import NamespaceDto = model.NamespaceDto
import { fetchGetIngresses, fetchGetServices } from '../services/network.service'
import ServiceDto = model.ServiceDto
import IngressDto = model.IngressDto
import { fetchGetPersistentVolumes } from '../services/storage.service'
import PersistentVolumeDto = model.PersistentVolumeDto
import NodeDto = model.NodeDto
import PodDto = model.PodDto
import DeploymentDto = model.DeploymentDto

interface GridResponse {
  fetchData: () => Promise<any>
  content?: {
    head: any
    body: any
  }
}

export function gridComposable(clusterCtx: string, k8sObject: string): GridResponse {
  switch (k8sObject) {
    case 'node': {
      console.log('Get Nodes')
      return gridBodyNodes(k8sObject)
    }
    case 'namespace': {
      console.log('Get Namespaces')
      return gridBodyNamespaces(k8sObject)
    }
    case 'service': {
      console.log('Get Services')
      return gridBodyServices(k8sObject)
    }
    case 'ingress': {
      console.log('Get Ingresses')
      return gridBodyIngresses(k8sObject)
    }
    case 'persistentVolumen': {
      console.log('Get Persistent Volumen')
      return gridBodyPersistentVolumes(k8sObject)
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
  const body = ref<
    Array<{
      name: string
      namespace: string
      replicas: number
      cpu: string
      memory: string
      age: string
      status: string
    }>
  >([])

  const fetchData = async () => {
    try {
      const pods = ref<PodDto[]>([])
      const header = ref<HeadParamsDto[]>([])

      pods.value = await fetchGetPods(clusterCtx)
      header.value = await fetchHeaderParams(k8sObject)

      console.log('PODS: ', pods.value)

      body.value = pods.value.map((p: any) => ({
        name: p.Name,
        namespace: p.Namespace,
        replicas: p.Replicas,
        cpu: `${p.Container.Limit.Cpu}/${p.Container.Request.Cpu}`,
        memory: `${p.Container.Limit.Memory}/${p.Container.Request.Memory}`,
        age: p.Age,
        status: p.Status,
      }))

      head.value = header.value.map((h: any) => ({
        title: h.Title,
        key: h.Key,
        align: h.Align,
        sortable: h.Sortable,
      }))
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }

  onMounted(async () => {
    console.log('onMounted triggered')
    // await fetchData()
  })

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
  const body = ref<
    Array<{
      name: string
      namespace: string
      status: string
      age: string
    }>
  >([])

  const fetchData = async (): Promise<any> => {
    try {
      const deployment = ref<DeploymentDto[]>([])
      const header = ref<HeadParamsDto[]>([])

      deployment.value = await fetchGetDeployments(clusterCtx)
      header.value = await fetchHeaderParams(k8sObject)

      console.log('DEPLOYMENTS', body.value)

      body.value = deployment.value.map((d: any) => ({
        name: d.Name,
        namespace: d.Namespace,
        status: d.Status,
        age: d.Age,
      }))

      head.value = header.value.map((h: any) => ({
        title: h.Title,
        key: h.Key,
        align: h.Align,
        sortable: h.Sortable,
      }))
    } catch (error) {
      console.log('Error fetching deployment data: ', error)
      throw error
    }
  }

  onMounted(async () => {
    console.log('onMounted triggered')
    await fetchData()
  })

  const response: GridResponse = {
    fetchData: fetchData,
    content: {
      head: head,
      body: body,
    },
  }

  return response
}

export function gridBodyServices(k8sObject: string) {
  const body = ref<ServiceDto[]>([])
  const header = ref<HeadParamsDto[]>([])

  const fetchData = async () => {
    try {
      body.value = await fetchGetServices('east') // TODO: replace with a dynamic value
      header.value = await fetchHeaderParams(k8sObject)
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }

  onMounted(async () => {
    console.log('onMounted triggered')
    await fetchData()
  })

  const response: GridResponse = {
    fetchData: fetchData,
    content: {
      head: header,
      body: body,
    },
  }

  return response
}

export function gridBodyIngresses(k8sObject: string) {
  const body = ref<IngressDto[]>([])
  const header = ref<HeadParamsDto[]>([])

  const fetchData = async () => {
    try {
      body.value = await fetchGetIngresses('east') // TODO: replace with a dynamic value
      header.value = await fetchHeaderParams(k8sObject)
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }

  onMounted(async () => {
    console.log('onMounted triggered')
    await fetchData()
  })

  const response: GridResponse = {
    fetchData: fetchData,
    content: {
      head: header,
      body: body,
    },
  }

  return response
}

export function gridBodyPersistentVolumes(k8sObject: string) {
  const body = ref<PersistentVolumeDto[]>([])
  const header = ref<HeadParamsDto[]>([])

  const fetchData = async () => {
    try {
      body.value = await fetchGetPersistentVolumes()
      header.value = await fetchHeaderParams(k8sObject)
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }

  onMounted(async () => {
    console.log('onMounted triggered')
    await fetchData()
  })

  const response: GridResponse = {
    fetchData: fetchData,
    content: {
      head: header,
      body: body,
    },
  }

  return response
}

export function gridBodyNamespaces(k8sObject: string) {
  const head = ref<Array<any>>([])
  const body = ref<
    Array<{
      name: string
      label: string
      age: string
      status: string
    }>
  >([])

  const fetchData = async () => {
    try {
      const namespaces = ref<NamespaceDto[]>([])
      const header = ref<HeadParamsDto[]>([])
      namespaces.value = await fetchGetNamespace()
      header.value = await fetchHeaderParams(k8sObject)

      console.log('NAMESPACES: ', namespaces.value)

      body.value = namespaces.value.map((p: any) => ({
        name: p.Name,
        label: '',
        age: p.Age,
        status: p.Status,
      }))

      head.value = header.value.map((h: any) => ({
        title: h.Title,
        key: h.Key,
        align: h.Align,
        sortable: h.Sortable,
      }))
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }

  onMounted(async () => {
    console.log('onMounted triggered')
    // await fetchData()
  })

  const response: GridResponse = {
    fetchData: fetchData,
    content: {
      head: head,
      body: body,
    },
  }

  return response
}

export function gridBodyNodes(namespace: string) {
  const body = ref<NodeDto[]>([])
  const header = ref<HeadParamsDto[]>([])

  const fetchData = async () => {
    try {
      body.value = await fetchGetNodes()
      header.value = await fetchHeaderParams(namespace)
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }

  onMounted(async () => {
    console.log('onMounted triggered')
    await fetchData()
  })

  const response: GridResponse = {
    fetchData: fetchData,
    content: {
      head: header,
      body: body,
    },
  }

  return response
}
