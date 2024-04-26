import { createRouter, createWebHistory } from 'vue-router'
import SignUp from "../components/Sign-up.vue"
import SignIn from "../components/Sign-in.vue"
import HomePage from "../views/Home.vue"

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
    
     {
      path:"/sign-up",
      component:SignUp
     },
     {
      path:"/",
      component:HomePage
     },
     {
        path:"/sign-in",
        component:SignIn
     }
    ]
})

export default router;
