import {
  LOGIN_REQUEST,
  LOGIN_SUCCESSFULL,
  LOGIN_FAILED
} from './constants'

export const loginRequest = ({dispatch}, state) => {
  dispatch(LOGIN_REQUEST, state)
  setTimeout(() => {
    if (Math.random() <= 0.5) {
      dispatch(LOGIN_SUCCESSFULL, {auth: true})
    } else {
      dispatch(LOGIN_FAILED, {auth: false})
    }
  }, 500)
}

export const loginSuccessfull = ({dispatch}, state) => {
  dispatch(LOGIN_SUCCESSFULL, state)
}

export const loginFailed = ({dispatch}, state) => {
  dispatch(LOGIN_FAILED, state)
}
