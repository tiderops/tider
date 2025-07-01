import { ref } from 'vue'
import { model } from '../../wailsjs/go/models'
import { fetchGetNamespaces } from '../services/general.service'
import NamespaceDto = model.NamespaceDto

interface GridResponse {
	fetchData: () => Promise<any>
	content?: {
		body: any
	}
}

export function useFilter(clusterCtx: string, k8sObject: string) {
	const body = ref<Array<any>>([])

	const fetchData = async () => {
		try {
			const namespaces = ref<NamespaceDto[]>([])

			namespaces.value = await fetchGetNamespaces(clusterCtx)

			console.log('NAMESPACES: ', namespaces.value)

			body.value = namespaces.value.map((n: NamespaceDto) => n.Name)

			console.log('body.value: ', body.value)
		} catch (error) {
			console.log('Error fetching namespace data: ', error)
			throw error
		}
	}

	const response: GridResponse = {
		fetchData: fetchData,
		content: {
			body: body,
		},
	}

	return response
}
