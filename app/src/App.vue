<template>
  <v-app dark>
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      clipped
      fixed
      app
    >
      <NavBar></NavBar>
    </v-navigation-drawer>

    <v-app-bar
      style="background-color:#343a40;"
      dark
      clipped-left
      fixed
      app
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-img
        src="/jasonkumpf.svg"
        height="50px"
        width="250px"
        contain
        position="left"
      />
      <div class="flex-grow-1" />
      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-btn icon href="https://blog.kumpf.io" target="_new" v-on="on">
            <v-icon>mdi-post</v-icon>
          </v-btn>
        </template>
        <span>My Blog</span>
      </v-tooltip>

      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-btn icon href="https://github.com/contd" target="_new" v-on="on">
            <v-icon>mdi-github-circle</v-icon>
          </v-btn>
        </template>
        <span>Github</span>
      </v-tooltip>

      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-btn icon href="https://instagram.com/kumpf.jason" target="_new" v-on="on">
            <v-icon>mdi-instagram</v-icon>
          </v-btn>
        </template>
        <span>Instagram</span>
      </v-tooltip>
    </v-app-bar>

    <v-content>
      <v-container fluid>
        <nuxt />
      </v-container>
    </v-content>

    <v-footer
      app
    >
      <span>&copy; 2019 {{ author }} </span>
    </v-footer>
  </v-app>
</template>

<script>
import NavBar from '@/components/NavBar'

export default {
  name: 'App',
  components: {
    NavBar
  },
  data () {
    return {
      author: 'Jason Kumpf',
      drawer: false,
      miniVariant: false
    }
  }
}
</script>
