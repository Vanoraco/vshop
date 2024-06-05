<template>
  <div class="shopping-cart-container">
    <h1 class="shopping-cart-title">Shopping Cart</h1>
    <div v-if="cart.length === 0" class="empty-cart-message">
      Your cart is empty.
    </div>
    <div v-else class="cart-items-container">
      <div v-for="(item, index) in cart" :key="item.Product_ID" class="cart-item-wrapper">
        <div class="cart-item-details">
          <img :src="item.image" alt="Product Image" class="product-image" />
          <div class="details">
            <div class="product-name">{{ item.product_name }}</div>
            <div class="product-price-quantity">
              <span class="product-price">ETB {{ item.price }}</span>
              <div class="quantity-controls">
                <button @click="decrementItem(item.Product_ID)">-</button>
                <span class="product-quantity">{{ item.quantity }}</span>
                <button @click="incrementItem(item.Product_ID)">+</button>
              </div>
            </div>
          </div>
        </div>
        <div class="cart-item-actions">
          <button class="remove-item-btn" @click="removeItem(item.Product_ID)">
            Remove
          </button>
        </div>
      </div>
      <div class="cart-total">
        <span class="total-label">Total:</span>
        <span class="total-amount">ETB {{ total }}</span>
      </div>
      <div class="cart-actions">
        <router-link to="/checkout">
          <button class="checkout-btn">Checkout</button>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { useCartStore } from "../stores/CartStore"; // Adjust the path to your store

export default {
  computed: {
    cart() {
      const cartStore = useCartStore();
      return cartStore.items;
    },
    total() {
      const cartStore = useCartStore();
      return cartStore.cartTotal;
    },
  },
  methods: {
    removeItem(itemId) {
      const cartStore = useCartStore();
      cartStore.removeFromCart(itemId);
    },
    incrementItem(itemId) {
      const cartStore = useCartStore();
      cartStore.incrementItemQuantity(itemId);
    },
    decrementItem(itemId) {
      const cartStore = useCartStore();
      cartStore.decrementItemQuantity(itemId);
    },
  },
};
</script>

<style scoped>
.shopping-cart-container {
  max-width: 800px;
  margin: 20px auto;
  padding: 20px;
  background-color: #fff;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
}

.shopping-cart-title {
  font-size: 24px;
  font-weight: 600;
  color: #333;
  margin-bottom: 20px;
  text-align: center;
  border-bottom: 2px solid #e0e0e0;
  padding-bottom: 10px;
}

.empty-cart-message {
  font-size: 16px;
  color: #888;
  text-align: center;
  margin-top: 20px;
}

.cart-items-container {
  margin-top: 20px;
}

.cart-item-wrapper {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px 0;
  border-bottom: 1px solid #ddd;
}

.cart-item-details {
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

.product-price-quantity {
  margin-top: 5px;
  font-size: 14px;
  color: #888;
  display: flex;
  align-items: center;
}

.quantity-controls {
  display: flex;
  align-items: center;
}

.quantity-controls button {
  background-color: #f0f0f0;
  border: none;
  padding: 5px 10px;
  margin: 0 5px;
  cursor: pointer;
}

.cart-item-actions {
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

.cart-total {
  display: flex;
  justify-content: flex-end;
  margin-top: 20px;
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.total-label {
  margin-right: 10px;
}

.cart-actions {
  display: flex;
  justify-content: flex-end;
  margin-top: 20px;
}

.checkout-btn {
  background-color: #4caf50;
  color: #fff;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease, transform 0.3s ease;
}

.checkout-btn:hover {
  background-color: #45a049;
  transform: scale(1.05);
}

@media (max-width: 600px) {
  .cart-item-wrapper {
    flex-direction: column;
    align-items: flex-start;
  }

  .cart-item-actions {
    margin-top: 10px;
  }

  .checkout-btn {
    width: 100%;
    text-align: center;
  }
}
</style>
