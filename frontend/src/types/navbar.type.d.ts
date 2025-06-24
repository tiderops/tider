export interface NavbarDto {
  Name: string
  IsVisible: boolean
  IsEditable: boolean
  K8sObject: K8sObjectDto[]
}
export interface K8sObjectDto {
  Name: string
  Link: string
  IsVisible: boolean
  IsEditable: boolean
}
