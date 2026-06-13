import { defineStore } from 'pinia'
import { ref } from 'vue'
import { fetchClusters } from '@/services/layout.service'
import { toAppError } from '@/services/apperror'
import type { AppError } from '@/services/apperror'
import { model } from '../../wailsjs/go/models'

export const useClusterStore = defineStore('cluster', () => {
	const clusters = ref<model.ClusterInfo[]>([])
	const currentCluster = ref('')
	const loading = ref(false)
	const error = ref<AppError | null>(null)
	const loaded = ref(false)

	async function loadClusters(force = false) {
		if (loaded.value && !force) return

		loading.value = true
		error.value = null

		try {
			clusters.value = (await fetchClusters()) ?? []
			loaded.value = true
		} catch (err) {
			error.value = toAppError(err)
			throw error.value
		} finally {
			loading.value = false
		}
	}

	function setCurrentCluster(name: string) {
		currentCluster.value = name
	}

	return { clusters, currentCluster, loading, error, loaded, loadClusters, setCurrentCluster }
})
