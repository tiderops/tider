import {
  DeleteNamespace, ExportNamespaceObjects,
  GetNamespaces,
  GetNamespace,
  GetNode,
  GetNodes,
  UpdateNamespace,
} from '../../wailsjs/go/middleware/GeneralMiddleware'
import { model } from '../../wailsjs/go/models'
import NamespaceDto = model.NamespaceDto

export const fetchGetNodes = async () => GetNodes()
export const fetchGetNode = async (name: string) => GetNode(name)

export const fetchGetNamespaces = async (): Promise<NamespaceDto[]> => GetNamespaces()
export const fetchGetNamespace = async (name: string) => GetNamespace(name)
export const fetchUpdateNamespace = async (name: string, dto: any) =>
  UpdateNamespace(name, dto)
export const fetchDeleteNamespace = async (name: string) => DeleteNamespace(name)

export const fetchExportNamespaceObject = async (name: string, directory: string) => ExportNamespaceObjects(name, directory)
