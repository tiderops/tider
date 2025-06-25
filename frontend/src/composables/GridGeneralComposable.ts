import { onMounted, ref } from 'vue'
import { database, model } from '../../wailsjs/go/models'
import { fetchHeaderParams } from '../services/layout.service'
import HeadParamsDto = database.HeadParamsDto
import NodeDto = model.NodeDto
import NamespaceDto = model.NamespaceDto
import { fetchGetNamespace, fetchGetNodes } from '../services/general.service'

export function gridGeneralComposable(k8sObject: string) {
	const nodes = ref<NodeDto[]>([])
	const namespaces = ref<NamespaceDto[]>([])
	const headers = ref<HeadParamsDto[]>([])

	const fetchData = async (): Promise<any> => {
		try {
			nodes.value = await fetchGetNodes()
			namespaces.value = await fetchGetNamespace()

			headers.value = await fetchHeaderParams(k8sObject)
		} catch (error) {
			console.log('Error fetching pod data: ', error)
			throw error
		}
	}

	onMounted(async () => {
		console.log('onMounted triggered')
		await fetchData()
	})

	return {
		nodes,
		namespaces,
		headers,
		fetchData,
	}
}
