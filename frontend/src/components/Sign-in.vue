<template>
  <div class="container">
    <div class="content">
      <div>
        <img src="../assets/images/logo.png" alt="" class="logo" />
      </div>
      <h2>Sign in to view your shop</h2>
      <form>
        <div class="form-group">
          <label for="email" class="label">Username</label>
          <input type="email" id="email" name="email" required />
        </div>
        <div class="form-group">
          <label for="password" class="label">Password</label>
          <input type="password" id="password" name="password" required />
        </div>
        <button type="submit" class="signin-btn">Sign In</button>
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
  methods: {
    signInWithGoogle() {
      // Implement your Google Sign In logic here
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
  background-color: #f1f1f1;
}

.logo {
  height: 40px;
  margin-bottom: 20px;
}

.content {
  width: 400px;
  background-color: #fff;
  border-radius: 4px;
  padding: 30px;
  box-shadow: 10px 5px 10px rgba(0, 0, 0, 0.1);
}

h2 {
  margin-bottom: 20px;
  font-size: 24px;
}

.form-group {
  margin-bottom: 20px;
}

.label {
  color: #555;
  font-size: 14px;
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