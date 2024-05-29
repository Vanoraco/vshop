import { createRouter, createWebHistory } from 'vue-router'
import SignUp from "../views/Sign-up.vue"
import SignIn from "../views/Sign-in.vue"
import HomePage from "../views/Home.vue"
import Shops from "../components/Shops.vue"
import ShopCard from '../components/ShopCard.vue'
import ManageSpace from "../views/Owner/ManageSpace.vue"
import DetailForm from "../views/Owner/DetailForm.vue"
import Login from "../views/Login.vue"
import OwnerLogin from "../views/Owner/Login.vue"
import Cart from "../components/Cart.vue"
import ShopDetail from "../components/ShopDetail.vue"
// import Admin from "../views/Admin/dashboard/Layout.vue"
import AdminMedias from "../views/Admin/pages/admin/medias.vue"
import AdminBin from "../views/Admin/pages/admin/recycle-bin.vue"
import AdminApp from "../views/Admin/AdminApp.vue"
import Profile from "../components/Profile.vue"
import AddProduct from "../views/Owner/AddItem.vue"
import Dashboard from "../views/Owner/Dashboard.vue"
import ProductDetail from '../components/ProductDetail.vue'

import DashboardLayout from '../components/Dashboard/DashboardLayout.vue'

import Dash from '../views/Dashboard/Dashboard.vue'
import Main from '../views/Dashboard/Main.vue'
import Forms from '../views/Dashboard/Forms.vue'
import Tables from '../views/Dashboard/Tables.vue'
import UIElements from '../views/Dashboard/UIElements.vue'
import AdLogin from '../views/Dashboard/Login.vue'
import Modal from '../views/Dashboard/Modal.vue'
import Card from '../views/Dashboard/Card.vue'
import Blank from '../views/Dashboard/Blank.vue'

const routes = [
  { 
    path: "/sign-up", 
    component: SignUp,
    meta: { layout: 'empty' }
   },
  { 
    path: "/", 
    component: HomePage,  
    meta: { layout: 'empty' } },
  { 
    path: "/shops", 
    component: Shops,
    meta: { layout: 'empty' },
  },
  { 
    path: "/sign-in", 
    component: SignIn,
    meta: { layout: 'empty' },
  },
  { 
    path: "/owner/manage-space", 
    component: ManageSpace,
    meta: { layout: 'empty' },
  },
  // { path: "/product-detail-page", component: productDetail },
  { 
    path: "/DetailForm", 
    component: DetailForm,
    meta: { layout: 'empty' },
  },
  { 
    path: "/login", 
    component: Login,
    meta: { layout: 'empty' },
  },
  {
     path: "/Cart", 
     component: Cart,
     meta: { layout: 'empty' },
    },
  { 
    path: "/ShopCard", 
   component: ShopCard,
   meta: { layout: 'empty' },
  },

  { 
    path: '/shop/:id', 
    name: 'shop', 
    component: ShopDetail,
     props: true ,
     meta: { layout: 'empty' }},

  { 
    path: "/admin", 
    component: AdminApp,
    meta: { layout: 'empty' },
 },
  { 
    path: "/admin/medias", 
    component: AdminMedias,
    meta: { layout: 'empty' }
   },
  { 
    path: "/admin/recycle-bin",
     component: AdminBin,
     meta: { layout: 'empty' }
     },
  { 
    path: '/user/profile',
     component: Profile,
     meta: { layout: 'empty' }
    },
  { 
    path:'/owner/add-product',
     component:AddProduct ,
     meta: { layout: 'empty' }
    },
  { path: '/dashboard',
   component:Dashboard ,
   meta: { layout: 'empty' }
  },
  {
    path:'/shops/shop/:id/product:id', 
    component:ProductDetail,
    meta: { layout: 'empty' }
  },
  {
    path:'/main', 
    component:Main,
    meta: { layout: 'empty' }},

    {
      path: '/owner-login',
      component:OwnerLogin,
      meta: { layout: 'empty' }
    },
  {
    path: '/owner',
    component:DashboardLayout,
    redirect: '/dash',
    meta: { layout: 'empty' }
  },

  {
    path: '/dash',
    name: 'Dashboard',
    component: Dash,
  },
  {
    path: '/forms',
    name: 'Forms',
    component: Forms,
  },
  {
    path: '/cards',
    name: 'Cards',
    component: Card,
  },
  {
    path: '/tables',
    name: 'Tables',
    component: Tables,
  },
  {
    path: '/ui-elements',
    name: 'UIElements',
    component: UIElements,
  },
  {
    path: '/modal',
    name: 'Modal',
    component: Modal,
  },
  {
    path: '/blank',
    name: 'Blank',
    component: Blank,
  },
  
  {
    path: '/',
    name: 'Login',
    component: AdLogin,
    meta: { layout: 'empty' },
  },
 
  
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
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
