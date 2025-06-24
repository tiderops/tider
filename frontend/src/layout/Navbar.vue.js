/// <reference types="../../node_modules/.vue-global-types/vue_3.5_0_0_0.d.ts" />
import { onMounted, reactive } from 'vue'
import { navbarComposable } from '@/composables/NavbarComposable'
import { K8sObjectDto, NavbarDto } from '@/types/navbar.type'
const { objects, fetchData } = navbarComposable()
const props = defineProps()
const state = reactive({
  menu: [],
  objects: [],
})
onMounted(async () => {
  await fetchData()
  console.log('OBJECTS', objects.value)
  const dto = objects.value.map((o) => ({
    Name: o.Name,
    IsVisible: o.IsVisible,
    IsEditable: o.IsEditable,
    K8sObject: o.K8sObject.map((k) => ({
      Name: k.Name,
      Link: k.Link,
      IsEditable: k.IsEditable,
      IsVisible: k.IsVisible,
    })),
  }))
  const type = dto.find((x) => x.Name === props.content)
  state.menu = dto
  state.objects = type?.K8sObject ?? []
})
debugger /* PartiallyEnd: #3632/scriptSetup.vue */
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
for (const [item, index] of __VLS_getVForSourceType(__VLS_ctx.state.objects)) {
  __VLS_asFunctionalElement(
    __VLS_intrinsicElements.div,
    __VLS_intrinsicElements.div,
  )({
    ...{ class: 'nav-links' },
    key: index,
  })
  const __VLS_0 = {}.RouterLink
  /** @type {[typeof __VLS_components.RouterLink, typeof __VLS_components.routerLink, typeof __VLS_components.RouterLink, typeof __VLS_components.routerLink, ]} */ // @ts-ignore
  const __VLS_1 = __VLS_asFunctionalComponent(
    __VLS_0,
    new __VLS_0({
      to: { name: item.Link },
      ...{ class: 'nav-link' },
    }),
  )
  const __VLS_2 = __VLS_1(
    {
      to: { name: item.Link },
      ...{ class: 'nav-link' },
    },
    ...__VLS_functionalComponentArgsRest(__VLS_1),
  )
  __VLS_3.slots.default
  item.Name
  var __VLS_3
}
/** @type {__VLS_StyleScopedClasses['navbar']} */ /** @type {__VLS_StyleScopedClasses['nav-content']} */ /** @type {__VLS_StyleScopedClasses['nav-links']} */ /** @type {__VLS_StyleScopedClasses['nav-link']} */ var __VLS_dollars
const __VLS_self = (await import('vue')).defineComponent({
  setup() {
    return {
      state: state,
    }
  },
  __typeProps: {},
})
export default (await import('vue')).defineComponent({
  setup() {
    return {}
  },
  __typeProps: {},
}) /* PartiallyEnd: #4569/main.vue */
