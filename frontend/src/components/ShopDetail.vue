<template>
  <!-- <router-link to="/shops/shop/:id/product:proid"> -->
    <div class="font-style">
    <h1 class="text-8xl font-bold mb-6 flex justify-center">{{ productList.shop_name }}</h1>
    <p class="text-lg mb-6 flex justify-center">{{ productList.shop_category }}</p>

    <div
      class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8 ml-9 "
    >
      <div
        v-for="product in productList.shop_products"
        :key="product.Product_ID"
        class="bg-white rounded-lg shadow-xl  hover:shadow-xl transition duration-300 bg-[radial-gradient(circle_at_bottom_left,_var(--tw-gradient-stops))] from-rose-600 via-sky-100 to-cyan-500"
      >
        <img
          :src="product.image"
          :alt="product.product_name"
          class="w-full h-56  rounded-md mb-4"
        />
        <h2 class="text-xl font-semibold mb-2">{{ product.product_name }}</h2>
        <p class="text-gray-800 font-semibold mt-2">ETB {{ product.price }}</p>
       
        <button
          @click="addToCart(product)"
          class="mr-2 bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
        >
          Add to cart
        </button>
        <router-link  :to="{ name: 'product', params: { proid: product._id }}" >
          <button
        
            class="mr-2 bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
          >
            Details
          </button>
        </router-link>
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
@import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");

.font-style {
  font-family: "Marcellus", sans-serif;
}
</style>
