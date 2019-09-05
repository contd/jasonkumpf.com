import axios from 'axios'

const apiClient = axios.create({
  baseURL: `http://localhost:8088`,
  withCredentials: false,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json'
  }
})

export default {
  getPhotos () {
    return apiClient.get('/original')
  },
  getPhotosByPath (path) {
    return apiClient.get('/original?path=' + path)
  },
  getPanoramas () {
    return apiClient.get('/panorama')
  },
  getPanoramasByPath (path) {
    return apiClient.get('/panorama?path=' + path)
  }
}
