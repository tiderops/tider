import { fetchRestartPod } from '../services/workload.service'

export function gridGeneralComposable(podName: string, namespace: string, clusterCtx: string) {
  return restartPod(podName, namespace, clusterCtx)
}

export function restartPod(podName: string, namespace: string, clusterCtx: string) {
  const fetchData = async () => {
    try {
      await fetchRestartPod(podName, namespace, clusterCtx)
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }

  return {
    fetchData,
  }
}
