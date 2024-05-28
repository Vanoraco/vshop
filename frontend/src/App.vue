

<template>
  <AppHeader />
  <br />
  <br />
  
  <component :is="layout">
    <router-view />
  </component>
 
</template>

<script setup>
import AppHeader from "./components/AppHeader.vue";
import Layout from "./views/Admin/dashboard/Layout.vue";
import Home from "./views/Home.vue";

import DashboardLayout from "./components/Dashboard/DashboardLayout.vue";
import { computed } from 'vue'
import { useRouter } from 'vue-router'

const defaultLayout = 'default'

const { currentRoute } = useRouter()

const layout = computed(
  () => `${currentRoute.value.meta.layout || defaultLayout}-layout`,
)
</script>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>