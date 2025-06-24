import { onMounted, reactive, ref } from 'vue'
import { sidebarComposable } from '@/composables/SidebarComposable'
import { ClusterLayout } from '@/types/layout.type'
import { database } from '../../wailsjs/go/models'
var CommonParameterDto = database.CommonParameterDto
const { commonParameters, kubernetesParameters, clusters, fetchData } = sidebarComposable()
const isCollapsed = ref(false)
const menu = ref([])
const setting = ref([])
const openDropdowns = reactive({})
const toggleSidebar = () => {
  isCollapsed.value = !isCollapsed.value
}
const toggleDropdown = (index) => {
  openDropdowns[index] = !openDropdowns[index]
}
const isDropdownOpen = (index) => {
  return !!openDropdowns[index]
}
onMounted(async () => {
  console.log('onMounted triggered sidebar')
  await fetchData()
  menu.value = clusters.value.map((e) => ({
    name: e.Name,
    cluster: e.Cluster,
    status: e.Status,
    options: kubernetesParameters.value.map((k) => ({
      name: k.Name,
      link: k.Link,
      icon: k.Icon,
    })),
  }))
  setting.value = commonParameters.value.map((c) => ({
    Name: c.Name,
    Link: c.Link,
    Icon: c.Icon,
  }))
})
debugger /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {}
let __VLS_components
let __VLS_directives
/** @type {__VLS_StyleScopedClasses['toggle-btn']} */ /** @type {__VLS_StyleScopedClasses['env-header']} */ /** @type {__VLS_StyleScopedClasses['env-header']} */ /** @type {__VLS_StyleScopedClasses['dropdown-arrow']} */ /** @type {__VLS_StyleScopedClasses['nav-link']} */ /** @type {__VLS_StyleScopedClasses['nav-link']} */ /** @type {__VLS_StyleScopedClasses['nav-link']} */ /** @type {__VLS_StyleScopedClasses['status-indicator']} */ /** @type {__VLS_StyleScopedClasses['nav-list']} */ /** @type {__VLS_StyleScopedClasses['nav-list']} */ /** @type {__VLS_StyleScopedClasses['nav-list']} */ /** @type {__VLS_StyleScopedClasses['sidebar']} */ /** @type {__VLS_StyleScopedClasses['sidebar']} */ // CSS variable injection
// CSS variable injection end
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.div,
  __VLS_intrinsicElements.div,
)({
  ...{ class: ['sidebar', { 'sidebar--collapsed': __VLS_ctx.isCollapsed }] },
})
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.div,
  __VLS_intrinsicElements.div,
)({
  ...{ class: 'sidebar-header' },
})
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.button,
  __VLS_intrinsicElements.button,
)({
  ...{ onClick: __VLS_ctx.toggleSidebar },
  ...{ class: 'toggle-btn' },
})
__VLS_ctx.isCollapsed ? '→' : 'Clusters ←'
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.ul,
  __VLS_intrinsicElements.ul,
)({
  ...{ class: 'nav-list' },
})
for (const [item, index] of __VLS_getVForSourceType(__VLS_ctx.menu)) {
  __VLS_asFunctionalElement(
    __VLS_intrinsicElements.li,
    __VLS_intrinsicElements.li,
  )({
    key: index,
    ...{ class: 'nav-item' },
  })
  __VLS_asFunctionalElement(
    __VLS_intrinsicElements.div,
    __VLS_intrinsicElements.div,
  )({
    ...{
      onClick: (...[$event]) => {
        __VLS_ctx.toggleDropdown(index)
      },
    },
    ...{ class: 'env-header' },
    ...{ class: { active: __VLS_ctx.isDropdownOpen(index) } },
  })
  __VLS_asFunctionalElement(
    __VLS_intrinsicElements.div,
    __VLS_intrinsicElements.div,
  )({
    ...{ class: 'env-info' },
  })
  __VLS_asFunctionalElement(
    __VLS_intrinsicElements.span,
    __VLS_intrinsicElements.span,
  )({
    ...{ class: 'icon' },
  })
  __VLS_asFunctionalElement(
    __VLS_intrinsicElements.span,
    __VLS_intrinsicElements.span,
  )({
    ...{ class: 'link-text' },
    ...{ class: { hidden: __VLS_ctx.isCollapsed } },
  })
  item.cluster
  __VLS_asFunctionalElement(
    __VLS_intrinsicElements.span,
    __VLS_intrinsicElements.span,
  )({
    ...{ class: 'status-indicator' },
    ...{ class: { 'status-active': item.status } },
  })
  if (!__VLS_ctx.isCollapsed) {
    __VLS_asFunctionalElement(
      __VLS_intrinsicElements.span,
      __VLS_intrinsicElements.span,
    )({
      ...{ class: 'dropdown-arrow' },
      ...{ class: { rotated: __VLS_ctx.isDropdownOpen(index) } },
    })
  }
  __VLS_asFunctionalElement(
    __VLS_intrinsicElements.ul,
    __VLS_intrinsicElements.ul,
  )({
    ...{ class: 'submenu' },
    ...{
      class: {
        'submenu-open': __VLS_ctx.isDropdownOpen(index),
        'submenu-collapsed': __VLS_ctx.isCollapsed,
      },
    },
  })
  for (const [option, index] of __VLS_getVForSourceType(item.options)) {
    __VLS_asFunctionalElement(
      __VLS_intrinsicElements.li,
      __VLS_intrinsicElements.li,
    )({
      key: index,
    })
    const __VLS_0 = {}.RouterLink
    /** @type {[typeof __VLS_components.RouterLink, typeof __VLS_components.routerLink, typeof __VLS_components.RouterLink, typeof __VLS_components.routerLink, ]} */ // @ts-ignore
    const __VLS_1 = __VLS_asFunctionalComponent(
      __VLS_0,
      new __VLS_0({
        to: { name: option.link, params: { cluster: item.cluster } },
        ...{ class: 'nav-link' },
      }),
    )
    const __VLS_2 = __VLS_1(
      {
        to: { name: option.link, params: { cluster: item.cluster } },
        ...{ class: 'nav-link' },
      },
      ...__VLS_functionalComponentArgsRest(__VLS_1),
    )
    __VLS_3.slots.default
    __VLS_asFunctionalElement(
      __VLS_intrinsicElements.span,
      __VLS_intrinsicElements.span,
    )({
      ...{ class: 'icon' },
    })
    option.icon
    __VLS_asFunctionalElement(
      __VLS_intrinsicElements.span,
      __VLS_intrinsicElements.span,
    )({
      ...{ class: 'link-text' },
      ...{ class: { hidden: __VLS_ctx.isCollapsed } },
    })
    option.name
    var __VLS_3
  }
}
__VLS_asFunctionalElement(
  __VLS_intrinsicElements.div,
  __VLS_intrinsicElements.div,
)({
  ...{ class: 'sidebar-footer' },
})
__VLS_asFunctionalElement(__VLS_intrinsicElements.ul, __VLS_intrinsicElements.ul)({})
for (const [item, index] of __VLS_getVForSourceType(__VLS_ctx.setting)) {
  __VLS_asFunctionalElement(
    __VLS_intrinsicElements.li,
    __VLS_intrinsicElements.li,
  )({
    key: index,
    ...{ class: 'nav-item' },
  })
  const __VLS_4 = {}.RouterLink
  /** @type {[typeof __VLS_components.RouterLink, typeof __VLS_components.routerLink, typeof __VLS_components.RouterLink, typeof __VLS_components.routerLink, ]} */ // @ts-ignore
  const __VLS_5 = __VLS_asFunctionalComponent(
    __VLS_4,
    new __VLS_4({
      to: { name: item.Link },
      ...{ class: 'nav-link' },
    }),
  )
  const __VLS_6 = __VLS_5(
    {
      to: { name: item.Link },
      ...{ class: 'nav-link' },
    },
    ...__VLS_functionalComponentArgsRest(__VLS_5),
  )
  __VLS_7.slots.default
  __VLS_asFunctionalElement(
    __VLS_intrinsicElements.span,
    __VLS_intrinsicElements.span,
  )({
    ...{ class: 'icon' },
  })
  item.Icon
  __VLS_asFunctionalElement(
    __VLS_intrinsicElements.span,
    __VLS_intrinsicElements.span,
  )({
    ...{ class: 'link-text' },
    ...{ class: { hidden: __VLS_ctx.isCollapsed } },
  })
  item.Name
  var __VLS_7
}
/** @type {__VLS_StyleScopedClasses['sidebar']} */ /** @type {__VLS_StyleScopedClasses['sidebar--collapsed']} */ /** @type {__VLS_StyleScopedClasses['sidebar-header']} */ /** @type {__VLS_StyleScopedClasses['toggle-btn']} */ /** @type {__VLS_StyleScopedClasses['nav-list']} */ /** @type {__VLS_StyleScopedClasses['nav-item']} */ /** @type {__VLS_StyleScopedClasses['env-header']} */ /** @type {__VLS_StyleScopedClasses['active']} */ /** @type {__VLS_StyleScopedClasses['env-info']} */ /** @type {__VLS_StyleScopedClasses['icon']} */ /** @type {__VLS_StyleScopedClasses['link-text']} */ /** @type {__VLS_StyleScopedClasses['hidden']} */ /** @type {__VLS_StyleScopedClasses['status-indicator']} */ /** @type {__VLS_StyleScopedClasses['status-active']} */ /** @type {__VLS_StyleScopedClasses['dropdown-arrow']} */ /** @type {__VLS_StyleScopedClasses['rotated']} */ /** @type {__VLS_StyleScopedClasses['submenu']} */ /** @type {__VLS_StyleScopedClasses['submenu-open']} */ /** @type {__VLS_StyleScopedClasses['submenu-collapsed']} */ /** @type {__VLS_StyleScopedClasses['nav-link']} */ /** @type {__VLS_StyleScopedClasses['icon']} */ /** @type {__VLS_StyleScopedClasses['link-text']} */ /** @type {__VLS_StyleScopedClasses['hidden']} */ /** @type {__VLS_StyleScopedClasses['sidebar-footer']} */ /** @type {__VLS_StyleScopedClasses['nav-item']} */ /** @type {__VLS_StyleScopedClasses['nav-link']} */ /** @type {__VLS_StyleScopedClasses['icon']} */ /** @type {__VLS_StyleScopedClasses['link-text']} */ /** @type {__VLS_StyleScopedClasses['hidden']} */ var __VLS_dollars
const __VLS_self = (await import('vue')).defineComponent({
  setup() {
    return {
      isCollapsed: isCollapsed,
      menu: menu,
      setting: setting,
      toggleSidebar: toggleSidebar,
      toggleDropdown: toggleDropdown,
      isDropdownOpen: isDropdownOpen,
    }
  },
})
export default (await import('vue')).defineComponent({
  setup() {
    return {}
  },
}) /* PartiallyEnd: #4569/main.vue */
