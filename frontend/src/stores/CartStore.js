import { defineStore } from 'pinia';

export const useCartStore = defineStore('cart', {
  state: () => ({
    items: [],
  }),
  actions: {
    LoadCartInstance() {
      const cs = localStorage.getItem('cart');
      if (!cs) {
        this.items = [];
      } else {
        this.items = JSON.parse(cs);
      }
    },
    addToCart(item) {
      const existingItem = this.items.find(i => i.Product_ID === item.Product_ID);
      if (existingItem) {
        existingItem.quantity += 1;
      } else {
        this.items.push({ ...item, quantity: 1 });
      }
      this.saveCart();
    },
    incrementItemQuantity(itemId) {
      const item = this.items.find(i => i.Product_ID === itemId);
      if (item) {
        item.quantity += 1;
        this.saveCart();
      }
    },
    decrementItemQuantity(itemId) {
      const item = this.items.find(i => i.Product_ID === itemId);
      if (item && item.quantity > 1) {
        item.quantity -= 1;
        this.saveCart();
      }
    },
    removeFromCart(itemId) {
      this.items = this.items.filter(item => item.Product_ID !== itemId);
      this.saveCart();
    },
    clearCart() {
      this.items = [];
      this.saveCart();
    },
    saveCart() {
      localStorage.setItem('cart', JSON.stringify(this.items));
    },
  },
  getters: {
    cartTotal() {
      return this.items.reduce((total, item) => total + (item.price * item.quantity), 0);
    },
  },
});
