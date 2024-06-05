import { defineStore } from 'pinia'
// import { auth, usersCollection } from '@/includes/firebase'
import axios from 'axios';

export default defineStore('user', {
  state: () => {
    return {
    userLoggedIn: false,
    ownerLoggedIn: false
    }
  },

  actions: {
    async registerUser(formdata) {
        
        const response = await axios.post('http://localhost:8000/users/signup', formdata, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        });

        console.log('User registered successfully:', response.data);
       
    },

    async registerOwner(formdata) {
      
      const response = await axios.post('http://localhost:8000/owner/signupowner', formdata, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        });
    },
    async login(formdata) {
      const response = await axios.post("http://localhost:8000/users/login", formdata);
      
      localStorage.setItem("token", response.data.token);
        localStorage.setItem("email", response.data.email);
        localStorage.setItem("firstname", response.data.first_name);
        localStorage.setItem("profile_img", response.data.profile_picture)
        localStorage.setItem("phone",response.data.phone)
      this.userLoggedIn = true
    },
    async loginOwner(formdata) {
      const response = await axios.post("http://localhost:8000/owners/login", formdata);
      
      localStorage.setItem("owner_token", response.data.token);
        localStorage.setItem("owner_email", response.data.owner_email);
        localStorage.setItem("owner_firstname", response.data.owner_firstname);
        localStorage.setItem("owner_img", response.data.owner_img)
        localStorage.setItem("owner_id", response.data._id)
      this.ownerLoggedIn = true
    },
     signOut() {
      this.userLoggedIn = false

      localStorage.removeItem("token");
        localStorage.removeItem("email");
        localStorage.removeItem("first_name");
        localStorage.removeItem("profile_img");
        window.location.reload()
    },
    signOutOwner() {
      this.ownerLoggedIn = false
      
      localStorage.removeItem("owner_token");
      localStorage.removeItem("owner_email");
      localStorage.removeItem("owner_firstname");
      localStorage.removeItem("owner_img")
      localStorage.removeItem("_id")

    }
  }
})
