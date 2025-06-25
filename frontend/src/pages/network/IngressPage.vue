<script setup lang="ts">
import {ref} from 'vue'
import {useRoute} from "vue-router";
import KsGridBodyV2 from "@/layouts/GridBody2.vue";

const k8sObject = ref<string>('ingress')
const namespace = ref<string>('')

const route = useRoute()
const clusterId = route.params.cluster as string
</script>

<template>
    <ks-grid-body-v2 :cluster="clusterId" :namespace="namespace" :k8sObject="k8sObject"></ks-grid-body-v2>
</template>

<style scoped></style>
