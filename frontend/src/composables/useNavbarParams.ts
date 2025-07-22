import { database } from '../../wailsjs/go/models'
import ObjectType = database.ObjectType
import { onMounted, ref } from 'vue'
import { fetchObjectsParameter } from '../services/navbar.service'

export function useNavbarParams() {
	const objects = ref<ObjectType[]>([])

	const fetchData = async (): Promise<ObjectType[]> => {
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
