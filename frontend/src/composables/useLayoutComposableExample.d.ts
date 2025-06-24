import { database } from '../../wailsjs/go/models'
import CommonParameterDto = database.CommonParameterDto
export declare function useLayoutComposableExample(): {
  result: import('vue').Ref<
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
  fetchData: () => Promise<CommonParameterDto[]>
}
