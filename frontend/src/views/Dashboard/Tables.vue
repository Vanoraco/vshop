<template>
  <div class="max-w-3xl mx-auto p-6 bg-white shadow-md rounded-lg">
    <h1 class="text-2xl font-semibold mb-6">Add New Product</h1>
    <form @submit.prevent="submitForm">
      <!-- Product Name -->
      <div class="mb-4">
        <label class="block text-gray-700">Product Name</label>
        <input
          type="text"
          v-model="product.name"
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
          required
        />
      </div>

      <!-- Description -->
      <div class="mb-4">
        <label class="block text-gray-700">Description</label>
        <textarea
          v-model="product.description"
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
          rows="4"
          required
        ></textarea>
      </div>

      <!-- Price -->
      <div class="mb-4">
        <label class="block text-gray-700">Price</label>
        <input
          type="number"
          v-model="product.price"
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
          required
        />
      </div>

      <!-- Stock -->
      <div class="mb-4">
        <label class="block text-gray-700">Available in Stock</label>
        <input
          type="number"
          v-model="product.stock"
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
          required
        />
      </div>

      <!-- Images -->
      <div class="mb-4">
        <label class="block text-gray-700">Images</label>
        <input
          type="file"
          @change="handleImageUpload"
          multiple
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md"
        />
        <div class="mt-2 grid grid-cols-3 gap-2">
          <div
            v-for="(image, index) in product.images"
            :key="index"
            class="relative"
          >
            <img
              :src="image.url"
              class="w-full h-24 object-cover rounded-md"
            />
            <button
              type="button"
              @click="removeImage(index)"
              class="absolute top-0 right-0 bg-red-500 text-white p-1 rounded-full"
            >
              &times;
            </button>
          </div>
        </div>
      </div>

      <!-- Submit Button -->
      <div class="flex justify-end">
        <button
          type="submit"
          class="px-6 py-2 bg-blue-500 text-white font-semibold rounded-md hover:bg-blue-600"
        >
          Add Product
        </button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      product: {
        name: '',
        description: '',
        price: 0,
        stock: 0,
        images: [],
      },
    };
  },
  methods: {
    handleImageUpload(event) {
      const files = event.target.files;
      for (let i = 0; i < files.length; i++) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.product.images.push({ url: e.target.result });
        };
        reader.readAsDataURL(files[i]);
      }
    },
    removeImage(index) {
      this.product.images.splice(index, 1);
    },
    submitForm() {
      // Handle form submission, e.g., send the data to a server
      console.log('Product submitted:', this.product);
    },
  },
};
</script>

<style scoped>
/* Add custom styles here if needed */
</style>
