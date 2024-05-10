import { defineStore } from 'pinia'
import { auth, usersCollection } from '@/includes/firebase'

export default defineStore('user', {
  state: () => ({
    userLoggedIn: false
  }),

  actions: {
    async registerUser(value) {
      const userCred = await auth.createUserWithEmailAndPassword(value.email, value.password)

      await usersCollection.doc(userCred.user.uid).set({
        name: value.name,
        email: value.email,
        age: value.age,
        country: value.country
      })

      await userCred.user.updateProfile({
        displayName: value.name
      })

      this.userLoggedIn = false
    },
    async authenticate(value) {
      await auth.signInWithEmailAndPassword(value.email, value.password)

      this.userLoggedIn = true
    },
    async signOut() {
      await auth.signOut()
      this.userLoggedIn = false
    }
  }
})
