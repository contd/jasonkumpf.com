import Vue from 'vue'
import axios from '@/plugins/axios'
import App from '@/App.vue'
import router from '@/router'
import store from '@/store'
import vuetify from '@/plugins/vuetify';

Vue.config.productionTip = false

new Vue({
  router,
  store,
  created () {
    const userString = localStorage.getItem('user')
    if (userString) {
      const userData = JSON.parse(userString)
      this.$store.commit('SET_USER_DATA', userData)
    }
    axios.interceptors.response.use(
      // simply return the response
      response => response,
      error => {
        // if we catch a 401 error
        if (error.response.status === 401) {
          // force a log out
          this.$store.dispatch('logout')
        }
        // reject the Promise, with the error as the reason
        return Promise.reject(error)
      }
    )
  },
  vuetify,
  render: h => h(App)
}).$mount('#app')
