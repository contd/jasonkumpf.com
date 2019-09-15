<template>
  <v-layout
    class="d-flex flex-row flex-wrap justify-space-around"
  >
    <v-container>
      <v-row
        no-gutters
      >
        <v-col>
          <MainCard
            title="Monterrey, Mexico"
            text="August 2014"
            imgheight="200px"
            width="300"
            imgsrc="/img/monterrey-8-2014.jpg"
            morelink="/maps/map1"
          />
          <MainCard
            title="Monterrey, Mexico"
            text="July 2015"
            imgheight="200px"
            width="300"
            imgsrc="/img/monterrey-7-2015.jpg"
            morelink="/maps/map2"
          />
        </v-col>
        <v-col>
          <googlemaps-map
            class="vgooglemap"
            :center.sync="center"
            :zoom.sync="zoom"
            :options="mapOptions"
            @click="centerOnDefault"
          >
            <googlemaps-user-position
              @update:position="setDefaultPosition"
            />
          </googlemaps-map>
        </v-col>
      </v-row>
    </v-container>
  </v-layout>
</template>

<script>
import MainCard from '@/components/MainCard'
// import * as monterrey2014 from '@/services/Monterrey-Sep-2014.json'

export default {
  name: 'Maps',
  head () {
    return {
      title: 'Maps'
    }
  },
  components: {
    MainCard
  },
  data () {
    return {
      zoom: 12,
      userPosition: null,
      center: {
        lat: 25.6793307,
        lng: -100.3081712
      },
      mapOptions: {}
    }
  },
  methods: {
    centerOnDefault () {
      if (this.userPosition) {
        this.center = {
          lat: 25.6793307,
          lng: -100.3081712
        }
      }
    },
    setDefaultPosition (position) {
      this.userPosition = position
    }
  }
}
</script>

<style lang="css">
.vgooglemap {
  height: 650px;
  width: 900px;
}
</style>
