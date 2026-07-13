<script setup lang="ts">
import { onMounted, reactive } from 'vue'
import { useNavbarParams } from '@/composables/useNavbarParams'
import type { K8sObjectDto, NavbarDto } from '@/types/navbar.type'

interface NavbarState {
	menu: NavbarDto[]
	objects: K8sObjectDto[]
}

const { objects, fetchData } = useNavbarParams()

const props = defineProps<{
	content: string
}>()

const state = reactive<NavbarState>({
	menu: [],
	objects: [],
})

onMounted(async () => {
	await fetchData()

	const dto: NavbarDto[] = objects.value.map((o) => ({
		Name: o.Name,
		IsVisible: o.IsVisible,
		IsEditable: o.IsEditable,
		K8sObject: o.K8sObject.map((k) => ({
			Name: k.Name,
			Link: k.Link,
			IsEditable: k.IsEditable,
			IsVisible: k.IsVisible,
		})),
	}))

	const type = dto.find((x) => x.Name === props.content)
	state.menu = dto
	state.objects = type?.K8sObject ?? []
})
</script>

<template>
	<nav class="navbar">
		<div class="nav-content">
			<div class="nav-links" v-for="(item, index) in state.objects" :key="index">
				<router-link :to="{ name: item.Link }" class="nav-link">{{ item.Name }}</router-link>
			</div>
		</div>
	</nav>
</template>

<style scoped>
.navbar {
	width: fit-content;
	position: sticky;
	display: flex;
	align-items: center;
	top: 0;
	left: 0;
	right: 0;
	background-color: rgba(255, 255, 255, 0.95);
	box-shadow: 0 2px 10px rgba(0, 0, 0, 0.15);
	transition: all 0.3s ease;
	z-index: 1000;
}

.navbar-scrolled {
	background-color: rgba(255, 255, 255, 0.98);
	box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.nav-content {
	width: 100%;
	padding: 0 20px;
	height: 100%;
	display: flex;
	justify-content: flex-start;
	align-items: center;
}
.nav-links {
	display: flex;
	gap: 1rem;
	padding: 5px 20px 5px 20px;
}

.nav-link {
	color: #333;
	text-decoration: none;
	font-weight: 500;
	padding: 0.5rem 1rem;
	transition: color 0.3s ease;
}

.nav-link:hover {
	color: #007bff;
}

.nav-link.router-link-active {
	color: #007bff;
}
</style>
