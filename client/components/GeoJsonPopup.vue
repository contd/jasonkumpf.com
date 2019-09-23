<template>
  <div>
    <div class="subtitle-2">
      {{ name }}
    </div>
    <div v-if="text" v-html="text" />
    <div v-else>
      <a :href="path" target="_new">
        <img :src="thumb" :alt="thumb" />
      </a>
    </div>
  </div>
</template>

<script>
export default {
  name: 'GeoJson2Popup',
  props: {
    name: {
      type: String,
      default: ''
    },
    thumb: {
      type: String,
      default: ''
    },
    taken: {
      type: String,
      default: ''
    },
    path: {
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
