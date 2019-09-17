<template>
  <client-only>
    <a-scene>
      <a-sky
        id="image-360"
        radius="10"
        :position="panoramaPosition"
        :rotation="panoramaRotation"
        :src="imgPath"
      />
    </a-scene>
  </client-only>
</template>

<script>
if (process.browser) {
  require('aframe')
}

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
      imgPath: '',
      panoramaPosition: '',
      panoramaRotation: ''
    }
  },
  created () {
    this.imgPath = (this.$route.query && this.$route.query.imgPath ? this.$route.query.imgPath : '')
  }
}
</script>
