<template>
  <v-layout
    fluid
  >
    <v-container
      fluid
    >
      <iframe
        src="https://www.google.com/maps/d/embed?mid=1j0XbxJDNzlkKL5xJ5jahy_qmXJQ"
      />
    </v-container>
  </v-layout>
</template>

<script>
export default {
  head () {
    return {
      title: 'Monterrey - August 2014'
    }
  }
}
</script>

<style lang="css">
iframe {
  width: 100vw;
  height: 90vh;
  min-width: 100%;
  min-height: 100%;
  border: none;
  padding: 0;
  margin: 0;
  overflow: hidden;
  z-index: 999999;
}
</style>
