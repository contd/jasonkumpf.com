<template>
  <v-layout
    fluid
  >
    <v-container
      fluid
    >
      <iframe
        :src="map[mapId]"
      />
    </v-container>
  </v-layout>
</template>

<script>
export default {
  head () {
    return {
      title: 'Monterrey - August 2014'
    }
  },
  validate ({ params }) {
    if (params && params.id) {
      return true
    }
  },
  data () {
    return {
      map: [
        { url: 'https://www.google.com/maps/d/embed?mid=13PagFei3Yz0lHY1qpvbN6TCFGnA' },
        { url: 'https://www.google.com/maps/d/embed?mid=1j0XbxJDNzlkKL5xJ5jahy_qmXJQ' }
      ],
      mapId: 0
    }
  },
  created () {
    this.mapId = (this.$route.param && this.$route.param.id ? this.$route.param.id : 0)
  }
}
</script>

<style lang="css">
iframe {
  width: 100vw;
  height: 90vh;
  min-width: 100%;
  min-height: 100%;
  border: none;
  padding: 0;
  margin: 0;
  overflow: hidden;
  z-index: 999999;
}
</style>
