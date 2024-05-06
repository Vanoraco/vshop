<template class="w-full bg-cyan-400">
  <div class="container ">
    <div class="content">
      <div>
        <img src="../assets/images/logo.png" alt="" class="logo" />
      </div>
      <h2>Create accunt to become shop owner</h2>
      <form>
        <div class="form-group">
          <label for="firstName" class="label">First name</label>
          <input type="text" id="firstName" name="firstName" required />
        </div>
        <div class="form-group">
          <label for="lastName" class="label">Last name</label>
          <input type="text" id="lastName" name="lastName" required />
        </div>

        <div class="form-group">
          <label for="password" class="label">Password</label>
          <input type="password" id="password" name="password" required />
        </div>
        <div class="form-group">
          <label for="confirmPassword" class="label">Confirm password</label>
          <input
            type="password"
            id="confirmPassword"
            name="confirmPassword"
            required
          />
        </div>
        <button type="submit" class="signup-btn">Next</button>
      </form>
      <div class="alternative-signup">
        <p>Or sign up with:</p>
        <button @click="signUpWithGoogle" class="google-btn">
          <img
            src="../assets/images/google.png"
            alt="Google Icon"
            class="google-icon"
          />
          Sign up with Google
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    signUpWithGoogle() {},
  },
};
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f1f1f1;
}

.logo {
  height: 40px;
  margin-bottom: 20px;
}

.content {

  background-color: #fff;
  border-radius: 4px;
  padding: 30px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

h2 {
  margin-bottom: 20px;
  font-size: 24px;
  color: #333; /* Customize the font color */
  font-family: Arial, sans-serif; /* Customize the font family */
}

.form-group {
  margin-bottom: 20px;
}

.label {
  color: #555; /* Customize the label color */
  font-size: 14px;
  font-family: Arial, sans-serif; /* Customize the font family */
}

input {
  width: 100%;
  padding: 10px;
  border: 1px solid #dfe1e5;
  border-radius: 4px;
  font-size: 14px;
  font-family: Arial, sans-serif; /* Customize the font family */
  color: #333; /* Customize the input text color */
}

.signup-btn {
  width: 100%;
  padding: 10px;
  background-color: #4285f4;
  color: #fff;
  border: none;
  border-radius: 4px;
  font-size: 14px;
  font-family: Arial, sans-serif; /* Customize the font family */
  cursor: pointer;
}

.google-btn {
  width: 100%;
  padding: 10px;
  background-color: #fff;
  color: #555; /* Customize the button text color */
  border: 1px solid #dfe1e5;
  border-radius: 4px;
  font-size: 14px;
  font-family: Arial, sans-serif; /* Customize the font family */
  cursor: pointer;
  display: flex;
  align-items: center;
}

.google-icon {
  height: 16px;
  margin-right: 8px;
}

.alternative-signup {
  margin-top: 20px;
  text-align: center;
  font-size: 14px;
  color: #555; /* Customize the alternative signup text color */
  font-family: Arial, sans-serif; /* Customize the font family */
}
</style>