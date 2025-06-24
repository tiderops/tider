/// <reference types="../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
import { computed, onMounted, reactive, ref } from 'vue'
import { gridComposable } from '@/composables/GridComposable'
import KsSidebarDetail from './SidebarDetail.vue'
import KsGridTable from '../components/GridTableComponent.vue'
import KsGridHeader from '../components/GridHeaderComponent.vue'
const snackbar = ref(false)
const text = ref('')
const timeout = ref(2000)
const props = defineProps()
// Composables & constants
const response = gridComposable(props.cluster, props.k8sObject)
const namespaces = ['ns-local', 'ns-dev']
const statuses = ['Running', 'Succeeded', 'Pending']
// Reactive state
const items = ref([])
const search = ref('')
const filterNamespace = ref('')
const filterStatus = ref('')
const sortBy = ref([{ key: 'name', order: 'asc' }])
const header = reactive({
  header: [],
})
// Filtering logic
const filteredItems = computed(() =>
  items.value.filter((item) => {
    const namespaceMatch = !filterNamespace.value || item.namespace.includes(filterNamespace.value)
    const statusMatch = !filterStatus.value || item.status.includes(filterStatus.value)
    const searchMatch =
      !search.value || item.name.toLowerCase().includes(search.value.toLowerCase())
    return namespaceMatch && statusMatch && searchMatch
  }),
)
// Sidebar logic
const isSidebarVisible = ref(false)
const selectedRow = ref(null)
const onRowClick = (cellData, item) => {
  selectedRow.value = item.item
  if (isSidebarVisible.value) {
    isSidebarVisible.value = false
    setTimeout(() => {
      isSidebarVisible.value = true
    }, 100)
  } else {
    isSidebarVisible.value = true
  }
}
const onEditItem = async (item) => {
  console.log('Parent received edit:', item)
  text.value = `Resource "${item.name}" was edited.`
  snackbar.value = true
  await response.fetchData()
  items.value = response.content?.body.value ?? []
}
const onDeleteItem = async (item) => {
  console.log('Parent received delete:', item)
  text.value = `Resource "${item.name}" was deleted.`
  snackbar.value = true
  await response.fetchData()
  items.value = response.content?.body.value ?? []
}
// Fetch data on mount
onMounted(async () => {
  await response.fetchData()
  header.header = response.content?.head.value ?? []
  items.value = response.content?.body.value ?? []
})
debugger /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {}
let __VLS_components
let __VLS_directives
const __VLS_0 = {}.VContainer
/** @type {[typeof __VLS_components.VContainer, typeof __VLS_components.vContainer, typeof __VLS_components.VContainer, typeof __VLS_components.vContainer, ]} */ // @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({}))
const __VLS_2 = __VLS_1({}, ...__VLS_functionalComponentArgsRest(__VLS_1))
var __VLS_4 = {}
__VLS_3.slots.default
const __VLS_5 = {}.VCard
/** @type {[typeof __VLS_components.VCard, typeof __VLS_components.vCard, typeof __VLS_components.VCard, typeof __VLS_components.vCard, ]} */ // @ts-ignore
const __VLS_6 = __VLS_asFunctionalComponent(__VLS_5, new __VLS_5({}))
const __VLS_7 = __VLS_6({}, ...__VLS_functionalComponentArgsRest(__VLS_6))
__VLS_8.slots.default
/** @type {[typeof KsGridHeader, ]} */ // @ts-ignore
const __VLS_9 = __VLS_asFunctionalComponent(
  KsGridHeader,
  new KsGridHeader({
    search: __VLS_ctx.search,
    filterNamespace: __VLS_ctx.filterNamespace,
    filterStatus: __VLS_ctx.filterStatus,
    namespaces: __VLS_ctx.namespaces,
    statuses: __VLS_ctx.statuses,
  }),
)
const __VLS_10 = __VLS_9(
  {
    search: __VLS_ctx.search,
    filterNamespace: __VLS_ctx.filterNamespace,
    filterStatus: __VLS_ctx.filterStatus,
    namespaces: __VLS_ctx.namespaces,
    statuses: __VLS_ctx.statuses,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_9),
)
/** @type {[typeof KsGridTable, ]} */ // @ts-ignore
const __VLS_12 = __VLS_asFunctionalComponent(
  KsGridTable,
  new KsGridTable({
    ...{ 'onClick:row': {} },
    ...{ onDelete: {} },
    ...{ onEdit: {} },
    cluster: props.cluster,
    headers: __VLS_ctx.header.header,
    items: __VLS_ctx.filteredItems,
    search: __VLS_ctx.search,
    sortBy: __VLS_ctx.sortBy,
  }),
)
const __VLS_13 = __VLS_12(
  {
    ...{ 'onClick:row': {} },
    ...{ onDelete: {} },
    ...{ onEdit: {} },
    cluster: props.cluster,
    headers: __VLS_ctx.header.header,
    items: __VLS_ctx.filteredItems,
    search: __VLS_ctx.search,
    sortBy: __VLS_ctx.sortBy,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_12),
)
let __VLS_15
let __VLS_16
let __VLS_17
const __VLS_18 = {
  'onClick:row': __VLS_ctx.onRowClick,
}
const __VLS_19 = {
  onDelete: __VLS_ctx.onDeleteItem,
}
const __VLS_20 = {
  onEdit: __VLS_ctx.onEditItem,
}
var __VLS_14
var __VLS_8
const __VLS_21 = {}.VSnackbar
/** @type {[typeof __VLS_components.VSnackbar, typeof __VLS_components.vSnackbar, typeof __VLS_components.VSnackbar, typeof __VLS_components.vSnackbar, ]} */ // @ts-ignore
const __VLS_22 = __VLS_asFunctionalComponent(
  __VLS_21,
  new __VLS_21({
    modelValue: __VLS_ctx.snackbar,
    timeout: __VLS_ctx.timeout,
  }),
)
const __VLS_23 = __VLS_22(
  {
    modelValue: __VLS_ctx.snackbar,
    timeout: __VLS_ctx.timeout,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_22),
)
__VLS_24.slots.default
__VLS_ctx.text
{
  const { actions: __VLS_thisSlot } = __VLS_24.slots
  const __VLS_25 = {}.VBtn
  /** @type {[typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, ]} */ // @ts-ignore
  const __VLS_26 = __VLS_asFunctionalComponent(
    __VLS_25,
    new __VLS_25({
      ...{ onClick: {} },
      color: 'blue',
      variant: 'text',
    }),
  )
  const __VLS_27 = __VLS_26(
    {
      ...{ onClick: {} },
      color: 'blue',
      variant: 'text',
    },
    ...__VLS_functionalComponentArgsRest(__VLS_26),
  )
  let __VLS_29
  let __VLS_30
  let __VLS_31
  const __VLS_32 = {
    onClick: (...[$event]) => {
      __VLS_ctx.snackbar = false
    },
  }
  __VLS_28.slots.default
  var __VLS_28
}
var __VLS_24
const __VLS_33 = {}.VCard
/** @type {[typeof __VLS_components.VCard, typeof __VLS_components.vCard, typeof __VLS_components.VCard, typeof __VLS_components.vCard, ]} */ // @ts-ignore
const __VLS_34 = __VLS_asFunctionalComponent(__VLS_33, new __VLS_33({}))
const __VLS_35 = __VLS_34({}, ...__VLS_functionalComponentArgsRest(__VLS_34))
__VLS_36.slots.default
/** @type {[typeof KsSidebarDetail, ]} */ // @ts-ignore
const __VLS_37 = __VLS_asFunctionalComponent(
  KsSidebarDetail,
  new KsSidebarDetail({
    ...{ onClose: {} },
    isVisible: __VLS_ctx.isSidebarVisible,
    selectedRow: __VLS_ctx.selectedRow,
  }),
)
const __VLS_38 = __VLS_37(
  {
    ...{ onClose: {} },
    isVisible: __VLS_ctx.isSidebarVisible,
    selectedRow: __VLS_ctx.selectedRow,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_37),
)
let __VLS_40
let __VLS_41
let __VLS_42
const __VLS_43 = {
  onClose: (...[$event]) => {
    __VLS_ctx.isSidebarVisible = false
  },
}
var __VLS_39
var __VLS_36
var __VLS_3
var __VLS_dollars
const __VLS_self = (await import('vue')).defineComponent({
  setup() {
    return {
      KsSidebarDetail: KsSidebarDetail,
      KsGridTable: KsGridTable,
      KsGridHeader: KsGridHeader,
      snackbar: snackbar,
      text: text,
      timeout: timeout,
      namespaces: namespaces,
      statuses: statuses,
      search: search,
      filterNamespace: filterNamespace,
      filterStatus: filterStatus,
      sortBy: sortBy,
      header: header,
      filteredItems: filteredItems,
      isSidebarVisible: isSidebarVisible,
      selectedRow: selectedRow,
      onRowClick: onRowClick,
      onEditItem: onEditItem,
      onDeleteItem: onDeleteItem,
    }
  },
  __typeProps: {},
})
export default (await import('vue')).defineComponent({
  setup() {
    return {}
  },
  __typeProps: {},
}) /* PartiallyEnd: #4569/main.vue */
