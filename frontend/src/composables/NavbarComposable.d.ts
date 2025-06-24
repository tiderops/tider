import { database } from '../../wailsjs/go/models'
import ObjectType = database.ObjectType
export declare function navbarComposable(): {
  objects: import('vue').Ref<
    {
      Name: string
      IsVisible: boolean
      IsEditable: boolean
      K8sObject: {
        Name: string
        Link: string
        IsVisible: boolean
        IsEditable: boolean
      }[]
      convertValues: (a: any, classs: any, asMap?: boolean) => any
    }[],
    | database.ObjectType[]
    | {
        Name: string
        IsVisible: boolean
        IsEditable: boolean
        K8sObject: {
          Name: string
          Link: string
          IsVisible: boolean
          IsEditable: boolean
        }[]
        convertValues: (a: any, classs: any, asMap?: boolean) => any
      }[]
  >
  fetchData: () => Promise<ObjectType[]>
}
