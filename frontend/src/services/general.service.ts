import {
  DeleteNamespace,
  ExportNamespaceObjects,
  GetNamespaces,
  GetNamespace,
  GetNode,
  GetNodes,
  UpdateNamespace,
} from '../../wailsjs/go/middleware/GeneralMiddleware'
import { model } from '../../wailsjs/go/models'
import NamespaceDto = model.NamespaceDto

export const fetchGetNodes = async (clusterCtx: string) => GetNodes(clusterCtx)
export const fetchGetNode = async (name: string, clusterCtx: string) => GetNode(name, clusterCtx)

export const fetchGetNamespaces = async (clusterCtx: string): Promise<NamespaceDto[]> =>
  GetNamespaces(clusterCtx)
export const fetchGetNamespace = async (name: string, clusterCtx: string) =>
  GetNamespace(name, clusterCtx)
export const fetchUpdateNamespace = async (name: string, dto: any, clusterCtx: string) =>
  UpdateNamespace(name, dto, clusterCtx)
export const fetchDeleteNamespace = async (name: string, clusterCtx: string) =>
  DeleteNamespace(name, clusterCtx)

export const fetchExportNamespaceObject = async (
  name: string,
  directory: string,
  clusterCtx: string,
) => ExportNamespaceObjects(name, directory, clusterCtx)
