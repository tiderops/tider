import { database } from '../../../wailsjs/go/models'
declare const _default: import('vue').DefineComponent<
  {},
  {
    callMiddleware: () => Promise<void>
    response: import('vue').Ref<
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
  },
  {},
  {},
  {},
  import('vue').ComponentOptionsMixin,
  import('vue').ComponentOptionsMixin,
  {},
  string,
  import('vue').PublicProps,
  Readonly<{}> & Readonly<{}>,
  {},
  {},
  {},
  {},
  string,
  import('vue').ComponentProvideOptions,
  true,
  {},
  any
>
export default _default
