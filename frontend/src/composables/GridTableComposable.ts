import { fetchRestartPod } from '../services/workload.service'

export function gridGeneralComposable(podName: string, namespace: string) {
  return restartPod(podName, namespace)
}

export function restartPod(podName: string, namespace: string) {
  const fetchData = async () => {
    try {
      await fetchRestartPod(podName, namespace)
    } catch (error) {
      console.log('Error fetching pod data: ', error)
      throw error
    }
  }

  return {
    fetchData
  }
}
