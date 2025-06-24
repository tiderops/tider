import { database, model } from '../../wailsjs/go/models'
export declare function sidebarComposable(): {
  commonParameters: import('vue').Ref<
    {
      Name: string
      Link: string
      Icon: string
    }[],
    | database.CommonParameterDto[]
    | {
        Name: string
        Link: string
        Icon: string
      }[]
  >
  kubernetesParameters: import('vue').Ref<
    {
      Name: string
      Link: string
      Icon: string
    }[],
    | database.CommonParameterDto[]
    | {
        Name: string
        Link: string
        Icon: string
      }[]
  >
  clusters: import('vue').Ref<
    {
      Name: string
      Cluster: string
      Server: string
      User: string
      Namespace: string
      Status: boolean
      Source: string
    }[],
    | model.ClusterInfo[]
    | {
        Name: string
        Cluster: string
        Server: string
        User: string
        Namespace: string
        Status: boolean
        Source: string
      }[]
  >
  fetchData: () => Promise<void>
}
