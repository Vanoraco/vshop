<template>
  <div class="flex justify-center gap-9 font-style text-2xl font-style">
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

export default {
  data() {
    return {
      reg_show_alert: false,
      reg_alert_variant: "",
      reg_alert_message: "",
      reg_in_submission: false,
      isLoggedIn: localStorage.getItem("email") !== null && localStorage.getItem("token") !== null,

    };
  },
  methods: {
    async loginUser(formdata) {
      this.reg_show_alert = true;
      this.reg_alert_variant = "bg-blue-500";
      this.reg_alert_message = "Please Wait! Logging In";

      try {
        const response = await axios.post("http://localhost:8000/users/login", formdata);
        localStorage.setItem("token", response.data.token);
        localStorage.setItem("email", response.data.email);
        localStorage.setItem("firstname", response.data.first_name);
        this.$router.push("/");
        this.isLoggedIn=false;
      } catch (error) {
        this.reg_in_submission = false;
        this.reg_alert_variant = "bg-red-500";
        this.reg_alert_message = "An unexpected error occurred. Please try again later";
        console.error("Error logging in:", error);
      }
    },
  },
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");

.font-style {
  font-family: "Marcellus", sans-serif;
}
</style>
