import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/views/Home.vue'
import About from '@/views/About.vue'
import Resume from '@/views/Resume.vue'
import Photos from '@/views/Photos.vue'
import Travels from '@/views/Travels.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      name: 'Home',
      path: '/',
      component: Home
    },
    {
      name: 'About',
      path: '/about',
      component: About
    },
    {
      name: 'Resume',
      path: '/resume',
      component: Resume
    },
    {
      name: 'Photos',
      path: '/photos',
      component: Photos
    },
    {
      name: 'Travels',
      path: '/travels',
      component: Travels
    }
  ]
})
