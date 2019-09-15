<template>
  <v-layout
    class="d-flex flex-row flex-wrap justify-space-around"
  >
    <MainCard
      title="About"
      text="Professional information about me; resume and portfolio."
      imgheight="200px"
      imgsrc="/img/me-icon.jpg"
      morelink="/about"
    />
    <MainCard
      title="Resume"
      text="Current and updated Resume and full CV. Also references to publications."
      imgheight="200px"
      imgsrc="/img/blog-front-icon.jpg"
      morelink="/resume"
    />
    <MainCard
      title="Photos"
      text="This is a place where I share old family photos scanned as well as new adventures."
      imgheight="200px"
      imgsrc="/img/delhi-house.jpg"
      morelink="/photos"
    />
    <MainCard
      title="Travels"
      text="Photos, maps and videos from my travels."
      imgheight="200px"
      imgsrc="/img/mexico-front.jpg"
      morelink="/travels"
    />
    <MainCard
      title="WebVR / Panoramas"
      text="WebVR and Panorama views from my travels and other places."
      imgheight="200px"
      imgsrc="/img/monterrey-3-2015.jpg"
      morelink="/panorama"
    />
    <MainCard
      title="Maps"
      text="Map with photo geolocated from my travels and other places."
      imgheight="200px"
      imgsrc="/img/monterrey-7-2015.jpg"
      morelink="/maps"
    />
  </v-layout>
</template>

<script>
import MainCard from '@/components/MainCard'

export default {
  components: {
    MainCard
  }
}
</script>
