import {
  DeleteIngress,
  DeleteService,
  GetIngress,
  GetIngresses,
  GetService,
  GetServices,
  UpdateIngress,
  UpdateService,
} from '../../wailsjs/go/middleware/NetworkMiddleware'

export const fetchGetServices = async (clusterCtx: string) => GetServices(clusterCtx)
export const fetchGetService = async (name: string, namespace: string, clusterCtx: string) =>
  GetService(name, namespace, clusterCtx)
export const fetchUpdateService = async (name: string, namespace: string, dto: any, clusterCtx: string) =>
  UpdateService(name, namespace, dto, clusterCtx)
export const fetchDeleteService = async (name: string, namespace: string, clusterCtx: string) =>
  DeleteService(name, namespace, clusterCtx)

export const fetchGetIngresses = async (clusterCtx: string) => GetIngresses(clusterCtx)
export const fetchGetIngress = async (name: string, namespace: string, clusterCtx: string) =>
  GetIngress(name, namespace, clusterCtx)
export const fetchUpdateIngress = async (name: string, namespace: string, dto: any, clusterCtx: string) =>
  UpdateIngress(name, namespace, dto, clusterCtx)
export const fetchDeleteIngress = async (name: string, namespace: string, clusterCtx: string) =>
  DeleteIngress(name, namespace, clusterCtx)
