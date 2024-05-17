import { defineStore } from 'pinia'

export const useCartStore = defineStore('cart', {
  state: () => ({
    Products: [
        {
        id:1,
        name: "Iphone ",
        price:800+"$",
        image:''
       },
       {
        id:2,
        name: "samsung s10 ",
        price:600+"$",
        image:''
       },
       {
        id:2,
        name: "Nokia x15 ",
        price:500+"$",
        image:''
       },
]
  }),
  actions: {
    addToCart(item) {
      this.items.push(item)
    },
    removeFromCart(index) {
      this.items.splice(index, 1)
    },
    clearCart() {
      this.items = []
    }
  },
  getters: {
    cartTotal() {
      return this.items.reduce((total, item) => total + item.price, 0)
    }
  }
})
