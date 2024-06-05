<template>
  <div>
    <!-- Header for desktop -->
    <div class="flex flex-row items-center mt-6 text-6xl h-14 align-middle hidden lg:flex" v-if="!ownerLoggedIn && !localOwnerToken">
      <div class="ml-4">
        <img src="../assets/images/logo.png" alt="Logo" class="h-12" />
      </div>
      <div class="font-serif ml-8">
        <ul class="flex flex-row gap-8 text-2xl">
          <router-link to="/">
            <li :class="{ 'border-b-4 px-2 py-1 rounded-lg text-3xl shadow-xl bg-gradient-to-tr from-teal-600 via-cyan-700 to-blue-800 text-white': $route.path === '/' }">Home</li>
          </router-link>
          <router-link to="/shops">
            <li :class="{ 'border-b-4 px-2 py-1 rounded-lg text-3xl shadow-xl bg-gradient-to-tr from-teal-600 via-cyan-700 to-blue-800 text-white': $route.path === '/shops' }">Shops</li>
          </router-link>
          <router-link to="/sign-up" v-if="!userLoggedIn && !localToken">
            <li :class="{ 'border-b-4 px-2 py-1 rounded-lg text-3xl shadow-xl bg-gradient-to-tr from-teal-600 via-cyan-700 to-blue-800 text-white': $route.path === '/sign-up' }">Sign Up</li>
          </router-link>
        </ul>
      </div>

      <div class="flex-grow"></div>
      <div class="content font-style" v-if="!userLoggedIn && !localToken">
        <div class="content__container">
          <p class="content__container__text">Own</p>
          <ul class="content__container__list">
            <li class="content__container__list__item">Shops !</li>
            <li class="content__container__list__item">Products !</li>
            <li class="content__container__list__item">Money !</li>
            <li class="content__container__list__item">Everything !</li>
          </ul>
        </div>
      </div>
      <div class="flex rounded-full border-2 border-blue-500 overflow-hidden max-w-md mx-auto font-[sans-serif] mr-7 mt-2">
        <input type="text" placeholder="Search..." class="w-full outline-none bg-white text-gray-600 text-sm px-4 py-3" />
        <button type="button" class="flex items-center justify-center bg-[#121314] px-5">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 192.904 192.904" width="16px" class="fill-white">
            <path d="m190.707 180.101-47.078-47.077c11.702-14.072 18.752-32.142 18.752-51.831C162.381 36.423 125.959 0 81.191 0 36.422 0 0 36.423 0 81.193c0 44.767 36.422 81.187 81.191 81.187 19.688 0 37.759-7.049 51.831-18.751l47.079 47.078a7.474 7.474 0 0 0 5.303 2.197 7.498 7.498 0 0 0 5.303-12.803zM15 81.193C15 44.694 44.693 15 81.191 15c36.497 0 66.189 29.694 66.189 66.193 0 36.496-29.692 66.187-66.189 66.187C44.693 147.38 15 117.689 15 81.193z"></path>
          </svg>
        </button>
      </div>
      <router-link to="/cart" class="">
        <i class="fas fa-shopping-cart text-gray-800 text-2xl mb-4 mr-2"></i>
      </router-link>
      <ul class="flex flex-row gap-9 mb-6 mr-6 mt-2">
        <li v-if="!userLoggedIn && !localToken">
          <router-link to="/login">
            <i class="pi pi-user text-2xl font-style rounded-md py-2 px-2 mr-3 bg-[#3cda5e] text-white hover:text-white hover:text-2xl hover:rounded-lg ml-2">
              Shop Now
            </i>
          </router-link>
        </li>
      </ul>
      <div class="w-[200px]" v-if="userLoggedIn || localToken">
        <div class="flex items-center justify-start space-x-4" @click="toggleDrop">
          <div class="font-semibold dark:text-white text-left cursor-pointer">
            <div class="text-xs text-gray-500 dark:text-gray-400"><img :src="profImage" alt="" class="rounded-full w-9 h-9"></div>
          </div>
        </div>
        <div v-show="showDropDown" class="absolute right-[75px] z-10 mt-2 w-44 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none" role="menu" aria-orientation="vertical" aria-labelledby="menu-button" tabindex="-1">
          <div class="py-1 text-left" role="none">
            <router-link to="/user/profile" class="hover:text-white"><p class="text-sm px-4 py-2 hover:bg-slate-900">Update Profile</p></router-link>
            <p class="text-sm px-4 py-2 hover:bg-slate-900 hover:text-white cursor-pointer" @click="logout">Logout</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Mobile Menu -->
    <div class="lg:hidden flex flex-row items-center mt-6 text-4xl h-14 align-middle">
      <div class="ml-4">
        <img src="../assets/images/logo.png" alt="Logo" class="h-12" />
      </div>
      <div class="flex-grow"></div>
      <button @click="toggleMobileMenu" class="mr-4">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" class="w-8 h-8 text-gray-800">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" />
        </svg>
      </button>
      <div v-if="isMobileMenuOpen" class="absolute top-0 left-0 w-full h-full bg-white z-50 flex flex-col items-center justify-center">
        <ul class="flex flex-col gap-8 text-2xl">
          <router-link to="/" @click="toggleMobileMenu">
            <li :class="{ 'border-b-4 border-black text-3xl shadow-xl': $route.path === '/' }">Home</li>
          </router-link>
          <router-link to="/shops" @click="toggleMobileMenu">
            <li :class="{ 'border-b-4 border-black text-3xl shadow-xl': $route.path === '/shops' }">Shops</li>
          </router-link>
          <router-link to="/sign-up" v-if="!userLoggedIn && !localToken" @click="toggleMobileMenu">
            <li :class="{ 'border-b-4 border-black text-3xl shadow-xl': $route.path === '/sign-up' }">Sign Up</li>
          </router-link>
          <router-link to="/login" v-if="!userLoggedIn && !localToken" @click="toggleMobileMenu">
            <li class="text-2xl bg-[#3cda5e] text-white rounded-md py-2 px-2">Shop Now</li>
          </router-link>
        </ul>
        <router-link to="/cart" class="mt-4" @click="toggleMobileMenu">
          <i class="fas fa-shopping-cart text-gray-800 text-2xl"></i>
        </router-link>
      </div>
    </div>
  </div>
</template>


<script>
import { mapState, mapStores } from 'pinia';
import useUserStore from '../stores/user'
import { mapActions } from 'pinia';

export default {
  name: "AppHeader",
  data() {
    return {
      showDropDown: false,
      localToken: localStorage.getItem("token") !== null,
      localOwnerToken: localStorage.getItem("owner_token") !== null,
      isMobileMenuOpen: false, // added for mobile menu toggle
    }
  },
  mounted() {
    console.log(this.localToken)
    console.log(this.userLoggedIn)
    console.log(this.ownerLoggedIn)
  },
  computed: {
    ...mapState(useUserStore, ['userLoggedIn']),
    ...mapState(useUserStore, ['ownerLoggedIn']),
    profImage() {
      return localStorage.getItem("profile_img")
    }
  },
  methods: {
    ...mapActions(useUserStore, ['signOut']),
    logout() {
      try {
        this.signOut();
        window.location.reload();
        this.$router.push("/");
      } catch (error) {
        console.error("Error logging out:", error);
      }
    },
    toggleDrop() {
      this.showDropDown = !this.showDropDown;
    },
    toggleMobileMenu() { // added for mobile menu toggle
      this.isMobileMenuOpen = !this.isMobileMenuOpen;
    }
  },
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");

.font-style {
  font-family: "Marcellus", sans-serif;
}

.content {
  position: absolute;
  bottom: 70%;
  left: 50%;
  transform: translate(-50%, -50%);
  height: 160px;
  overflow: hidden;
  font-family: 'Marcellus', sans-serif;
  font-size: 35px;
  line-height: 40px;
  color: #000000;
}

.content__container {
  position: relative;
  font-weight: 600;
  overflow: hidden;
  height: 40px;
  padding: 0 40px;
}

.content__container:before,
.content__container:after {
  position: absolute;
  top: 0;
  color: #16a085;
  font-size: 42px;
  line-height: 40px;
  -webkit-animation-name: opacity;
  -webkit-animation-duration: 2s;
  -webkit-animation-iteration-count: infinite;
  animation-name: opacity;
  animation-duration: 2s;
  animation-iteration-count: infinite;
}

.content__container:before {
  content: '[';
  left: 0;
}

.content__container:after {
  content: ']';
  right: 0;
}

.content__container__text {
  display: inline;
  float: left;
  margin: 0;
}

.content__container__list {
  margin-top: 0;
  padding-left: 110px;
  text-align: left;
  list-style: none;
  -webkit-animation-name: change;
  -webkit-animation-duration: 10s;
  -webkit-animation-iteration-count: infinite;
  animation-name: change;
  animation-duration: 10s;
  animation-iteration-count: infinite;
}

.content__container__list__item {
  line-height: 40px;
  margin: 0;
}

@-webkit-keyframes opacity {
  0%, 100% {
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
}

@-webkit-keyframes change {
  0%, 12.66%, 100% {
    transform: translate3d(0, 0, 0);
  }
  16.66%, 29.32% {
    transform: translate3d(0, -25%, 0);
  }
  33.32%, 45.98% {
    transform: translate3d(0, -50%, 0);
  }
  49.98%, 62.64% {
    transform: translate3d(0, -75%, 0);
  }
  66.64%, 79.3% {
    transform: translate3d(0, -50%, 0);
  }
  83.3%, 95.96% {
    transform: translate3d(0, -25%, 0);
  }
}

@keyframes opacity {
  0%, 100% {
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
}

@keyframes change {
  0%, 12.66%, 100% {
    transform: translate3d(0, 0, 0);
  }
  16.66%, 29.32% {
    transform: translate3d(0, -25%, 0);
  }
  33.32%, 45.98% {
    transform: translate3d(0, -50%, 0);
  }
  49.98%, 62.64% {
    transform: translate3d(0, -75%, 0);
  }
  66.64%, 79.3% {
    transform: translate3d(0, -50%, 0);
  }
  83.3%, 95.96% {
    transform: translate3d(0, -25%, 0);
  }
}
</style>
