<template>
  <v-layout
    class="d-flex flex-row flex-wrap justify-space-around"
  >
    <MainCard
      title="Monterrey, Mexico"
      text="August 2014"
      imgheight="200px"
      imgsrc="/img/monterrey-8-2014.jpg"
      morelink="/travels/map1"
    />
    <MainCard
      title="Monterrey, Mexico"
      text="July 2015"
      imgheight="200px"
      imgsrc="/img/monterrey-7-2015.jpg"
      morelink="/travels/map2"
    />
  </v-layout>
</template>

<script>
import MainCard from '@/components/MainCard'

export default {
  name: 'Travels',
  head () {
    return {
      title: 'Travels'
    }
  },
  components: {
    MainCard
  }
}
</script>
