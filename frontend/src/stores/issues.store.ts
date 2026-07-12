import { defineStore } from 'pinia'
import { ref } from 'vue'
import { fetchIssues } from '@/data/issues.data'
import { toAppError } from '@/services/apperror'
import type { AppError } from '@/services/apperror'
import type { Issue } from '@/types/issue'

export const useIssuesStore = defineStore('issues', () => {
	const items = ref<Issue[]>([])
	const loading = ref(false)
	const error = ref<AppError | null>(null)
	const loaded = ref(false)

	async function load(force = false) {
		if (loaded.value && !force) {
			return
		}
		loading.value = true
		error.value = null
		try {
			items.value = await fetchIssues()
			loaded.value = true
		} catch (err) {
			error.value = toAppError(err)
		} finally {
			loading.value = false
		}
	}

	return { items, loading, error, loaded, load }
})
