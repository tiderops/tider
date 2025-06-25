import { ref, onMounted } from 'vue'
import { database } from '../../wailsjs/go/models'
import CommonParameterDto = database.CommonParameterDto
import { fetchCommonParameters } from '../services/layout.service'

export function useLayoutComposableExample() {
	const result = ref<CommonParameterDto[]>([])

	const fetchData = async (): Promise<CommonParameterDto[]> => {
		try {
			result.value = await fetchCommonParameters()

			return result.value
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
		result,
		fetchData,
	}
}
