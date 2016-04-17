import Vue from './apiBase'
import config from '../config'

export const getLogin = (email, password) => {
  const accEndPoints = config.DBRoot + 'lvo-accounts/' +
    '/lvo-accounts/_design/email_finder/_view/email_finder'

  const opts = {
    url: accEndPoints,
    method: 'GET',
    params: {
      key: email
    }
  }

  return new Promise((resolve, reject) => {
    Vue.http.get(opts)
      .then((resp) => {
        console.log('respdata', resp)
        resolve(resp)
      })
      .catch((err) => {
        reject(err)
      })
  })
}
