<script setup lang="ts">
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useClusterStore } from '@/stores/cluster.store'
import KsGridBodyV2 from '@/components/shared/ResourceGrid.vue'

const k8sObject = ref<string>('node')
const namespace = ref<string>('')

const { currentCluster: clusterId } = storeToRefs(useClusterStore())
</script>

<template>
	<ks-grid-body-v2 :cluster="clusterId" :namespace="namespace" :k8sObject="k8sObject"></ks-grid-body-v2>
</template>

<style scoped></style>
