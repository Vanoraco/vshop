import { defineStore } from 'pinia'
// import { auth, usersCollection } from '@/includes/firebase'
import axios from 'axios';

export default defineStore('user', {
  state: () => {
    return {
    userLoggedIn: false
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
      this.userLoggedIn = true
    },
     signOut() {
      this.userLoggedIn = false

      localStorage.removeItem("token");
        localStorage.removeItem("email");
        localStorage.removeItem("first_name");
        localStorage.removeItem("profile_img");
    }
  }
})
