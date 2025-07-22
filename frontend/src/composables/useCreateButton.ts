import { fetchDeleteDeployment, fetchRestartPod, fetchUpdateDeployment, fetchUpdatePod } from '../services/workload.service'
import { ref } from 'vue'
import { fetchDeleteIngress, fetchDeleteService, fetchUpdateService } from '../services/network.service'

export function useCreateButton() {
	const loading = ref(false)
	const error = ref<unknown>(null)
	const success = ref(false)

	async function create(name: string, namespace: string, clusterCtx: string, k8sObject: string) {
		loading.value = true
		error.value = null
		success.value = false

		try {
			switch (k8sObject) {
				case 'pod': {
					// await fetchUpdatePod(name, namespace, clusterCtx)
					break
				}
				case 'deployment': {
					// await fetchUpdateDeployment(name, namespace, clusterCtx)
					break
				}
				case 'service': {
					// await fetchUpdateService(name, namespace, clusterCtx)
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

	return { create, loading, error, success }
}
