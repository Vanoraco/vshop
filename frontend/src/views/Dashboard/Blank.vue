<template>
  <div class="max-w-7xl mx-auto p-6">
    <h1 class="text-3xl font-semibold mb-6">My Products</h1>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="product in products"
        :key="product.id"
        class="bg-white shadow-md rounded-lg overflow-hidden"
      >
        <img
          :src="product.images[0]"
          alt="Product Image"
          class="w-full h-48 object-cover"
        />
        <div class="p-4">
          <h2 class="text-xl font-semibold">{{ product.name }}</h2>
          <p class="text-gray-700 mt-2">{{ product.description }}</p>
          <div class="mt-4 flex justify-between items-center">
            <span class="text-lg font-bold text-blue-600">${{ product.price }}</span>
            <span
              class="text-sm font-medium text-gray-600"
              :class="{
                'text-green-600': product.stock > 0,
                'text-red-600': product.stock === 0,
              }"
            >
              {{ product.stock > 0 ? 'In Stock' : 'Out of Stock' }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      products: [
        {
          id: 1,
          name: 'Product 1',
          description: 'This is a description of product 1.',
          price: 29.99,
          stock: 10,
          images: [
            'https://via.placeholder.com/400x300',
          ],
        },
        {
          id: 2,
          name: 'Product 2',
          description: 'This is a description of product 2.',
          price: 49.99,
          stock: 5,
          images: [
            'https://via.placeholder.com/400x300',
          ],
        },
        {
          id: 3,
          name: 'Product 3',
          description: 'This is a description of product 3.',
          price: 19.99,
          stock: 0,
          images: [
            'https://via.placeholder.com/400x300',
          ],
        },
        // Add more products as needed
      ],
    };
  },
};
</script>

<style scoped>
/* Add custom styles here if needed */
</style>

