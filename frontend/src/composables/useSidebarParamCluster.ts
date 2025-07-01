import { ref, onMounted } from 'vue'
import { fetchCommonParameters, fetchKubernetesParameters, fetchClusters } from '../services/layout.service'
import { database, model } from '../../wailsjs/go/models'
import CommonParameterDto = database.CommonParameterDto
import ClusterInfo = model.ClusterInfo

export function useSidebarParamCluster() {
	const commonParameters = ref<CommonParameterDto[]>([])
	const kubernetesParameters = ref<CommonParameterDto[]>([])
	const clusters = ref<ClusterInfo[]>([])

	const fetchData = async () => {
		try {
			commonParameters.value = await fetchCommonParameters()
			kubernetesParameters.value = await fetchKubernetesParameters()
			clusters.value = await fetchClusters()
			console.log('clusters.value: ', clusters.value)
		} catch (error) {
			console.error('Error fetching environment data:', error)
			throw error
		}
	}

	onMounted(async () => {
		console.log('onMounted triggered')
		await fetchData()
	})

	return {
		commonParameters,
		kubernetesParameters,
		clusters,
		fetchData,
	}
}
