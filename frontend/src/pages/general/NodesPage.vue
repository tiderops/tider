<script lang="ts">
import { defineComponent } from 'vue'
import KsGridBody from '../../layout/GridBody.vue'
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

<!--  <ks-grid-body :namespace="namespace" :k8sObject="k8sObject"></ks-grid-body>-->
</template>

<style scoped></style>
