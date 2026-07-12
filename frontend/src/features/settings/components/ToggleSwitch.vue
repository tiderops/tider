<script setup lang="ts">
const model = defineModel<boolean>({ required: true })
defineProps<{ label?: string }>()
</script>

<template>
	<button
		type="button"
		role="switch"
		:aria-checked="model"
		:aria-label="label"
		class="sw"
		:class="{ on: model }"
		@click="model = !model"
	>
		<span class="knob"></span>
	</button>
</template>

<style scoped>
.sw {
	width: 40px;
	height: 23px;
	border-radius: 999px;
	border: 1px solid var(--border);
	background: var(--surface-2);
	padding: 2px;
	cursor: pointer;
	flex: none;
	transition: background 0.16s ease, border-color 0.16s ease;
}
.sw.on {
	background: var(--brand);
	border-color: transparent;
}
.knob {
	display: block;
	width: 17px;
	height: 17px;
	border-radius: 50%;
	background: #fff;
	box-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
	transition: transform 0.16s ease;
}
.sw.on .knob {
	transform: translateX(17px);
}
</style>
