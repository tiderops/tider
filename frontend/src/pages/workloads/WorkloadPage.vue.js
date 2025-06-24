/// <reference types="../../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import KsNavBar from '../../layout/Navbar.vue'
const followingPage = ref('Workload')
const route = useRoute()
const clusterId = route.params.cluster
console.log('Cluster ID:', clusterId)
debugger /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {}
let __VLS_components
let __VLS_directives
// CSS variable injection
// CSS variable injection end
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({})
/** @type {[typeof KsNavBar, typeof KsNavBar, ]} */ // @ts-ignore
const __VLS_0 = __VLS_asFunctionalComponent(
  KsNavBar,
  new KsNavBar({
    content: __VLS_ctx.followingPage,
  }),
)
const __VLS_1 = __VLS_0(
  {
    content: __VLS_ctx.followingPage,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_0),
)
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({})
const __VLS_3 = {}.RouterView
/** @type {[typeof __VLS_components.RouterView, typeof __VLS_components.routerView, typeof __VLS_components.RouterView, typeof __VLS_components.routerView, ]} */ // @ts-ignore
const __VLS_4 = __VLS_asFunctionalComponent(__VLS_3, new __VLS_3({}))
const __VLS_5 = __VLS_4({}, ...__VLS_functionalComponentArgsRest(__VLS_4))
__VLS_asFunctionalElement(__VLS_intrinsicElements.h2, __VLS_intrinsicElements.h2)({})
var __VLS_dollars
const __VLS_self = (await import('vue')).defineComponent({
  setup() {
    return {
      KsNavBar: KsNavBar,
      followingPage: followingPage,
    }
  },
})
export default (await import('vue')).defineComponent({
  setup() {
    return {}
  },
}) /* PartiallyEnd: #4569/main.vue */
