<template>
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
</template>
  
  <script>

    import axios from 'axios';
    export default {
      mounted() {
          this.getShop()
      },
      data() {
          return {
              shops: [
                {
                  name: 'Siqala',
                  Rating: 5,
                  Categories: 'Electronics',
                  phone_number:"0983315117",
                  ShopImg: 'https://www.suninternational.com/content/dam/suninternational/sun-city-resort/shopping/SCR5r0250-sun-city-sun-central-retail-billabong-XS_5072-1.jpg.sunimage.800.400.jpg'
                },
                {
                  name: 'Siqala',
                  Rating: 5,
                  Categories: 'Electronics',
                  ShopImg: 'https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w'
                },
                {
                  name: 'Siqala',
                  Rating: 5,
                  Categories: 'Electronics',
                  ShopImg: 'https://www.suninternational.com/content/dam/suninternational/sun-city-resort/shopping/SCR5r0250-sun-city-sun-central-retail-billabong-XS_5072-1.jpg.sunimage.800.400.jpg'
                },
                {
                  name: 'Siqala',
                  Rating: 5,
                  Categories: 'Electronics',
                  ShopImg: 'https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w'
                },
                {
                  name: 'Siqala',
                  Rating: 5,
                  Categories: 'Electronics',
                  ShopImg: 'https://www.suninternational.com/content/dam/suninternational/sun-city-resort/shopping/SCR5r0250-sun-city-sun-central-retail-billabong-XS_5072-1.jpg.sunimage.800.400.jpg'
                },
                {
                  name: 'Siqala',
                  Rating: 5,
                  Categories: 'Electronics',
                  ShopImg: 'https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w'
                }
            ],
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