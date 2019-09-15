<template>
  <a-scene>
    <a-sky id="image-360" radius="10" :src="imgPath" />
  </a-scene>
</template>

<script>
import 'aframe'

export default {
  layout: 'aframe',
  head () {
    return {
      title: '360° Viewer'
    }
  },
  validate ({ params, query, error }) {
    if (query && query.imgPath) {
      return true
    }
    return true
  },
  data () {
    return {
      imgPath: ''
    }
  },
  created () {
    this.imgPath = (this.$route.query && this.$route.query.imgPath ? this.$route.query.imgPath : '')
  }
}
</script>
