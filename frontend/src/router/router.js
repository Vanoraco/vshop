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
import ShopDetail from "../components/ShopDetail.vue"
// import Admin from "../views/Admin/dashboard/Layout.vue"
import AdminMedias from "../views/Admin/pages/admin/medias.vue"
import AdminBin from "../views/Admin/pages/admin/recycle-bin.vue"
import AdminApp from "../views/Admin/AdminApp.vue"
import Profile from "../components/Profile.vue"

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
      path:"/shops",
      component:Shops
     },
     {
        path:"/sign-in",
        component:SignIn
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
     },
     {
      path:"/Shop detail",
      component: ShopDetail
     },
     {
      path: '/shop-detail/:shopName',
      name: 'ShopDetail',
      component: ShopDetail,
      props: true
     } ,
     {
      path:"/admin",
      component: AdminApp
     },
     {
      path:"/admin/medias",
      component: AdminMedias
     },
     {
      path:"/admin/recycle-bin",
      component: AdminBin
     },
     {
      path:'/user/profile',
      component:Profile
     }
    ]
})

export default router;
