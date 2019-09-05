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
      />
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
// import { mapState } from 'vuex'
import FolderCard from '@/components/FolderCard'
import PhotoCard from '@/components/PhotoCard'

export default {
  head () {
    return {
      title: 'Photos'
    }
  },
  components: {
    FolderCard,
    PhotoCard
  },
  asyncData ({ $axios, error }) {
    $axios.get('http://localhost:8088/original')
      .then((resp) => {
        return {
          photos: resp.data.Pictures,
          folders: resp.data.Directories
        }
      })
      .catch((e) => {
        error({
          statusCode: 503,
          message: `Unable to fetch photos at this time: ${e}`
        })
      })
  }
}
</script>
