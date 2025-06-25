<script lang="ts">
import { defineComponent } from 'vue'
import KsGridBodyV2 from '../../layout/GridBody2.vue'
import KsNavBar from '../../layout/Navbar.vue'

export default defineComponent({
	name: 'NamespacePage',
	components: { KsNavBar, KsGridBody, KsGridBodyV2 },
	data() {
		return {
			k8sObject: 'namespace',
			namespace: 'mock',
		}
	},
})
console.log('PRUEBA CARGA NS')
</script>

<template>
	<ks-grid-body-v2 :namespace="namespace" :k8sObject="k8sObject"></ks-grid-body-v2>
</template>

<style scoped></style>
