import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue';
import SignUp from '../views/SignUp.vue';

const router = new createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
          path: '',
          component: Home
        },
        {
        path: '/signup',
        name: 'Sign Up',
        component: SignUp
    },
    ]
})

export default router