<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import KsNavBar from '../../layout/Navbar.vue'

// TODO: Revisar el funciona wails events o usar SSE
export default defineComponent({
	name: 'NetworkPage',
	components: { KsNavBar },
	data() {
		return {
			followingPage: 'Network',
		}
	},
	// setup() {
	//     const messages = ref<string[]>([]);
	//
	//     // onMounted(() => {
	//     //     // Listen for the "new-message" event emitted from the backend
	//     //     window.runtime.EventsOn("new-message", (message: string) => {
	//     //         messages.value.push(message);
	//     //         console.log("Received:", message);
	//     //     });
	//     //
	//     //     // Optionally call the backend to start emitting messages
	//     //     window.runtime.Call("EmitMessages").catch((err: any) => {
	//     //         console.error("Error calling EmitMessages:", err);
	//     //     });
	//     // });
	//
	//     onMounted(() => {
	//         window.runtime.EventsOn("new-message", (message: string) => {
	//             console.log("Received message:", message);
	//         });
	//
	//         window.runtime.Call("EmitMessages").catch((err: any) => {
	//             console.error("Error:", err);
	//         });
	//     });
	//
	//     return {
	//         messages,
	//     };
	// },
})
</script>

<template>
	<div>
		<ks-nav-bar :content="followingPage"></ks-nav-bar>
	</div>
	<div>
		<router-view></router-view>
		<h2>Network page</h2>
	</div>
	<!--    <div>-->
	<!--        <h1>Messages</h1>-->
	<!--        <ul>-->
	<!--            <li v-for="(message, index) in messages" :key="index">-->
	<!--                {{ message }}-->
	<!--            </li>-->
	<!--        </ul>-->
	<!--    </div>-->
</template>

<style scoped>
h1 {
	color: #42b983;
}
</style>
