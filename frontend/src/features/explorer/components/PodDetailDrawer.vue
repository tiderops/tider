<script setup lang="ts">
import StatusChip from '@/components/shared/StatusChip.vue'
import UsageBar from '@/components/shared/UsageBar.vue'
import type { PodDetail } from '../types'

const props = defineProps<{ pod: PodDetail | null; open: boolean }>()

const emit = defineEmits<{
	(e: 'close'): void
	(e: 'diagnose', pod: PodDetail): void
	(e: 'action', label: string, pod: PodDetail): void
}>()
</script>

<template>
	<Transition name="drawer">
		<div v-if="open && props.pod" class="drawer-root">
			<div class="scrim" @click="emit('close')"></div>
			<aside class="drawer">
				<header class="drawer-head">
					<div class="head-main">
						<div class="head-row">
							<StatusChip :status="props.pod.status" />
							<span v-if="props.pod.restartWindow" class="faint">{{ props.pod.restartWindow }}</span>
						</div>
						<div class="name">{{ props.pod.name }}</div>
						<div class="faint sub">{{ props.pod.namespace }} · {{ props.pod.node }}</div>
					</div>
					<button class="icon-btn" @click="emit('close')">✕</button>
				</header>

				<div class="drawer-body">
					<button class="btn ai full" @click="emit('diagnose', props.pod)">🩺 Diagnose with Troubleshoot Assistant</button>

					<div class="section-t">Overview</div>
					<dl class="kv">
						<dt>Controller</dt>
						<dd>{{ props.pod.controller }}</dd>
						<dt>Image</dt>
						<dd>{{ props.pod.image }}</dd>
						<dt>QoS class</dt>
						<dd>{{ props.pod.qosClass }}</dd>
						<dt>Pod IP</dt>
						<dd>{{ props.pod.podIP }}</dd>
						<dt>Service account</dt>
						<dd>{{ props.pod.serviceAccount }}</dd>
						<dt>Started</dt>
						<dd>{{ props.pod.started }}</dd>
					</dl>

					<div class="section-t">Resources</div>
					<div class="usage-grid">
						<div class="usage-card">
							<div class="usage-top"><span class="dim">CPU</span><span class="mono">{{ props.pod.cpu.label }}</span></div>
							<UsageBar :pct="props.pod.cpu.pct" />
							<div class="usage-note">{{ props.pod.cpu.pct >= 95 ? 'throttled · at limit' : 'within limit' }}</div>
						</div>
						<div class="usage-card">
							<div class="usage-top"><span class="dim">Memory</span><span class="mono">{{ props.pod.memory.label }}</span></div>
							<UsageBar :pct="props.pod.memory.pct" />
							<div class="usage-note">{{ props.pod.memory.pct >= 90 ? 'near limit' : 'within limit' }}</div>
						</div>
					</div>

					<template v-if="props.pod.lastTermination">
						<div class="section-t">Last termination</div>
						<pre class="codeblock">
                            <span class="c"># container exit · reason {{ props.pod.lastTermination.reason }}</span>
                            <span class="k">Reason</span>: {{ props.pod.lastTermination.reason }}
                            <span class="k">Message</span>: {{ props.pod.lastTermination.message }}
                            <span class="c"># last log lines</span>
                            <template v-for="(line, i) in props.pod.lastTermination.log" :key="i">{{ line }}</template>
                        </pre>
					</template>

					<div class="actions">
						<button class="btn" @click="emit('action', 'Logs', props.pod)">📜 Logs</button>
						<button class="btn" @click="emit('action', 'Exec', props.pod)">⌨ Exec</button>
						<button class="btn" @click="emit('action', 'Restart', props.pod)">↻ Restart</button>
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
.drawer-head {
	padding: 18px 20px;
	border-bottom: 1px solid var(--border-soft);
	display: flex;
	align-items: flex-start;
	gap: 12px;
}
.head-main {
	flex: 1;
	min-width: 0;
}
.head-row {
	display: flex;
	align-items: center;
	gap: 8px;
}
.name {
	font-family: var(--mono);
	font-size: 15px;
	font-weight: 600;
	margin-top: 8px;
	word-break: break-all;
}
.sub {
	font-size: 12px;
	margin-top: 2px;
}
.faint {
	color: var(--text-faint);
	font-size: 11px;
}
.dim {
	color: var(--text-dim);
}
.mono {
	font-family: var(--mono);
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
.drawer-body {
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
.kv {
	display: grid;
	grid-template-columns: 124px 1fr;
	gap: 10px 14px;
	font-size: 12.5px;
	margin: 0;
}
.kv dt {
	color: var(--text-faint);
}
.kv dd {
	margin: 0;
	color: var(--text);
	font-family: var(--mono);
	font-size: 12px;
	word-break: break-all;
}
.usage-grid {
	display: grid;
	grid-template-columns: 1fr 1fr;
	gap: 14px;
}
.usage-card {
	background: var(--surface);
	border: 1px solid var(--border-soft);
	border-radius: var(--r-lg);
	padding: 13px 14px;
}
.usage-top {
	display: flex;
	justify-content: space-between;
	font-size: 12px;
	margin-bottom: 8px;
}
.usage-note {
	font-size: 10.5px;
	color: var(--text-faint);
	margin-top: 6px;
}
.codeblock {
	background: #0a0e14;
	border: 1px solid var(--border-soft);
	border-radius: var(--r-md);
	padding: 14px 16px;
	font-family: var(--mono);
	font-size: 12px;
	line-height: 1.7;
	color: #c9d4e0;
	overflow: auto;
	margin: 0;
	white-space: pre-wrap;
}
.codeblock .k {
	color: #6ea8ff;
}
.codeblock .c {
	color: var(--text-faint);
}
.actions {
	display: flex;
	gap: 8px;
	margin-top: 24px;
}
.btn {
	display: inline-flex;
	align-items: center;
	justify-content: center;
	gap: 8px;
	font-size: 13px;
	font-weight: 600;
	padding: 9px 14px;
	border-radius: var(--r-sm);
	border: 1px solid var(--border);
	background: var(--surface-2);
	color: var(--text);
	cursor: pointer;
	flex: 1;
}
.btn:hover {
	border-color: #3a465a;
}
.btn.full {
	width: 100%;
	flex: none;
}
.btn.ai {
	background: linear-gradient(180deg, var(--accent), var(--accent-deep));
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
