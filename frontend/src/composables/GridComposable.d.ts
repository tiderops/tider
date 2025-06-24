interface GridResponse {
  fetchData: () => Promise<any>
  content?: {
    head: any
    body: any
  }
}
export declare function gridComposable(clusterCtx: string, k8sObject: string): GridResponse
export declare function gridBodyPods(clusterCtx: string, k8sObject: string): GridResponse
export declare function gridBodyDeployments(clusterCtx: string, k8sObject: string): GridResponse
export declare function gridBodyServices(k8sObject: string): GridResponse
export declare function gridBodyIngresses(k8sObject: string): GridResponse
export declare function gridBodyPersistentVolumes(k8sObject: string): GridResponse
export declare function gridBodyNamespaces(k8sObject: string): GridResponse
export declare function gridBodyNodes(namespace: string): GridResponse
export {}
