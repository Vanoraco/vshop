<template>
  <div class="font-style container mx-auto py-8">
    <h1 class="text-4xl md:text-6xl lg:text-8xl font-bold mb-6 text-center">{{ productList.shop_name }}</h1>
    <p class="text-base md:text-lg lg:text-xl mb-6 text-center">{{ productList.shop_category }}</p>
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8 px-4">
      <div
        v-for="product in productList.shop_products"
        :key="product.Product_ID"
        class="bg-white rounded-lg shadow-md hover:shadow-lg transition duration-300 bg-gradient-to-bl from-green-300 via-yellow-300 to-pink-300 p-4"
      >
      
        <img
          :src="product.image"
          :alt="product.product_name"
          class="w-full h-56 rounded-md mb-4 object-cover"
        />
        <h2 class="text-xl font-semibold mb-2">{{ product.product_name }}</h2>
        <p class="text-gray-800 font-semibold mt-2">ETB {{ product.price }}</p>
        <div class="flex justify-between mt-4">
          <button
            @click="addToCart(product)"
            class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
          >
            Add to <i class="fas fa-star  bg-grey"></i>
          </button>
          <router-link :to="{ name: 'product', params: { proid: product._id } }">
            <button
              class="bg-gradient-to-r  from-teal-500 via-cyan-600 to-blue-700 ml-3 text-white font-bold py-2 px-4 rounded-lg shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out hover:bg-red-700 text-white font-bold py-2 px-4 rounded shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
            >
              Details
            </button>
          </router-link>
          
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import { useRoute } from "vue-router";
import { useCartStore } from "../stores/CartStore";

export default {
  mounted() {
    this.ProductList();
  },
  data() {
    return {
      shopId: this.$route.params.id,
      productList: [],
    };
  },
  methods: {
    async ProductList() {
      const shopID = this.shopId;
      const products = await axios.get(`http://localhost:8000/shops/viewproducts?shop_id=${shopID}`);
      console.log(products.data);
      this.productList = products.data;
    },
    addToCart(product) {
      const cartStore = useCartStore();
      cartStore.addToCart(product);
    },
  },
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");

.font-style {
  font-family: "Marcellus", sans-serif;
}

.container {
  max-width: 1200px;
  margin: auto;
}

.grid > div {
  transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
}

.grid > div:hover {
  transform: scale(1.05);
  box-shadow: 0 10px 15px rgba(0, 0, 0, 0.1);
}

@media (min-width: 768px) {
  .text-center {
    text-align: left;
  }
}
</style>
