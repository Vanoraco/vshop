import { createRouter, createWebHistory } from 'vue-router'
import SignUp from "../views/Sign-up.vue"
import SignIn from "../views/Sign-in.vue"
import HomePage from "../views/Home.vue"
import Shops from "../components/Shops.vue"
import ManageSpace from "../views/owner/ManageSpace.vue"
import productDetail from "../views/product_detail.vue"
import DetailForm from "../views/owner/DetailForm.vue"
import BuyerProfile from "../components/BuyerProfileCard.Vue"


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
      path:"/BuyerProfile",
      component:BuyerProfile
     }
    ]
})

export default router;
