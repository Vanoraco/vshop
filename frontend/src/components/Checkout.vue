<template>
  <div class="checkout-form">
    <form @submit.prevent="submitForm">
      <label for="name">Name:</label>
      <input type="text" id="name" v-model="formData.name" required />

      <label for="email">Email:</label>
      <input type="email" id="email" v-model="formData.email" required />
      
      <label for="phone_number">Phone number:</label>
      <input type="text" id="phone_number" v-model="formData.phone_number" required />
      
      <label for="address">Address:</label>
      <input type="text" id="address" v-model="formData.address" required />

      <label for="country">Select City:</label>
      <select id="country" v-model="formData.country" required>
        <option value="">Select a city</option>
        <option value="Addis Abeba">Addis Abeba</option>
        <option value="Hawasa">Hawasa</option>
        <option value="Shashemene">Shashemene</option>
        <option value="Mekele">Mekele</option>
        <option value="Bale Robe">Bale Robe</option>
        <option value="Arbaminch">Arbaminch</option>
      </select>

      <div class="product-table mt-6">
        <table>
          <thead>
            <tr>
              <th>Product</th>
              <th>Price</th>
              <th>Quantity</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>{{ product.product_name }}</td>
              <td>{{ product.price }}</td>
              <td>
                <button @click="decrementQuantity">-</button>
                {{ quantity }}
                <button @click="incrementQuantity">+</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <input type="submit" value="Place Order" />
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      formData: {
        name: "",
        email: "",
        address: "",
        city: "",
        country: "",
        phone_number: ""
      },
      product: this.$route.params.product || {},
      quantity: 1
    };
  },
  methods: {
    incrementQuantity() {
      this.quantity += 1;
    },
    decrementQuantity() {
      if (this.quantity > 1) {
        this.quantity -= 1;
      }
    },
    async submitForm() {
      // Handle form submission here
      const orderData = {
        ...this.formData,
        product: this.product,
        quantity: this.quantity
      };
      console.log("Order Data:", orderData);

      try {
        const response = await axios.post('http://localhost:8000/orders', orderData);
        console.log("Order submitted successfully!", response.data);
        this.$router.push({ name: 'shopOwnerDashboard' });
      } catch (error) {
        console.error("There was an error submitting the order:", error);
      }
    }
  }
};
</script>

<style scoped>
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

.product-table {
  margin-top: 20px;
}

.product-table table {
  width: 100%;
  border-collapse: collapse;
}

.product-table th,
.product-table td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

.product-table th {
  background-color: #f2f2f2;
}

.product-table button {
  padding: 5px 10px;
}
</style>
