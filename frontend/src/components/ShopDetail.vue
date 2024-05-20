<template>
  <div>
    <h1 class="text-4xl font-bold mb-6">{{ shop.name }}</h1>
    <p class="text-lg mb-6">{{ shop.description }}</p>

    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8">
      <div
        v-for="product in products"
        :key="product.id"
        class="bg-white rounded-lg shadow-md p-4 hover:shadow-xl transition duration-300"
      >
        <img :src="product.image" :alt="product.name" class="w-full h-48 object-cover rounded-md mb-4" />
        <h2 class="text-xl font-semibold mb-2">{{ product.name }}</h2>
        <p class="text-gray-500">{{ product.description }}</p>
        <p class="text-gray-800 font-semibold mt-2">{{ product.price }}</p>
        <button class="mr-2 bg-red-600">3D</button>
        <button class="ml-2 bg-red-600" @click="addToCart(product)">Add to cart</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      shop: {},
      products: [],
    };
  },
  async mounted() {
    console.log('Route Params on Mount:', this.$route.params);
    await this.getShop();
    await this.getProducts();
  },
 
  methods: {
    async getShop() {
      try {
        const shopId = this.$route.params.id;
        console.log('Shop ID in getShop:', shopId);  // Debug line
        const response = await axios.get(`http://localhost:8000/shops/viewshop/${shopId}`);
        this.shop = response.data;
      } catch (error) {
        console.error('Error fetching shop details:', error);
      }
    },
    async getProducts() {
      try {
        const shopId = this.$route.params.id;
        const response = await axios.get(`http://localhost:8000/shops/viewproducts/${shopId}`);
        this.products = response.data;
      } catch (error) {
        console.error('Error fetching products:', error);
      }
    },
    // addToCart(product) {
    //   this.$emit('add-to-cart', product);
    // },
  },
};
</script>

<style>
</style>