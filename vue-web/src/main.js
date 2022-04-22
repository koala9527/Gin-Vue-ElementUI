import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'

Vue.config.productionTip = false
Vue.prototype.$axios = axios
if (process.env.ENV == 'production') {
  axios.defaults.baseURL = '/api'  //关键代码
} else {
  axios.defaults.baseURL = process.env.VUE_APP_BASE_API //关键代码
}

axios.defaults.headers.post['Content-Type'] = 'application/json';

new Vue({
  render: h => h(App),
}).$mount('#app')
