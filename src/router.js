import Vue from 'vue'
import VueRouter from 'vue-router'
import DefaultLayout from '@/layouts/default.vue'

Vue.use(VueRouter)

import Home from './views/Home.vue'
const routes = [
  {
    path: '/',
    name: 'DefaultLayout',
    component: DefaultLayout,
    children: [
      {
        path: ':id',
        name: 'Home',
        component: Home,
      },
    ],
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
})
export default router
