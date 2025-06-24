/// <reference types="../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
const props = defineProps()
const emit = defineEmits()
debugger /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {}
let __VLS_components
let __VLS_directives
// CSS variable injection
// CSS variable injection end
const __VLS_0 = {}.VToolbar
/** @type {[typeof __VLS_components.VToolbar, typeof __VLS_components.vToolbar, typeof __VLS_components.VToolbar, typeof __VLS_components.vToolbar, ]} */ // @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(
  __VLS_0,
  new __VLS_0({
    flat: true,
  }),
)
const __VLS_2 = __VLS_1(
  {
    flat: true,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_1),
)
var __VLS_4 = {}
__VLS_3.slots.default
const __VLS_5 = {}.VToolbarTitle
/** @type {[typeof __VLS_components.VToolbarTitle, typeof __VLS_components.vToolbarTitle, typeof __VLS_components.VToolbarTitle, typeof __VLS_components.vToolbarTitle, ]} */ // @ts-ignore
const __VLS_6 = __VLS_asFunctionalComponent(__VLS_5, new __VLS_5({}))
const __VLS_7 = __VLS_6({}, ...__VLS_functionalComponentArgsRest(__VLS_6))
__VLS_8.slots.default
var __VLS_8
const __VLS_9 = {}.VSpacer
/** @type {[typeof __VLS_components.VSpacer, typeof __VLS_components.vSpacer, typeof __VLS_components.VSpacer, typeof __VLS_components.vSpacer, ]} */ // @ts-ignore
const __VLS_10 = __VLS_asFunctionalComponent(__VLS_9, new __VLS_9({}))
const __VLS_11 = __VLS_10({}, ...__VLS_functionalComponentArgsRest(__VLS_10))
const __VLS_13 = {}.VTextField
/** @type {[typeof __VLS_components.VTextField, typeof __VLS_components.vTextField, typeof __VLS_components.VTextField, typeof __VLS_components.vTextField, ]} */ // @ts-ignore
const __VLS_14 = __VLS_asFunctionalComponent(
  __VLS_13,
  new __VLS_13({
    ...{ 'onUpdate:modelValue': {} },
    modelValue: props.search,
    label: 'Search',
    clearable: true,
    ...{ class: 'mx-4' },
  }),
)
const __VLS_15 = __VLS_14(
  {
    ...{ 'onUpdate:modelValue': {} },
    modelValue: props.search,
    label: 'Search',
    clearable: true,
    ...{ class: 'mx-4' },
  },
  ...__VLS_functionalComponentArgsRest(__VLS_14),
)
let __VLS_17
let __VLS_18
let __VLS_19
const __VLS_20 = {
  'onUpdate:modelValue': (...[$event]) => {
    __VLS_ctx.emit('update:search', $event)
  },
}
var __VLS_16
const __VLS_21 = {}.VSelect
/** @type {[typeof __VLS_components.VSelect, typeof __VLS_components.vSelect, typeof __VLS_components.VSelect, typeof __VLS_components.vSelect, ]} */ // @ts-ignore
const __VLS_22 = __VLS_asFunctionalComponent(
  __VLS_21,
  new __VLS_21({
    ...{ 'onUpdate:modelValue': {} },
    modelValue: props.filterNamespace,
    items: props.namespaces,
    label: 'Namespace',
    clearable: true,
    ...{ class: 'mx-5' },
  }),
)
const __VLS_23 = __VLS_22(
  {
    ...{ 'onUpdate:modelValue': {} },
    modelValue: props.filterNamespace,
    items: props.namespaces,
    label: 'Namespace',
    clearable: true,
    ...{ class: 'mx-5' },
  },
  ...__VLS_functionalComponentArgsRest(__VLS_22),
)
let __VLS_25
let __VLS_26
let __VLS_27
const __VLS_28 = {
  'onUpdate:modelValue': (...[$event]) => {
    __VLS_ctx.emit('update:filterNamespace', $event)
  },
}
var __VLS_24
const __VLS_29 = {}.VSelect
/** @type {[typeof __VLS_components.VSelect, typeof __VLS_components.vSelect, typeof __VLS_components.VSelect, typeof __VLS_components.vSelect, ]} */ // @ts-ignore
const __VLS_30 = __VLS_asFunctionalComponent(
  __VLS_29,
  new __VLS_29({
    ...{ 'onUpdate:modelValue': {} },
    modelValue: props.filterStatus,
    items: props.statuses,
    label: 'Status',
    clearable: true,
    ...{ class: 'mx-5' },
  }),
)
const __VLS_31 = __VLS_30(
  {
    ...{ 'onUpdate:modelValue': {} },
    modelValue: props.filterStatus,
    items: props.statuses,
    label: 'Status',
    clearable: true,
    ...{ class: 'mx-5' },
  },
  ...__VLS_functionalComponentArgsRest(__VLS_30),
)
let __VLS_33
let __VLS_34
let __VLS_35
const __VLS_36 = {
  'onUpdate:modelValue': (...[$event]) => {
    __VLS_ctx.emit('update:filterStatus', $event)
  },
}
var __VLS_32
var __VLS_3
/** @type {__VLS_StyleScopedClasses['mx-4']} */ /** @type {__VLS_StyleScopedClasses['mx-5']} */ /** @type {__VLS_StyleScopedClasses['mx-5']} */ var __VLS_dollars
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
