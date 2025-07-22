<script setup lang="ts">
import KsGridBodyV2 from '../../layouts/GridBody2.vue'
import { useRoute } from 'vue-router'
import { ref } from 'vue'

const k8sObject = ref<string>('pod')
const namespace = ref<string>('')

const route = useRoute()
const clusterId = route.params.cluster as string
console.log('POD CLUSTER_ID:', clusterId)
</script>

<template>
	<ks-grid-body-v2 :cluster="clusterId" :namespace="namespace" :k8sObject="k8sObject" :hasNamespace="true"></ks-grid-body-v2>
</template>

<style scoped></style>
