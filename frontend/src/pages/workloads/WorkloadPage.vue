<script setup lang="ts">
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import KsNavBar from '../../layout/Navbar.vue'

const followingPage = ref<string>('Workload')

const route = useRoute()
const clusterId = route.params.cluster as string
console.log("Cluster ID:", clusterId)
</script>

<template>
  <div>
    <ks-nav-bar :content="followingPage"></ks-nav-bar>
  </div>
  <div>
    <router-view></router-view>
    <h2>Workload page</h2>
  </div>
</template>

<style scoped></style>
