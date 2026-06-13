<script setup lang="ts">
import KsSidebarV2 from '@/layouts/SidebarV2.vue'
</script>

<template>
	<div class="app-container">
		<ks-sidebar-v2></ks-sidebar-v2>
		<div class="content">
			<router-view></router-view>
		</div>
	</div>
</template>

<style>
.app-container {
	display: flex;
	min-height: 100vh;
}

.content {
	width: 100%;
}
</style>
