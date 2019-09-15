<template>
  <v-card
    class="mx-auto my-2"
    raised
    width="300"
    height="300"
    @click="showImg"
  >
    <v-img
      class="white--text"
      height="240"
      :src="photo.thumb"
    >
      <v-card-title class="align-end fill-height">
        {{ photo.name | fixName }}
      </v-card-title>
    </v-img>

    <v-card-actions>
      <div>
        <v-icon>mdi-calendar</v-icon>
        {{ createdAt }}
      </div>
      <div class="flex-grow-1" />
      <div>
        <v-icon>mdi-harddisk</v-icon>
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
    },
    fixName (title) {
      return title.replace(/\.jpg$/, '').replace(/_/g, ' ').replace(/-/g, ' ').replace(/\b[a-z]/g, char => char.toUpperCase())
    }
  },
  props: {
    photo: {
      type: Object,
      default: undefined
    },
    panoram: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    createdAt () {
      if (this.photo.exif && this.photo.exif.taken) {
        return moment(this.photo.exif.taken).format('ddd, MMM Do YYYY')
      } else {
        return moment(this.photo.modified).format('ddd, MMM Do YYYY')
      }
    }
  },
  methods: {
    showImg () {
      if (this.panoram) {
        this.$emit('goto-pano')
      } else {
        this.$emit('show-img')
      }
    }
  }
}
</script>
