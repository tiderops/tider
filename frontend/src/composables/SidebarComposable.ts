import { ref, onMounted } from 'vue'
import {
  fetchCommonParameters,
  fetchKubernetesParameters,
  fetchClusters,
} from '../services/layout.service'
import { database, model } from '../../wailsjs/go/models'
import EnvironmentDto = model.EnvironmentDto
import CommonParameterDto = database.CommonParameterDto
import ClusterInfo = model.ClusterInfo

export function sidebarComposable() {
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

  // TODO: No es necesario agregar async, porque luego se tiene un async
  // onMounted es parte del lifecycle de vuejs
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
