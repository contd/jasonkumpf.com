import axios from 'axios'

const apiClient = axios.create({
  baseURL: process.env.API_BASE_URL,
  withCredentials: false,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json'
  }
})

export default {
  getGeoJson (path) {
    return apiClient.get('/travels/' + path)
      .then((resp) => {
        return resp.data.Pictures
      })
      .then((pictures) => {
        return pictures.map((picture) => {
          return {
            type: 'Feature',
            geometry: {
              type: 'Point',
              coordinates: [
                picture.exif.long,
                picture.exif.lat
              ]
            },
            properties: {
              name: picture.name,
              taken: picture.exif.taken,
              thumb: picture.thumb,
              path: picture.path
            }
          }
        })
      })
      .then((features) => {
        return {
          type: 'FeatureCollection',
          features
        }
      })
  }
}
