import { createApp } from 'vue'
import App from './App.vue'
import Router from './plugins/router'
import Store from './store/index'
import AxiosPlugin from './plugins/axios'
import './index.css'

const root = createApp(App)
root.use(Router)
root.use(AxiosPlugin, {})
root.use(Store)
root.mount('#app')
