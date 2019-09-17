<template>
  <v-layout
    class="fill-height"
  >
    <v-container
      fluid
      class="d-flex flex-wrap mx-2 justify-space-around fill-height"
      flat
      tile
    >
      <PhotoCard
        v-for="(panorama, idx) in panoramas"
        :key="idx"
        :panoram="pano"
        :photo="panorama"
        @goto-pano="gotoPano(panorama)"
      />
      </v-card>
    </v-container>
  </v-layout>
</template>

<script>
import { mapState } from 'vuex'
import PhotoCard from '@/components/PhotoCard'

export default {
  head () {
    return {
      title: 'Panoramas'
    }
  },
  components: {
    PhotoCard
  },
  data () {
    return {
      overlay: false,
      pano: true,
      aPhoto: ''
    }
  },
  computed: mapState({
    panoramas: state => state.photos.panoramas
  }),
  async fetch ({ store, params, error }) {
    try {
      await store.dispatch('photos/fetchPanoramasByPath', params.album)
    } catch (e) {
      error({
        statusCode: 503,
        message: `Unable to fetch albums (panorama) at this time: ${e}`
      })
    }
  },
  methods: {
    gotoPano (panorama) {
      if (panorama.path.startsWith('/photos/panorama/puerto_vallarta_2014')) {
        this.$router.push(`/panorama/vuepano/?imgPath=${panorama.path}`)
      } else {
        this.$router.push(`/panorama/aframe/?imgPath=${panorama.path}`)
      }
    }
  }
}
</script>
