<<<<<<< HEAD
=======
<!-- <script setup lang="ts">
import { useTableData } from '../../composables/useTableData'

const {
  simpleTableData,
  paginatedTableData,
  wideTableData,
} = useTableData()
</script>

>>>>>>> cad65aee50fdeb782aa2c6e5a8e7dd0968f06047
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
<<<<<<< HEAD

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
=======
 -->

 <template>
  <div class="text-xl">
    <vee-form :validation-schema="productschema" @submit="addProduct"   class="shadow-2xl py-9 px-9 bg-slate-50 font-style rounded">
          <p class="text-5xl flex justify-center">Add Product</p>
          <br>
          
    <!-- Name -->
    <div class="mb-3">
      <label class="inline-block mb-2">Product Name</label>
      <vee-field
        type="text"
        name="product_name"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter Product Name"
      />
      <ErrorMessage class="text-red-600" name="product_name" />
    </div>

    <div class="mb-3">
      <label class="inline-block mb-2">Product Description</label>
      <vee-field
        type="text"
        name="product_description"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter Product Description"
      />
      <ErrorMessage class="text-red-600" name="product_description" />
    </div>

    <div class="mb-3">
      <label class="inline-block mb-2">Product Price</label>
      <vee-field
        type="number"
        name="price"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter Price"
      />
      <ErrorMessage class="text-red-600" name="price" />
    </div>

    <div class="mb-3">
      <label class="inline-block mb-2">Product Quantity</label>
      <vee-field
        type="number"
        name="quantity"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter Quantity"
      />
      <ErrorMessage class="text-red-600" name="quantity" />
    </div>
    <!-- Email -->
   

    <div class="mb-3">
      <label class="inline-block mb-2">Product Image</label>
      
      <vee-field
        name="file"
        type="file">

      </vee-field>
    </div>



    <button
      type="submit"
      class="block w-full bg-black text-white py-1.5 px-3 rounded-full transition hover:bg-purple-700 cursor-pointer"
      :disabled="reg_in_submission"
    >
      Create my Shop
    </button>
         </vee-form>
  </div>
</template>
<script>
import axios from 'axios';

export default {
  data() {
    return {
      owner_id: localStorage.getItem('owner_id'),
      reg_in_submission: false,
      reg_alert_message: 'Please Wait! Your Account is being created',
      reg_alert_variant: 'bg-blue-500',
      reg_show_alert: false,

      selectedFile: null,
      selectedImage: "",

    
      productschema: {
        product_name: 'required|min:3|max:100|alpha_spaces',
        product_description: 'required|min:3|max:100',
        price: 'required|min:0',
        quantity: 'required|min_value:1',
      },
    };
  },

  methods: {
    add() {
     console.log("Hi")
    },
    async addProduct(formdata) {
      console.log(formdata)
      const ownerID = this.owner_id
      try {
       const response = await axios.post(`http://localhost:8000/owners/addproducttoshop?owner_id=${ownerID}`, formdata, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        })
        console.log(response)
        this.$toast.success('Successfully Added Product.')
      } catch(error) {
        console.log(error)
        this.$toast.error('Adding Product failed, please try again in minute')
      }

    }
  },
}
</script>
<style>
   @import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");

.font-style {
  font-family: "Marcellus", sans-serif;
}
</style>
>>>>>>> cad65aee50fdeb782aa2c6e5a8e7dd0968f06047
