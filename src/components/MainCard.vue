<template>
  <v-card
    class="mx-auto my-2"
    :raised="raised"
    :width="width"
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

    <v-card-actions>
      <div class="flex-grow-1" />
      <v-btn
        text
        color="deep-purple accent-4"
        :to="morelink"
        nuxt
      >
        Read
      </v-btn>
    </v-card-actions>
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
      type: Number,
      default: 350
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
