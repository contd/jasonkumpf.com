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
  components: {
    FolderCard,
    PhotoCard
  },
  created () {
    this.axios.get('/albums')
      .then((resp) => {
        return {
          photos: resp.data.Pictures,
          folders: resp.data.Directories
        }
      })
      .catch((e) => {
        console.log(`Error Getting Photo Albums: ${e}`)
      })
  }
}
</script>
