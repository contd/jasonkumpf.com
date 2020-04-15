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
      <img
        src="/logo.svg"
        height="50px"
        width="50px"
        contain
        position="left"
      />
      <span style="font-weight: 700;font-size: 3em;font-family: Liberation Mono;">jason</span><span style="font-weight: 400;font-size: 3em;font-family: Liberation Mono;">kumpf</span>
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
            <v-icon>mdi-github</v-icon>
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
        <transition name="slide-fade" mode="out-in">
          <nuxt />
        </transition>
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
      drawer: true,
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
        },
        {
          icon: 'mdi-panorama-horizontal',
          title: 'WebVR',
          to: '/panorama'
        },
        {
          icon: 'mdi-map',
          title: 'Maps',
          to: '/maps'
        }
      ]
    }
  }
}
</script>

<style lang="css">
.fade-enter {
  opacity: 0;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease-out;
}

.fade-leave-to {
  opacity: 0;
}

.slide-fade-enter {
  transform: translateX(10px);
  opacity: 0;
}

.slide-fade-enter-active,
.slide-fade-leave-active {
  transition: all 0.2s ease;
}

.slide-fade-leave-to {
  transform: translateX(-10px);
  opacity: 0;
}
</style>
