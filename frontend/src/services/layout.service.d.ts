export declare const fetchCommonParameters: () => Promise<
  import('../../wailsjs/go/models').database.CommonParameterDto[]
>
export declare const fetchKubernetesParameters: () => Promise<
  import('../../wailsjs/go/models').database.CommonParameterDto[]
>
export declare const fetchClusters: () => Promise<
  import('../../wailsjs/go/models').model.ClusterInfo[]
>
export declare const fetchHeaderParams: (
  k8sObject: string,
) => Promise<import('../../wailsjs/go/models').database.HeadParamsDto[]>
