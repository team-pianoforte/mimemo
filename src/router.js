import Vue from 'vue'
import VueRouter from 'vue-router'
import DefaultLayout from '@/layouts/DefaultLayout.vue'
import Board from './views/Board.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'DefaultLayout',
    component: DefaultLayout,
    children: [
      {
        path: ':id',
        name: 'Home',
        component: Board,
        props: true,
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
