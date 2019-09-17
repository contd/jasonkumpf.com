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
        v-for="(photo, idx) in photos"
        :key="idx"
        :photo="photo"
        @show-img="showPhoto(photo)"
      />
      <fullscreen ref="fullscreen" @change="fullscreenChange">
        <v-card
          @click="closeOverlay"
        >
          <v-img
            :src="aPhoto.path"
            :alt="aPhoto.name"
            :aspect-ratio="getAspect(aPhoto)"
            max-height="100vh"
            contain
          />
        </v-card>
      </fullscreen>
    </v-container>
  </v-layout>
</template>

<script>
import { mapState } from 'vuex'
import fullscreen from 'vue-fullscreen'
import Vue from 'vue'
import PhotoCard from '@/components/PhotoCard'

Vue.use(fullscreen)

export default {
  head () {
    return {
      title: 'Photos'
    }
  },
  components: {
    PhotoCard
  },
  data () {
    return {
      fullscreen: false,
      aPhoto: ''
    }
  },
  computed: mapState({
    photos: state => state.photos.photos
  }),
  async fetch ({ store, params, error }) {
    try {
      await store.dispatch('photos/fetchPhotosByPath', params.album)
    } catch (e) {
      error({
        statusCode: 503,
        message: `Unable to fetch photos at this time: ${e}`
      })
    }
  },
  methods: {
    getAspect (photo) {
      return photo.width / photo.height
    },
    showPhoto (photo) {
      this.aPhoto = photo
      this.$refs.fullscreen.toggle()
    },
    closeOverlay () {
      this.$refs.fullscreen.toggle()
    },
    fullscreenChange (fullscreen) {
      this.fullscreen = fullscreen
    }
  }
}
</script>
