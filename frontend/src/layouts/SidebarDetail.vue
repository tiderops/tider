<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'

const props = defineProps<{
	isVisible: boolean
	selectedRow: {
		name?: string
		namespace?: string
		replicas?: number
		cpu?: string
		memory?: string
		age?: string
		status?: string
	}
}>()

const emit = defineEmits<{
	(e: 'close'): void
}>()
</script>

<template>
	<v-layout>
		<v-navigation-drawer v-model="props.isVisible" location="right" :width="750" temporary class="custom-sidebar">
			<v-card flat>
				<template v-slot:prepend>
					<v-btn icon="mdi-chevron-left" variant="text" @click="emit('close')"></v-btn>
					<v-card-title style="color: black">Object Details</v-card-title>
				</template>
				<v-divider></v-divider>
				<v-card-text v-if="props.selectedRow">
					<p>Name: {{ props.selectedRow.name }}</p>
					<p>Namespace: {{ props.selectedRow.namespace }}</p>
					<p>Replicas: {{ props.selectedRow.replicas }}</p>
					<p>CPU: {{ props.selectedRow.cpu }}</p>
					<p>Memory: {{ props.selectedRow.memory }}</p>
					<p>Age: {{ props.selectedRow.age }}</p>
					<p>Status: {{ props.selectedRow.status }}</p>
				</v-card-text>
			</v-card>
		</v-navigation-drawer>
	</v-layout>
</template>

<style scoped>
.custom-sidebar {
	position: fixed; /* Ensure it's above all other content */
	top: 0;
	right: 0;
	height: 100%;
	z-index: 1100; /* Higher than navbar */
}
</style>
