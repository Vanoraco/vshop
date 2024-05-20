<template>
  <div class="flex flex-row items-center mt-6 text-6xl">
    <div class="ml-4">
      <img src="../assets/images/logo.png" alt="Logo" class="h-12" />
    </div>
    <div class="font-serif ml-8">
      <ul class="flex flex-row gap-8 text-lg">
        <router-link to="/">
          <li :class="{ 'text-xl': $route.path === '/' }">Home</li>
        </router-link>
        <router-link to="/shops">
          <li :class="{ 'text-xl': $route.path === '/shops' }">Shops</li>
        </router-link>
        <router-link to="/sign-up" v-if="!isLoggedIn">
          <li :class="{ 'text-xl': $route.path === '/sign-up' }">Sign Up</li>
        </router-link>
        <div class="ml-auto">
      <router-link to="/cart">
        <i class="fas fa-shopping-cart text-gray-800"></i>
      </router-link>
    </div>
     
      </ul>
    </div>

    <div class="flex-grow"></div>
    <div class="flex rounded-md border-2 border-blue-500 overflow-hidden max-w-md mx-auto font-[sans-serif] mr-9">
      <input
        type="text"
        placeholder="Search..."
        class="w-full outline-none bg-white text-gray-600 text-sm px-4 py-3"
      />
      <button type="button" class="flex items-center justify-center bg-[#007bff] px-5">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 192.904 192.904"
          width="16px"
          class="fill-white"
        >
          <path
            d="m190.707 180.101-47.078-47.077c11.702-14.072 18.752-32.142 18.752-51.831C162.381 36.423 125.959 0 81.191 0 36.422 0 0 36.423 0 81.193c0 44.767 36.422 81.187 81.191 81.187 19.688 0 37.759-7.049 51.831-18.751l47.079 47.078a7.474 7.474 0 0 0 5.303 2.197 7.498 7.498 0 0 0 5.303-12.803zM15 81.193C15 44.694 44.693 15 81.191 15c36.497 0 66.189 29.694 66.189 66.193 0 36.496-29.692 66.187-66.189 66.187C44.693 147.38 15 117.689 15 81.193z"
          ></path>
        </svg>
      </button>
    </div>
    <ul class="flex flex-row gap-9 mb-6 mr-4">
      <li v-if="!isLoggedIn">
        <router-link to="/login">
          <i class="pi pi-user text-2xl font-style hover:bg-black hover:text-white hover:rouned hover:px-3 hover:py-1">
            Login
          </i>
        </router-link>
      </li>
      <li v-if="isLoggedIn">
        <router-link to="/user/profile">
          <i class="pi pi-user-edit text-xl">Update Profile</i>
        </router-link>
      </li>
      <li v-if="isLoggedIn">
        <router-link @click="logout">
          <i class="pi pi-sign-out text-xl">Sign out</i>
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "AppHeader",
  data() {
    return {
      isLoggedIn: localStorage.getItem("email") !== null,
    };
  },
  methods: {
    logout() {
      try {
        localStorage.removeItem("email");
        localStorage.removeItem("token");
        localStorage.removeItem("firstname");
        this.isLoggedIn = false;
        this.$router.push("/");
      } catch (error) {
        console.error("Error logging out:", error);
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
