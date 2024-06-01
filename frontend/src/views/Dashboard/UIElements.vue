<template>
  <div class="text-xl">
    <vee-form :validation-schema="shopschema" @submit="addShop"   class="shadow-2xl py-9 px-9 bg-slate-50 font-style rounded">
          <p class="text-5xl flex justify-center">Create your shop</p>
          <br>
          
    <!-- Name -->
    <div class="mb-3">
      <label class="inline-block mb-2">Shop Name</label>
      <vee-field
        type="text"
        name="shop_name"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter Shop Name"
      />
      <ErrorMessage class="text-red-600" name="shop_name" />
    </div>

    <div class="mb-3">
      <label class="inline-block mb-2">Shop Category</label>
     <vee-field name="shop_category" as="select" class="py-3 px-4 pe-9 block w-full border shadow-2xl border-black rounded-lg text-lg">
        <option value="cosmetics">Cosmetics</option>
        <option value="electronics">Electronics</option>
        <option value="shoe">Shoes/Clothes</option>
      </vee-field>
      
    </div>
    <!-- Email -->
   

    <div class="mb-3">
      <label class="inline-block mb-2">Shop Image</label>
      
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

    
      shopschema: {
        shop_name: 'required|min:3|max:100|alpha_spaces',
        
      },
    };
  },

  methods: {
    add() {
     console.log("Hi")
    },
    async addShop(formdata) {
      console.log(formdata)
      const ownerID = this.owner_id
      try {
       const response = await axios.post(`http://localhost:8000/owners/addshop?owner_id=${ownerID}`, formdata, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        })
        console.log(response)
        this.$toast.success('Successfully Created.')
      } catch(error) {
        console.log(error)
        this.$toast.error('Creating Shop failed, please try again in minute')
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