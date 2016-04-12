import Vue from 'vue'
import VueRouter from 'vue-router'

import MainApp from './App'
import Routes from './routes'

Vue.use(VueRouter)
const router = new VueRouter()
router.map(Routes)

router.start(MainApp, '#main-app')
