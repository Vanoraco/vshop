<template>
    <div>
      <h1>Shopping Cart</h1>
      <div v-if="cart.length === 0">Your cart is empty.</div>
      <div v-else>
        <div v-for="(item, index) in cart" :key="index" class="cart-item">
          <div>{{ item.name }}</div>
          <div>{{ item.price }}</div>
          <div>
            <button @click="removeItem(index)">Remove</button>
          </div>
        </div>
        <div>Total: {{ total }}</div>
        <button @click="checkout">Checkout</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        cart: [],
      };
    },
    computed: {
      total() {
        return this.cart.reduce((acc, item) => acc + item.price, 0);
      },
    },
    methods: {
      removeItem(index) {
        this.cart.splice(index, 1);
      },
      checkout() {
        // Implement checkout logic
      },
    },
  };
  </script>
  
  <style scoped>
  .cart-item {
    display: flex;
    justify-content: space-between;
    margin-bottom: 10px;
  }
  </style>
  