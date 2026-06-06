<script setup lang="ts">

</script>

<!-- src/components/SettingsViewUI.vue -->
<template>
    <section class="settings">
        <h1 class="title">Settings</h1>
        <p class="subtitle">Configure &amp; Personalize your KubExplorer Tool</p>

        <!-- Tabs (static UI: “App” looks active) -->
        <div class="tabs">
            <button class="tab active">App</button>
            <button class="tab">Kubernetes</button>
        </div>

        <!-- Panel: App -->
        <div class="panel">
            <!-- Dark mode -->
            <div class="row">
                <div class="left">
                    <div class="label">Dark mode</div>
                </div>
                <div class="right">
                    <!-- Pure CSS toggle (no bindings) -->
                    <label class="switch">
                        <input type="checkbox" checked />
                        <span class="slider"></span>
                    </label>
                </div>
            </div>

            <!-- Clear cache -->
            <div class="row">
                <div class="left">
                    <div class="label">Clear cache</div>
                </div>
                <div class="right">
                    <button class="pill">Clean</button>
                </div>
            </div>

            <!-- App data location -->
            <div class="row">
                <div class="left">
                    <div class="label">App data location</div>
                </div>
                <div class="right path-field">
                    <input
                        class="path-input"
                        type="text"
                        value="/Users/beto/Library/Application Support/Lens/binaries"
                        readonly
                    />
                    <button class="browse">Browse</button>
                </div>
            </div>
        </div>
    </section>
</template>

<!-- No <script> block on purpose -->

<style scoped>
/* Base palette */
.settings {
    color: #e7ecf3;
    background: #0b0e12;
    min-height: 100%;
    padding: 24px;
    box-sizing: border-box;
}
.title {
    margin: 0 0 4px;
    font-weight: 700;
    font-size: 1.6rem;
}
.subtitle {
    margin: 0 0 18px;
    color: #8a95a3;
}

/* Tabs */
.tabs {
    display: flex;
    gap: 10px;
    border-bottom: 1px solid #28303a;
    margin-bottom: 10px;
}
.tab {
    background: transparent;
    color: #bfc9d6;
    border: none;
    padding: 8px 10px;
    border-bottom: 2px solid transparent;
    cursor: default; /* UI only */
    font-weight: 600;
}
.tab.active {
    color: #ffffff;
    border-bottom-color: #6d5cff;
}

/* Panel & rows */
.panel { padding-top: 10px; }
.row {
    display: grid;
    grid-template-columns: 1fr auto;
    align-items: start;
    padding: 14px 0;
    border-bottom: 1px solid #1c232c;
}
.left .label {
    color: #cbd5e1;
    font-size: 0.98rem;
}
.right { display: flex; align-items: center; gap: 10px; }

/* Toggle switch (CSS only) */
.switch {
    position: relative; display: inline-block;
    width: 54px; height: 28px;
}
.switch input { display: none; }
.slider {
    position: absolute; inset: 0;
    background: #2a3340;
    border-radius: 99px;
    transition: background .2s ease;
    box-shadow: inset 0 0 0 1px #39404b;
}
.slider::after {
    content: ""; position: absolute; top: 3px; left: 3px;
    width: 22px; height: 22px; border-radius: 50%;
    background: #cfd7df;
    transition: transform .2s ease, background .2s ease;
}
input:checked + .slider { background: #6d5cff; }
input:checked + .slider::after { transform: translateX(26px); background: #fff; }

/* Buttons */
.pill {
    border: none; border-radius: 999px;
    padding: 6px 14px;
    background: #2d3650; color: #e7ecf3; font-weight: 600;
}
.browse {
    border: none; border-radius: 10px;
    padding: 8px 12px;
    background: #6d5cff; color: #fff; font-weight: 700;
}

/* Path field */
.path-field { display: flex; gap: 10px; align-items: center; }
.path-input {
    min-width: 360px;
    padding: 8px 10px;
    border-radius: 10px;
    border: 1px solid #313b47;
    background: #141a22;
    color: #e7ecf3;
}
</style>