/// <reference types="../../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
import { defineComponent, ref, onMounted } from 'vue'
import KsNavBar from '../../layout/Navbar.vue'
// TODO: Revisar el funciona wails events o usar SSE
export default defineComponent({
  name: 'NetworkPage',
  components: { KsNavBar },
  data() {
    return {
      followingPage: 'Network',
    }
  },
  // setup() {
  //     const messages = ref<string[]>([]);
  //
  //     // onMounted(() => {
  //     //     // Listen for the "new-message" event emitted from the backend
  //     //     window.runtime.EventsOn("new-message", (message: string) => {
  //     //         messages.value.push(message);
  //     //         console.log("Received:", message);
  //     //     });
  //     //
  //     //     // Optionally call the backend to start emitting messages
  //     //     window.runtime.Call("EmitMessages").catch((err: any) => {
  //     //         console.error("Error calling EmitMessages:", err);
  //     //     });
  //     // });
  //
  //     onMounted(() => {
  //         window.runtime.EventsOn("new-message", (message: string) => {
  //             console.log("Received message:", message);
  //         });
  //
  //         window.runtime.Call("EmitMessages").catch((err: any) => {
  //             console.error("Error:", err);
  //         });
  //     });
  //
  //     return {
  //         messages,
  //     };
  // },
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
__VLS_asFunctionalElement(__VLS_intrinsicElements.h2, __VLS_intrinsicElements.h2)({})
var __VLS_dollars
let __VLS_self
