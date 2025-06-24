import { database, model } from '../../wailsjs/go/models'
export declare function gridGeneralComposable(k8sObject: string): {
  nodes: import('vue').Ref<any, any>
  namespaces: import('vue').Ref<
    {
      Name: string
      Version: string
      CreationTime: string
      Labels: Record<string, string>
      Status: string
    }[],
    | model.NamespaceDto[]
    | {
        Name: string
        Version: string
        CreationTime: string
        Labels: Record<string, string>
        Status: string
      }[]
  >
  headers: import('vue').Ref<
    {
      Title: string
      Key: string
      Align: string
      Sortable: boolean
    }[],
    | database.HeadParamsDto[]
    | {
        Title: string
        Key: string
        Align: string
        Sortable: boolean
      }[]
  >
  fetchData: () => Promise<any>
}
