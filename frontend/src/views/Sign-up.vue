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
              <div class="text-center color-white">
                <button type="submit" class="btn btn-primary">Sign Up</button>
              </div>
            </form>
            <hr />
            <div class="text-center">
              <p>Already have an account? <a href="#">Log In</a></p>
              <p>Or sign up with:</p>
              <button class="btn btn-primary btn-sm m-3">Google</button>
              <button class="btn btn-primary btn-sm m-3">Twitter</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      name: "",
      email: "",
      phone: "",
      role: "",
      country: "",
    };
  },
  methods: {
    submitForm() {
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
}
.btn-primary {
  background-color: #4285f4;
  border-color: #4285f4;
}
.btn-primary:hover {
  background-color: #357ae8;
  border-color: #357ae8;
}
.btn-outline-primary {
  color: #4285f4;
  border-color: #4285f4;
}
.btn-outline-primary:hover {
  color: #ffffff;
  background-color: #4285f4;
  border-color: #4285f4;
}
</style>
