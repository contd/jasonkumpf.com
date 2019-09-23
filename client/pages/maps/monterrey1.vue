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
import GeoJsonPopup from '@/components/GeoJsonPopup'
const monterrey2014 = require('@/services/monterrey-09-2014.json')

function onEachFeature (feature, layer) {
  const GeoJsonPop = Vue.extend(GeoJsonPopup)
  const popup = new GeoJsonPop({
    propsData: {
      name: feature.properties.name,
      text: feature.properties.description
    }
  })
  layer.bindPopup(popup.$mount().$el)
}

export default {
  name: 'Monterrey2014',
  head () {
    return {
      title: 'Monterrey - August 2014'
    }
  },
  data () {
    return {
      url: 'http://{s}.tile.osm.org/{z}/{x}/{y}.png',
      zoom: 11,
      center: [25.6848304, -100.3126147],
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
  created () {
    this.geojson = monterrey2014
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
