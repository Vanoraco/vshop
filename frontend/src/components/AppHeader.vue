<template>
  <div class="flex flex-row">
    <div class="mt-9 ml-20">
      <img src="../assets/images/logo.png" alt="" class="h-12" />
    </div>
    <div
      class="font-serif mt-12 flex justify-center font-medium mx-auto mr-96 header-main"
    >
      <ul class="flex flex-row gap-20 text-xl mx-auto">
        <router-link to="/">
          <li class="text-2xl">Home</li>
        </router-link>

        <router-link to="/shops">
          <li class="text-2xl">Shops</li>
        </router-link>

        <router-link to="/sign-in">
          <li class="text-2xl">Sign in</li>
        </router-link>

        <router-link to="/sign-up">
          <li class="text-2xl">create a shop</li>
        </router-link>
      </ul>
    </div>

    <div class="mt-11 mr-36">
      <ul class="mr-36 flex flex-row">
        <li>
          <span class="mr-4 mt-3"
            ><i class="pi pi-search text-xl mt-3"></i></span
          ><input
            type="text"
            class="border-black bg-slate-100 rounded-full h-10 placeholder:pl-[45px] placeholder:text-xl p-3"
            placeholder="Search"
          />
        </li>
        <li>
          <span class="mt-12"
            ><i class="pi pi-shopping-cart mt-3 ml-3 text-xl"></i
          ></span>
        </li>
      </ul>
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