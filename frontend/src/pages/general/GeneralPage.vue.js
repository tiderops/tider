/// <reference types="../../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
import { defineComponent } from 'vue'
import KsNavBar from '../../layout/Navbar.vue'
export default defineComponent({
  name: 'GeneralPage',
  components: { KsNavBar },
  data() {
    return {
      followingPage: 'General',
    }
  },
  setup() {},
})
debugger /* PartiallyEnd: #3632/script.vue */
const __VLS_ctx = {}
const __VLS_componentsOption = { KsNavBar }
let __VLS_components
let __VLS_directives
// CSS variable injection
// CSS variable injection end
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({})
const __VLS_0 = {}.KsNavBar
/** @type {[typeof __VLS_components.KsNavBar, typeof __VLS_components.ksNavBar, typeof __VLS_components.KsNavBar, typeof __VLS_components.ksNavBar, ]} */ // @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(
  __VLS_0,
  new __VLS_0({
    content: __VLS_ctx.followingPage,
  }),
)
const __VLS_2 = __VLS_1(
  {
    content: __VLS_ctx.followingPage,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_1),
)
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({})
const __VLS_4 = {}.RouterView
/** @type {[typeof __VLS_components.RouterView, typeof __VLS_components.routerView, typeof __VLS_components.RouterView, typeof __VLS_components.routerView, ]} */ // @ts-ignore
const __VLS_5 = __VLS_asFunctionalComponent(__VLS_4, new __VLS_4({}))
const __VLS_6 = __VLS_5({}, ...__VLS_functionalComponentArgsRest(__VLS_5))
var __VLS_dollars
let __VLS_self
