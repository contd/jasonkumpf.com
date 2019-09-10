<template>
  <v-layout
    class="fill-height"
  >
    </v-container>
    <v-container
      fluid
      class="d-flex flex-wrap mx-2 justify-space-around fill-height"
      flat
      tile
    >
      <PhotoCard
        v-for="(panorama, idx) in panoramas"
        :key="idx"
        :photo="panorama"
        @click="showPhoto(panorama)"
      />
    </v-container>
    <v-overlay :value="overlay">
      <v-btn
        icon
        @click="closeOverlay"
      >
        <v-img
          :src="aPhoto.original"
          :alt="aPhoto.name"
        />
        <v-icon>mdi-close</v-icon>
      </v-btn>
    </v-overlay>
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
      aPhoto: ''
    }
  },
  computed: mapState({
    panoramas: state => state.panoramas.panoramas
  }),
  async fetch ({ store, params, error }) {
    try {
      await store.dispatch('panoramas/fetchPanoramasByPath', params.album)
    } catch (e) {
      error({
        statusCode: 503,
        message: `Unable to fetch albums (panorama) at this time: ${e}`
      })
    }
  },
  methods: {
    showPhoto (photo) {
      this.aPhoto = photo
      this.overlay = true
    },
    closeOverlay () {
      this.overlay = false
    }
  }
}
</script>
