import { createRouter, createWebHashHistory } from 'vue-router'
import Start from '../views/Start.vue'
import Test from '../views/Test.vue'
import Editor from '../components/Editor.vue'
import Login from '../components/Login.vue'
require('../style.css')
const routes = [
  {
    path: '/',
    name: 'Start',
    component: Start
  },
  {path:'/login',
    name: 'Login',
    component:Login
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/test/:id',
    name: 'Test',
    component: Test,
  },
  {
    path: '/editor',
    name: 'Editor',
    component: Editor,
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
