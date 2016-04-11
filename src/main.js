import Vue from 'vue'
import VueRouter from 'vue-router'

import MainApp from './App'

import Home from './containers/home'
import ErrorPage from './containers/404'

Vue.use(VueRouter)
const router = new VueRouter()

router.map({
  '/': {
    component: Home
  },
  '*': {
    component: ErrorPage
  }
})

router.start(MainApp, '#main-app')
