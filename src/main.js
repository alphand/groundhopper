import Vue from 'vue'
import VueRouter from 'vue-router'
import VueValidator from 'vue-validator'

import MainApp from './App'
import Routes from './routes'

Vue.use(VueRouter)
Vue.use(VueValidator)

const router = new VueRouter()
router.map(Routes)

router.start(MainApp, '#main-app')
