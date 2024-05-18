<!-- <template>
  <div class="grid grid-cols-3 gap-5">
    <div
      v-for="(shop, index) in shops"
      :key="index"
      class="rounded-md max-w-96 shadow-sm shop_container"
    >
      <div class="w-full px-6 pt-4 flex">
        <p class="text-lg font-semibold">{{ shop.ShopName }}</p>
      </div>
      <div class="grid gap-4 p-6 grid-cols-2">
        <figure
          v-for="(image, imgIndex) in shop.image"
          :key="imgIndex"
          class="ring-1 ring-blue-100 overflow-hidden rounded-sm col-span-1 h-32 flex items-center justify-center"
        >
          <img
            :src="image"
            alt="product"
            class="rounded-sm h-full w-full hover:scale-110 object-cover transition-transform duration-500 ease-in-out"
          />
        </figure>
      </div>
      <div class="p-4">
        <router-link
          to="/shop detail"
          class="px-4 font-semibold hover:cursor-pointer text-sm text-indigo-700"
        >
          View the shop
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      shops: [
        {
          ShopName: "Bokacha Electronics",
          ShopOwner: "Abebe",
          Rating: 5,
          Categories: "Electronics",
          image: [
            "https://cdn.pixabay.com/photo/2022/09/25/22/25/iphones-7479304_1280.jpg",
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
          ],
        },
        {
          ShopName: "Bokacha Electronics",
          ShopOwner: "Abebe",
          Rating: 5,
          Categories: "Electronics",
          image: [
            "https://www.suninternational.com/content/dam/suninternational/sun-city-resort/shopping/SCR5r0250-sun-city-sun-central-retail-billabong-XS_5072-1.jpg.sunimage.800.400.jpg",
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
          ],
        },
        {
          ShopName: "Emaye",
          ShopOwner: "Mola",
          Rating: 4,
          Categories: "Clothes",
          image: [
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
          ],
        },
      ],
    };
  },
};
</script>
<style scoped>
.shop_container {
  background: #c3dfe0;
}
</style> -->
<template>
  <ul class="grid grid-cols-4 gap-5">
    <li  v-for="shop in shops" class="rounded-lg shadow-2xl mt-9 ml-20 shop-list">
      <img :src="shop.ShopImg" alt="" class="rounded lg w-[100%] h-[250px]"/>
          <p class="flex justify-center font-bold">{{ shop.name }} </p>
          <p class="flex justify-center font-bold"> Categories - {{ shop.Categories }} </p>
          <p class="flex justify-center font-bold text-gold">Owner phone number- {{ shop.phone_number}} </p>
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
