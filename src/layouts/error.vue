<template>
  <v-app dark>
    <v-content>
      <v-container fluid>
        <v-row>
          <v-col>
            <NuxtLink to="/">
              Home
            </NuxtLink>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <h1 v-if="error.statusCode === 404">
              {{ pageNotFound }}
            </h1>
            <h1 v-else>
              {{ otherError }}
            </h1>
            <pre>{{ errorDetails }}</pre>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
export default {
  layout: 'empty',
  props: {
    error: {
      type: Object,
      default: null
    }
  },
  head () {
    const title =
      this.error.statusCode === 404 ? this.pageNotFound : this.otherError
    return {
      title
    }
  },
  data () {
    return {
      pageNotFound: '404 Not Found',
      otherError: 'An error occurred'
    }
  },
  computed: {
    errorDetails () {
      return JSON.stringify(this.error)
    }
  }
}
</script>

<style scoped>
h1 {
  font-size: 20px;
}
</style>
