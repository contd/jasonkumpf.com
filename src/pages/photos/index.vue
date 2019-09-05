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
      <div
        v-for="(folder, i) in folders"
        :key="i"
      >
        <nuxt-link :to="`/photos?path=${folder.name}`">
          {{ folder.name }}
        </nuxt-link>
      </div>
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
        :name="photo.name"
        :path="photo.path"
        :thumb="photo.thumb"
        :size="photo.size"
        :width="photo.width"
        :modified="photo.modified"
      />
    </v-container>
  </v-layout>
</template>

<script>
// import { mapState } from 'vuex'
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
      links: [],
      folders: []
    }
  },
  async asyncData ({ $axios, error }) {
    try {
      const files = await $axios.get('/original')
      return {
        photos: files.Pictures,
        folders: files.Directories
      }
    } catch (e) {
      error({
        statusCode: 503,
        message: `Unable to fetch photos at this time: ${e}`
      })
    }
  }
}
</script>
