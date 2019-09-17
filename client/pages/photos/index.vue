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
      <FolderCard
        v-for="(folder, i) in folders"
        :key="i"
        :folder="folder"
        base-path="photos"
      />
    </v-container>
  </v-layout>
</template>

<script>
import { mapState } from 'vuex'
import FolderCard from '@/components/FolderCard'

export default {
  head () {
    return {
      title: 'Photos'
    }
  },
  components: {
    FolderCard
  },
  computed: mapState({
    folders: state => state.photos.folders
  }),
  async fetch ({ store, params, error }) {
    try {
      await store.dispatch('photos/fetchPhotos')
    } catch (e) {
      error({
        statusCode: 503,
        message: `Unable to fetch photos at this time: ${e}`
      })
    }
  }
}
</script>
