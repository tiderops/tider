<script setup lang="ts">
import { ref } from 'vue'
import KsNavBar from '@/layouts/NavbarMenu.vue'

const followingPage = ref<string>('Workload')
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
