import { model } from '../../wailsjs/go/models'
import NamespaceDto = model.NamespaceDto
export declare const fetchGetNodes: (clusterCtx: string) => Promise<model.NodeDtoV2[]>
export declare const fetchGetNode: (name: string, clusterCtx: string) => Promise<model.NodeDtoV2>
export declare const fetchGetNamespaces: (clusterCtx: string) => Promise<NamespaceDto[]>
export declare const fetchGetNamespace: (
  name: string,
  clusterCtx: string,
) => Promise<model.NamespaceDto>
export declare const fetchUpdateNamespace: (
  name: string,
  dto: any,
  clusterCtx: string,
) => Promise<void>
export declare const fetchDeleteNamespace: (name: string, clusterCtx: string) => Promise<void>
export declare const fetchExportNamespaceObject: (
  name: string,
  directory: string,
  clusterCtx: string,
) => Promise<void>
