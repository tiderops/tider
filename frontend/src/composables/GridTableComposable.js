import { fetchRestartPod } from '../services/workload.service'
export function gridGeneralComposable(podName, namespace, clusterCtx) {
  return restartPod(podName, namespace, clusterCtx)
}
export function restartPod(podName, namespace, clusterCtx) {
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
