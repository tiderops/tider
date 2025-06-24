/// <reference types="../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
import { reactive } from 'vue'
// import {Greet} from '../../wailsjs/go/main/App'
const data = reactive({
  name: '',
  resultText: 'Please enter your name below 👇',
})
function greet() {
  // Greet(data.name).then(result => {
  //   data.resultText = result
  // })
}
debugger /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {}
let __VLS_components
let __VLS_directives
/** @type {__VLS_StyleScopedClasses['input-box']} */ /** @type {__VLS_StyleScopedClasses['btn']} */ /** @type {__VLS_StyleScopedClasses['input-box']} */ /** @type {__VLS_StyleScopedClasses['input-box']} */ /** @type {__VLS_StyleScopedClasses['input']} */ /** @type {__VLS_StyleScopedClasses['input-box']} */ /** @type {__VLS_StyleScopedClasses['input']} */ // CSS variable injection
// CSS variable injection end
__VLS_asFunctionalElement(__VLS_intrinsicElements.main, __VLS_intrinsicElements.main)({})
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.div,
  __VLS_intrinsicElements.div,
)({
  id: 'result',
  ...{ class: 'result' },
})
__VLS_ctx.data.resultText
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.div,
  __VLS_intrinsicElements.div,
)({
  id: 'input',
  ...{ class: 'input-box' },
})
__VLS_asFunctionalElement(__VLS_intrinsicElements.input)({
  id: 'name',
  value: __VLS_ctx.data.name,
  autocomplete: 'off',
  ...{ class: 'input' },
  type: 'text',
})
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.button,
  __VLS_intrinsicElements.button,
)({
  ...{ onClick: __VLS_ctx.greet },
  ...{ class: 'btn' },
})
/** @type {__VLS_StyleScopedClasses['result']} */ /** @type {__VLS_StyleScopedClasses['input-box']} */ /** @type {__VLS_StyleScopedClasses['input']} */ /** @type {__VLS_StyleScopedClasses['btn']} */ var __VLS_dollars
const __VLS_self = (await import('vue')).defineComponent({
  setup() {
    return {
      data: data,
      greet: greet,
    }
  },
})
export default (await import('vue')).defineComponent({
  setup() {
    return {}
  },
}) /* PartiallyEnd: #4569/main.vue */
