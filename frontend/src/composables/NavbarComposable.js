import { database } from '../../wailsjs/go/models'
var ObjectType = database.ObjectType
import { onMounted, ref } from 'vue'
import { fetchObjectsParameter } from '../services/navbar.service'
export function navbarComposable() {
  const objects = ref([])
  const fetchData = async () => {
    try {
      objects.value = await fetchObjectsParameter()
      return objects.value
    } catch (error) {
      console.error('Error fetching data', error)
      throw error
    }
  }
  onMounted(async () => {
    console.log('onMounted Trigger Navbar')
    await fetchData()
  })
  return {
    objects,
    fetchData,
  }
}
