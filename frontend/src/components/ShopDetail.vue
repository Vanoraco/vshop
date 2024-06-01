<template>
  <router-link to="/shops/shop/:id/product:proid">
    <div>
    <h1 class="text-4xl font-bold mb-6">{{ productList.shop_name }}</h1>
    <p class="text-lg mb-6">{{ productList.shop_category }}</p>

    <div
      class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8"
    >
      <div
        v-for="product in productList.shop_products"
        :key="product.Product_ID"
        class="bg-white rounded-lg shadow-md p-4 hover:shadow-xl transition duration-300"
      >
        <img
          :src="product.image"
          :alt="product.product_name"
          class="w-full h-48 object-cover rounded-md mb-4"
        />
        <h2 class="text-xl font-semibold mb-2">{{ product.product_name }}</h2>
        <p class="text-gray-800 font-semibold mt-2">ETB {{ product.price }}</p>
        <button
          class="mr-2 bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
        >
          3D
        </button>
        <button
          @click="addToCart(product)"
          class="mr-2 bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
        >
          Add to cart
        </button>
      </div>
    </div>
  </div>
  </router-link>
 
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
      const products = await axios.get(
        `http://localhost:8000/shops/viewproducts?shop_id=${shopID}`
      );
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
/* Add your styles here */
</style>
