<template>
    <div class="favorites-container">
      <h1 class="favorites-title">Favorite Items</h1>
      <div v-if="favorites.length === 0" class="empty-favorites-message">
        You have no favorite items.
      </div>
      <div v-else class="favorites-items-container">
        <div v-for="(item, index) in favorites" :key="item.id" class="favorites-item-wrapper">
          <div class="favorites-item-details">
            <img :src="item.image" alt="Product Image" class="product-image" />
            <div class="details">
              <div class="product-name">{{ item.product_name }}</div>
              <div class="shop-name">price: {{ item.price }}</div>
              <div class="date-added">Added on: {{ formatDate(item.date_added) }}</div>
            </div>
          </div>
          <div class="favorites-item-actions">
            <button class="remove-item-btn" @click="removeItem(item.id)">
              Remove
            </button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import { useCartStore } from "../stores/CartStore"; 
  import { format } from "date-fns"; 
  
  export default {
    computed: {
      favorites() {
        const favoritesStore = useCartStore();
        return favoritesStore.items;
      },
    },
    methods: {
      removeItem(itemId) {
        const favoritesStore = useCartStore();
        favoritesStore.removeFromCart(itemId);
      },
      formatDate(date) {
        
      }
    },
  };
  </script>
  
  <style scoped>
  .favorites-container {
    max-width: 800px;
    margin: 20px auto;
    padding: 20px;
    background-color: #fff;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
  }
  
  .favorites-title {
    font-size: 24px;
    font-weight: 600;
    color: #333;
    margin-bottom: 20px;
    text-align: center;
    border-bottom: 2px solid #e0e0e0;
    padding-bottom: 10px;
  }
  
  .empty-favorites-message {
    font-size: 16px;
    color: #888;
    text-align: center;
    margin-top: 20px;
  }
  
  .favorites-items-container {
    margin-top: 20px;
  }
  
  .favorites-item-wrapper {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 15px 0;
    border-bottom: 1px solid #ddd;
  }
  
  .favorites-item-details {
    display: flex;
    align-items: center;
  }
  
  .product-image {
    width: 50px;
    height: 50px;
    object-fit: cover;
    border-radius: 8px;
    margin-right: 15px;
  }
  
  .details {
    display: flex;
    flex-direction: column;
  }
  
  .product-name {
    font-size: 16px;
    font-weight: 600;
    color: #333;
  }
  
  .shop-name, .date-added {
    font-size: 14px;
    color: #888;
    margin-top: 5px;
  }
  
  .favorites-item-actions {
    display: flex;
    align-items: center;
  }
  
  .remove-item-btn {
    background-color: #ff4d4d;
    color: #fff;
    border: none;
    padding: 8px 12px;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s ease, transform 0.3s ease;
  }
  
  .remove-item-btn:hover {
    background-color: #e60000;
    transform: scale(1.05);
  }
  
  @media (max-width: 600px) {
    .favorites-item-wrapper {
      flex-direction: column;
      align-items: flex-start;
    }
  
    .favorites-item-actions {
      margin-top: 10px;
    }
  }
  </style>
  