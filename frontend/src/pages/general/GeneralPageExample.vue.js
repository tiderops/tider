import { defineComponent, ref } from 'vue'
import {
  GetCommonParameters,
  GetKubernetesParameters,
} from '../../../wailsjs/go/middleware/ParameterMiddleware'
import { database } from '../../../wailsjs/go/models'
var CommonParameterDto = database.CommonParameterDto
import { Init } from '@wailsapp/runtime'
import { useLayoutComposableExample } from '../../composables/useLayoutComposableExample'
import { sidebarComposable } from '../../composables/SidebarComposable'
console.log('HOLA GENERAL')
// let response : CommonParameterDto[] = []
// async function getCommonParameters() {
//     return GetCommonParameters().then((res) => {
//         console.log("GENERAL", res);
//         return res.map((r) => {
//             const dto: CommonParameterDto = {
//                 Name : r.Name,
//                 Link : r.Link,
//                 Icon : r.Icon
//             }
//
//             response.push(dto);
//         })
//
//     })
// }
// await getCommonParameters()
// console.log("response FINAL:", response)
export default defineComponent({
  name: 'GeneralPage',
  setup() {
    const response = ref([])
    const { result } = useLayoutComposableExample()
    const callMiddleware = async () => {
      try {
        response.value = result.value.map((r) => ({
          Name: r.Name,
          Link: r.Link,
          Icon: r.Icon,
        }))
        console.log('Result from Go:', response.value)
      } catch (error) {
        console.error('Error calling Go function:', error)
      }
    }
    return {
      callMiddleware,
      response,
    }
  },
})
// const isScrolled = ref(false)
//
// const handleScroll = () => {
//     isScrolled.value = window.scrollY > 50
// }
// Lifecycle hooks
// onMounted(() => {
//     window.addEventListener('scroll', handleScroll)
// })
//
// onUnmounted(() => {
//     window.removeEventListener('scroll', handleScroll)
// })
////////
// const {commonParameters, fetchData} = useLayoutComposableV2();
// await fetchData()
// console.log("commonParameters ", commonParameters.value);
debugger /* PartiallyEnd: #3632/script.vue */
const __VLS_ctx = {}
let __VLS_components
let __VLS_directives
/** @type {__VLS_StyleScopedClasses['nav-link']} */ /** @type {__VLS_StyleScopedClasses['nav-link']} */ // CSS variable injection
// CSS variable injection end
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.nav,
  __VLS_intrinsicElements.nav,
)({
  ...{ class: 'navbar' },
})
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.div,
  __VLS_intrinsicElements.div,
)({
  ...{ class: 'nav-content' },
})
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.div,
  __VLS_intrinsicElements.div,
)({
  ...{ class: 'nav-links' },
})
const __VLS_0 = {}.RouterLink
/** @type {[typeof __VLS_components.RouterLink, typeof __VLS_components.routerLink, typeof __VLS_components.RouterLink, typeof __VLS_components.routerLink, ]} */ // @ts-ignore
const __VLS_1 = __VLS_asFunctionalComponent(
  __VLS_0,
  new __VLS_0({
    to: '/pod',
    ...{ class: 'nav-link' },
  }),
)
const __VLS_2 = __VLS_1(
  {
    to: '/pod',
    ...{ class: 'nav-link' },
  },
  ...__VLS_functionalComponentArgsRest(__VLS_1),
)
__VLS_3.slots.default
var __VLS_3
const __VLS_4 = {}.RouterLink
/** @type {[typeof __VLS_components.RouterLink, typeof __VLS_components.routerLink, typeof __VLS_components.RouterLink, typeof __VLS_components.routerLink, ]} */ // @ts-ignore
const __VLS_5 = __VLS_asFunctionalComponent(
  __VLS_4,
  new __VLS_4({
    to: '/deployment',
    ...{ class: 'nav-link' },
  }),
)
const __VLS_6 = __VLS_5(
  {
    to: '/deployment',
    ...{ class: 'nav-link' },
  },
  ...__VLS_functionalComponentArgsRest(__VLS_5),
)
__VLS_7.slots.default
var __VLS_7
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({})
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.button,
  __VLS_intrinsicElements.button,
)({
  ...{ onClick: __VLS_ctx.callMiddleware },
})
__VLS_asFunctionalElement(__VLS_intrinsicElements.p, __VLS_intrinsicElements.p)({})
__VLS_ctx.response
__VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({})
__VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({})
__VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({})
__VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({})
__VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({})
__VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({})
__VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({})
__VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({})
__VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({})
__VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({})
__VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({})
__VLS_asFunctionalElement(__VLS_intrinsicElements.h1, __VLS_intrinsicElements.h1)({})
/** @type {__VLS_StyleScopedClasses['navbar']} */ /** @type {__VLS_StyleScopedClasses['nav-content']} */ /** @type {__VLS_StyleScopedClasses['nav-links']} */ /** @type {__VLS_StyleScopedClasses['nav-link']} */ /** @type {__VLS_StyleScopedClasses['nav-link']} */ var __VLS_dollars
let __VLS_self
