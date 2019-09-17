import PhotoService from '@/services/PhotoService.js'

export const state = () => ({
  photos: [],
  folders: [],
  photo: {},
  panoramas: [],
  panorama: {},
  travels: [],
  travel: {}
})

export const mutations = {
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
  },
  SET_TRAVELS (state, travels) {
    state.travels = travels
  },
  SET_TRAVEL (state, travel) {
    state.travel = travel
  }
}

export const actions = {
  fetchPhotos ({ commit }) {
    return PhotoService.getPhotos().then((resp) => {
      commit('SET_PHOTOS', resp.data.Pictures)
      commit('SET_FOLDERS', resp.data.Directories)
    })
  },
  fetchPhotosByPath ({ commit }, path) {
    return PhotoService.getPhotosByPath(path).then((resp) => {
      commit('SET_PHOTOS', resp.data.Pictures)
    })
  },
  fetchPanoramas ({ commit }) {
    return PhotoService.getPanoramas().then((resp) => {
      commit('SET_PANORAMAS', resp.data.Pictures)
      commit('SET_FOLDERS', resp.data.Directories)
    })
  },
  fetchPanoramasByPath ({ commit }, path) {
    return PhotoService.getPanoramasByPath(path).then((resp) => {
      commit('SET_PANORAMAS', resp.data.Pictures)
    })
  },
  fetchTravels ({ commit }) {
    return PhotoService.getTravels().then((resp) => {
      commit('SET_TRAVELS', resp.data.Pictures)
      commit('SET_FOLDERS', resp.data.Directories)
    })
  },
  fetchTravelsByPath ({ commit }, path) {
    return PhotoService.getTravelsByPath(path).then((resp) => {
      commit('SET_TRAVELS', resp.data.Pictures)
    })
  }
}
