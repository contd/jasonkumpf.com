<template>
  <v-container
    fluid
    class="d-flex flex-wrap mx-2 justify-space-around fill-height"
  >
    <MainCard
      title="Monterrey, Mexico"
      text="August 2014"
      imgheight="300px"
      width="400"
      imgsrc="/img/monterrey-8-2014.jpg"
      morelink="/maps/map1"
    />
    <MainCard
      title="Monterrey, Mexico"
      text="July 2015"
      imgheight="300px"
      width="400"
      imgsrc="/img/monterrey-7-2015.jpg"
      morelink="/maps/map2"
    />
  </v-container>
</template>

<script>
import MainCard from '@/components/MainCard'

export default {
  name: 'Maps',
  head () {
    return {
      title: 'Maps'
    }
  },
  components: {
    MainCard
  }
}
</script>
