<template>
  <v-card
    class="mx-auto my-2"
    raised
    width="300"
    height="300"
  >
    <v-img
      class="white--text"
      height="240"
      :src="photo.thumb"
    />

    <v-card-actions>
      <div>
        {{ photo.modified | humanDate }}
      </div>
      <div class="flex-grow-1" />
      <div>
        {{ photo.size | humanSize }}
      </div>
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
    humanSize (num, precision = 3, addSpace = true) {
      const UNITS = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']

      if (Math.abs(num) < 1) {
        return num + (addSpace ? ' ' : '') + UNITS[0]
      }

      const exponent = Math.min(Math.floor(Math.log10(num < 0 ? -num : num) / 3), UNITS.length - 1)
      const n = Number(((num < 0 ? -num : num) / 1000 ** exponent).toPrecision(precision))

      return (num < 0 ? '-' : '') + n + (addSpace ? ' ' : '') + UNITS[exponent]
    }
  },
  props: {
    photo: {
      type: Object,
      default: undefined
    }
  }
}
</script>
