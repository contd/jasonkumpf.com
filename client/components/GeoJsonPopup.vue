<template>
  <div>
    <div class="subtitle-2">
      {{ icnNum }}. {{ name }}
    </div>
    <div v-html="text" />
  </div>
</template>

<script>
export default {
  name: 'GeoJson2Popup',
  props: {
    icnNum: {
      type: String,
      default: ''
    },
    name: {
      type: String,
      default: ''
    },
    text: {
      type: String,
      default: ''
    }
  }
}
</script>
