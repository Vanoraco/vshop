<template>
    <div class="product-detail">
      <div class="product-image">
        <img :src="product.imageUrl" :alt="product.name" />
      </div>
      <div class="product-info">
        <h1>{{ product.name }}</h1>
        <p class="product-price">{{ product.price }}</p>
        <p>{{ product.description }}</p>
        <button @click="addToCart">Add to Cart</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        product: {
          name: "Product Name",
          price: "$19.99",
          description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          imageUrl: "product-image-url.jpg"
        }
      };
    },
    methods: {
      
      addToCart() {
        // Add product to cart logic
        console.log("Product added to cart!");
      }
    }
  };
  </script>
  
  <style>
  .product-detail {
    display: flex;
    justify-content: space-around;
    align-items: center;
    margin: 20px;
  }
  
  .product-image img {
    max-width: 400px;
    max-height: 400px;
  }
  
  .product-info {
    flex: 1;
    margin-left: 20px;
  }
  
  .product-price {
    font-size: 1.5em;
    color: #B12704;
  }
  
  button {
    padding: 10px 20px;
    background-color: #FF9900;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 1em;
  }
  
  button:hover {
    background-color: #FFA500;
  }
  </style>
  