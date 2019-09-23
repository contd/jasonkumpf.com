import GeoJsonService from '@/services/GeoJsonService'

export const state = () => ({
  geojson: {}
})

export const mutations = {
  SET_GEOJSON (state, geojson) {
    state.geojson = geojson
  }
}

export const actions = {
  fetchGeoJson ({ commit }, path) {
    return GeoJsonService.getGeoJson(path).then((resp) => {
      commit('SET_GEOJSON', resp)
    })
  }
}
