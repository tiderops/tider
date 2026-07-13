import { ref } from 'vue'
import type { ResourceDescriptor, GridHeader, GridRow } from '@/resources/types'
import { fetchGridHeaders } from '@/resources/headers'
import { toAppError } from '@/services/apperror'
import type { AppError } from '@/services/apperror'

export function useResourceGrid(descriptor: ResourceDescriptor | undefined, cluster: string) {
	const headers = ref<GridHeader[]>([])
	const rows = ref<GridRow[]>([])
	const loading = ref(false)
	const error = ref<AppError | null>(null)

	const fetchData = async () => {
		if (!descriptor) {
			return
		}

		loading.value = true
		error.value = null

		try {
			const [fetchedRows, fetchedHeaders] = await Promise.all([
				descriptor.fetchRows(cluster),
				fetchGridHeaders(descriptor.kind),
			])
			rows.value = fetchedRows
			headers.value = fetchedHeaders
		} catch (err) {
			error.value = toAppError(err)
			rows.value = []
		} finally {
			loading.value = false
		}
	}

	return { headers, rows, loading, error, fetchData }
}
