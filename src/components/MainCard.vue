<template>
  <v-card
    class="mx-auto my-2"
    :raised="raised"
    :width="width"
    :to="morelink"
    nuxt
  >
    <v-img
      class="white--text"
      :height="imgheight"
      :src="imgsrc"
    />

    <v-card-title>
      {{ title }}
    </v-card-title>

    <v-card-text>
      {{ text }}
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      default: 'Title'
    },
    text: {
      type: String,
      default: ''
    },
    morelink: {
      type: String,
      default: ''
    },
    raised: {
      type: Boolean,
      default: true
    },
    width: {
      type: String,
      default: '350'
    },
    imgsrc: {
      type: String,
      default: 'https://cdn.vuetifyjs.com/images/cards/docks.jpg'
    },
    imgheight: {
      type: String,
      default: '200px'
    }
  }
}
</script>
