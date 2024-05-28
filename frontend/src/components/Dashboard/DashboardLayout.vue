<script setup lang="ts">
import Sidebar from './Sidebar.vue'
import Header from './Header.vue'

console.log(Sidebar)
</script>

<template>
  <div class="flex h-screen bg-gray-200 font-roboto -mt-12">
    <Sidebar />
    
    <div class="flex-1 flex flex-col overflow-hidden">
      <Header />

      <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-200">
        <div class="container mx-auto px-6 py-8">
          <slot />
        </div>
      </main>
    </div>
  </div>
</template>

<style>
  ::-webkit-scrollbar {
  width: 7px;
}

::-webkit-scrollbar-track{
  background: #2D3748;
}

::-webkit-scrollbar-thumb{
  background: #CBD5E0;
}

::-webkit-scrollbar-thumb:hover{
  background: #718096;
}
</style>
