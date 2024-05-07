<template>
  <div class="container">
    <div class="left-section">
      <img
        src="../assets/images/ecommerce.jpg"
        alt="Other Image"
        class="other-image"
      />
    </div>
    <div class="content">
      <div class="right-section">
        <img src="../assets/images/logo.png" alt="Logo" class="logo" />
      </div>
      <h2>Sign in to view your shop</h2>
      <form>
        <div class="form-group">
          <label for="email" class="label">Username</label>
          <input v-model="username" type="email" id="email" name="email" required />
        </div>
        <div class="form-group">
          <label for="password" class="label">Password</label>
          <input v-model="password" type="password" id="password" name="password" required />
        </div>
        <button @click="signIn" type="submit" class="signin-btn">
          Sign In
        </button>
      </form>
      <div class="alternative-signin">
        <p>Or sign in with:</p>
        <button @click="signInWithGoogle" class="google-btn">
          <img
            src="../assets/images/google.png"
            alt="Google Icon"
            class="google-icon"
          />
          Sign in with Google
        </button>
      </div>
      <div class="forgot-password">
        <a href="#" class="forgot-password-link">Forgot password?</a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    signInWithGoogle() {
      // Implement Google sign-in functionality
    },
    signIn() {
      // Implement sign-in functionality
      if (this.username !== "" && this.password !== "") {
        this.$router.push("/manage-space");
      }
    },
  },
};
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: center;
  height: 100vh;
  background-color: #6BBAEC;
}

.left-section,
.content {
  /* flex: 1; */
}

.left-section {
  display: flex;
  justify-content: center;
}

.other-image {
  max-width: 100%;
  max-height: 100vh;
}

.logo {
  max-width: 100%;
  height: auto;
}

.content {
  display: flex;
  flex-direction: column;
  align-items: center;
}

h2 {
  margin-bottom: 20px;
  font-size: 24px;
}

.form-group {
  margin-bottom: 20px;
  border: 1px solid #dfe1e5;
  padding: 10px;
}

.label {
  color: #555;
  font-size: 14px;
  font-weight: bold;
}

input {
  width: 100%;
  padding: 10px;
  border: 1px solid #dfe1e5;
  border-radius: 4px;
  font-size: 14px;
}

.signin-btn {
  width: 100%;
  padding: 10px;
  background-color: #4285f4;
  color: #fff;
  border: none;
  border-radius: 4px;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.signin-btn:hover {
  background-color: #3367d6;
}

.google-btn {
  width: 100%;
  padding: 10px;
  background-color: #fff;
  color: #555;
  border: 1px solid #dfe1e5;
  border-radius: 4px;
  font-size: 14px;
  cursor: pointer;
  display: flex;
  align-items: center;
  transition: background-color 0.3s ease;
}

.google-btn:hover {
  background-color: #f2f2f2;
}

.google-icon {
  height: 16px;
  margin-right: 8px;
}

.alternative-signin {
  margin-top: 20px;
  text-align: center;
  font-size: 14px;
  color: #555;
}

.forgot-password {
  margin-top: 20px;
  text-align: center;
}

.forgot-password-link {
  color: #4285f4;
  text-decoration: none;
  font-size: 14px;
}
</style>