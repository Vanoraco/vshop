<template>
  <ul class="grid grid-cols-4 gap-5">
    <li  v-for="shop in shops" class="rounded-lg shadow-2xl mt-9 ml-20 shop-list">
      <img :src="shop.ShopImg" alt="" class="rounded lg w-[100%] h-[250px]"/>
          <p class="flex justify-center font-bold">Shop-Name -  {{ shop.ShopName }} </p>
          <p class="flex justify-center font-bold"> Categories - {{ shop.Categories }} </p>
          <p class="flex justify-center font-bold text-gold">Owner - {{ shop.ShopOwner }} </p>
           <p class="flex justify-center text-amber-600">Rating {{ shop.Rating }} </p>

    </li>
  </ul>
</template>

<script>
  export default {
    data() {
        return {
            shops: [
              {
                ShopName: 'Siqala',
                ShopOwner: 'Abebe',
                Rating: 5,
                Categories: 'Electronics',
                ShopImg: 'https://www.suninternational.com/content/dam/suninternational/sun-city-resort/shopping/SCR5r0250-sun-city-sun-central-retail-billabong-XS_5072-1.jpg.sunimage.800.400.jpg'
              },
              {
                ShopName: 'Emaye',
                ShopOwner: 'Mola',
                Rating: 4,
                Categories: 'Clothes',
                ShopImg: 'https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w'
              },
              {
                ShopName: 'Siqala',
                ShopOwner: 'Abebe',
                Rating: 5,
                Categories: 'Electronics',
                ShopImg: 'https://www.suninternational.com/content/dam/suninternational/sun-city-resort/shopping/SCR5r0250-sun-city-sun-central-retail-billabong-XS_5072-1.jpg.sunimage.800.400.jpg'
              },
              {
                ShopName: 'Emaye',
                ShopOwner: 'Mola',
                Rating: 4,
                Categories: 'Clothes',
                ShopImg: 'https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w'
              },
              {
                ShopName: 'Siqala',
                ShopOwner: 'Abebe',
                Rating: 5,
                Categories: 'Electronics',
                ShopImg: 'https://www.suninternational.com/content/dam/suninternational/sun-city-resort/shopping/SCR5r0250-sun-city-sun-central-retail-billabong-XS_5072-1.jpg.sunimage.800.400.jpg'
              },
              {
                ShopName: 'Emaye',
                ShopOwner: 'Mola',
                Rating: 4,
                Categories: 'Clothes',
                ShopImg: 'https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w'
              }
          ]
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