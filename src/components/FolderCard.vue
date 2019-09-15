<template>
  <v-card
    class="mx-auto my-2"
    raised
    width="300"
    height="300"
    :to="`/${basePath}/${folder.name}`"
    nuxt
  >
    <v-img
      class="white--text"
      height="180"
      :src="`/photos/${folder.thumb}`"
    />

    <v-card-title>
      {{ folder.name | fixName }}
    </v-card-title>

    <v-card-text class="d-flex justify-space-between">
      <div class="d-inline-flex grey--text subtitle-1">
        Items: {{ folder.items }}
      </div>
      <div class="d-inline-flex grey--text subtitle-1">
        <v-icon>mdi-calendar</v-icon>
        {{ folder.modified | humanDate }}
      </div>
    </v-card-text>
  </v-card>
</template>

<script>
import moment from 'moment'

export default {
  filters: {
    humanDate (date) {
      return moment(date).format('ddd, MMM Do YYYY')
    },
    fixName (title) {
      return title.replace(/_/g, ' ').replace(/\b[a-z]/g, char => char.toUpperCase())
    }
  },
  props: {
    folder: {
      type: Object,
      default: undefined
    },
    basePath: {
      type: String,
      default: 'photos'
    }
  }
}
</script>
