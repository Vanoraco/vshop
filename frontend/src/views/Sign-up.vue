<template>
  <div class="container-fluid py-5 signup-page">
    <div class="row justify-content-center align-items-center">
      <div class="col-md-6 col-lg-4">
        <div class="card shadow">
          <div class="card-body">
            <h2 class="text-center mb-4">Sign Up</h2>
            <form @submit.prevent="submitForm">
              <div class="form-group">
                <label for="name" class="form-label">Name</label>
                <input
                  type="text"
                  id="name"
                  class="form-control"
                  v-model="name"
                  required
                />
              </div>
              <div class="form-group">
                <label for="email" class="form-label">Email address</label>
                <input
                  type="email"
                  id="email"
                  class="form-control"
                  v-model="email"
                  required
                />
              </div>
              <div class="form-group">
                <label for="phone" class="form-label">Phone Number</label>
                <input
                  type="tel"
                  id="phone"
                  class="form-control"
                  v-model="phone"
                  required
                />
              </div>
              <div class="form-group">
                <label for="role" class="form-label">Role</label>
                <select id="role" class="form-control" v-model="role" required>
                  <option value="">Select Role</option>
                  <option value="buyer">Buyer</option>
                  <option value="shopowner">Shop Owner</option>
                </select>
              </div>
              <div class="form-group">
                <label for="country" class="form-label">Country</label>
                <select
                  id="country"
                  class="form-control"
                  v-model="country"
                  required
                >
                  <option value="">Select Country</option>
                  <option value="Ethiopia">Ethiopia</option>
                </select>
              </div>
              <div class="form-group">
                <label for="phone" class="form-label">Password</label>
                <input
                  type="tel"
                  id="phone"
                  class="form-control"
                  v-model="Password.password"
                  required
                />
              </div>
              <div class="form-group">
                <label for="phone" class="form-label">Confirm Password</label>
                <input
                  type="tel"
                  id="phone"
                  class="form-control"
                  v-model="Password.confirm"
                  required
                />
              </div>
              <div class="text-center">
                <button type="submit" class="btn btn-primary btn-block">
                  Sign Up
                </button>
              </div>
            </form>
            <hr />
            <div class="text-center">
              <p>Already have an account? <a href="#">Log In</a></p>
              <p>Or sign up with:</p>
              <div>
                <button class="btn btn-google btn-block">
                  <i class="fab fa-google"></i> Sign up with Google
                </button>
                <button class="btn btn-twitter btn-block">
                  <i class="fab fa-twitter"></i> Sign up with Twitter
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { useVuelidate } from '@vuelidate/core'
import { required, email } from '@vuelidate/validators'
export default {
  
  data() {
    return {
      v$:useVuelidate(),
      name: "",
      email: "",
      phone: "",
      role: "",
      country: "",
      Password:{
        password:"",
        confirm:""
      }
    };
  },
  validations(){
    return{
      name: {required},
      email: {required},
      phone: {required},
      role: {required},
      country: {required},
      Password:{
        password:{required},
        confirm:{required}
    }
   
  }},
  methods: {
    
    submitForm() {
      console.log(this.v$)
      if (this.role === "buyer") {
        this.$router.push("/");
      } else if (this.role === "shopowner") {
        this.$router.push("/manage-space");
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