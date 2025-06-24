import { onMounted, ref } from 'vue'
import { database, model } from '../../wailsjs/go/models'
import { fetchHeaderParams } from '../services/layout.service'
var HeadParamsDto = database.HeadParamsDto
var NodeDto = model.NodeDto
var NamespaceDto = model.NamespaceDto
import { fetchGetNamespace, fetchGetNodes } from '../services/general.service'
export function gridGeneralComposable(k8sObject) {
  const nodes = ref([])
  const namespaces = ref([])
  const headers = ref([])
  const fetchData = async () => {
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
