<script setup lang="ts">
import { fetchAbout } from '../settings.data'
import SettingsCard from './SettingsCard.vue'

const about = fetchAbout()
</script>

<template>
	<SettingsCard title="About Tider" subtitle="Open-source multi-cluster Kubernetes management.">
		<div class="grid">
			<div class="kv"><span>Version</span><b class="mono">{{ about.version }}</b></div>
			<div class="kv"><span>Runtime</span><b>{{ about.build }}</b></div>
			<div class="kv"><span>License</span><b>{{ about.license }}</b></div>
			<div class="kv">
				<span>Backend</span>
				<b :class="about.wailsRuntime ? 'ok' : 'idle'">{{ about.wailsRuntime ? 'connected' : 'not connected (mock data)' }}</b>
			</div>
		</div>

		<div class="links">
			<a class="link" :href="about.repo" target="_blank" rel="noopener">★ GitHub repository</a>
			<a class="link" :href="about.docs" target="_blank" rel="noopener">📖 Documentation</a>
		</div>

		<p class="note">Version and build metadata are placeholders until wired to the Wails build pipeline.</p>
	</SettingsCard>
</template>

<style scoped>
.grid {
	display: grid;
	grid-template-columns: repeat(2, 1fr);
	gap: 1px;
	background: var(--border-soft);
	border: 1px solid var(--border-soft);
	border-radius: var(--r-md);
	overflow: hidden;
	margin: 10px 0 18px;
}
.kv {
	display: flex;
	flex-direction: column;
	gap: 4px;
	padding: 13px 15px;
	background: var(--surface);
}
.kv span {
	font-size: 11px;
	color: var(--text-faint);
	text-transform: uppercase;
	letter-spacing: 0.05em;
}
.kv b {
	font-size: 13.5px;
	font-weight: 600;
	color: var(--text);
}
.kv b.ok {
	color: var(--ok);
}
.kv b.idle {
	color: var(--text-dim);
}
.mono {
	font-family: var(--mono);
}
.links {
	display: flex;
	gap: 10px;
	flex-wrap: wrap;
}
.link {
	font-size: 12.5px;
	font-weight: 600;
	padding: 8px 14px;
	border-radius: var(--r-sm);
	border: 1px solid var(--border);
	background: var(--surface-2);
	color: var(--text);
	text-decoration: none;
}
.link:hover {
	border-color: var(--brand);
	color: var(--brand);
}
.note {
	margin: 16px 0 2px;
	font-size: 11.5px;
	color: var(--text-faint);
}
</style>
