<template>
  <div class="fixed inset-0 z-50 flex items-center justify-center">
    <div class="bg-white rounded-lg shadow-lg p-8">
      <button
        class="bg-gradient-to-r from-purple-500 via-pink-500 to-red-500 text-white font-bold py-2 px-4 rounded-full shadow-md mb-4"
        @click="$emit('closeForm')"
      >
        Contact Admins for 3D View
      </button>
      <div class="mb-4">
        <label for="productName" class="block font-bold mb-2">Product Name</label>
        <input
          id="productName"
          type="text"
          class="w-full border border-gray-300 rounded-md px-4 py-2"
          placeholder="Enter product name"
          v-model="productName"
        />
      </div>
      <div class="mb-4">
        <label for="productPrice" class="block font-bold mb-2">Price</label>
        <div class="flex items-center">
          <span class="mr-2 text-gray-500">ETB</span>
          <input
            id="productPrice"
            type="number"
            class="w-full border border-gray-300 rounded-md px-4 py-2"
            placeholder="Enter price"
            v-model="productPrice"
          />
        </div>
      </div>
      <div class="mb-4">
        <label for="productImages" class="block font-bold mb-2">Product Images</label>
        <input
          id="productImages"
          type="file"
          class="w-full border border-gray-300 rounded-md px-4 py-2"
          @change="handleImageUpload"
          multiple
        />
        <div class="mt-4 grid grid-cols-4 gap-4">
          <div v-for="(image, index) in productImages" :key="index" class="relative">
            <img :src="image" class="w-full h-auto rounded-md" />
            <button
              class="absolute top-0 right-0 m-2 bg-red-500 hover:bg-red-600 text-white font-bold py-1 px-2 rounded-full"
              @click="removeImage(index)"
            >
              &times;
            </button>
          </div>
        </div>
      </div>
      <div class="mb-4">
        <label for="productQuantity" class="block font-bold mb-2">Quantity</label>
        <input
          id="productQuantity"
          type="number"
          class="w-full border border-gray-300 rounded-md px-4 py-2"
          placeholder="Enter quantity"
          v-model="productQuantity"
        />
      </div>
      <button
        class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-md"
        @click="addProduct"
      >
        Add Product
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AddProductForm',
  data() {
    return {
      productName: '',
      productPrice: null,
      productImages: [],
      productQuantity: null,
    };
  },
  methods: {
    handleImageUpload(event) {
      const files = event.target.files;
      const imagesArray = Array.from(files);
      const previews = imagesArray.map((file) => URL.createObjectURL(file));
      this.productImages.push(...previews);
    },
    removeImage(index) {
      this.productImages.splice(index, 1);
    },
    addProduct() {
      const productData = {
        name: this.productName,
        price: this.productPrice,
        images: this.productImages,
        quantity: this.productQuantity,
      };

      // Here, you can emit an event or make an API call to add the product
      this.$emit('addProduct', productData);

      // Reset form fields after adding the product
      this.productName = '';
      this.productPrice = null;
      this.productImages = [];
      this.productQuantity = null;
    },
  },
};
</script>