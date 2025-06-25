<script setup lang="ts">
import { ref } from 'vue'
import KsGridBodyV2 from '@/layouts/GridBody2.vue'
import { useRoute } from 'vue-router'

const k8sObject = ref<string>('namespace')
const namespace = ref<string>('')

const route = useRoute()
const clusterId = route.params.cluster as string
console.log('PRUEBA CARGA NS')
</script>

<template>
	<ks-grid-body-v2 :cluster="clusterId" :namespace="namespace" :k8sObject="k8sObject"></ks-grid-body-v2>
</template>

<style scoped></style>
