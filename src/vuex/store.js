import Vue from 'vue'
import Vuex from 'vuex'

import {
  LOGIN_REQUEST,
  LOGIN_SUCCESSFULL,
  LOGIN_FAILED
} from './constants'

Vue.use(Vuex)

const setLoader = (state, val = false) => {
  state.isLoading = val
}

const state = {
  user: {
    authenticated: false
  },
  isLoading: false
}

const mutations = {
  [LOGIN_REQUEST] (state, newstate) {
    setLoader(state, true)
    console.log('handling login request', state, newstate)
  },
  [LOGIN_SUCCESSFULL] (state, newstate) {
    setLoader(state)
    state = {state, ...newstate}
    console.log('handling login successfull', state, newstate)
  },
  [LOGIN_FAILED] (state, newstate) {
    setLoader(state)
    state = {state, ...newstate}
    console.log('handling login failed', state, newstate)
  }
}

export default new Vuex.Store({
  state,
  mutations
})
