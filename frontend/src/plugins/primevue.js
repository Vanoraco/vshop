import PrimeVue from 'primevue/config';

export default {
  install: (app) => {
    app.config.globalProperties.$primevue = PrimeVue;
    app.use(PrimeVue);
  }
}