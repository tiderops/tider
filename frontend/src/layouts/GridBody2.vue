<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { useGridTable } from '@/composables/useGridTable'
import { useFilter } from '@/composables/useFilter'
import KsSidebarDetail from './SidebarDetail.vue'
import KsGridTable from '../components/GridTableComponent.vue'
import KsGridHeader from '../components/GridHeaderComponent.vue'

const snackbar = ref(false)
const text = ref('')
const timeout = ref(2000)

interface GridItem {
	name: string
	namespace: string
	replicas: number
	cpu: string
	memory: string
	age: string
	status: string
}

interface HeadState {
	header: Array<any>
}

// Props
const props = defineProps<{
	cluster?: string
	k8sObject: string
	namespace: string
}>()

// Composables & constants
const response = useGridTable(props.cluster, props.k8sObject)
const response2 = useFilter(props.cluster, props.k8sObject)
const statuses = ['Running', 'Succeeded', 'Pending']

// Reactive state
const items = ref<GridItem[]>([])
const search = ref('')
const filterNamespace = ref('')
const filterStatus = ref('')
const sortBy = ref([{ key: 'name', order: 'asc' }])
const ns = ref<Array<any>>([])

const header = reactive<HeadState>({
	header: [],
})

// Filtering logic
const filteredItems = computed(() =>
	items.value.filter((item) => {
		const namespaceMatch = !filterNamespace.value || item.namespace.includes(filterNamespace.value)
		const statusMatch = !filterStatus.value || item.status.includes(filterStatus.value)
		const searchMatch = !search.value || item.name.toLowerCase().includes(search.value.toLowerCase())
		return namespaceMatch && statusMatch && searchMatch
	}),
)

// Sidebar logic
const isSidebarVisible = ref(false)
const selectedRow = ref<any>(null)

const onRowClick = (cellData: any, item: any) => {
	selectedRow.value = item.item

	if (isSidebarVisible.value) {
		isSidebarVisible.value = false
		setTimeout(() => {
			isSidebarVisible.value = true
		}, 100)
	} else {
		isSidebarVisible.value = true
	}
}

const onEditItem = async (item: any) => {
	console.log('Parent received edit:', item)
	text.value = `Resource "${item.name}" was edited.`
	snackbar.value = true

	await response.fetchData()
	items.value = response.content?.body.value ?? []
}

const onDeleteItem = async (item: any) => {
	console.log('Parent received delete:', item)
	text.value = `Resource "${item.name}" was deleted.`
	snackbar.value = true

	await response.fetchData()
	items.value = response.content?.body.value ?? []
}

// Fetch data on mount
onMounted(async () => {
	await response.fetchData()
	await response2.fetchData()

	header.header = response.content?.head.value ?? []
	items.value = response.content?.body.value ?? []

	ns.value = response2.content?.body.value ?? []
})
</script>

<template>
	<v-container>
		<v-card>
			<KsGridHeader
				v-model:search="search"
				v-model:filterNamespace="filterNamespace"
				v-model:filterStatus="filterStatus"
				:namespaces="ns"
				:statuses="statuses"
				:namespaceFilterEnable="true"
			/>
			<KsGridTable
				:cluster="props.cluster"
				:headers="header.header"
				:items="filteredItems"
				:search="search"
				:sortBy="sortBy"
				@click:row="onRowClick"
				@delete="onDeleteItem"
				@edit="onEditItem"
				:k8sObject="props.k8sObject"
			/>
		</v-card>
		<v-snackbar v-model="snackbar" :timeout="timeout">
			{{ text }}
			<template v-slot:actions>
				<v-btn color="blue" variant="text" @click="snackbar = false">Close</v-btn>
			</template>
		</v-snackbar>
		<v-card>
			<KsSidebarDetail :isVisible="isSidebarVisible" :selectedRow="selectedRow" @close="isSidebarVisible = false" />
		</v-card>
	</v-container>
</template>
