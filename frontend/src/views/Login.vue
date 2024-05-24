<template>
  <div class="flex justify-center gap-9 font-style text-2xl font-style" v-if="!signOwner">
    <vee-form @submit="loginUser" class="shadow-2xl py-9 px-9">
      <p class="text-5xl">Sign in to buy products</p>
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
          name="email"
          class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
          placeholder="Enter Email"
        />
        <ErrorMessage class="text-red-600" name="email" />
      </div>

      <!-- Password -->
      <div class="mb-3">
        <label class="inline-block mb-2">Password</label>
        <vee-field name="password" :bails="false" v-slot="{ field, errors }">
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
        Are you a shop owner? <p class="ml-2 cursor-pointer text-blue-500"> Login as owner</p>
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
    <vee-form @submit="login" class="shadow-2xl py-9 px-9" v-if="signOwner">
      <p class="text-5xl">Sign in to Manage your shop</p>
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
          name="email"
          class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
          placeholder="Enter Email"
        />
        <ErrorMessage class="text-red-600" name="email" />
      </div>

      <!-- Password -->
      <div class="mb-3">
        <label class="inline-block mb-2">Password</label>
        <vee-field name="password" :bails="false" v-slot="{ field, errors }">
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
        Are you a shop owner? <a class="ml-2 cursor-pointer text-blue-500" @click="signAsOwner"> Login as owner</a>
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
import axios from "axios";

import { mapState } from "pinia";
import { mapActions } from "pinia";
import useUserStore from "../stores/user"

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
    ...mapActions(useUserStore, ['login']),
    async loginUser(formdata) {
      this.reg_show_alert = true;
      this.reg_alert_variant = "bg-blue-500";
      this.reg_alert_message = "Please Wait! Logging In";

      try {
         await this.login(formdata)
        this.$router.push("/");
        this.isLoggedIn=false;
      } catch (error) {
        this.reg_in_submission = false;
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

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");

.font-style {
  font-family: "Marcellus", sans-serif;
}
</style>