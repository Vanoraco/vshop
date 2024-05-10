<template>
  <v-container fluid py-5 class="signup-page">
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="6" lg="4">
        <v-card class="shadow">
          <v-card-text>
            <h2 class="text-center mb-4">Sign Up</h2>
            <v-form @submit.prevent="submitForm">
              <v-text-field v-model="name" label="Name" required></v-text-field>
              <v-text-field
                v-model="email"
                label="Email address"
                type="email"
                required
              ></v-text-field>
              <v-text-field
                v-model="phone"
                label="Phone Number"
                type="tel"
                required
              ></v-text-field>
              <v-select
                v-model="role"
                label="Role"
                :items="roleItems"
                required
              ></v-select>
              <v-select
                v-model="country"
                label="Country"
                :items="countryItems"
                required
              ></v-select>
              <v-text-field
                v-model="password"
                label="Password"
                type="password"
                required
              ></v-text-field>
              <v-text-field
                v-model="confirmPassword"
                label="Confirm Password"
                type="password"
                required
              ></v-text-field>
              <v-btn type="submit" color="primary" block> Sign Up </v-btn>
            </v-form>
            <hr />
            <div class="text-center">
              <p>Already have an account? <a href="#">Log In</a></p>
              <p>Or sign up with:</p>
              <div>
                <v-btn color="google" block>
                  <v-icon left>mdi-google</v-icon> Sign up with Google
                </v-btn>
                <v-btn color="twitter" block>
                  <v-icon left>mdi-twitter</v-icon> Sign up with Twitter
                </v-btn>
              </div>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { useVuelidate } from "@vuelidate/core";
import { required, email } from "@vuelidate/validators";
export default {
  data() {
    return {
      v$: useVuelidate(),
      name: "",
      email: "",
      phone: "",
      role: "",
      country: "",
      Password: {
        password: "",
        confirm: "",
      },
      roleItems: ["Buyer", "Shop Owner"],
      countryItems: ["Ethiopia", "Djibouti", "Eritrea", "Somalia", "Kenya"],
    };
  },
  validations() {
    return {
      name: { required },
      email: { required },
      phone: { required },
      role: { required },
      country: { required },
      Password: {
        password: { required },
        confirm: { required },
      },
    };
  },
  methods: {
    submitForm() {
      console.log(this.v$);
      if (this.role === "buyer") {
        this.$router.push("/");
      } else if (this.role === "Shop Owner") {
        this.$router.push("/DetailForm");
      }
    },
  },
};
</script>

<style scoped>
.signup-page {
  background-image: url("https://source.unsplash.com/1600x900/?ecommerce");
  background-size: cover;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.card {
  background-color: rgba(255, 255, 255, 0.9);
  border-radius: 8px;
  padding: 20px;
}

.card-body {
  padding: 0;
}

.form-group {
  margin-bottom: 20px;
}

.form-label {
  font-weight: bold;
}

.form-control {
  border-radius: 5px;
  border: none;
  background-color: #f4f4f4;
  padding: 12px;
  font-size: 16px;
  box-shadow: none;
}

.btn-primary {
  background-color: #4285f4;
  border-color: #4285f4;
  border-radius: 5px;
  padding: 12px;
  font-size: 16px;
  width: 100%;
}

.btn-primary:hover {
  background-color: #357ae8;
  border-color: #357ae8;
}

.btn-google {
  background-color: #dd4b39;
  border-color: #dd4b39;
  border-radius: 5px;
  padding: 12px;
  font-size: 16px;
  width: 100%;
}

.btn-google:hover {
  background-color: #c23321;
  border-color: #c23321;
}

.btn-twitter {
  background-color: #1da1f2;
  border-color: #1da1f2;
  border-radius: 5px;
  padding: 12px;
  font-size: 16px;
  width: 100%;
}

.btn-twitter:hover {
  background-color: #0c85d0;
  border-color: #0c85d0;
}
</style>