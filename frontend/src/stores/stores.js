/*
import { defineStore } from 'pinia'
// import { auth, usersCollection } from '@/includes/firebase'
import axios from 'axios';

export default defineStore('shop', {
  state: () => {
    return {
        owner_id: localStorage.getItem('owner_id')
    }
  },

  actions: {
    async createShop(formdata) {
        const ownerID = this.owner_id

        const response = await axios.post(`http://localhost:8000/owners/addshop?owner_id=${ownerID}`, formdata, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        });

        console.log('Shop registered successfully:', response.data);
       
    },

   
})
*/
