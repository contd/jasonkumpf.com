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
      />
    </v-container>
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
  }
}
</script>
