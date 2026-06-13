<script setup lang="ts">
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useClusterStore } from '@/stores/cluster.store'
import KsGridBodyV2 from '@/components/shared/ResourceGrid.vue'

const k8sObject = ref<string>('persistentVolumeClaim')
const namespace = ref<string>('')

const { currentCluster: clusterId } = storeToRefs(useClusterStore())
</script>

<template>
	<ks-grid-body-v2 :cluster="clusterId" :namespace="namespace" :k8sObject="k8sObject" :hasNamespace="true"></ks-grid-body-v2>
</template>

<style scoped></style>
