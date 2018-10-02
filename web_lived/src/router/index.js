import Vue from 'vue'
import Router from 'vue-router'
import index from '@/pages/index'
import Home from '@/pages/home/Home'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: index
    },
    {
      path: '/home',
      name: 'Home',
      component: Home
    }
  ]
})
