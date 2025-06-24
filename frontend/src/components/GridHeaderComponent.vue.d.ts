type __VLS_Props = {
  search?: string
  filterNamespace?: string
  filterStatus?: string
  namespaces?: string[]
  statuses?: string[]
}
declare const _default: import('vue').DefineComponent<
  __VLS_Props,
  {},
  {},
  {},
  {},
  import('vue').ComponentOptionsMixin,
  import('vue').ComponentOptionsMixin,
  {} & {
    'update:search': (value: string) => any
    'update:filterNamespace': (value: string) => any
    'update:filterStatus': (value: string) => any
  },
  string,
  import('vue').PublicProps,
  Readonly<__VLS_Props> &
    Readonly<{
      'onUpdate:search'?: ((value: string) => any) | undefined
      'onUpdate:filterNamespace'?: ((value: string) => any) | undefined
      'onUpdate:filterStatus'?: ((value: string) => any) | undefined
    }>,
  {},
  {},
  {},
  {},
  string,
  import('vue').ComponentProvideOptions,
  false,
  {},
  any
>
export default _default
