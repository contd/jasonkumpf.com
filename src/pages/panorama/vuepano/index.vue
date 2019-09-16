<template>
  <div
    class="vuepano"
  >
    <Panorama
      :source="imgPath"
      @viewchange="onChange"
    />
  </div>
</template>

<script>
import Panorama from '@/components/Panorama'

export default {
  layout: 'aframe',
  head () {
    return {
      title: 'Vue Panorama Viewer'
    }
  },
  validate ({ params, query, error }) {
    if (query && query.imgPath) {
      return true
    }
    return true
  },
  components: {
    Panorama
  },
  data () {
    return {
      imgPath: ''
    }
  },
  created () {
    this.imgPath = (this.$route.query && this.$route.query.imgPath ? this.$route.query.imgPath : '')
  },
  methods: {
    onChange (e, angle) {
      console.log(e, angle)
    }
  }
}
</script>

<style>
.vuepano {
  height: 100vh;
  width: 100vw;
}
</style>
