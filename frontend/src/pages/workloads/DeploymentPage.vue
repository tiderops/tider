<script setup lang="ts">
import KsGridBodyV2 from '../../layout/GridBody2.vue'
import {useRoute} from "vue-router";
import { ref } from 'vue'

const k8sObject = ref<string>('deployment')
const namespace = ref<string>('east')

const route = useRoute()
const clusterId = route.params.cluster as string
console.log("DEPLOYMENT CLUSTER_ID:", clusterId)

</script>

<template>
  <ks-grid-body-v2 :cluster = "clusterId" :namespace="namespace" :k8sObject="k8sObject"></ks-grid-body-v2>
</template>

<style scoped></style>

