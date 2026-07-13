import { fileURLToPath } from 'node:url';
import { mergeConfig, defineConfig, configDefaults } from 'vitest/config';
import viteConfig from './vite.config';
export default mergeConfig(viteConfig, defineConfig({
    test: {
        globals: true,
        environment: 'jsdom',
        exclude: [...configDefaults.exclude, 'e2e/**'],
        root: fileURLToPath(new URL('./', import.meta.url)),
        setupFiles: ['./src/test-setup.ts'],
        // Vuetify ships .css imports in its ESM build; inline it so
        // component tests can mount Vuetify components under jsdom.
        server: {
            deps: {
                inline: ['vuetify'],
            },
        },
    },
}));
