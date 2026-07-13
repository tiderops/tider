<script setup lang="ts">
import { useRoute } from 'vue-router'
import KsSidebarV2 from '@/layouts/SidebarV2.vue'
import AppShell from '@/shell/AppShell.vue'

const route = useRoute()
</script>

<template>
    <!-- kx-shell screens (Home, Clusters, Explorer, …) render in the shared app shell. -->
    <AppShell v-if="route.meta.shell === 'kx'">
        <router-view></router-view>
    </AppShell>

    <!-- Legacy Vuetify pages keep the old sidebar until they are migrated. -->
    <div v-else class="app-container">
        <ks-sidebar-v2></ks-sidebar-v2>
        <div class="content">
            <router-view></router-view>
        </div>
    </div>
</template>

<style>
.app-container {
    display: flex;
    min-height: 100vh;
}

.content {
    width: 100%;
}
</style>
