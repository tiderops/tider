/// <reference types="../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
import { ref } from 'vue'
const count = ref(0)
const increment = () => {
  count.value++
}
debugger /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {}
let __VLS_components
let __VLS_directives
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({})
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.p,
  __VLS_intrinsicElements.p,
)({
  'data-testid': 'count',
})
__VLS_ctx.count
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.button,
  __VLS_intrinsicElements.button,
)({
  ...{ onClick: __VLS_ctx.increment },
})
var __VLS_dollars
const __VLS_self = (await import('vue')).defineComponent({
  setup() {
    return {
      count: count,
      increment: increment,
    }
  },
})
export default (await import('vue')).defineComponent({
  setup() {
    return {}
  },
}) /* PartiallyEnd: #4569/main.vue */
