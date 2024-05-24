<template>
    <div class="flex flex-col md:flex-row justify-between items-start mt-8">
      <div class="w-full md:w-1/2 mr-8">
        <div class="product-gallery">
          <div class="main-image">
            <img :src="product.image" :alt="product.name" class="w-full  rounded-lg" />
          </div>
          <div class="thumbnails flex justify-between mt-4">
            <div v-for="image in product.images" :key="image" class="thumbnail w-1/4 mr-2">
              <img :src="image" :alt="product.name" class="w-full h-auto rounded-lg" />
            </div>
          </div>
        </div>
      </div>
  
      <div class="w-full md:w-1/2">
        <h1 class="text-3xl font-bold mb-2">{{ product.name }}</h1>
        <div class="flex items-center mb-4">
          <div class="rating text-yellow-400 mr-2">
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star"></i>
            <i class="fas fa-star-half-alt"></i>
          </div>
          <span class="text-gray-500">({{ product.reviews.length }} reviews)</span>
        </div>
        <p class="text-2xl font-bold mb-4">ETB {{ product.price }}</p>
        <p class="mb-4">{{ product.description }}</p>
        <div class="flex items-center mb-4">
          <button class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded mr-4">
            Buy Now
          </button>
          <button class="bg-gray-200 hover:bg-gray-300 text-gray-700 font-bold py-2 px-4 rounded">
            Add to Cart
          </button>
        </div>
        <div class="product-3d-view">
          <h2 class="text-xl font-bold mb-2">3D View</h2>
          <iframe
            :src="product.threeDUrl"
            frameborder="0"
            allowfullscreen
            class="w-full h-64 rounded-lg"
          ></iframe>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
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
          threeDUrl: "https://github.com/Nathyhac/3D-models/phone.glb",
          reviews: [
            { rating: 5, comment: "Great product!" },
            { rating: 4, comment: "Good product, but could be better." },
            { rating: 3, comment: "Average product." },
          ],
        },
      };
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