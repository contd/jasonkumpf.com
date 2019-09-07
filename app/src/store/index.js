import Vue from 'vue'
import Vuex from 'vuex'
import axios from '@/plugins/axios'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    user: null,
    isNewUser: true,
    photos: [],
    folders: [],
    photo: {},
    panoramas: [],
    panorama: {}
  },
  mutations: {
    SET_USER_DATA (state, userData) {
      localStorage.setItem('user', JSON.stringify(userData))
      axios.defaults.headers.common['Authorization'] = `Bearer ${
        userData.token
      }`
      state.user = userData
    },
    LOGOUT () {
      localStorage.removeItem('user')
      location.reload()
    },
    IS_NEW_USER (state, isNewUser) {
      state.isNewUser = isNewUser
    },
    SET_PHOTOS (state, photos) {
      state.photos = photos
    },
    SET_FOLDERS (state, folders) {
      state.folders = folders
    },
    SET_PHOTO (state, photo) {
      state.photo = photo
    },
    SET_PANORAMAS (state, panoramas) {
      state.panoramas = panoramas
    },
    SET_PANORAMA (state, panorama) {
      state.panorama = panorama
    }
  },
  actions: {
    register ({ commit }, credentials) {
      return axios.post('/register', credentials)
        .then(({ data }) => {
          commit('SET_USER_DATA', data)
        })
    },
    login ({ commit }, credentials) {
      return axios
        .post('/login', credentials)
        .then(({ data }) => {
          commit('SET_USER_DATA', data)
        })
    },
    logout ({ commit }) {
      commit('LOGOUT')
    },
    isNewUser ({ commit }, isNewUser) {
      commit('IS_NEW_USER', isNewUser)
    },
    getAlubms ({ commit }) {
      return axios.get('/albums').then((resp) => {
        commit('SET_FOLDERS', resp.data.Directories)
      })
    },
    getPhotos ({ commit }, album) {
      return axios.get(`/albums?path=${album}`).then((resp) => {
        commit('SET_PHOTOS', resp.data)
      })
    },
    fetchPanoramas ({ commit }) {
      return axios.getP('/panorama').then((resp) => {
        commit('SET_FOLDERS', resp.data.Directories)
      })
    },
    fetchPanoramasByPath ({ commit }, path) {
      return axios.get(`/panorama?path=${path}`).then((resp) => {
        commit('SET_PANORAMAS', resp.data)
      })
    }
  },
  getters: {
    loggedIn (state) {
      return !!state.user
    }
  }
})
