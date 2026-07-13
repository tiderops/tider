import { createApp } from 'vue'
import '@/styles/main.css'
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { createPinia } from 'pinia'
import router from './router'
import App from './App.vue'
import '@mdi/font/css/materialdesignicons.css'

const vuetify = createVuetify({
	components,
	directives,
	icons: {
		defaultSet: 'mdi',
	},
	theme: {
		defaultTheme: 'kxLight',
		themes: {
			kxLight: {
				dark: false,
				colors: {
					primary: '#1583E6',
					secondary: '#10B6C2',
					background: '#EEF4F9',
					surface: '#FFFFFF',
					error: '#CF222E',
					info: '#0C93A0',
					success: '#1A7F37',
					warning: '#9A6700',
				},
			},
		},
	},
})

const app = createApp(App)
app.use(createPinia())
app.use(vuetify)
app.use(router)
app.mount('#app')
