import {
	fetchUpdateDeployment,
	fetchUpdatePod
} from '../services/workload.service'
import { ref } from 'vue'
import {fetchUpdateService} from '../services/network.service'

export function useGridUpdateButton() {
	const loading = ref(false)
	const error = ref<unknown>(null)
	const success = ref(false)

	async function update(name: string, namespace: string, dto: any, clusterCtx: string, k8sObject: string) {
		loading.value = true
		error.value = null
		success.value = false

		try {
			switch (k8sObject) {
				case 'pod': {
					await fetchUpdatePod(name, namespace, dto, clusterCtx)
					break
				}
				case 'deployment': {
					await fetchUpdateDeployment(name, namespace, dto, clusterCtx)
					break
				}
				case 'service': {
					await fetchUpdateService(name, namespace, dto, clusterCtx)
					break
				}
				default: {
					error.value = err
					console.log('K8s object not found')
					break
				}
			}

			success.value = true
		} catch (err) {
			error.value = err
			throw err
		} finally {
			loading.value = false
		}
	}

	return { update, loading, error, success }
}
