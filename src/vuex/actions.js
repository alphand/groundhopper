import {
  LOGIN_REQUEST,
  LOGIN_SUCCESSFULL,
  LOGIN_FAILED
} from './constants'

import {AccountsAPI} from '../api'

export const loginRequest = ({dispatch}, state) => {
  dispatch(LOGIN_REQUEST, state)

  AccountsAPI.getLogin(state.email, state.password)
    .then((data) => {
      console.log('login data', data)
    })
    .catch((err) => {
      console.log('err login data', err)
    })
}

export const loginSuccessfull = ({dispatch}, state) => {
  dispatch(LOGIN_SUCCESSFULL, state)
}

export const loginFailed = ({dispatch}, state) => {
  dispatch(LOGIN_FAILED, state)
}
