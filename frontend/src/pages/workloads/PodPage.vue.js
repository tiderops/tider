/// <reference types="../../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
import KsGridBodyV2 from '../../layout/GridBody2.vue'
import { useRoute } from 'vue-router'
import { ref } from 'vue'
const k8sObject = ref('pod')
const namespace = ref('east')
const route = useRoute()
const clusterId = route.params.cluster
console.log('POD CLUSTER_ID:', clusterId)
debugger /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {}
let __VLS_components
let __VLS_directives
// CSS variable injection
// CSS variable injection end
/** @type {[typeof KsGridBodyV2, typeof KsGridBodyV2, ]} */ // @ts-ignore
const __VLS_0 = __VLS_asFunctionalComponent(
  KsGridBodyV2,
  new KsGridBodyV2({
    cluster: __VLS_ctx.clusterId,
    namespace: __VLS_ctx.namespace,
    k8sObject: __VLS_ctx.k8sObject,
  }),
)
const __VLS_1 = __VLS_0(
  {
    cluster: __VLS_ctx.clusterId,
    namespace: __VLS_ctx.namespace,
    k8sObject: __VLS_ctx.k8sObject,
  },
  ...__VLS_functionalComponentArgsRest(__VLS_0),
)
var __VLS_3 = {}
var __VLS_2
var __VLS_dollars
const __VLS_self = (await import('vue')).defineComponent({
  setup() {
    return {
      KsGridBodyV2: KsGridBodyV2,
      k8sObject: k8sObject,
      namespace: namespace,
      clusterId: clusterId,
    }
  },
})
export default (await import('vue')).defineComponent({
  setup() {
    return {}
  },
}) /* PartiallyEnd: #4569/main.vue */
