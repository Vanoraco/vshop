import { createApp } from 'vue'
import './assets/base.css'
import App from './App.vue'

import PrimeVuePlugin from './plugins/primevue';

createApp(App).mount('#app')

app.config.globalProperties.$primevue = PrimeVue;
app.use(PrimeVuePlugin);
