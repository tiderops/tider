export declare const fetchGetServices: (
  clusterCtx: string,
) => Promise<import('../../wailsjs/go/models').model.ServiceDto[]>
export declare const fetchGetService: (
  name: string,
  namespace: string,
  clusterCtx: string,
) => Promise<import('../../wailsjs/go/models').model.ServiceDto>
export declare const fetchUpdateService: (
  name: string,
  namespace: string,
  dto: any,
  clusterCtx: string,
) => Promise<void>
export declare const fetchDeleteService: (
  name: string,
  namespace: string,
  clusterCtx: string,
) => Promise<void>
export declare const fetchGetIngresses: (
  clusterCtx: string,
) => Promise<import('../../wailsjs/go/models').model.IngressDto[]>
export declare const fetchGetIngress: (
  name: string,
  namespace: string,
  clusterCtx: string,
) => Promise<import('../../wailsjs/go/models').model.IngressDto>
export declare const fetchUpdateIngress: (
  name: string,
  namespace: string,
  dto: any,
  clusterCtx: string,
) => Promise<void>
export declare const fetchDeleteIngress: (
  name: string,
  namespace: string,
  clusterCtx: string,
) => Promise<void>
