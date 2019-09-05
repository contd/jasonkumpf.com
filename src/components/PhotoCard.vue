<template>
  <v-card
    class="mx-auto my-2"
    :raised="raised"
    width="240"
  >
    <v-img
      class="white--text"
      height="240"
      :src="thumb"
    />

    <v-card-title>
      <v-link
        :to="path"
        nuxt
      >
        {{ name }}
      </v-link>
    </v-card-title>

    <v-card-text>
      {{ size | humanSize }}
      {{ modified | humanDate }}
    </v-card-text>

    </v-card-actions>
  </v-card>
</template>

<script>
import moment from 'moment'

export default {
  filters: {
    humanDate (date) {
      return moment(date).format('ddd, MMM Do YYYY')
    },
    humanSize (num) {
      return this.prettyBytes(num)
    }
  },
  props: {
    name: {
      type: String,
      default: 'Photo Name'
    },
    path: {
      type: String,
      default: ''
    },
    thumb: {
      type: String,
      default: ''
    },
    size: {
      type: Number,
      default: 0
    },
    width: {
      type: Number,
      default: 240
    },
    modified: {
      type: String,
      default: ''
    }
  },
  methods: {
    prettyBytes (num, precision = 3, addSpace = true) {
      const UNITS = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']

      if (Math.abs(num) < 1) {
        return num + (addSpace ? ' ' : '') + UNITS[0]
      }

      const exponent = Math.min(Math.floor(Math.log10(num < 0 ? -num : num) / 3), UNITS.length - 1)
      const n = Number(((num < 0 ? -num : num) / 1000 ** exponent).toPrecision(precision))

      return (num < 0 ? '-' : '') + n + (addSpace ? ' ' : '') + UNITS[exponent]
    }
  }
}
</script>
