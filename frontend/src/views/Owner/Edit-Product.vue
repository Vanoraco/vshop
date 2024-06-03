<template>
    <div class="text-xl">
      <vee-form :validation-schema="productschema" @submit="editProduct"   class="shadow-2xl py-9 px-9 bg-slate-50 font-style rounded">
            <p class="text-5xl flex justify-center">Edit Product</p>
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
        shopId: this.$route.params.id,
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
      async editProduct(formdata) {
        console.log(formdata)
        const ownerID = this.owner_id
        try {
         const response = await axios.put(`http://localhost:8000/products/edit?shop_id=${ownerID}&product_id=${this.shopId}`, formdata, {
            headers: {
              'Content-Type': 'multipart/form-data',
            },
          })
          console.log(response)
          this.$toast.success('Successfully Edited Product.')
          this.$router.push('/dash')
        } catch(error) {
          console.log(error)
          this.$toast.error('Editing Product failed, please try again in minute')
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