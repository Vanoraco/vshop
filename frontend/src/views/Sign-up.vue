<template>
  <div>
  <div class="container" v-if="selectedForm == ''">
        <div class="top-text-wrapper flex justify-center font-style text-3xl ">
          <h1>Join as User or Shop Owner</h1>
        </div>
        <div class="grid-wrapper grid-col-auto">
          <label for="role-1" class="role">
            <input type="radio" name="role" id="role-1" value="buyer" v-model="selectedRole"/>
            <div class="card-content-wrapper">
              <span class="check-icon"></span>
              <div class="card-content">
                <img
                  src="https://img.freepik.com/free-vector/niche-service-marketplace-concept-illustration_114360-7303.jpg?w=740&t=st=1715291257~exp=1715291857~hmac=2347f01320250f22ebe38ed5a8ee1730956da1772be0df2090a7204eff010e77"
                  alt=""
                />
                <p class="font-style text-2xl">I want to join as a Buyer, I want to buy Products</p>
              </div>
            </div>
          </label>
          <!-- /.role -->

          <label for="role-2" class="role">
            <input type="radio" name="role" id="role-2" value="seller" v-model="selectedRole"/>
            <div class="card-content-wrapper">
              <span class="check-icon"></span>
              <div class="card-content">
                <img
                  src="https://img.freepik.com/free-vector/we-are-open-concept-illustration_114360-9780.jpg?t=st=1715291318~exp=1715294918~hmac=32f948381897bfede3a1ac8f1f6bbd27c71f1cdec54e9c9a470df167fd8c5d0b&w=740"
                  alt=""
                />
                <p class="font-style text-2xl">I want to join as a Shop owner, I want to sell Products</p>
              </div>
            </div>
          </label>
          <!-- /.role -->
        </div>
        <!-- /.grid-wrapper -->
        <div class="flex justify-center " >
          <button class="mt-12 bg-black font-style text-xl px-4 py-2" v-if="selectedRole === 'buyer'" @click="selectForm('buyer')">{{ buttonText }}</button>
          <button class="mt-12 bg-black font-style text-xl px-4 py-2" v-else-if="selectedRole === 'seller'" @click="selectForm('seller')"> {{ buttonText }} </button>
          <button class="mt-12 bg-grey font-style text-xl px-4 py-2 " v-else @click="selectForm('seller')" disabled> {{ buttonText }} </button>
        </div>

        <div class="flex justify-center mt-9 font-style ">
          Already have an account ? <router-link to="/login" class="ml-2 text-green"> Log in</router-link>  
        </div>
      </div>

      <div class=" flex justify-center gap-9 font-style text-2xl">
       
        
        
        <vee-form :validation-schema="userschema" @submit="registerUser" :initial-values="userData" v-if="selectedForm == 'buyer'" class="shadow-2xl py-9 px-9">
          <p class="text-5xl">Sign up to buy products</p>
          <br>
          
          <div
    class="text-white  font-bold p-4 rounded mb-4"
    v-if="reg_show_alert"
    :class="reg_alert_variant"
  >
    {{ reg_alert_message }}
  </div>
    <!-- Name -->
    <div class="mb-3">
      <label class="inline-block mb-2">First Name</label>
      <vee-field
        type="text"
        name="first_name"
        class="block w-full py-1.5 px-3 text-gray-800 border  border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter First Name"
      />
      <ErrorMessage class="text-red-600" name="name" />
    </div>

    <div class="mb-3">
      <label class="inline-block mb-2">Last Name</label>
      <vee-field
        type="text"
        name="last_name"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter Last Name"
      />
      <ErrorMessage class="text-red-600" name="lname" />
    </div>
    <!-- Email -->
    <div class="mb-3">
      <label class="inline-block mb-2">Email</label>
      <vee-field
        type="email"
        name="email"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter Email"
      />
      <ErrorMessage class="text-red-600" name="email" />
    </div>
    
    <!-- Password -->
    <div class="mb-3">
      <label class="inline-block mb-2">Password</label>
      <vee-field name="password" :bails="false" v-slot="{ field, errors }">
        <input
          type="password"
          class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
          placeholder="Password"
          v-bind="field"
        />
        <div class="text-red-600" v-for="err in errors" :key="err">
          {{ err }}
        </div>
      </vee-field>
    </div>
    <!-- Confirm Password -->
    <div class="mb-3">
      <label class="inline-block mb-2">Confirm Password</label>
      <vee-field
        type="password"
        name="confirm_password"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Confirm Password"
      />
      <ErrorMessage class="text-red-600" name="confirm_password" />
    </div>

    <div class="mb-3">
      <label class="inline-block mb-2">Phone Number</label>
      <vee-field
        type="number"
        name="phone"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter Phone Number"
      />
      <ErrorMessage class="text-red-600" name="phone" />
    </div>

    <div class="mb-3 pl-6">
      <vee-field
        name="tos"
        value="1"
        type="checkbox"
        class="w-4 h-4 float-left -ml-6 mt-1 rounded"
      />
      <label class="inline-block">Accept terms of service</label>
      <ErrorMessage class="text-red-600 block" name="tos" />
    </div>
    <button
      type="submit"
      class="block w-full bg-black text-white py-1.5 px-3 rounded transition hover:bg-purple-700"
      :disabled="reg_in_submission"
    >
      Create my account
    </button>
         </vee-form>
      </div>
      <div class="flex justify-center font-style text-xl">
         <vee-form :validation-schema="ownerschema" @submit="registerOwner" :initial-values="userData" v-if="selectedForm == 'seller'" class="shadow-2xl py-9 px-9">
          <p class="text-5xl">Sign up to sell products</p>
          <br>
    <!-- Name -->
    <div class="mb-3">
      <label class="inline-block mb-2">First Name</label>
      <vee-field
        type="text"
        name="first_name"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter First Name"
      />
      <ErrorMessage class="text-red-600" name="name" />
    </div>

    <div class="mb-3">
      <label class="inline-block mb-2">Last Name</label>
      <vee-field
        type="text"
        name="last_name"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter Last Name"
      />
      <ErrorMessage class="text-red-600" name="lname" />
    </div>
    <!-- Email -->
    <div class="mb-3">
      <label class="inline-block mb-2">Email</label>
      <vee-field
        type="email"
        name="email"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter Email"
      />
      <ErrorMessage class="text-red-600" name="email" />
    </div>
    
    <!-- Password -->
    <div class="mb-3">
      <label class="inline-block mb-2">Password</label>
      <vee-field name="password" :bails="false" v-slot="{ field, errors }">
        <input
          type="password"
          class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
          placeholder="Password"
          v-bind="field"
        />
        <div class="text-red-600" v-for="err in errors" :key="err">
          {{ err }}
        </div>
      </vee-field>
    </div>
    <!-- Confirm Password -->
    <div class="mb-3">
      <label class="inline-block mb-2">Confirm Password</label>
      <vee-field
        type="password"
        name="confirm_password"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Confirm Password"
      />
      <ErrorMessage class="text-red-600" name="confirm_password" />
    </div>

    <div class="mb-3">
      <label class="inline-block mb-2">Phone No.</label>
      <vee-field
        type="number"
        name="phone"
        class="block w-full py-1.5 px-3 text-gray-800 border border-gray-300 transition duration-500 focus:outline-none focus:border-black rounded"
        placeholder="Enter Phone Number"
      />
      <ErrorMessage class="text-red-600" name="phone" />
    </div>

    <div class="mb-3">
      <label class="inline-block mb-2">Profile.</label>
      <input v-validate="'image'" data-vv-as="image" name="image_field" type="file">

    </div>

    <div class="mb-3 pl-6">
      <vee-field
        name="tos"
        value="1"
        type="checkbox"
        class="w-4 h-4 float-left -ml-6 mt-1 rounded"
      />
      <label class="inline-block"> Accept terms of service</label>
      <ErrorMessage class="text-red-600 block" name="tos" />
    </div>

    <button
      type="submit"
      class="block w-full bg-black text-white py-1.5 px-3 rounded transition hover:bg-purple-700"
      :disabled="reg_in_submission"
    >
      Create my account
    </button>
         </vee-form>
        </div>
      </div>
      <!-- /.container -->
</template>

<script>

import { mapActions } from 'pinia'
import axios from 'axios'

// import { Vuelidate } from 'vuelidate';

import 'vue-toast-notification/dist/theme-bootstrap.css'
import { useToast } from 'vue-toastification'


export default {
 
  data() {
    return {

      // v$: Vuelidate(),


      reg_in_submission: false,
      reg_alert_message: 'Please Wait! Your Account is being created',
      reg_alert_variant: 'bg-blue-500',
      reg_show_alert: false,
      name: "",
      email: "",
      phone: "",
      role: "",
      country: "",

      selectedRole: '',
      selectedForm: '',
      schema: {
        name: 'required|min:3|max:100|alpha_spaces',
        lname: 'required|min:3|max:100|alpha_spaces',
        email: 'required|min:3|max:100|email',
        password: 'required|min:9|max:100|excluded:password',
        phone: 'required|max:12',
        confirm_password: 'passwords_mismatched:@password',
        tos: 'tos'
      },
      ownerschema: {
        name: 'required|min:3|max:100|alpha_spaces',
        lname: 'required|min:3|max:100|alpha_spaces',
        email: 'required|min:3|max:100|email',
        password: 'required|min:9|max:100|excluded:password',
        phone: 'required|max:12',
        confirm_password: 'passwords_mismatched:@password',
        tos: 'tos'
      },
    };
  },
  methods: {

    // submitForm() {
    //   console.log(this.v$);
    //   if (this.role === "buyer") {
    //     this.$router.push("/");
    //   } else if (this.role === "shopowner") {
    //     this.$router.push("/manage-space");
    //   }
    // },
    selectForm(role) {
      this.selectedForm = role;
    },

    async registerUser(formdata) {
      console.log(formdata)
     
      try {
        this.reg_show_alert = true
        this.reg_alert_variant = 'bg-blue-500'
        this.reg_alert_message = 'Please Wait! Your Account is being created'
        // Make a POST request to your Golang backend API endpoint
        const response = await axios.post('http://localhost:8000/users/signup', formdata);
        this.reg_in_submission = true
      
        this.$router.push('login' );
        console.log('User registered successfully:', response.data);
        // Optionally, you can redirect the user to another page or display a success message
      } catch (error) {
        this.reg_in_submission = false 
        this.reg_alert_variant = 'bg-red-500'
        this.reg_alert_message = 'An unexpected error occurred. Please try again Later'
        //this.$toast.error('An unexpected error occurred. Please try again Later');
        this.reg_alert_message = 'Redirecting you to shops...'
        
        setTimeout(() =>  {
         
          this.$router.push( '/shops' )
        }, 2000
      )
        console.error('Error registering user:', error.response ? error.response.data : error.message);
        return
        // Optionally, you can display an error message to the user
      }
      this.reg_alert_variant = 'bg-green-500'
      this.reg_alert_message = 'Success! Your Account is created'
      window.location.reload()
    }
  }
,

  computed: {
    buttonText() {
      return this.selectedRole == '' ? 'Create Account' : (this.selectedRole == 'buyer' ? 'Buy Products' : 'Own a Shop')
    }
  }
};
</script>

<style scoped>

@import url("https://fonts.googleapis.com/css2?family=Jacquarda+Bastarda+9+Charted&family=Marcellus&family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");

.font-style {
  font-family: "Marcellus", sans-serif;
}

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

.back-img {
  background: url('https://img.freepik.com/free-photo/abstract-blur-defocused-shopping-mall-center-department-store_1203-9225.jpg?t=st=1715297182~exp=1715300782~hmac=93c8c2a2e1eb2d650a8202ae417ccbb50966fd64d1e85bedf0896d3af16a6b6e&w=1380');
  background-size: cover;
  height: 100%;
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

.radio-item [type="radio"] {
	display: none;
}

.radio-item label {
	display: block;
	padding: 20px 60px;
	background: #fbfbff;
	border: 2px solid rgba(32, 29, 29, 0.1);
	border-radius: 8px;
	cursor: pointer;
	font-size: 18px;
	font-weight: 400;
	min-width: 250px;
	white-space: nowrap;
	position: relative;
	transition: 0.4s ease-in-out 0s;
}
.radio-item label:after,
.radio-item label:before {
	content: "";
	position: absolute;
	border-radius: 50%;
}
.radio-item label:after {
	height: 19px;
	width: 19px;
	border: 2px solid #524eee;
	left: 19px;
	top: calc(50% - 12px);
}
.radio-item label:before {
	background: #524eee;
	height: 20px;
	width: 20px;
	left: 18px;
	top: 26px;
	transform: scale(5);
	opacity: 0;
	visibility: hidden;
	transition: 0.4s ease-in-out 0s;
}
.radio-item [type="radio"]:checked ~ label {
	border-color: #524eee;
}
.radio-item [type="radio"]:checked ~ label::before {
	opacity: 1;
	visibility: visible;
	transform: scale(1);
}


 * {
	 margin: 0;
	 padding: 0;
	 box-sizing: border-box;
}
 body {
	 font-family: 'Open Sans', sans-serif;
	 font-size: 15px;
	 line-height: 1.5;
	 font-weight: 400;
	 background: #f0f3f6;
	 color: #3a3a3a;
}
 hr {
	 margin: 20px 0;
	 border: none;
	 border-bottom: 1px solid #d9d9d9;
}
 label, input {
	 cursor: pointer;
}
 h2, h3, h4, h5 {
	 font-weight: 600;
	 line-height: 1.3;
	 color: #1f2949;
}
 h2 {
	 font-size: 24px;
}
 h3 {
	 font-size: 18px;
}
 h4 {
	 font-size: 14px;
}
 h5 {
	 font-size: 12px;
	 font-weight: 400;
}
 img {
	 max-width: 100%;
	 display: block;
	 vertical-align: middle;
}
 .container {
	 max-width: 99vw;
	 margin: 15px auto;
	 padding: 0 15px;
}
 .top-text-wrapper {
	 margin: 20px 0 30px 0;
}
 .top-text-wrapper h4 {
	 font-size: 24px;
	 margin-bottom: 10px;
}
 .top-text-wrapper code {
	 font-size: 0.85em;
	 background: linear-gradient(90deg, #fce3ec, #ffe8cc);
	 color: #f20;
	 padding: 0.1rem 0.3rem 0.2rem;
	 border-radius: 0.2rem;
}
 .tab-section-wrapper {
	 padding: 30px 0;
}
 .grid-wrapper {
	 display: grid;
	 grid-gap: 30px;
	 place-items: center;
	 place-content: center;
}
 .grid-col-auto {
	 grid-auto-flow: column;
	 grid-template-rows: auto;
}
/* ******************* Main Styeles : Radio Card */
 label.role {
	 cursor: pointer;
}
 label.role .card-content-wrapper {
	 background: #fff;
	 border-radius: 5px;
	 max-width: 280px;
	 min-height: 330px;
	 padding: 15px;
	 display: grid;
	 box-shadow: 0 2px 4px 0 rgba(219, 215, 215, 0.04);
	 transition: 200ms linear;
}
 label.role .check-icon {
	 width: 20px;
	 height: 20px;
	 display: inline-block;
	 border: solid 2px #e3e3e3;
	 border-radius: 50%;
	 transition: 200ms linear;
	 position: relative;
}
 label.role .check-icon:before {
	 content: '';
	 position: absolute;
	 inset: 0;
	 background-image: url("data:image/svg+xml,%3Csvg width='12' height='9' viewBox='0 0 12 9' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M0.93552 4.58423C0.890286 4.53718 0.854262 4.48209 0.829309 4.42179C0.779553 4.28741 0.779553 4.13965 0.829309 4.00527C0.853759 3.94471 0.889842 3.88952 0.93552 3.84283L1.68941 3.12018C1.73378 3.06821 1.7893 3.02692 1.85185 2.99939C1.91206 2.97215 1.97736 2.95796 2.04345 2.95774C2.11507 2.95635 2.18613 2.97056 2.2517 2.99939C2.31652 3.02822 2.3752 3.06922 2.42456 3.12018L4.69872 5.39851L9.58026 0.516971C9.62828 0.466328 9.68554 0.42533 9.74895 0.396182C9.81468 0.367844 9.88563 0.353653 9.95721 0.354531C10.0244 0.354903 10.0907 0.369582 10.1517 0.397592C10.2128 0.425602 10.2672 0.466298 10.3112 0.516971L11.0651 1.25003C11.1108 1.29672 11.1469 1.35191 11.1713 1.41247C11.2211 1.54686 11.2211 1.69461 11.1713 1.82899C11.1464 1.88929 11.1104 1.94439 11.0651 1.99143L5.06525 7.96007C5.02054 8.0122 4.96514 8.0541 4.90281 8.08294C4.76944 8.13802 4.61967 8.13802 4.4863 8.08294C4.42397 8.0541 4.36857 8.0122 4.32386 7.96007L0.93552 4.58423Z' fill='white'/%3E%3C/svg%3E%0A");
	 background-repeat: no-repeat;
	 background-size: 12px;
	 background-position: center center;
	 transform: scale(1.6);
	 transition: 200ms linear;
	 opacity: 0;
}
 label.role input[type='radio'] {
	 appearance: none;
	 -webkit-appearance: none;
	 -moz-appearance: none;
}
 label.role input[type='radio']:checked + .card-content-wrapper {
	 box-shadow: 0 2px 4px 0 rgba(219, 215, 215, 0.5), 0 0 0 2px #3057d5;
}
 label.role input[type='radio']:checked + .card-content-wrapper .check-icon {
	 background: #3057d5;
	 border-color: #3057d5;
	 transform: scale(1.2);
}
 label.role input[type='radio']:checked + .card-content-wrapper .check-icon:before {
	 transform: scale(1);
	 opacity: 1;
}
 label.role input[type='radio']:focus + .card-content-wrapper .check-icon {
	 box-shadow: 0 0 0 4px rgba(48, 86, 213, 0.2);
	 border-color: #3056d5;
}
 label.role .card-content img {
	 margin-bottom: 10px;
}
 label.role .card-content h4 {
	 font-size: 16px;
	 letter-spacing: -0.24px;
	 text-align: center;
	 color: #1f2949;
	 margin-bottom: 10px;
}
 label.role .card-content h5 {
	 font-size: 14px;
	 line-height: 1.4;
	 text-align: center;
	 color: #686d73;
}
 

</style>