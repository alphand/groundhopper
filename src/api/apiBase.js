import Vue from 'vue'
import VueResource from 'vue-resource'

Vue.use(VueResource)

Vue.http.headers.common['Access-Control-Allow-Origin'] = '*'

export default Vue
