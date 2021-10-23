import { createApp } from 'vue'
import App from './App.vue'
import Router from './plugins/router'
import AxiosPlugin from './plugins/axios'
import './index.css'

const root = createApp(App)
root.use(Router)
root.use(AxiosPlugin, {})
root.mount('#app')
