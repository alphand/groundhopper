import Vue from './apiBase'

export const getLogin = (email, password) => {
  // const accEndPoints = '/api/proxy/192.168.99.100:5984/lvo-accounts/' +
  //   '_design/email_finder/_view/email_finder'

  const createSessionEP = '/api/sessions/create'

  const opts = {
    url: createSessionEP,
    method: 'POST',
    data: {
      email: email,
      password: password
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
