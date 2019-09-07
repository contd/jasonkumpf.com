<template>
  <v-list>
    <v-list-item
      to="/"
      router
      exact
    >
      <v-list-item-action>
        <v-icon>mdi-apps</v-icon>
      </v-list-item-action>
      <v-list-item-content>
        <v-list-item-title v-text="Home" />
      </v-list-item-content>
    </v-list-item>
    <v-list-item
      to="/about"
      router
      exact
    >
      <v-list-item-action>
        <v-icon>mdi-account-circle</v-icon>
      </v-list-item-action>
      <v-list-item-content>
        <v-list-item-title v-text="About" />
      </v-list-item-content>
    </v-list-item>
    <v-list-item
      to="/resume"
      router
      exact
    >
      <v-list-item-action>
        <v-icon>mdi-file-account</v-icon>
      </v-list-item-action>
      <v-list-item-content>
        <v-list-item-title v-text="Resume" />
      </v-list-item-content>
    </v-list-item>
    <v-list-item
      to="/photos"
      router
      exact
    >
      <v-list-item-action>
        <v-icon>mdi-camera</v-icon>
      </v-list-item-action>
      <v-list-item-content>
        <v-list-item-title v-text="Photos" />
      </v-list-item-content>
    </v-list-item>
    <v-list-item
      to="/travel"
      router
      exact
    >
      <v-list-item-action>
        <v-icon>mdi-airplane</v-icon>
      </v-list-item-action>
      <v-list-item-content>
        <v-list-item-title v-text="Travel" />
      </v-list-item-content>
    </v-list-item>
  </v-list>
</template>

<script>
export default {
  name: 'NavBar'
}
</script>
