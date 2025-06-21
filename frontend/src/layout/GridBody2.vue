<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { gridComposable } from '@/composables/GridComposable'
import KsSidebarDetail from './SidebarDetail.vue'
import KsGridTable from '../components/GridTableComponent.vue'
import KsGridHeader from '../components/GridHeaderComponent.vue'

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
defineProps<{
    cluster?: string
    k8sObject: string
    namespace: string
}>()

// Composables & constants
const props = defineProps()
const response = gridComposable(props.cluster, props.k8sObject)
const namespaces = ['ns-local', 'ns-dev']
const statuses = ['Alive', 'Inactive']

// Reactive state
const items = ref<GridItem[]>([])
const search = ref('')
const filterNamespace = ref('')
const filterStatus = ref('')
const sortBy = ref([{ key: 'name', order: 'asc' }])

const header = reactive<HeadState>({
    header: [],
})

// Filtering logic
const filteredItems = computed(() =>
    items.value.filter((item) => {
        const namespaceMatch =
            !filterNamespace.value || item.namespace.includes(filterNamespace.value)
        const statusMatch = !filterStatus.value || item.status.includes(filterStatus.value)
        const searchMatch =
            !search.value || item.name.toLowerCase().includes(search.value.toLowerCase())
        return namespaceMatch && statusMatch && searchMatch
    })
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

// Fetch data on mount
onMounted(async () => {
    await response.fetchData()
    header.header = response.content?.head.value ?? []
    items.value = response.content?.body.value ?? []
})
</script>

<template>
  <v-container>
    <v-card>
      <KsGridHeader
        v-model:search="search"
        v-model:filterNamespace="filterNamespace"
        v-model:filterStatus="filterStatus"
        :namespaces="namespaces"
        :statuses="statuses"
      />
      <KsGridTable
        :headers="header.header"
        :items="filteredItems"
        :search="search"
        :sortBy="sortBy"
        @click:row="onRowClick"
      />
    </v-card>
    <v-card>
      <KsSidebarDetail
        :isVisible="isSidebarVisible"
        :selectedRow="selectedRow"
        @close="isSidebarVisible = false"
      />
    </v-card>
  </v-container>
</template>
