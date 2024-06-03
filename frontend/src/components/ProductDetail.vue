<template >
  <div class="flex flex-col md:flex-row justify-between items-start max-w-[96vw] mt-8">
    <div class="w-full md:w-1/2 mr-8">
      <div class="product-gallery">
        <div class="main-image h-[590px] w-full p-10 overflow-hidden">
          <img
            :src="productList.image"
            :alt="productList.name"
            class="w-full rounded-lg object-fill h-[500px]  "
          />
        </div>
        <div class="thumbnails flex justify-between px-8 mt-4">
          <div
            v-for="image in product.images"
            :key="image"
            class="thumbnail w-1/4 mr-2"
          >
            <img
              :src="image"
              :alt="product.name"
              class="w-full h-auto rounded-lg"
            />
          </div>
        </div>
      </div>
    </div>

    <div class="w-full md:w-1/2">
      <h1 class="text-4xl font-bold mb-6">{{ productList.shop_name }}</h1>
      <h1 class="text-3xl font-bold mb-2">{{ productList.product_name }}</h1>
      <div class="flex items-center mb-4">
        <div class="rating text-yellow-400 mr-2">
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
          <i class="fas fa-star-half-alt"></i>
        </div>
        <span class="text-gray-500"
          >({{ product.reviews.length }} reviews)</span
        >
      </div>
      <p class="text-2xl font-bold mb-4">ETB {{ productList.price }}</p>
      <p class="mb-4">{{ productList.product_description }}</p>
      <div class="flex items-center mb-4">
        <button
          class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded mr-4"
        >
          Buy Now
        </button>
        <button
          class="mr-2 bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
        >
          Add to cart
        </button>
        <router-link to="/3Dview">
          <button
            class="mr-2 bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
          >
            3D view
          </button>
        </router-link>
      </div>
      <!-- <div class="product-3d-view">
       
      </div> -->
    </div>
  </div>
</template>
  
  <script>
import axios from 'axios';

export default {
  // mounted() {
  //   this.ProductList();
  // },

  data() {
    return {
      shopId: this.$route.params.id,
      prodId: this.$route.params.proid,
      productList: [],
      product: {
        name: "Product Name",
        price: 49.99,
        description: "This is a description of the product.",
        image: "https://via.placeholder.com/500x500",
        images: [
          "https://via.placeholder.com/200x200",
          "https://via.placeholder.com/200x200",
          "https://via.placeholder.com/200x200",
          "https://via.placeholder.com/200x200",
        ],
        threeDUrl: "./assets/images/ecommerce.jpg",
        reviews: [
          { rating: 5, comment: "Great product!" },
          { rating: 4, comment: "Good product, but could be better." },
          { rating: 3, comment: "Average product." },
        ],
      },
    };
  },
  methods: {

    async ProductList() {
      const shopID = this.shopId;
      const products = await axios.get(
        `http://localhost:8000/shop/product?shop_id=${shopID}&product_id=${this.prodId}`
      );
      console.log(products.data);
      this.productList = products.data;
    },

    // async ProductList() {
    //   const shopID = this.shopId;
    //   const products = await axios.get(
    //     `http://localhost:8000/shops/viewproducts?shop_id=${shopID}`
    //   );
    //   console.log(products.data);
    //   this.productList = products.data;
    // },

  },
};
</script>
  
  <style>
.product-gallery {
  display: flex;
  flex-direction: column;
}

.thumbnails {
  display: flex;
  justify-content: space-between;
}

.thumbnail {
  cursor: pointer;
}
</style>