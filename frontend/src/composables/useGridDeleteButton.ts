import { fetchDeleteDeployment, fetchRestartPod } from '../services/workload.service'
import { ref } from 'vue'
import { fetchDeleteIngress, fetchDeleteService } from '../services/network.service'

export function useGridDeleteButton() {
	const loading = ref(false)
	const error = ref<unknown>(null)
	const success = ref(false)

	async function restart(name: string, namespace: string, clusterCtx: string, k8sObject: string) {
		loading.value = true
		error.value = null
		success.value = false

		const k8sComponents = new Map<string, Function>();
		k8sComponents.set('pod', fetchRestartPod(name, namespace, clusterCtx))
		k8sComponents.set('deployment', fetchDeleteDeployment(name, namespace, clusterCtx))
		k8sComponents.set('service', fetchDeleteService(name, namespace, clusterCtx))
		k8sComponents.set('ingress', fetchDeleteIngress(name, namespace, clusterCtx))

		try {
			await k8sComponents.get(k8sObject)
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
