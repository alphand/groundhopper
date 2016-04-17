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

  // var AccsDB = CouchDB.use('lvo-accounts')
  // AccsDB.view(
  //   'email_finder',
  //   'email_finder',
  //   state.email,
  //   (err, body) => {
  //     console.log('couch result', err, body)
  //
  //     if (err) {
  //       return dispatch(LOGIN_FAILED, {err, auth: false})
  //     }
  //
  //     dispatch(LOGIN_SUCCESSFULL, {body, auth: true})
  //   })

  // FirebaseRef
  //   .child('/users/' + state.email)
  //   .once('value', (snapshot) => {
  //     console.log('firebase snapshot', snapshot.val())
  //   }, (err) => {
  //     console.log('firebase err snapshot', err)
  //   })
}

export const loginSuccessfull = ({dispatch}, state) => {
  dispatch(LOGIN_SUCCESSFULL, state)
}

export const loginFailed = ({dispatch}, state) => {
  dispatch(LOGIN_FAILED, state)
}
