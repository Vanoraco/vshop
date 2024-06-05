<template>
    <div class="checkout-form">
      <form @submit.prevent="submitForm">
        <label for="name">Name:</label>
        <input type="text" id="name" v-model="formData.name" required>
        
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="formData.email" required>
        
        <label for="address">Address:</label>
        <input type="text" id="address" v-model="formData.address" required>
        
        <label for="city">City:</label>
        <input type="text" id="city" v-model="formData.city" required>
        
        <label for="country">Country:</label>
        <select id="country" v-model="formData.country" required>
          <option value="">Select a country</option>
          <option value="Hawasa">United States</option>
          <option value="Addis ababa">Canada</option>
          <option value="Arbaminch">United Kingdom</option>
          <!-- Add more country options here -->
        </select>
        
        <input type="submit" value="Place Order">
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        formData: {
          name: '',
          email: '',
          address: '',
          city: '',
          country: ''
        }
      };
    },
    methods: {
      submitForm(data) {
        // Handle form submission here
        console.log('Form submitted!');
        console.log(data)
        // You can perform further actions like sending the data to a server for processing
      }
    }
  };
  </script>
  
  <style scoped>
  /* Add your component-specific CSS styling here */
  .checkout-form {
    background-color: #f7f7f7;
    border: 1px solid #ddd;
    padding: 20px;
  }
  
  .checkout-form label {
    display: block;
    margin-bottom: 10px;
    font-weight: bold;
  }
  
  .checkout-form input[type="text"],
  .checkout-form input[type="email"],
  .checkout-form select {
    width: 100%;
    padding: 10px;
    margin-bottom: 20px;
    border: 1px solid #ddd;
    border-radius: 4px;
  }
  
  .checkout-form input[type="submit"] {
    background-color: #f90;
    color: #fff;
    border: none;
    padding: 10px 20px;
    cursor: pointer;
  }
  </style>