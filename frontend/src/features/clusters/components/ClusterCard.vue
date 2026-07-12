<script setup lang="ts">
import { computed } from 'vue'
import StatusChip from '@/components/shared/StatusChip.vue'
import UsageBar from '@/components/shared/UsageBar.vue'
import { useSettingsStore } from '@/stores/settings.store'
import type { ChipTone } from '@/types/status'
import type { ClusterEnvironment, ClusterSummary } from '@/types/fleet'

const props = defineProps<{ card: ClusterSummary }>()

const emit = defineEmits<{
	(e: 'open', card: ClusterSummary): void
	(e: 'troubleshoot', card: ClusterSummary): void
}>()

const settings = useSettingsStore()

const DOT: Record<ChipTone, string> = {
	ok: 'var(--ok)',
	warn: 'var(--warn)',
	err: 'var(--err)',
	info: 'var(--info)',
	idle: 'var(--idle)',
}
const dotColor = computed(() => DOT[props.card.statusTone])

const ENV_CYCLE: ClusterEnvironment[] = ['prod', 'staging', 'dev', 'none']
const env = computed(() => settings.envFor(props.card.name))

function cycleEnv() {
	const next = ENV_CYCLE[(ENV_CYCLE.indexOf(env.value) + 1) % ENV_CYCLE.length]
	settings.setEnv(props.card.name, next)
}
</script>

<template>
	<div class="ccard" :class="[{ degraded: card.issues }, env !== 'none' ? `env-${env}` : '']" @click="emit('open', card)">
		<div class="top">
			<div class="id">
				<span class="hdot" :style="{ background: dotColor, boxShadow: `0 0 0 4px color-mix(in srgb, ${dotColor} 18%, transparent)` }"></span>
				<div>
					<div class="cname">
						{{ card.name }}
						<button type="button" class="env-pill" :class="env" title="Tag environment (prod / staging / dev)" @click.stop="cycleEnv()">
							{{ env === 'none' ? '+ env' : env }}
						</button>
					</div>
					<div class="csrc">{{ card.source }}</div>
				</div>
			</div>
			<StatusChip :status="card.statusLabel" :tone="card.statusTone" />
		</div>

		<template v-if="card.metricsAvailable">
			<div class="metrics">
				<div class="m">
					<div class="mhead"><span class="dim">CPU</span><span class="mono">{{ card.cpu }}%</span></div>
					<UsageBar :pct="card.cpu" />
				</div>
				<div class="m">
					<div class="mhead"><span class="dim">Memory</span><span class="mono">{{ card.memory }}%</span></div>
					<UsageBar :pct="card.memory" />
				</div>
			</div>

			<div class="foot">
				<template v-if="card.issues">
					<span class="attn">🩺 {{ card.issues }} workloads need attention</span>
					<a class="link" @click.stop="emit('troubleshoot', card)">Open Troubleshoot →</a>
				</template>
				<template v-else>
					<span class="tag">{{ card.nodes }} nodes</span>
					<span class="tag">{{ card.pods }} pods</span>
					<span class="tag">{{ card.namespaces }} namespaces</span>
				</template>
			</div>
		</template>

		<div v-else class="no-metrics">Live metrics unavailable — metrics-server integration pending.</div>
	</div>
</template>

<style scoped>
.ccard {
	background: var(--surface);
	border: 1px solid var(--border-soft);
	border-radius: var(--r-lg);
	padding: 18px;
	cursor: pointer;
	transition: border-color 0.15s ease;
}
.ccard:hover {
	border-color: var(--brand);
}
.ccard.degraded {
	border-color: rgba(248, 81, 73, 0.4);
}
.ccard.env-prod {
	border-left: 4px solid var(--env-prod);
}
.ccard.env-staging {
	border-left: 4px solid var(--env-staging);
}
.ccard.env-dev {
	border-left: 4px solid var(--env-dev);
}
.top {
	display: flex;
	align-items: center;
	justify-content: space-between;
}
.id {
	display: flex;
	align-items: center;
	gap: 12px;
}
.hdot {
	width: 9px;
	height: 9px;
	border-radius: 50%;
	flex: none;
}
.cname {
	font-family: var(--mono);
	font-size: 15px;
	font-weight: 600;
	color: var(--text);
	display: flex;
	align-items: center;
	gap: 8px;
}
.env-pill {
	font-family: var(--sans);
	font-size: 9.5px;
	font-weight: 700;
	letter-spacing: 0.05em;
	text-transform: uppercase;
	padding: 1px 8px;
	border-radius: 999px;
	border: 1px dashed var(--border);
	background: none;
	color: var(--text-faint);
	cursor: pointer;
}
.env-pill.prod {
	border: 1px solid transparent;
	background: var(--env-prod-bg);
	color: var(--env-prod);
}
.env-pill.staging {
	border: 1px solid transparent;
	background: var(--env-staging-bg);
	color: var(--env-staging);
}
.env-pill.dev {
	border: 1px solid transparent;
	background: var(--env-dev-bg);
	color: var(--env-dev);
}
.csrc {
	font-size: 11.5px;
	color: var(--text-faint);
	margin-top: 2px;
}
.metrics {
	display: grid;
	grid-template-columns: 1fr 1fr;
	gap: 18px;
	margin-top: 16px;
}
.mhead {
	display: flex;
	justify-content: space-between;
	font-size: 12px;
	margin-bottom: 8px;
}
.dim {
	color: var(--text-dim);
}
.mono {
	font-family: var(--mono);
}
.foot {
	display: flex;
	align-items: center;
	gap: 10px;
	margin-top: 16px;
	font-size: 12px;
}
.tag {
	font-family: var(--mono);
	font-size: 11px;
	padding: 2px 7px;
	border-radius: 5px;
	background: var(--surface-2);
	border: 1px solid var(--border-soft);
	color: var(--text-dim);
}
.attn {
	color: var(--err);
}
.link {
	color: var(--brand);
	text-decoration: none;
	cursor: pointer;
}
.link:hover {
	text-decoration: underline;
}
.no-metrics {
	margin-top: 16px;
	font-size: 12px;
	color: var(--text-faint);
	font-style: italic;
}
</style>
