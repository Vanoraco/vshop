<template>
  <div class="grid grid-cols-4 gap-5">
    <div
      v-for="(shop, index) in shops"
      :key="index"
      class="rounded-md max-w-96 bg-grey shadow-sm"
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
          :to="'/Shop detail/' + shop.ShopName"
          class="px-4 font-semibold hover:cursor-pointer text-sm text-indigo-700"
          >See more</router-link
        >
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
            "https://www.suninternational.com/content/dam/suninternational/sun-city-resort/shopping/SCR5r0250-sun-city-sun-central-retail-billabong-XS_5072-1.jpg.sunimage.800.400.jpg",
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
            "https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w",
            'https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w',
            'https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w',
            'https://images.squarespace-cdn.com/content/v1/62b85870cef97862d6324f88/51e920da-a832-4290-960e-f5b051ef0cae/Labour+and+Wait+london+interior.jpg?format=2500w'

          ]
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
