<template>
  <v-card
    class="mx-auto my-2"
    raised
    width="300"
    height="300"
  >
    <v-img
      class="white--text"
      height="180"
      :src="folder.thumb"
    >
    </v-img>

    <v-card-title>
      <nuxt-link
        :to="`/photos/${folder.name}`"
      >
        {{ folder.name | fixName }}
      </nuxt-link>
    </v-card-title>

    <v-card-text>
      <span class="grey--text subtitle-1">
        Created: {{ folder.modified | humanDate }}
      </span>
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
      return title.replace(/_/g, ' ')
    }
  },
  props: {
    folder: {
      // eslint-disable-next-line vue/require-default-prop
      type: Object
    }
  }
}
</script>
