/// <reference types="../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
import { defineProps, defineEmits } from 'vue'
const props = defineProps()
const emit = defineEmits()
debugger /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {}
let __VLS_components
let __VLS_directives
// CSS variable injection
// CSS variable injection end
const __VLS_0 = {}.VLayout
/** @type {[typeof __VLS_components.VLayout, typeof __VLS_components.vLayout, typeof __VLS_components.VLayout, typeof __VLS_components.vLayout, ]} */ // @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(__VLS_0, new __VLS_0({}))
const __VLS_2 = __VLS_1({}, ...__VLS_functionalComponentArgsRest(__VLS_1))
var __VLS_4 = {}
__VLS_3.slots.default
const __VLS_5 = {}.VNavigationDrawer
/** @type {[typeof __VLS_components.VNavigationDrawer, typeof __VLS_components.vNavigationDrawer, typeof __VLS_components.VNavigationDrawer, typeof __VLS_components.vNavigationDrawer, ]} */ // @ts-ignore
const __VLS_6 = __VLS_asFunctionalComponent(
  __VLS_5,
  new __VLS_5({
    modelValue: props.isVisible,
    location: 'right',
    width: 750,
    temporary: true,
    ...{ class: 'custom-sidebar' },
  }),
)
const __VLS_7 = __VLS_6(
  {
    modelValue: props.isVisible,
    location: 'right',
    width: 750,
    temporary: true,
    ...{ class: 'custom-sidebar' },
  },
  ...__VLS_functionalComponentArgsRest(__VLS_6),
)
__VLS_8.slots.default
const __VLS_9 = {}.VCard
/** @type {[typeof __VLS_components.VCard, typeof __VLS_components.vCard, typeof __VLS_components.VCard, typeof __VLS_components.vCard, ]} */ // @ts-ignore
const __VLS_10 = __VLS_asFunctionalComponent(
  __VLS_9,
  new __VLS_9({
    flat: true,
  }),
)
const __VLS_11 = __VLS_10(
  {
    flat: true,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_10),
)
__VLS_12.slots.default
{
  const { prepend: __VLS_thisSlot } = __VLS_12.slots
  const __VLS_13 = {}.VBtn
  /** @type {[typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, typeof __VLS_components.VBtn, typeof __VLS_components.vBtn, ]} */ // @ts-ignore
  const __VLS_14 = __VLS_asFunctionalComponent(
    __VLS_13,
    new __VLS_13({
      ...{ onClick: {} },
      icon: 'mdi-chevron-left',
      variant: 'text',
    }),
  )
  const __VLS_15 = __VLS_14(
    {
      ...{ onClick: {} },
      icon: 'mdi-chevron-left',
      variant: 'text',
    },
    ...__VLS_functionalComponentArgsRest(__VLS_14),
  )
  let __VLS_17
  let __VLS_18
  let __VLS_19
  const __VLS_20 = {
    onClick: (...[$event]) => {
      __VLS_ctx.emit('close')
    },
  }
  var __VLS_16
  const __VLS_21 = {}.VCardTitle
  /** @type {[typeof __VLS_components.VCardTitle, typeof __VLS_components.vCardTitle, typeof __VLS_components.VCardTitle, typeof __VLS_components.vCardTitle, ]} */ // @ts-ignore
  const __VLS_22 = __VLS_asFunctionalComponent(
    __VLS_21,
    new __VLS_21({
      ...{ style: {} },
    }),
  )
  const __VLS_23 = __VLS_22(
    {
      ...{ style: {} },
    },
    ...__VLS_functionalComponentArgsRest(__VLS_22),
  )
  __VLS_24.slots.default
  var __VLS_24
}
const __VLS_25 = {}.VDivider
/** @type {[typeof __VLS_components.VDivider, typeof __VLS_components.vDivider, typeof __VLS_components.VDivider, typeof __VLS_components.vDivider, ]} */ // @ts-ignore
const __VLS_26 = __VLS_asFunctionalComponent(__VLS_25, new __VLS_25({}))
const __VLS_27 = __VLS_26({}, ...__VLS_functionalComponentArgsRest(__VLS_26))
if (props.selectedRow) {
  const __VLS_29 = {}.VCardText
  /** @type {[typeof __VLS_components.VCardText, typeof __VLS_components.vCardText, typeof __VLS_components.VCardText, typeof __VLS_components.vCardText, ]} */ // @ts-ignore
  const __VLS_30 = __VLS_asFunctionalComponent(__VLS_29, new __VLS_29({}))
  const __VLS_31 = __VLS_30({}, ...__VLS_functionalComponentArgsRest(__VLS_30))
  __VLS_32.slots.default
  __VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
  props.selectedRow.name
  __VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
  props.selectedRow.namespace
  __VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
  props.selectedRow.replicas
  __VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
  props.selectedRow.cpu
  __VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
  props.selectedRow.memory
  __VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
  props.selectedRow.age
  __VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
  props.selectedRow.status
  var __VLS_32
}
var __VLS_12
var __VLS_8
var __VLS_3
/** @type {__VLS_StyleScopedClasses['custom-sidebar']} */ var __VLS_dollars
const __VLS_self = (await import('vue')).defineComponent({
  setup() {
    return {
      emit: emit,
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
