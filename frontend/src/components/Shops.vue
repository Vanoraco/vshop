<template>
  <div>
    <div class="flex justify-center text-7xl font-bold shop-list">
      <p :class="{ 'hidden': $route.path !== '/shops' }">All Shops</p>
     
    </div>
    <hr :class="{ 'hidden': $route.path !== '/shops' }">
  <ul class="grid grid-cols-3 gap-1">
    
      <li class="rounded-lg shadow-xl  mt-9 ml-20 shop-list w-96  bg-gradient-to-r from-blue-300 to-yellow-300"  v-for="shop in shop" :key="shop.shop_id">
        <router-link  :to="{ name: 'shop', params: { id: shop.shop_id } }"
           class="h-[350px] w-full hover:shadow-2xl hover:text-2xl">
        <img :src="shop.image" alt="" class="rounded lg w-[100%] h-[250px]" />
        <p class="flex justify-center font-bold">{{ shop.shop_name }} </p>
        <p class="flex justify-center font-bold"> Categories - {{ shop.shop_category }} </p>
      </router-link>
       <!--  <p class="flex justify-center font-bold text-gold">Owner phone number- {{ shop.phone_number}} </p> -->
       <!--  <p class="flex justify-center text-amber-600">Rating {{ shop.Rating }} </p> -->
      </li>
   
    
  </ul>
</div>
</template>
  
  <script>

    import axios from 'axios';
    export default {
      mounted() {
          this.getShop()
      },
      data() {
          return {
            shop: []
          }
      },
      methods: {
         async getShop() {
            const response = await axios.get('http://localhost:8000/users/viewshops');
            this.shop = response.data;
            console.log(this.shop);
        }
      }
    }
  </script>
  
  <style>
   @import url('https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap');
     .shop-list {
      font-family: "Marcellus", sans-serif;
     }
  </style>