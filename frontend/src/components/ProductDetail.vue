<template >
  <div class="flex flex-col md:flex-row justify-between items-start max-w-[96vw] mt-8 font-style">
    <div class="w-full md:w-1/2 mr-8">
      <div class="product-gallery">
        <div class="main-image h-[590px] w-full p- px-3 overflow-hidden shadow-2xl rounded-lg bg-clip-padding backdrop-filter backdrop-blur-md">
          <img
            :src="productList.image"
            :alt="productList.name"
            class="w-full rounded-lg object-fill h-[500px] border-black "
          />
        </div>
        <div class="thumbnails flex justify-between px-8 mt-4">
          <div
            v-for="image in product.images"
            :key="image"
            class="thumbnail w-1/4 mr-2 shadow-xl"
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
        
        <button @click="initializeTransaction"
          class="bg-gradient-to-tr from-teal-600 via-cyan-700 to-blue-800 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded mr-4 hover:text-xl hover:shadow-xl "
        >
          Buy Now 
        </button>
      
        <button  @click="addToCart(productList)" v-if="userLoggedIn"
          class="mr-2 bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded shadow-lg transform hover:scale-105 transition-transform duration-300 ease-in-out"
        >
          Add to cart
        </button>
      </div>
      <div class="product-3d-view">
        <h2 class="text-xl font-bold mb-2">3D View</h2>
        <img
          src="../assets//images/ecommerce.jpg"
          frameborder="0"
          allowfullscreen
          class="w-full h-64 rounded-lg"
        />
      </div>
    </div>
  </div>
</template>
  
  <script>
import axios from 'axios';
import { useCartStore } from "../stores/CartStore";
import  useUserStore  from "../stores/user"
import { mapState } from 'pinia';

export default {
  mounted() {
    this.ProductList();
  },

  data() {
    return {
      shopId: this.$route.params.id,
      prodId: this.$route.params.proid,
      paymentRequest: {
  "amount": "",
  "currency": "",
  "email": "",
  "first_name": "",
  "last_name": "",
  "phone_number": "",
  "tx_ref": new Date().getMinutes,
  "callback_url": "https://webhook.site/077164d6-29cb-40df-ba29-8a00e59a7e60",
  "return_url": "https://www.google.com/",
  "customization[title]": "Payment for my favourite merchant",
  "customization[description]": "I love online payments"
},
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
  computed: {
    ...mapState(useUserStore, ['userLoggedIn']),
    
    },
  methods: {
    async ProductList() {
      const shopID = this.shopId;
      const products = await axios.get(
        `http://localhost:8000/shop/product?shop_id=${shopID}&product_id=${this.prodId}`
      );
      console.log(products.data);
      this.productList = products.data;
      this.paymentRequest.amount = String( products.data.price); // Assign price to amount
        this.paymentRequest.currency = "ETB"; // Set currency
        this.paymentRequest.email = localStorage.getItem("email"); // Set email (replace with actual logic to get user's email)
        this.paymentRequest.first_name = localStorage.getItem('firstname'); // Set first name (replace with actual logic to get user's first name)
        this.paymentRequest.last_name = localStorage.getItem('firstname'); // Set last name (replace with actual logic to get user's last name)
        this.paymentRequest.phone_number = localStorage.getItem('phone'); // Set phone number (replace with actual logic to get user's phone number)
        this.paymentRequest.tx_ref = `tx-${Date.now()}`;
    },
    async initializeTransaction() {
        try {
            console.log(this.paymentRequest)
          const response = await axios.post("http://localhost:8000/intialize-chapa", this.paymentRequest, {
            headers: {
              'Content-Type': 'application/json'
            }
          });
          
          const data = response.data; 
          const responseData = JSON.parse(response.data.response);
          console.log(responseData.data)
          // Correctly access the data field of the response
          if (data && responseData.data.checkout_url) {
            this.checkoutUrl = data.checkout_url;
            window.location.href = responseData.data.checkout_url
          } else {
            console.error("Checkout URL not found in the response:", data);
          }
        } catch (error) {
          console.error("Error initializing transaction:", error);
        }
      },
    addToCart(product) {
      const cartStore = useCartStore();
      cartStore.addToCart(product);
    },
  },
};
</script>
  
  <style>
  @import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");

.font-style {
  font-family: "Marcellus", sans-serif;
}

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