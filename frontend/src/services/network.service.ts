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

export const fetchGetServices = async (namespace: string) => GetServices(namespace)
export const fetchGetService = async (name: string, namespace: string) =>
  GetService(name, namespace)
export const fetchUpdateService = async (name: string, namespace: string, dto: any) =>
  UpdateService(name, namespace, dto)
export const fetchDeleteService = async (name: string, namespace: string) =>
  DeleteService(name, namespace)

export const fetchGetIngresses = async (namespace: string) => GetIngresses(namespace)
export const fetchGetIngress = async (name: string, namespace: string) =>
  GetIngress(name, namespace)
export const fetchUpdateIngress = async (name: string, namespace: string, dto: any) =>
  UpdateIngress(name, namespace, dto)
export const fetchDeleteIngress = async (name: string, namespace: string) =>
  DeleteIngress(name, namespace)
