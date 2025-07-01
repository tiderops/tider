<script setup lang="ts">
import { useGridButton } from '@/composables/useGridButton'

const props = defineProps<{
	cluster?: string
	headers: any[]
	items: any[]
	search?: string
	sortBy?: any[]
	k8sObject?: string
}>()

const emit = defineEmits<{
	(e: 'rowClick', item: any): void
	(e: 'edit', item: any): void
	(e: 'delete', item: any): void
}>()

const { restart, loading, error, success } = useGridButton()

const editPod = (item: any) => {
	console.log('EDIT', item)

	emit('edit', item)
}

const deletePod = async (item: any) => {
	console.log('DELETE', item)
	console.log('DELETE - NAME', item.name)
	console.log('DELETE - NS', item.namespace)
	console.log('DELETE - CLUSTER ID', props.cluster)

	await restart(item.name, item.namespace, props.cluster, props.k8sObject)
	emit('delete', item)
}
</script>

<template>
	<v-data-table-virtual
		:headers="props.headers"
		:items="props.items"
		:search="props.search"
		:sort-by="props.sortBy"
		height="720"
		item-value="name"
		density="compact"
		fixed-header
		hover
	>
		<template v-slot:item.status="{ item }">
			<v-chip :color="item.status === 'Running' ? 'green' : 'red'" dark>
				{{ item.status }}
			</v-chip>
		</template>

		<template v-slot:item.actions="{ item }">
			<v-btn @click.stop="editPod(item)" icon>
				<v-icon icon="mdi-pencil" />
			</v-btn>

			<v-btn :loading="loading" @click.stop="deletePod(item)" icon>
				<v-icon icon="mdi-delete" />
			</v-btn>

			<!--            <v-alert v-if="success" type="success"></v-alert>-->
			<!--            <v-alert v-if="error"   type="error"> {{ error?.message }} </v-alert>-->
		</template>
	</v-data-table-virtual>
</template>

<style scoped></style>
