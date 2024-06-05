<template>
    <div>
      <form @submit.prevent="initializeTransaction">
        <input type="text" v-model="paymentRequest.amount" placeholder="Amount"  />
        <input type="text" v-model="paymentRequest.currency" placeholder="Currency"  />
        <input type="email" v-model="paymentRequest.email" placeholder="Email"  />
        <input type="text" v-model="paymentRequest.first_name" placeholder="First Name"  />
        <input type="text" v-model="paymentRequest.last_name" placeholder="Last Name" />
        <input type="text" v-model="paymentRequest.phone_number" placeholder="Phone Number"  />
        <input type="text" v-model="paymentRequest.tx_ref" placeholder="Transaction Reference"  />
        <button type="submit">Initialize Transaction</button>
      </form>
      <p v-if="checkoutUrl">Checkout URL: <a :href="checkoutUrl" target="_blank">{{ checkoutUrl }}</a></p>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        paymentRequest: {
  "amount": "",
  "currency": "",
  "email": "",
  "first_name": "",
  "last_name": "",
  "phone_number": "",
  "tx_ref": "",
  "callback_url": "https://webhook.site/077164d6-29cb-40df-ba29-8a00e59a7e60",
  "return_url": "https://www.google.com/",
  "customization[title]": "Payment for my favourite merchant",
  "customization[description]": "I love online payments"
}
,
        checkoutUrl: ""
      };
    },
    methods: {
      async initializeTransaction() {
        try {
            console.log(this.paymentRequest)
          const response = await axios.post("http://localhost:8000/intialize-chapa", this.paymentRequest, {
            headers: {
              'Content-Type': 'application/json'
            }
          });
          const data = response.data; 
          const responseData = JSON.parse(response.data.response);
          console.log(responseData.data)
          // Correctly access the data field of the response
          if (data && responseData.data.checkout_url) {
            this.checkoutUrl = data.checkout_url;
          } else {
            console.error("Checkout URL not found in the response:", data);
          }
        } catch (error) {
          console.error("Error initializing transaction:", error);
        }
      }
    }
  };
  </script>
  
  <style scoped>
  /* Add any styles you need here */
  </style>
  