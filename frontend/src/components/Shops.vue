<template>
  <div class="container mx-auto">
    <div class="flex justify-center text-5xl font-bold mt-8 mb-4">
      <p :class="{ 'hidden': $route.path !== '/shops' }">All Shops</p>
    </div>
    <hr :class="{ 'hidden': $route.path !== '/shops' }" class="mb-8">
    <ul class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
      <li
        class="rounded-lg shadow-xl bg-gradient-to-r from-blue-300 to-yellow-300 p-4 transform transition duration-300 hover:scale-105"
        v-for="shop in shop"
        :key="shop.shop_id"
      >
        <router-link :to="{ name: 'shop', params: { id: shop.shop_id } }">
          <img :src="shop.image" alt="" class="rounded-lg w-full h-48 object-cover mb-4" />
          <div class="text-center">
            <p class="font-bold text-lg mb-2">{{ shop.shop_name }}</p>
            <p class="text-gray-700">Categories - {{ shop.shop_category }}</p>
          </div>
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  mounted() {
    this.getShop();
  },
  data() {
    return {
      shop: []
    };
  },
  methods: {
    async getShop() {
      const response = await axios.get('http://localhost:8000/users/viewshops');
      this.shop = response.data;
    }
  }
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Marcellus&display=swap');

.container {
  max-width: 96vw;
  margin-top: 2rem;
}

.shop-list {
  font-family: 'Marcellus', sans-serif;
}

img {
  transition: transform 0.3s;
}

img:hover {
  transform: scale(1.05);
}
</style>
