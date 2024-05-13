import { createRouter, createWebHistory } from 'vue-router'
import SignUp from "../views/Sign-up.vue"
import SignIn from "../views/Sign-in.vue"
import HomePage from "../views/Home.vue"
import Shops from "../components/Shops.vue"
import ShopCard from '../components/ShopCard.vue'
import ManageSpace from "../views/owner/ManageSpace.vue"
import productDetail from "../views/product_detail.vue"
import DetailForm from "../views/owner/DetailForm.vue"
import Login from "../views/Login.vue"
import Cart from "../components/Cart.vue"

// import BuyerProfile from "../components/BuyerProfileCard.Vue"



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
        path:"/Shops",
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

     },
     {

      path: "/Cart",
      component: Cart

     },
     {
      path:"/ShopCard",
      component: ShopCard
     }
    ]
})

export default router;
