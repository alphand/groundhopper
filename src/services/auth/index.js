import {AccountsAPI} from '../api/accounts'

export default {

  checkAuth() {
    var jwt = this.getUserToken()
    if (jwt) {
      this.setUserAuth(true)
    } else {
      this.setUserAuth(false)
    }
  }

  getUserToken() {
    return localStorage.getItem('id_token_gh')
  }

  getAuthHeader() {
    return {
      'Authorization': 'Bearer ' + this.getUserToken()
     }
  }

  setUserAuth(status) {
    this.user.authenticated = status
  }

  setToken(token) {
    localStorage.setItem('id_token_gh', token)
  }

}
