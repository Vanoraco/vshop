<template>
  <div>
    <div class="flex justify-center text-7xl font-bold shop-list">
      <p :class="{ 'hidden': $route.path !== '/shops' }">All Shops</p>
     
    </div>
    <hr :class="{ 'hidden': $route.path !== '/shops' }">
  <ul class="grid grid-cols-4 gap-5">
    <router-link v-for="shop in shop" :key="shop.name" :to="{ name: 'shop', params: { id: shop.shop_id } }"
           class="rounded-lg shadow-2xl mt-9 ml-20 shop-list">
      <li>
        <img :src="shop.image" alt="" class="rounded lg w-[100%] h-[250px]" />
        <p class="flex justify-center font-bold">{{ shop.shop_name }} </p>
        <p class="flex justify-center font-bold"> Categories - {{ shop.shop_category }} </p>
       <!--  <p class="flex justify-center font-bold text-gold">Owner phone number- {{ shop.phone_number}} </p> -->
       <!--  <p class="flex justify-center text-amber-600">Rating {{ shop.Rating }} </p> -->
      </li>
    </router-link>
    
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