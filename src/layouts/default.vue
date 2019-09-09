<template>
  <v-app dark>
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      clipped
      fixed
      app
    >
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
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
export default {
  head () {
    return {
      title: this.title,
      meta: [
        { hid: 'description', name: 'description', content: process.env.npm_package_description }
      ]
    }
  },
  data () {
    return {
      title: 'Jason Kumpf',
      author: 'Jason Kumpf',
      drawer: false,
      miniVariant: false,
      items: [
        {
          icon: 'mdi-apps',
          title: 'Home',
          to: '/'
        },
        {
          icon: 'mdi-account-circle',
          title: 'About',
          to: '/about'
        },
        {
          icon: 'mdi-file-account',
          title: 'Resume',
          to: '/resume'
        },
        {
          icon: 'mdi-camera',
          title: 'Photos',
          to: '/photos'
        },
        {
          icon: 'mdi-airplane',
          title: 'Travels',
          to: '/travels'
        }
      ]
    }
  }
}
</script>
