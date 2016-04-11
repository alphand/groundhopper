import Home from './containers/home'
import ErrorPage from './containers/404'

import Login from './containers/accounts/login'

export default {
  '/': {
    component: Home
  },
  '/accounts/login': {
    component: Login
  },
  '/*any': {
    component: ErrorPage
  }
}

