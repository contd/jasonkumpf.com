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
        v-for="(photo, idx) in photos"
        :key="idx"
        :photo="photo"
        @click="showPhoto(photo)"
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
      title: 'Photos'
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
