import { createApp } from 'vue'
import './assets/base.css'
import App from './App.vue'
import { createPinia } from 'pinia'

import 'primevue/resources/primevue.min.css';

import 'primeicons/primeicons.css';


//import PrimeVuePlugin from './plugins/primevue';
import PrimeVue from "primevue/config";
import vuetify from './plugins/vuetify';
import router from './router/router'; 
import VeeValidateplugin from './includes/validation'


const pinia = createPinia()

const app = createApp(App)


//app.config.globalProperties.$primevue = PrimeVue;
app.use(PrimeVue);

app.use(vuetify);
 

app.use(router)

app.use(VeeValidateplugin)
app.use(pinia)

app.mount('#app')
