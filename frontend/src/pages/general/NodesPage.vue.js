/// <reference types="../../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
import { defineComponent } from 'vue'
import KsNavBar from '../../layout/Navbar.vue'
export default defineComponent({
  name: 'NodesPage',
  components: { KsNavBar, KsGridBody },
  data() {
    return {
      k8sObject: 'node',
      namespace: 'mock',
    }
  },
})
debugger /* PartiallyEnd: #3632/script.vue */
const __VLS_ctx = {}
const __VLS_componentsOption = { KsNavBar, KsGridBody }
let __VLS_components
let __VLS_directives
// CSS variable injection
// CSS variable injection end
__VLS_asFunctionalElement(__VLS_intrinsicElements.h2, __VLS_intrinsicElements.h2)({})
var __VLS_dollars
let __VLS_self
