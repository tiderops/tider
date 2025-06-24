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
var NamespaceDto = model.NamespaceDto
export const fetchGetNodes = async (clusterCtx) => GetNodes(clusterCtx)
export const fetchGetNode = async (name, clusterCtx) => GetNode(name, clusterCtx)
export const fetchGetNamespaces = async (clusterCtx) => GetNamespaces(clusterCtx)
export const fetchGetNamespace = async (name, clusterCtx) => GetNamespace(name, clusterCtx)
export const fetchUpdateNamespace = async (name, dto, clusterCtx) =>
  UpdateNamespace(name, dto, clusterCtx)
export const fetchDeleteNamespace = async (name, clusterCtx) => DeleteNamespace(name, clusterCtx)
export const fetchExportNamespaceObject = async (name, directory, clusterCtx) =>
  ExportNamespaceObjects(name, directory, clusterCtx)
