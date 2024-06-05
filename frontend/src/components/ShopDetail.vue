<template>
  <!-- <router-link to="/shops/shop/:id/product:proid"> -->
    <div class="font-style">
    <h1 class="text-8xl font-bold mb-6 flex justify-center">{{ productList.shop_name }}</h1>
    <p class=" mb-6 flex justify-center text-5xl" >{{ productList.shop_category }}</p>

    <div
      class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8 ml-9 "
    >
      <div
        v-for="product in productList.shop_products"
        :key="product.Product_ID"
        class="bg-white rounded-xl w-96 shadow-xl  hover:shadow-2xl hover:text-2xl transition duration-300 bg-gradient-to-bl from-green-300 via-yellow-300 to-pink-300"
      >
        <img
          :src="product.image"
          :alt="product.product_name"
          class="w-full h-56  rounded-md mb-4"
        />
        <h2 class="text-xl font-semibold mb-2 ml-3">{{ product.product_name }}</h2>
        <p class="text-gray-800 font-semibold mt-2 ml-3">ETB {{ product.price }}</p>
       
        <button
          @click="addToCart(product)" v-if="userLoggedIn"
          class="mr-2 ml-3 mt-3 mb-1 bg-gradient-to-b from-cyan-500 via-blue-600 to-indigo-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
        >
          Add to cart
        </button>
        <router-link  :to="{ name: 'product', params: { proid: product._id }}" >
          <button
        
            class="mr-2 mb-1 ml-3 mt-2 bg-gradient-to-r  from-teal-500 via-cyan-600 to-blue-700 hover:bg-red-700 text-white font-bold py-2 px-4 rounded-lg shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
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
import  useUserStore  from "../stores/user"
import { mapState } from 'pinia';
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
  computed: {
    ...mapState(useUserStore, ['userLoggedIn']),
    
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
