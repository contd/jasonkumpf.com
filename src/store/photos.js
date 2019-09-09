import PhotoService from '@/services/PhotoService.js'

export const state = () => ({
  photos: [],
  folders: [],
  photo: {},
  panoramas: [],
  panorama: {}
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
      commit('SET_PHOTOS', resp.data.Pictures)
      commit('SET_FOLDERS', resp.data.Directories)
    })
  },
  fetchPanoramasByPath ({ commit }, path) {
    return PhotoService.getPanoramasByPath(path).then((resp) => {
      commit('SET_PHOTOS', resp.data)
    })
  }
}
