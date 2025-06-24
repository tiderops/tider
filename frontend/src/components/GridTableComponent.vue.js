/// <reference types="../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
import { gridGeneralComposable } from '@/composables/GridTableComposable'
const props = defineProps()
const emit = defineEmits()
const editPod = (item) => {
  console.log('EDIT', item)
  emit('edit', item)
}
const deletePod = async (item) => {
  console.log('DELETE', item)
  console.log('DELETE - NAME', item.name)
  console.log('DELETE - NS', item.namespace)
  console.log('DELETE - CLUSTER ID', props.cluster)
  const { fetchData } = gridGeneralComposable(item.name, item.namespace, props.cluster)
  await fetchData()
  emit('delete', item)
}
debugger /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {}
let __VLS_components
let __VLS_directives
// CSS variable injection
// CSS variable injection end
const __VLS_0 = {}.VDataTableVirtual
/** @type {[typeof __VLS_components.VDataTableVirtual, typeof __VLS_components.vDataTableVirtual, typeof __VLS_components.VDataTableVirtual, typeof __VLS_components.vDataTableVirtual, ]} */ // @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(
  __VLS_0,
  new __VLS_0({
    headers: props.headers,
    items: props.items,
    search: props.search,
    sortBy: props.sortBy,
    height: '720',
    itemValue: 'name',
    density: 'compact',
    fixedHeader: true,
    hover: true,
  }),
)
const __VLS_2 = __VLS_1(
  {
    headers: props.headers,
    items: props.items,
    search: props.search,
    sortBy: props.sortBy,
    height: '720',
    itemValue: 'name',
    density: 'compact',
    fixedHeader: true,
    hover: true,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_1),
)
var __VLS_4 = {}
__VLS_3.slots.default
{
  const { 'item.status': __VLS_thisSlot } = __VLS_3.slots
  const [{ item }] = __VLS_getSlotParams(__VLS_thisSlot)
  const __VLS_5 = {}.VChip
  /** @type {[typeof __VLS_components.VChip, typeof __VLS_components.vChip, typeof __VLS_components.VChip, typeof __VLS_components.vChip, ]} */ // @ts-ignore
  const __VLS_6 = __VLS_asFunctionalComponent(
    __VLS_5,
    new __VLS_5({
      color: item.status === 'Running' ? 'green' : 'red',
      dark: true,
    }),
  )
  const __VLS_7 = __VLS_6(
    {
      color: item.status === 'Running' ? 'green' : 'red',
      dark: true,
    },
    ...__VLS_functionalComponentArgsRest(__VLS_6),
  )
  __VLS_8.slots.default
  item.status
  var __VLS_8
}
{
  const { 'item.actions': __VLS_thisSlot } = __VLS_3.slots
  const [{ item }] = __VLS_getSlotParams(__VLS_thisSlot)
  const __VLS_9 = {}.VBtn
  /** @type {[typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, ]} */ // @ts-ignore
  const __VLS_10 = __VLS_asFunctionalComponent(
    __VLS_9,
    new __VLS_9({
      ...{ onClick: {} },
      icon: true,
    }),
  )
  const __VLS_11 = __VLS_10(
    {
      ...{ onClick: {} },
      icon: true,
    },
    ...__VLS_functionalComponentArgsRest(__VLS_10),
  )
  let __VLS_13
  let __VLS_14
  let __VLS_15
  const __VLS_16 = {
    onClick: (...[$event]) => {
      __VLS_ctx.editPod(item)
    },
  }
  __VLS_12.slots.default
  const __VLS_17 = {}.VIcon
  /** @type {[typeof __VLS_components.VIcon, typeof __VLS_components.vIcon, ]} */ // @ts-ignore
  const __VLS_18 = __VLS_asFunctionalComponent(
    __VLS_17,
    new __VLS_17({
      icon: 'mdi-pencil',
    }),
  )
  const __VLS_19 = __VLS_18(
    {
      icon: 'mdi-pencil',
    },
    ...__VLS_functionalComponentArgsRest(__VLS_18),
  )
  var __VLS_12
  const __VLS_21 = {}.VBtn
  /** @type {[typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, ]} */ // @ts-ignore
  const __VLS_22 = __VLS_asFunctionalComponent(
    __VLS_21,
    new __VLS_21({
      ...{ onClick: {} },
      icon: true,
    }),
  )
  const __VLS_23 = __VLS_22(
    {
      ...{ onClick: {} },
      icon: true,
    },
    ...__VLS_functionalComponentArgsRest(__VLS_22),
  )
  let __VLS_25
  let __VLS_26
  let __VLS_27
  const __VLS_28 = {
    onClick: (...[$event]) => {
      __VLS_ctx.deletePod(item)
    },
  }
  __VLS_24.slots.default
  const __VLS_29 = {}.VIcon
  /** @type {[typeof __VLS_components.VIcon, typeof __VLS_components.vIcon, ]} */ // @ts-ignore
  const __VLS_30 = __VLS_asFunctionalComponent(
    __VLS_29,
    new __VLS_29({
      icon: 'mdi-delete',
    }),
  )
  const __VLS_31 = __VLS_30(
    {
      icon: 'mdi-delete',
    },
    ...__VLS_functionalComponentArgsRest(__VLS_30),
  )
  var __VLS_24
}
var __VLS_3
var __VLS_dollars
const __VLS_self = (await import('vue')).defineComponent({
  setup() {
    return {
      editPod: editPod,
      deletePod: deletePod,
    }
  },
  __typeEmits: {},
  __typeProps: {},
})
export default (await import('vue')).defineComponent({
  setup() {
    return {}
  },
  __typeEmits: {},
  __typeProps: {},
}) /* PartiallyEnd: #4569/main.vue */
