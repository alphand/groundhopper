import Vue from 'vue'
import VueRouter from 'vue-router'

import MainApp from './App'
import Home from './containers/home'

Vue.use(VueRouter)
const router = new VueRouter()

router.map({
  '/': {
    component: MainApp
  },
  '/home': {
    component: Home
  }
})

/* eslint-disable no-new */
// new Vue({
//   el: '#main-app',
//   components: { App }
// })

router.start(MainApp, '#main-app')
