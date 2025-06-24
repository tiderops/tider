type __VLS_Props = {
  isVisible: boolean
  selectedRow: {
    name?: string
    namespace?: string
    replicas?: number
    cpu?: string
    memory?: string
    age?: string
    status?: string
  }
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
    close: () => any
  },
  string,
  import('vue').PublicProps,
  Readonly<__VLS_Props> &
    Readonly<{
      onClose?: (() => any) | undefined
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
