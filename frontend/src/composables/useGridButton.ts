import { fetchDeleteDeployment, fetchRestartPod } from '../services/workload.service'
import { ref } from 'vue'
import { fetchDeleteIngress, fetchDeleteService } from '../services/network.service'

export function useGridButton() {
	const loading = ref(false)
	const error = ref<unknown>(null)
	const success = ref(false)

	async function restart(name: string, namespace: string, clusterCtx: string, k8sObject: string) {
		loading.value = true
		error.value = null
		success.value = false

		try {
			switch (k8sObject) {
				case 'pod': {
					await fetchRestartPod(name, namespace, clusterCtx)
					break
				}
				case 'deployment': {
					await fetchDeleteDeployment(name, namespace, clusterCtx)
					break
				}
				case 'service': {
					await fetchDeleteService(name, namespace, clusterCtx)
					break
				}
				case 'ingress': {
					await fetchDeleteIngress(name, namespace, clusterCtx)
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

	return { restart, loading, error, success }
}
