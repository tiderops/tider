<script lang="ts">
import { defineComponent, onMounted, reactive, toRefs } from 'vue'
import { sidebarComposable } from '../composables/SidebarComposable'
import { EnvironmentLayout } from '../types/layout.type'
import { database } from '../../wailsjs/go/models'
import CommonParameterDto = database.CommonParameterDto

interface SidebarState {
  isCollapsed: boolean
  menu: EnvironmentLayout[]
  setting: CommonParameterDto[]
  openDropdowns: { [key: number]: boolean }
}

export default defineComponent({
  name: 'ks-sidebar-v2',

  setup() {
    const { commonParameters, kubernetesParameters, environments, fetchData } = sidebarComposable()

    const state = reactive<SidebarState>({
      isCollapsed: false,
      menu: [],
      setting: [],
      openDropdowns: {},
    })

    onMounted(async () => {
      console.log('onMounted triggered sidebar')
      await fetchData()

      const env: EnvironmentLayout[] = environments.value.map((e) => {
        return {
          name: e.Name,
          description: e.Server,
          env: e.Cluster,
          status: e.Status,
          options: kubernetesParameters.value.map((k) => {
            return {
              name: k.Name,
              link: k.Link,
              icon: k.Icon,
            }
          }),
        }
      })

      console.log('commonParameters.value', commonParameters.value)
      const set: CommonParameterDto[] = commonParameters.value.map((c) => {
        return {
          Name: c.Name,
          Link: c.Link,
          Icon: c.Icon,
        }
      })

      state.menu = env
      state.setting = set
    })

    const toggleSidebar = (): void => {
      state.isCollapsed = !state.isCollapsed
    }

    const toggleDropdown = (index: number): void => {
      state.openDropdowns[index] = !state.openDropdowns[index]
    }

    const isDropdownOpen = (index: number): boolean => {
      return !!state.openDropdowns[index]
    }

    return {
      ...toRefs(state),
      toggleSidebar,
      toggleDropdown,
      isDropdownOpen,
    }
  },
})
</script>

<template>
  <div :class="['sidebar', { 'sidebar--collapsed': isCollapsed }]">
    <div class="sidebar-header">
      <button class="toggle-btn" @click="toggleSidebar">
        {{ isCollapsed ? '→' : 'Clusters ←' }}
      </button>
    </div>
    <ul class="nav-list">
      <li v-for="(item, index) in menu" :key="index" class="nav-item">
        <div
          class="env-header"
          @click="toggleDropdown(index)"
          :class="{ active: isDropdownOpen(index) }"
        >
          <div class="env-info">
            <span class="icon">📊</span>
            <span class="link-text" :class="{ hidden: isCollapsed }">
              {{ item.name }} - {{ item.env }}
              <span class="status-indicator" :class="{ 'status-active': item.status }"></span>
            </span>
          </div>
          <span
            v-if="!isCollapsed"
            class="dropdown-arrow"
            :class="{ rotated: isDropdownOpen(index) }"
          >
            ▼
          </span>
        </div>

        <ul
          class="submenu"
          :class="{
            'submenu-open': isDropdownOpen(index),
            'submenu-collapsed': isCollapsed,
          }"
        >
          <li v-for="(option, index) in item.options" :key="index">
            <router-link
              :to="{ name: option.link, params: { cluster: item.name } }"
              class="nav-link"
            >
              <span class="icon">{{ option.icon }}</span>
              <span class="link-text" :class="{ hidden: isCollapsed }">
                {{ option.name }}
              </span>
            </router-link>
          </li>
        </ul>
      </li>
    </ul>
    <div class="sidebar-footer">
      <ul>
        <li v-for="(item, index) in setting" :key="index" class="nav-item">
          <router-link :to="{ name: item.Link }" class="nav-link">
            <span class="icon">{{ item.Icon }}</span>
            <span class="link-text" :class="{ hidden: isCollapsed }">{{ item.Name }}</span>
          </router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
.sidebar {
  width: 250px;
  background-color: #2c3e50;
  color: #ecf0f1;
  transition: all 0.3s ease;
  display: flex;
  flex-direction: column;
  position: relative;
  height: 100vh;
}

.sidebar--collapsed {
  width: 60px;
}

.sidebar-header {
  padding: 15px;
  display: flex;
  justify-content: flex-end;
  border-bottom: 1px solid #34495e;
}

.sidebar-footer {
  margin-bottom: 12px;
  display: flex;
  flex-direction: column;
  border-bottom: 1px solid #34495e;
}

.toggle-btn {
  background: none;
  border: none;
  color: #ecf0f1;
  cursor: pointer;
  padding: 5px 10px;
  font-size: 1.2em;
  border-radius: 4px;
  width: 100%;
  text-align: left;
}

.toggle-btn:hover {
  background-color: #34495e;
}

.nav-list {
  list-style: none;
  padding: 0;
  margin: 0;
  overflow-y: auto;
  flex-grow: 1;
}

.nav-item {
  border-bottom: 1px solid #34495e;
}

.env-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 15px;
  cursor: pointer;
  transition: background-color 0.2s;
  user-select: none;
}

.env-header:hover {
  background-color: #34495e;
}

.env-header.active {
  background-color: #34495e;
}

.env-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.dropdown-arrow {
  font-size: 0.8em;
  transition: transform 0.3s ease;
}

.dropdown-arrow.rotated {
  transform: rotate(180deg);
}

.submenu {
  height: 0;
  opacity: 0;
  visibility: hidden;
  transition: all 0.3s ease;
  background-color: #243342;
}

.submenu-open {
  height: auto;
  opacity: 1;
  visibility: visible;
  padding: 8px 0;
}

.submenu-collapsed {
  padding-left: 0;
}

.nav-link {
  color: #ecf0f1;
  text-decoration: none;
  padding: 10px 15px 10px 45px;
  display: flex;
  align-items: center;
  transition: all 0.2s ease;
  position: relative;
}

.nav-link:hover {
  background-color: #34495e;
}

.nav-link:active {
  background-color: #2c3e50;
}

.nav-link .icon {
  width: 24px;
  text-align: center;
  margin-right: 10px;
  position: absolute;
  left: 15px;
}

.link-text {
  transition: opacity 0.3s ease;
  white-space: nowrap;
}

.hidden {
  opacity: 0;
  width: 0;
  display: none;
}

.status-indicator {
  display: inline-block;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  margin-left: 8px;
  background-color: #dc3545;
}

.status-indicator.status-active {
  background-color: #28a745;
}

/* Scrollbar Styles */
.nav-list::-webkit-scrollbar {
  width: 5px;
}

.nav-list::-webkit-scrollbar-track {
  background: #2c3e50;
}

.nav-list::-webkit-scrollbar-thumb {
  background: #34495e;
  border-radius: 3px;
}

/* Prevent text selection */
.sidebar {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.sidebar::-webkit-scrollbar-thumb:hover {
  background: #1abc9c;
}
</style>
