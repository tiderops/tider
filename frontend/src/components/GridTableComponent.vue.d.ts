type __VLS_Props = {
  cluster?: string
  headers: any[]
  items: any[]
  search?: string
  sortBy?: any[]
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
    delete: (item: any) => any
    rowClick: (item: any) => any
    edit: (item: any) => any
  },
  string,
  import('vue').PublicProps,
  Readonly<__VLS_Props> &
    Readonly<{
      onDelete?: ((item: any) => any) | undefined
      onRowClick?: ((item: any) => any) | undefined
      onEdit?: ((item: any) => any) | undefined
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
