import axios from 'axios'

const apiClient = axios.create({
  baseURL: process.env.API_BASE_URL ? process.env.API_BASE_URL : 'http://localhost:8088',
  withCredentials: false,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json'
  }
})

export default {
  getPhotos () {
    return apiClient.get('/albums')
  },
  getPhotosByPath (path) {
    return apiClient.get('/albums/' + path)
  },
  getPanoramas () {
    return apiClient.get('/panorama')
  },
  getPanoramasByPath (path) {
    return apiClient.get('/panorama/' + path)
  },
  getTravels () {
    return apiClient.get('/travels')
  },
  getTravelsByPath (path) {
    return apiClient.get('/travels/' + path)
  }
}