<script setup lang="ts">
const props = defineProps<{
	search?: string
	filterNamespace?: string
	filterStatus?: string
	namespaces?: string[]
	statuses?: string[]
	namespaceFilterEnable: boolean
}>()

const emit = defineEmits<{
	(e: 'update:search', value: string): void
	(e: 'update:filterNamespace', value: string): void
	(e: 'update:filterStatus', value: string): void
}>()
</script>

<template>
	<v-toolbar flat>
		<v-toolbar-title>Filters</v-toolbar-title>
		<v-spacer></v-spacer>
		<v-text-field :model-value="props.search" label="Search By Name" clearable class="mx-4" @update:model-value="emit('update:search', $event)"></v-text-field>
		<v-select
			v-show="props.namespaceFilterEnable"
			:model-value="props.filterNamespace"
			:items="props.namespaces"
			label="Namespace"
			clearable
			class="mx-5"
			@update:model-value="emit('update:filterNamespace', $event)"
		></v-select>
		<v-select
			:model-value="props.filterStatus"
			:items="props.statuses"
			label="Status"
			clearable
			class="mx-5"
			@update:model-value="emit('update:filterStatus', $event)"
		></v-select>
        <v-btn><v-icon icon="mdi-plus"/></v-btn>
	</v-toolbar>
</template>

<style scoped></style>
