<template>
    <v-container>
      <v-row>
        <v-col cols="12">
          <v-form @submit.prevent="submitForm">
            <v-text-field v-model="shopOwnerName" label="Shop Owner Name"></v-text-field>
            <v-text-field v-model="shopName" label="Shop Name"></v-text-field>
            <v-select v-model="shopCategory" :items="categories" label="Shop Category"></v-select>
            <v-btn type="submit" color="primary">Submit</v-btn>
          </v-form>
        </v-col>
      </v-row>
    </v-container>
  </template>
  
  <script>


  export default {
    data() {
      return {
        shopOwnerName: '',
        shopName: '',
        shopCategory: null,
        categories: ['Clothing', 'Electronics', 'Books', 'Home & Garden'] // Add more categories as needed
      };
    },
    methods: {
      submitForm() {
      if (this.categories!="") {
        this.$router.push("/manage-space");
      }
      }
    }
  };
  </script>
  