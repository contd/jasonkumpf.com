<template>
  <v-container
    fluid
    class="d-flex flex-wrap mx-2 justify-space-around fill-height"
  >
    <MainCard
      title="Monterrey, Mexico"
      text="August 2014"
      imgheight="200px"
      width="300"
      imgsrc="/img/monterrey-8-2014.jpg"
      morelink="/maps/monterrey1"
    />
    <MainCard
      title="Puerto Vallarta, Mexico"
      text="December 2014"
      imgheight="200px"
      width="300"
      imgsrc="/photos/travels/PuertoVallarta/PuertoVallarta-001.jpg"
      morelink="/maps/puertovallarta"
    />
    <MainCard
      title="Sayulita, Mexico"
      text="December 2014"
      imgheight="200px"
      width="300"
      imgsrc="/photos/travels/Sayulita/Sayulita-001.jpg"
      morelink="/maps/sayulita"
    />
    <MainCard
      title="Monterrey, Mexico"
      text="March 2015"
      imgheight="200px"
      width="300"
      imgsrc="/img/monterrey-3-2015.jpg"
      morelink="/maps/monterrey2"
    />
    <MainCard
      title="Monterrey, Mexico"
      text="July 2015"
      imgheight="200px"
      width="300"
      imgsrc="/img/monterrey-7-2015.jpg"
      morelink="/maps/monterrey3"
    />
    <MainCard
      title="Colombia"
      text="July 2015"
      imgheight="200px"
      width="300"
      imgsrc="/photos/travels/Colombia/Colombia-001.jpg"
      morelink="/maps/colombia"
    />
    <MainCard
      title="Mexico City, Mexico"
      text="October 2016"
      imgheight="200px"
      width="300"
      imgsrc="/photos/travels/MexicoCity/MexicoCity-001.jpg"
      morelink="/maps/mexicocity"
    />
    <MainCard
      title="Colima, Mexico"
      text="March 2018"
      imgheight="200px"
      width="300"
      imgsrc="/photos/travels/Colima/Colima-001.jpg"
      morelink="/maps/colima"
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
