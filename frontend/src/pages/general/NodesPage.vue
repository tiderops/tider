<script lang="ts">
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
</script>

<template>
  <h2>qweqwe</h2>
</template>

<style scoped></style>
