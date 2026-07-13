<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { defaultRestoreScope } from '../backup.data'
import type { RestoreScopeItem, Snapshot } from '../types'

const props = defineProps<{ open: boolean; snapshot: Snapshot | null; clusters: string[] }>()

const emit = defineEmits<{
	(e: 'close'): void
	(e: 'restore', payload: { snapshot: Snapshot; target: string; policy: string; count: number }): void
}>()

type Policy = 'skip' | 'overwrite' | 'fail'

const target = ref('')
const policy = ref<Policy>('skip')
const scope = ref<RestoreScopeItem[]>([])

watch(
	() => props.open,
	(open) => {
		if (open) {
			scope.value = defaultRestoreScope.map((s) => ({ ...s }))
			target.value = props.clusters.find((c) => c !== props.snapshot?.source) ?? props.clusters[0] ?? ''
			policy.value = 'skip'
		}
	},
)

const includedCount = computed(() => scope.value.filter((s) => s.included).reduce((sum, s) => sum + s.resources, 0))
const crossCluster = computed(() => !!props.snapshot && target.value !== props.snapshot.source)

function submit() {
	if (props.snapshot) {
		emit('restore', { snapshot: props.snapshot, target: target.value, policy: policy.value, count: includedCount.value })
	}
}
</script>

<template>
	<Transition name="drawer">
		<div v-if="open && snapshot" class="drawer-root">
			<div class="scrim" @click="emit('close')"></div>
			<aside class="drawer">
				<header class="head">
					<div class="hmain">
						<div class="hrow"><span class="chip ok">{{ snapshot.status }}</span><span class="faint">{{ snapshot.size }} · {{ snapshot.resources.toLocaleString() }} resources</span></div>
						<div class="name">Restore {{ snapshot.name }}</div>
						<div class="faint sub">captured from {{ snapshot.source }}</div>
					</div>
					<button class="icon-btn" @click="emit('close')">✕</button>
				</header>

				<div class="body">
					<div class="section-t">Target cluster</div>
					<select v-model="target" class="select">
						<option v-for="c in clusters" :key="c" :value="c">Restore to {{ c }}</option>
					</select>
					<div v-if="crossCluster" class="cross">⇄ Cross-cluster restore — clone {{ snapshot.source }} state into {{ target }}.</div>

					<div class="section-t">Scope · {{ scope.length }} namespaces</div>
					<label v-for="s in scope" :key="s.namespace" class="scope-row">
						<input v-model="s.included" type="checkbox" />
						<span class="mono">{{ s.namespace }}</span>
						<span class="faint">{{ s.resources ? s.resources + ' resources' : 'excluded by default' }}</span>
					</label>

					<div class="section-t">Conflict policy</div>
					<div class="policy">
						<button v-for="p in (['skip', 'overwrite', 'fail'] as Policy[])" :key="p" type="button" class="pill" :class="{ on: policy === p }" @click="policy = p">
							{{ p === 'skip' ? 'Skip existing' : p === 'overwrite' ? 'Overwrite' : 'Fail on conflict' }}
						</button>
					</div>

					<div class="warn">⚠ {{ includedCount.toLocaleString() }} resources will be created in {{ target }}. Run a dry run to preview the diff first.</div>

					<div class="actions">
						<button class="btn" @click="emit('restore', { snapshot, target, policy, count: includedCount })">Dry run</button>
						<button class="btn primary" @click="submit">Restore {{ includedCount.toLocaleString() }} resources</button>
					</div>
				</div>
			</aside>
		</div>
	</Transition>
</template>

<style scoped>
.drawer-root {
	position: fixed;
	inset: 0;
	z-index: 50;
}
.scrim {
	position: absolute;
	inset: 0;
	background: var(--overlay);
	backdrop-filter: blur(1.5px);
}
.drawer {
	position: absolute;
	top: 0;
	right: 0;
	bottom: 0;
	width: 480px;
	max-width: 92vw;
	background: var(--surface);
	border-left: 1px solid var(--border);
	box-shadow: 0 24px 60px rgba(0, 0, 0, 0.55);
	display: flex;
	flex-direction: column;
}
.head {
	padding: 18px 20px;
	border-bottom: 1px solid var(--border-soft);
	display: flex;
	align-items: flex-start;
	gap: 12px;
}
.hmain {
	flex: 1;
	min-width: 0;
}
.hrow {
	display: flex;
	align-items: center;
	gap: 8px;
}
.chip {
	font-size: 11px;
	font-weight: 600;
	padding: 3px 9px;
	border-radius: 999px;
}
.chip.ok {
	color: var(--ok);
	background: var(--ok-bg);
}
.name {
	font-family: var(--mono);
	font-size: 14px;
	font-weight: 600;
	margin-top: 8px;
	word-break: break-all;
}
.faint {
	color: var(--text-faint);
	font-size: 11px;
}
.sub {
	font-size: 12px;
	margin-top: 2px;
}
.icon-btn {
	width: 32px;
	height: 32px;
	display: grid;
	place-items: center;
	border-radius: var(--r-sm);
	color: var(--text-dim);
	background: none;
	border: 1px solid transparent;
	cursor: pointer;
}
.icon-btn:hover {
	background: var(--surface-2);
	border-color: var(--border);
}
.body {
	padding: 20px;
	overflow: auto;
	flex: 1;
}
.section-t {
	font-size: 11px;
	font-weight: 600;
	letter-spacing: 0.06em;
	text-transform: uppercase;
	color: var(--text-faint);
	margin: 22px 0 12px;
}
.section-t:first-child {
	margin-top: 0;
}
.select {
	width: 100%;
	background: var(--surface-2);
	border: 1px solid var(--border);
	border-radius: var(--r-sm);
	padding: 10px 12px;
	font-size: 13px;
	color: var(--text);
	font-family: var(--sans);
	cursor: pointer;
}
.cross {
	font-size: 11.5px;
	color: var(--info);
	margin-top: 8px;
}
.scope-row {
	display: flex;
	align-items: center;
	gap: 8px;
	font-size: 12.5px;
	padding: 5px 0;
}
.scope-row input {
	accent-color: var(--brand);
}
.mono {
	font-family: var(--mono);
}
.policy {
	display: flex;
	gap: 8px;
	flex-wrap: wrap;
}
.pill {
	font-size: 12px;
	font-weight: 600;
	padding: 6px 12px;
	border-radius: 999px;
	border: 1px solid var(--border);
	background: var(--surface);
	color: var(--text);
	cursor: pointer;
}
.pill.on {
	background: linear-gradient(180deg, var(--brand), var(--brand-deep));
	border-color: transparent;
	color: #fff;
}
.warn {
	margin-top: 24px;
	padding: 13px 14px;
	border-radius: var(--r-lg);
	background: var(--warn-bg);
	border: 1px solid rgba(210, 153, 34, 0.4);
	font-size: 12.5px;
	color: var(--warn);
}
.actions {
	display: flex;
	gap: 8px;
	margin-top: 24px;
}
.btn {
	flex: 1;
	display: inline-flex;
	align-items: center;
	justify-content: center;
	font-size: 13px;
	font-weight: 600;
	padding: 10px 14px;
	border-radius: var(--r-sm);
	border: 1px solid var(--border);
	background: var(--surface-2);
	color: var(--text);
	cursor: pointer;
}
.btn:hover {
	border-color: var(--brand);
}
.btn.primary {
	flex: 1.4;
	background: linear-gradient(180deg, var(--brand), var(--brand-deep));
	border-color: transparent;
	color: #fff;
}

.drawer-enter-active .drawer,
.drawer-leave-active .drawer {
	transition: transform 0.22s ease;
}
.drawer-enter-from .drawer,
.drawer-leave-to .drawer {
	transform: translateX(100%);
}
.drawer-enter-active .scrim,
.drawer-leave-active .scrim {
	transition: opacity 0.22s ease;
}
.drawer-enter-from .scrim,
.drawer-leave-to .scrim {
	opacity: 0;
}
</style>
