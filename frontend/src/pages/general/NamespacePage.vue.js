/// <reference types="../../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
import { defineComponent } from 'vue'
import KsGridBodyV2 from '../../layout/GridBody2.vue'
import KsNavBar from '../../layout/Navbar.vue'
export default defineComponent({
  name: 'NamespacePage',
  components: { KsNavBar, KsGridBody, KsGridBodyV2 },
  data() {
    return {
      k8sObject: 'namespace',
      namespace: 'mock',
    }
  },
})
console.log('PRUEBA CARGA NS')
debugger /* PartiallyEnd: #3632/script.vue */
const __VLS_ctx = {}
const __VLS_componentsOption = { KsNavBar, KsGridBody, KsGridBodyV2 }
let __VLS_components
let __VLS_directives
// CSS variable injection
// CSS variable injection end
const __VLS_0 = {}.KsGridBodyV2
/** @type {[typeof __VLS_components.KsGridBodyV2, typeof __VLS_components.ksGridBodyV2, typeof __VLS_components.KsGridBodyV2, typeof __VLS_components.ksGridBodyV2, ]} */ // @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(
  __VLS_0,
  new __VLS_0({
    namespace: __VLS_ctx.namespace,
    k8sObject: __VLS_ctx.k8sObject,
  }),
)
const __VLS_2 = __VLS_1(
  {
    namespace: __VLS_ctx.namespace,
    k8sObject: __VLS_ctx.k8sObject,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_1),
)
var __VLS_4 = {}
var __VLS_3
var __VLS_dollars
let __VLS_self
