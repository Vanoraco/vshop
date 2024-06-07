import { createRouter, createWebHistory } from 'vue-router'
import SignUp from "../views/Sign-up.vue"
import SignIn from "../views/Sign-in.vue"
import HomePage from "../views/Home.vue"
import Shops from "../components/Shops.vue"
import Checkout from "../components/Checkout.vue"
import Login from "../views/Login.vue"
import OwnerLogin from "../views/Owner/Login.vue"
import Cart from "../components/Cart.vue"
import ShopDetail from "../components/ShopDetail.vue"
import AdminApp from "../views/Admin/AdminApp.vue"
import AdminMedias from "../views/Admin/pages/admin/medias.vue"
import AdminBin from "../views/Admin/pages/admin/recycle-bin.vue"
import Profile from "../components/Profile.vue"
import ProductDetail from '../components/ProductDetail.vue'
import DashboardLayout from '../components/Dashboard/DashboardLayout.vue'
import Dash from '../views/Dashboard/Dashboard.vue'
import Main from '../views/Dashboard/Main.vue'
import Forms from '../views/Dashboard/Forms.vue'
import Tables from '../views/Dashboard/Tables.vue'
import UIElements from '../views/Dashboard/UIElements.vue'
import Modal from '../views/Dashboard/Modal.vue'
import Card from '../views/Dashboard/Card.vue'
import Blank from '../views/Dashboard/Blank.vue'
import ThreeD from '../components/3DView.vue'
import EditProduct from '../views/Owner/Edit-Product.vue'
import FavouritePage from '../components/fav.vue'

const routes = [
  { path: "/sign-up", component: SignUp, meta: { layout: 'empty' } },
  { path: "/", component: HomePage, meta: { layout: 'empty' } },
  { path: "/shops", component: Shops, meta: { layout: 'empty' } },
  { path: "/sign-in", component: SignIn, meta: { layout: 'empty' } },
  { path: "/login", component: Login, meta: { layout: 'empty' } },
  { path: "/cart", component: Cart, meta: { layout: 'empty' } },
  { path: "/shop/:id", name: 'shop', component: ShopDetail, props: true, meta: { layout: 'empty' } },
  { path: '/edit-product/:id', name: 'editprod', component: EditProduct, props: true },
  { path: "/admin", component: AdminApp, meta: { layout: 'empty' } },
  { path: "/admin/medias", component: AdminMedias, meta: { layout: 'empty' } },
  { path: "/admin/recycle-bin", component: AdminBin, meta: { layout: 'empty' } },
  { path: '/user/profile', component: Profile, meta: { layout: 'empty' } },
  { path: '/shop/:id/product/:proid', name: 'product', component: ProductDetail, meta: { layout: 'empty' } },
  { path: '/main', component: Main, meta: { layout: 'empty' } },
  { path: '/owner-login', component: OwnerLogin, meta: { layout: 'empty' } },
  { path: '/owner', component: DashboardLayout, redirect: '/owner/dash', meta: { layout: 'empty' } },
  { path: '/owner/dash', name: 'Dashboard', component: Dash },
  { path: '/owner/forms', name: 'Forms', component: Forms },
  { path: '/owner/cards', name: 'Cards', component: Card },
  { path: '/owner/tables', name: 'Tables', component: Tables },
  { path: '/owner/ui-elements', name: 'UIElements', component: UIElements },
  { path: '/owner/modal', name: 'Modal', component: Modal },
  { path: '/blank', name: 'Blank', component: Blank },
  { path: '/checkout', name: 'Checkout', component: Checkout, meta: { layout: 'empty' } },
  { path: '/3Dview', name: "3D view", component: ThreeD, meta: { layout: 'empty' } },
  {path: '/favourites', name:"FavouritePage", component:FavouritePage,  meta: { layout: 'empty' }  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem('token'); // Logic to check if user is authenticated

  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (isAuthenticated) {
      next(); // User is authenticated, allow access
    } else {
      next('/login'); // User is not authenticated, redirect to login page
    }
  } else {
    next(); // Route does not require authentication, allow access
  }
});

export default router;
