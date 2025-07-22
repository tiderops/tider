<script lang="ts" setup>
import { defineEmits, watch, ref } from 'vue'
import { useGridUpdateButton } from '@/composables/useGridUpdateButton'

const props = defineProps<{
	isVisible: boolean
	item: Record<string, any>
}>()

const formValid = ref(false)
const submitted = ref(false)
const form = ref<Record<string, string>>({})
const { update } = useGridUpdateButton()

// Emit
const emit = defineEmits<{
	(e: 'close'): void
}>()

const submitForm = async () => {
	if (formValid.value) {
		submitted.value = true
		console.log('Form Data:', form.value)
		// await update()
	}
}

// Flatten object helper
const flattenObject = (obj: any, prefix = ''): Record<string, string> => {
	const flat: Record<string, string> = {}
	let i = 0
	for (const key in obj) {
		const value = obj[key]
		const newKey = prefix ? `${prefix}.${key}` : key
		console.log(i, value, newKey)
		i++
		if (typeof value === 'object' && value !== null && !Array.isArray(value)) {
			Object.assign(flat, flattenObject(value, newKey))
		} else {
			flat[newKey] = String(value ?? '')
		}
	}
	return flat
}

// Watch item prop and re-populate the form
watch(
	() => props.item,
	(newItem) => {
		form.value = flattenObject(newItem)
	},
	{ immediate: true },
)
</script>

<template>
	<v-layout>
		<v-navigation-drawer v-model="props.isVisible" location="right" :width="750" temporary class="custom-sidebar">
			<v-card flat>
				<template v-slot:prepend>
					<v-btn icon="mdi-chevron-left" variant="text" @click="emit('close')"></v-btn>
					<v-card-title style="color: black">Object Details</v-card-title>
				</template>
				<v-form v-model="formValid" @submit.prevent="submitForm">
					<v-text-field
						v-for="(value, key) in form"
						:key="key"
						v-model="form[key]"
						:label="key"
						:rules="[(v) => !!v || 'Required']"
						hide-details="auto"
						density="compact"
					></v-text-field>

					<v-btn type="submit" :disabled="!formValid" color="primary">Submit</v-btn>
				</v-form>

				<div v-if="submitted">
					<h3>Form Submitted!</h3>
					<pre>{{ form }}</pre>
				</div>
			</v-card>
		</v-navigation-drawer>
	</v-layout>
</template>

<style scoped>
.custom-sidebar {
	position: fixed;
	top: 0;
	right: 0;
	height: 100%;
	z-index: 1100;
}
</style>
