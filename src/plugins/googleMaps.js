import Vue from 'vue'
import 'vue-googlemaps/dist/vue-googlemaps.css'
import VueGoogleMaps from 'vue-googlemaps'

Vue.use(VueGoogleMaps, {
  load: {
    apiKey: 'AIzaSyDME0PPZpPsgair2k9U4iz9a0AP_T7w73A',
    libraries: ['places'],
    useBetaRenderer: false
  }
})
