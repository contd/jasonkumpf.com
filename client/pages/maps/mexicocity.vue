<template>
  <v-container
    fluid
  >
    <v-row no-gutters>
      <v-col cols="12">
        <client-only>
          <l-map
            style="height: 80vh; width: 100%"
            :zoom="zoom"
            :center="center"
            @update:zoom="zoomUpdated"
            @update:center="centerUpdated"
            @update:bounds="boundsUpdated"
          >
            <l-tile-layer :url="url" />
            <l-geo-json
              :geojson="geojson"
              :options="options"
            />
          </l-map>
        </client-only>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Vue from 'vue'
import GeoJsonService from '@/services/GeoJsonService'
import GeoJsonPopup from '@/components/GeoJsonPopup'

function onEachFeature (feature, layer) {
  const GeoJsonPop = Vue.extend(GeoJsonPopup)
  const popup = new GeoJsonPop({
    propsData: {
      name: feature.properties.name.replace(/\.jpg$/, '').replace(/\.jpeg$/, ''),
      taken: feature.properties.taken,
      thumb: feature.properties.thumb,
      path: feature.properties.path
    }
  })
  layer.bindPopup(popup.$mount().$el)
}

export default {
  name: 'MexicoCity',
  head () {
    return {
      title: 'Mexico City - October 2016'
    }
  },
  data () {
    return {
      url: 'http://{s}.tile.osm.org/{z}/{x}/{y}.png',
      zoom: 8,
      center: [20.373829, -103.576461],
      bounds: null,
      enableTooltip: true,
      geojson: null
    }
  },
  computed: {
    options () {
      return {
        onEachFeature
      }
    }
  },
  async asyncData ({ error }) {
    try {
      const resp = await GeoJsonService.getGeoJson('MexicoCity')
      return {
        geojson: resp
      }
    } catch (e) {
      error({
        statusCode: 503,
        message: `Unable to fetch geojson at this time: ${e}`
      })
    }
  },
  methods: {
    zoomUpdated (zoom) {
      this.zoom = zoom
    },
    centerUpdated (center) {
      this.center = center
    },
    boundsUpdated (bounds) {
      this.bounds = bounds
    }
  }
}
</script>
