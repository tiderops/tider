<script lang="ts">
import { defineComponent } from 'vue'
import ksSidebarV2 from './layout/SidebarV2.vue'

export default defineComponent({
	name: 'App',
	components: {
		ksSidebarV2,
	},
})
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
