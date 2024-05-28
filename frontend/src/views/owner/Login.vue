<template>
    <div class="flex justify-center gap-9 font-style text-2xl font-style">
      <vee-form @submit="loginAsOwner" class="shadow-2xl py-9 px-9">
        <p class="text-5xl">Sign in to sell Products</p>
        <br />
        <div
          class="text-white font-bold p-4 rounded mb-4"
          v-if="reg_show_alert"
          :class="reg_alert_variant"
        >
          {{ reg_alert_message }}
        </div>
  
        <!-- Email -->
        <div class="mb-3">
          <label class="inline-block mb-2">Email</label>
          <vee-field
            type="email"
            name="owner_email"
            class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
            placeholder="Enter Email"
          />
          <ErrorMessage class="text-red-600" name="owner_email" />
        </div>
  
        <!-- Password -->
        <div class="mb-3">
          <label class="inline-block mb-2">Password</label>
          <vee-field name="owner_password" :bails="false" v-slot="{ field, errors }">
            <input
              type="password"
              class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
              placeholder="Password"
              v-bind="field"
            />
            <div class="text-red-600" v-for="err in errors" :key="err">
              {{ err }}
            </div>
          </vee-field>
        </div>
  
        <button
          type="submit"
          class="block w-full bg-black text-white py-1.5 px-3 rounded transition hover:bg-purple-700"
          :disabled="reg_in_submission"
        >
          Login
        </button>
   
        <div class="flex justify-center mt-12">
        <router-link to="login" class="ml-2 cursor-pointer text-green-500 shadow-xl rounded hover:text-3xl hover:bg-gradient-to-b from-lime-600 to-lime-800 hover:text-white">I am a buyer</router-link>
      </div>
  

        <div class="flex justify-center mt-24">
          ---- Don't have an account ? ----
        </div>
        <div>
          <router-link to="/sign-up" class="flex justify-center cursor-pointer bg-green rounded">
            Sign Up
          </router-link>
        </div>
      </vee-form>
     
        
    </div>
  </template>

  <script>

import { mapActions } from "pinia";
import useUserStore from "../../stores/user"

    export default {
  data() {
    return {
      reg_show_alert: false,
      reg_alert_variant: "",
      reg_alert_message: "",
      reg_in_submission: false,
      signOwner: false
    };
  },
  computed: {
    
    
  },
  methods: {
    ...mapActions(useUserStore, ['loginOwner']),
    async loginAsOwner(formdata) {
      this.reg_show_alert = true;
      this.reg_alert_variant = "bg-blue-500";
      this.reg_alert_message = "Please Wait! Logging In";

      try {
         await this.loginOwner(formdata)
         this.$toast.success('Successfully Logged In.')
        this.$router.push("/owner");
        this.isLoggedIn=false;
      } catch (error) {
        this.reg_in_submission = false;
        this.$toast.error('Please Try again')
        this.reg_alert_variant = "bg-red-500";
        this.reg_alert_message = "An unexpected error occurred. Please try again later";
        console.error("Error logging in:", error);
      }
      
    },
    signAsOwner() {
      this.signOwner = true
      this.$router.push('/shops')
      console.log(this.signOwner)
    }
  },
};
 </script>

 <style>
  @import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");

.font-style {
  font-family: "Marcellus", sans-serif;
}
</style>