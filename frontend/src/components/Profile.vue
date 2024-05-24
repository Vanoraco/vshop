<template>
  <div
    class="container mx-auto p-4 flex justify-center items-center min-h-screen"
  >
    <div class="bg-white shadow rounded-lg p-6 w-full max-w-lg">
      <div class="text-center relative">
        <div
          class="relative inline-block"
          @mouseover="showCameraIcon = true"
          @mouseout="showCameraIcon = false"
        >
          <img
            :src="profile.image"
            alt="Profile Image"
            class="rounded-full h-32 w-32 object-cover cursor-pointer"
            @click="triggerFileInput"
          />
          <div
            v-if="showCameraIcon"
            class="absolute bottom-0 right-0 bg-gray-700 text-white p-1 rounded-full cursor-pointer"
            @click="triggerFileInput"
          >
            <i class="fas fa-camera"></i>
          </div>
          <input
            type="file"
            ref="fileInput"
            class="hidden"
            @change="updateProfileImage"
            accept="image/*"
          />
        </div>
        <div>Nati@vshop</div>
      </div>
      <div class="mt-4 text-center">
        <input
          v-model="profile.name"
          class="text-xl font-semibold border-b focus:outline-none focus:border-blue-500 text-center"
        />
      </div>
      <form @submit.prevent="updateProfile" class="mt-6 space-y-4">
        <div class="space-y-2">
          <label class="block text-gray-700">Phone Number</label>
          <input
            v-model="profile.phone"
            type="tel"
            class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
          />
        </div>
        <div class="space-y-2">
          <label class="block text-gray-700">Shipping Address</label>
          <button
            type="button"
            class="w-full px-4 py-2 bg-blue-500 text-white rounded-lg"
            @click="toggleShippingForm"
          >
            Edit Shipping Address
          </button>
          <div v-if="showShippingForm" class="mt-2 space-y-2">
            <input
              placeholder="Town"
              v-model="profile.shippingAddress.town"
              type="text"
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
            />
            <input
              placeholder="Zip Code"
              v-model="profile.shippingAddress.zipCode"
              type="text"
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
            />
            <input
              placeholder="City"
              v-model="profile.shippingAddress.city"
              type="text"
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
            />
            <input
              placeholder="Country"
              v-model="profile.shippingAddress.country"
              type="text"
              class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
            />
          </div>
        </div>
        <div class="space-y-2">
          <label class="block text-gray-700">Location</label>
          <input
            v-model="profile.location"
            type="text"
            class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
          />
        </div>
        <div class="space-y-2">
          <label class="block text-gray-700">Buy History</label>
          <textarea
            v-model="profile.history"
            class="w-full px-4 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
          >
You have no purchase history</textarea
          >
        </div>
        <button
          type="submit"
          class="w-full py-2 px-4 bg-blue-500 text-white rounded-lg"
        >
          Update Profile
        </button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  // mounted() {
  //   this.getUserInfo();
  // },
  data() {
    return {
      profile: {
        image: "https://via.placeholder.com/150",
        name: "",
        phone: "",
        shippingAddress: {
          town: "",
          zipCode: "",
          city: "",
          country: "",
        },
        location: "",
        history: "",
      },
      showCameraIcon: false,
      showShippingForm: false,
    };
  },
  methods: {
    triggerFileInput() {
      this.$refs.fileInput.click();
    },
    updateProfileImage(event) {
      const file = event.target.files[0];
      const reader = new FileReader();
      reader.onload = (e) => {
        this.profile.image = e.target.result;
      };
      reader.readAsDataURL(file);
    },
    editProfile() {
      alert("Edit profile.");
    },
    toggleShippingForm() {
      this.showShippingForm = !this.showShippingForm;
    },
    updateProfile() {
      alert("Profile updated successfully!");
    },
    // async getUserInfo() {
    //   const response = await axios.get("http://localhost:8000/users/user");
    //   this.user = response.data;
    //   console.log(this.user);
    // },
  },
};
</script>

<style scoped>
</style>
