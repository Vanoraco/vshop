import { createApp } from 'vue'
import './assets/base.css'
import App from './App.vue'
import { createPinia } from 'pinia'

import 'primevue/resources/primevue.min.css';

import 'primeicons/primeicons.css';



import ToastPlugin from 'vue-toast-notification';
// Import one of the available themes
//import 'vue-toast-notification/dist/theme-default.css';
import 'vue-toast-notification/dist/theme-bootstrap.css';

import DashboardLayout from './components/Dashboard/DashboardLayout.vue'

//import PrimeVuePlugin from './plugins/primevue';
import PrimeVue from "primevue/config";
import vuetify from './plugins/vuetify';
import router from './router/router'; 
import VeeValidateplugin from './includes/validation'
// import '@fortawesome/fontawesome-free/css/all.css';
import '@fortawesome/fontawesome-free/css/all.min.css';




const pinia = createPinia()

const app = createApp(App)


//app.config.globalProperties.$primevue = PrimeVue;
app.use(PrimeVue);

app.use(vuetify);
 
app.use(ToastPlugin)

app.use(router)

app.use(VeeValidateplugin)
app.use(pinia)

app.component('DefaultLayout', DashboardLayout)

app.mount('#app')
