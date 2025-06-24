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
export const fetchGetServices = async (clusterCtx) => GetServices(clusterCtx)
export const fetchGetService = async (name, namespace, clusterCtx) =>
  GetService(name, namespace, clusterCtx)
export const fetchUpdateService = async (name, namespace, dto, clusterCtx) =>
  UpdateService(name, namespace, dto, clusterCtx)
export const fetchDeleteService = async (name, namespace, clusterCtx) =>
  DeleteService(name, namespace, clusterCtx)
export const fetchGetIngresses = async (clusterCtx) => GetIngresses(clusterCtx)
export const fetchGetIngress = async (name, namespace, clusterCtx) =>
  GetIngress(name, namespace, clusterCtx)
export const fetchUpdateIngress = async (name, namespace, dto, clusterCtx) =>
  UpdateIngress(name, namespace, dto, clusterCtx)
export const fetchDeleteIngress = async (name, namespace, clusterCtx) =>
  DeleteIngress(name, namespace, clusterCtx)
