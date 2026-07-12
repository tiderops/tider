import { ref } from 'vue'
import type { Ref } from 'vue'
import { toAppError } from '@/services/apperror'
import type { AppError } from '@/services/apperror'

export function useAsyncData<T>(fetcher: () => Promise<T>, initial: T) {
	const data = ref(initial) as Ref<T>
	const loading = ref(false)
	const error = ref<AppError | null>(null)

	async function reload() {
		loading.value = true
		error.value = null
		try {
			data.value = await fetcher()
		} catch (err) {
			error.value = toAppError(err)
		} finally {
			loading.value = false
		}
	}

	return { data, loading, error, reload }
}
