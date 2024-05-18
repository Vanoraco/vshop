import { createRouter, createWebHistory } from 'vue-router'
import SignUp from "../views/Sign-up.vue"
import SignIn from "../views/Sign-in.vue"
import HomePage from "../views/Home.vue"
import Shops from "../components/Shops.vue"
import ManageSpace from "../views/owner/ManageSpace.vue"
import productDetail from "../views/product_detail.vue"
import DetailForm from "../views/owner/DetailForm.vue"
import Login from "../views/Login.vue"


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
     },
     {
        path:"/shops",
        component:Shops
     },
     {
        path:"/manage-space",
        component:ManageSpace
     },
     {
      path:"/product-detail-page",
      component:productDetail
     },
     {
      path:"/DetailForm",
      component:DetailForm
     },
     {
      path: "/login",
      component: Login
     }
    ]
})

router.beforeEach((to, from, next) => {
   const isAuthenticated =  localStorage.getItem('token')/* Logic to check if user is authenticated */;
 
   if (to.matched.some(record => record.meta.requiresAuth)) {
     // Route requires authentication
     if (isAuthenticated) {
       // User is authenticated, allow access
       next('/');
     } else {
       // User is not authenticated, redirect to login page
       next('/login');
     }
   } else {
     // Route does not require authentication, allow access
     next();
   }
 });
 

export default router;
