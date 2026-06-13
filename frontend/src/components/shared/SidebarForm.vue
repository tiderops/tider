<script lang="ts" setup>
import { computed, watch, ref } from 'vue'
import type { ResourceDescriptor, ResourceDetail, UpdatePayload } from '@/resources/types'
import { toAppError } from '@/services/apperror'
import type { AppError } from '@/services/apperror'
import { model } from '../../../wailsjs/go/models'

const props = defineProps<{
	isVisible: boolean
	item?: ResourceDetail
	cluster: string
	descriptor: ResourceDescriptor
}>()

const formValid = ref(false)
const submitted = ref(false)
const submitError = ref<AppError | null>(null)
const form = ref<Record<string, string>>({})
const data = ref<Record<string, string>>({})

const emit = defineEmits<{
	(e: 'close'): void
}>()

const visible = computed({
	get: () => props.isVisible,
	set: (open: boolean) => {
		if (!open) emit('close')
	},
})

const submitForm = async () => {
	if (!formValid.value || !props.descriptor.update) return

	const request = model.PodUpdate.createFrom({
		App: form.value.name,
		Container: {
			Image: form.value.image,
			PullPolicy: form.value.pullPolicy,
			Port: form.value.port,
		},
	}) as UpdatePayload

	submitError.value = null
	try {
		await props.descriptor.update(
			{ cluster: props.cluster, namespace: form.value.namespace ?? '', name: form.value.name },
			request,
		)
		submitted.value = true
	} catch (err) {
		submitError.value = toAppError(err)
	}
}

interface FlattenedItem {
	edit: Record<string, string>
	flat: Record<string, string>
}

const flattenItem = (item: ResourceDetail): FlattenedItem => {
	const editableKeys = 'editable' in item ? (item.editable ?? []) : []
	const edit: Record<string, string> = {}
	const flat: Record<string, string> = {}

	const walk = (obj: Record<string, unknown>, prefix: string) => {
		for (const key in obj) {
			if (key === 'editable') continue

			const value = obj[key]
			const newKey = prefix ? `${prefix}.${key}` : key

			if (typeof value === 'object' && value !== null && !Array.isArray(value)) {
				walk(value as Record<string, unknown>, newKey)
			} else if (editableKeys.includes(newKey)) {
				edit[newKey] = String(value ?? '')
			} else {
				flat[newKey] = String(value ?? '')
			}
		}
	}

	walk(item as unknown as Record<string, unknown>, '')
	return { edit, flat }
}

watch(
	() => props.item,
	(newItem) => {
		if (!newItem) {
			form.value = {}
			data.value = {}
			return
		}
		const { edit, flat } = flattenItem(newItem)
		form.value = edit
		data.value = flat
	},
	{ immediate: true },
)
</script>

<template>
	<v-layout>
		<v-navigation-drawer v-model="visible" location="right" :width="750" temporary class="custom-sidebar">
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

					<v-text-field disabled v-for="(value, key) in data" :key="key" v-model="data[key]" :label="key" hide-details="auto" density="compact"></v-text-field>

					<v-btn type="submit" :disabled="!formValid" color="primary">Submit</v-btn>
				</v-form>

				<div v-if="submitted">
					<h3>Form Submitted!</h3>
					<pre>{{ form }}</pre>
				</div>
				<v-alert v-if="submitError" type="error" variant="tonal" class="ma-2" :text="submitError.message" />
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
