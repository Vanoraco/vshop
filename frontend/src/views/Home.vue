<template>
   
    <Carousels />
    <br>
    <br>

    <div class="ml-3 shop-header">
      <p class="text-3xl font-extrabold">Top Shops</p>
    </div>
    <ShopCard />
    <div class="flex justify-center mt-9 ">
      <p class="bg-[#DB4444] text-white px-5 py-1 shop-header">View All Shops</p>
    </div>

    <div class="ml-3 shop-header">
      <p class="text-3xl font-extrabold">Explore based on Category</p>
    </div>

    <div class="text-8xl flex justify-center">
      <p class="pi pi-desktop rounded-lg shadow-2xl mt-9 ml-20 h-[180px] w-[200px] flex p-12">
        
      </p>
      
      <p class="pi pi-shopping-bag rounded-lg shadow-2xl mt-9 ml-20 h-[180px] w-[200px] p-12"></p>
      <p class="pi pi-palette rounded-lg shadow-2xl mt-9 ml-20 h-[180px] w-[200px] p-12"></p>
    </div>
</template>

<script>
  import Carousels from '../components/Carousels.vue';
  import ShopCard from '../components/ShopCard.vue'
  export default {
    name: 'home',
    components: {
        Carousels,
        ShopCard
    }
  }
</script>

<style scoped>
   @import url('https://fonts.googleapis.com/css2?family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap');
   @import url('https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap');
   @import url('https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap');
   .shop-header {
    font-family: "Marcellus", sans-serif;
   }
</style>