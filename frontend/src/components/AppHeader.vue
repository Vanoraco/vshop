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

          <li :class="{ 'text-xl': $route.path === '/sign-up' }" v-if="!isUserLoggedIn">
            Sign Up
          </li>

          <li :class="{ 'text-xl': $route.path === '/sign-up' }">Sign Up</li>

        </router-link>
      </ul>
    </div>

    <div class="flex-grow"></div>
    <div class="flex rounded-md border-2 border-blue-500 overflow-hidden max-w-md mx-auto font-[sans-serif] mr-9">
        <input type="email" placeholder="Search..."
          class="w-full outline-none bg-white text-gray-600 text-sm px-4 py-3" />
        <button type='button' class="flex items-center justify-center bg-[#007bff] px-5">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 192.904 192.904" width="16px" class="fill-white">
            <path
              d="m190.707 180.101-47.078-47.077c11.702-14.072 18.752-32.142 18.752-51.831C162.381 36.423 125.959 0 81.191 0 36.422 0 0 36.423 0 81.193c0 44.767 36.422 81.187 81.191 81.187 19.688 0 37.759-7.049 51.831-18.751l47.079 47.078a7.474 7.474 0 0 0 5.303 2.197 7.498 7.498 0 0 0 5.303-12.803zM15 81.193C15 44.694 44.693 15 81.191 15c36.497 0 66.189 29.694 66.189 66.193 0 36.496-29.692 66.187-66.189 66.187C44.693 147.38 15 117.689 15 81.193z">
            </path>
          </svg>
        </button>
      </div>

    <div class="mr-4 flex items-center relative gap-6 ">
     
     <!--  <span class="absolute inset-y-0 left-0 flex items-center pl-3"
        ><i class="pi pi-search text-xl"></i
      ></span> -->
    <ul class="flex flex-row gap-9 mb-6 mr-4" v-if="isUserLoggedIn">
     <li>
      <router-link to="/cart">
        <i class="pi pi-shopping-cart text-xl"></i>
      </router-link>
     </li>
    
     <li>
      <router-link to="/profile">
        <i class="pi 
pi-user-edit text-xl"> Update Profile</i>
      </router-link>
     </li>
      <li>
      <router-link @click="logout()">
        <i class="
   pi pi-sign-out text-xl"> Sign out</i>
      </router-link>
     </li>
     <li>

     </li>
    </ul>

    <ul class="flex flex-row gap-9 mb-6 mr-4" v-if="!isUserLoggedIn">
      <li>
      <router-link to="/cart">
        <i class="pi pi-shopping-cart text-xl"></i>
      </router-link>
     </li>

     <li>
      <router-link to="/login">
        <i class="pi pi-user text-2xl font-style hover:bg-black hover:text-white hover:rouned hover:px-3 hover:py-1"> Login</i>
      </router-link>
     </li>
      <li>

      </li>
    </ul>
     
      <!-- <span class="text-xl cursor-pointer ml-2 "
        ><i class="pi pi-shopping-cart"></i
      ></span>

      <span class="text-xl cursor-pointer ml-2"
        ><i class="pi pi-user"></i -->
        
      <!-- ></span> -->

     
        <span class="text-xl cursor-pointer ml-2">
         <router-link>
          <i class="pi pi-user"></i>
         </router-link> 
        </span>
      

    </div>
  </div>
</template>

<script>

export default {
  name: "AppHeader",
  computed: {
    isUserLoggedIn() {
      // Check if the user's email exists in localStorage
      return localStorage.getItem('email') !== null;
    },
  },
  methods: {
    logout() {
    try {
      // Attempt to remove the user's email from localStorage
      localStorage.removeItem('email');

      // Redirect the user to the login page
      
      window.location.reload();
    } catch (error) {
      // Log any errors that occur during logout
      console.error('Error logging out:', error);
    }
  }
  },
  components: {},
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");

.header-main {
  font-family: "Marcellus", sans-serif;
}
.menu-container {
    position: relative;
}

.menu {
    position: absolute;
    z-index: 1;
    background: whitesmoke;
    display: none;
}

.menu-container:active .menu {
    display: block;
}

.menu-container:hover .menu {
    display: block;
}

.menu a {
    width: 100%;
    min-width: 8rem;
}
</style>
