<template>
   
    <Carousels />
    <Carousels />

    
    <br>
    <br>

    <div class="ml-3 shop-header">
      <p class="text-3xl font-extrabold">Top Shops</p>
    </div>
    
    <shop/>

     
    <div class="flex justify-center mt-9 ">
      <router-link to="/shops"><p class="bg-[#DB4444] text-white px-5 py-1 shop-header">View All Shops</p></router-link> 
    </div>

    <div class="ml-3 shop-header mt-12">
      <p class="text-3xl font-extrabold">Explore based on Category</p>
    </div>

    <div class="text-8xl flex justify-center shop-header">
      <div class="relative">
        <i class="pi pi-desktop rounded-lg shadow-2xl mt-9 ml-20 h-[180px] w-[200px] flex p-12"></i>
        <span class="text-xl mt-5 absolute ml-32 flex justify-center">Electronics</span>
      </div>

      <div class="relative">
        <i class="pi pi-shopping-bag rounded-lg shadow-2xl mt-9 ml-20 h-[180px] w-[200px] flex p-12"></i>
        <span class="text-xl mt-5 absolute ml-32 flex justify-center">Shoe / Clothes</span>
      </div>

      <div class="relative">
        <i class="pi pi-palette rounded-lg shadow-2xl mt-9 ml-20 h-[180px] w-[200px] flex p-12"></i>
        <span class="text-xl mt-5 absolute ml-32 flex justify-center">Cosmetics</span>
      </div>
     
    </div>

    <div class="mt-36 bg-slate-900 text-white flex justify-center h-60 rounded">
      <p class="flex mt-28 text-4xl shop-header">Enhance Your Shopping Experience</p>
    </div>

    <div class="flex justify-center shop-header">
      <div class="flex flex-col">
      <i class="pi pi-truck rounded-full  text-4xl mt-9 ml-20  w-[200px] flex p-12 "></i>
        <span class="text-xl relative ml-24 -mt-5 ">Fast Delivery</span>
      </div>

      <div class="flex flex-col">
      <i class="pi pi-headphones rounded-full  text-4xl mt-9 ml-28  flex p-12 "></i>
        <span class="text-xl relative ml-24 -mt-5 ">24/7 Customer Service</span>
      </div>

      <div class="flex flex-col">
      <i class="pi pi-money-bill rounded-full  text-4xl mt-9 ml-28 w-[200px] flex p-12 "></i>
        <span class="text-xl relative ml-24 -mt-5 ">Money Back Guarantee</span>
      </div>
    </div>

    <div class="bg-slate-900 text-white shop-header mt-20 shop-header">
      <div class="pt-12 pb-9 flex flex-row gap-56 justify-center">
      <div class="">
        <p class="font-bold text-xl">Accounts</p>
        <ul class="flex gap-6 flex-col mt-6 text-sm">
          <li>Login / Register</li>
          <li>Cart</li>
          <li>Shop</li>
        </ul>
      </div>

      <div>
        <p class="font-bold text-xl">Quick Link</p>
        <ul class="flex gap-6 flex-col mt-6 text-sm">
          <li>Privacy Policy</li>
          <li>Terms Of Use</li>
          <li>Contact</li>
        </ul>
      </div>

      <div>
        <p class="font-bold text-xl">Download App</p>
        <img src="../assets/images/qrcode.png" class="bg-white h-36" alt="">
      </div>
    </div>
  </div>
</template>

<script>
  import Carousels from '../components/Carousels.vue';
  // import Carousel  from '../components/Carousal.vue';
  // import ShopCard from '../components/ShopCard.vue'
  import shop from '../components/Shops.vue'
  export default {
    
   
    data() {
      return {
       pageReloaded: false
      }
    },
    name: 'home',
    components: {
        Carousels,
        shop
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