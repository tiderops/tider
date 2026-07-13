import { ref } from 'vue'
import { fetchCommonParameters, fetchKubernetesParameters } from '../services/layout.service'
import { config } from '../../wailsjs/go/models'
import CommonParameterDto = config.CommonParameterDto

export function useSidebarParamCluster() {
	const commonParameters = ref<CommonParameterDto[]>([])
	const kubernetesParameters = ref<CommonParameterDto[]>([])

	const fetchData = async () => {
		commonParameters.value = await fetchCommonParameters()
		kubernetesParameters.value = await fetchKubernetesParameters()
	}

	return {
		commonParameters,
		kubernetesParameters,
		fetchData,
	}
}
