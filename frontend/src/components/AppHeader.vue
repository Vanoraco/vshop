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

        <router-link to="/sign-up">
          <li :class="{ 'text-xl': $route.path === '/sign-up' }">Sign Up</li>
        </router-link>
      </ul>
    </div>

    <div class="flex-grow"></div>

    <div class="mr-4 flex items-center relative">
      <input
        type="text"
        class="border border-gray-400 rounded-full h-10 pl-8 pr-4 py-1 placeholder-gray-400 text-sm"
        placeholder="Search"
      />
      <span class="absolute inset-y-0 left-0 flex items-center pl-3"
        ><i class="pi pi-search text-xl"></i
      ></span>
      <span class="text-xl cursor-pointer ml-2"
        ><i class="pi pi-shopping-cart"></i
      ></span>
     
        <span class="text-xl cursor-pointer ml-2">
        
          <i class="pi pi-user"></i>
        
        </span>
      
    </div>
  </div>
</template>

<script>
export default {
  name: "AppHeader",
  components: {},
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");

.header-main {
  font-family: "Marcellus", sans-serif;
}
</style>
