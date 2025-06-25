<script setup lang="ts">
import { ref } from 'vue'
import KsNavBar from '../../layouts/Navbar.vue'

const followingPage = ref<string>('Network')
</script>

<template>
	<div>
		<ks-nav-bar :content="followingPage"></ks-nav-bar>
	</div>
	<div>
		<router-view></router-view>
		<h2>Network page</h2>
	</div>
</template>

<style scoped></style>
