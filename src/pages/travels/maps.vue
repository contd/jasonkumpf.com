<template>
  <v-layout
    class="d-flex flex-row flex-wrap justify-space-around"
  >
    <v-row>
      <v-col>
        <MainCard
          title="Monterrey, Mexico"
          text="August 2014"
          imgheight="300px"
          width="400"
          imgsrc="/img/monterrey-8-2014.jpg"
          morelink="/travels/map1"
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <MainCard
          title="Monterrey, Mexico"
          text="July 2015"
          imgheight="300px"
          width="400"
          imgsrc="/img/monterrey-7-2015.jpg"
          morelink="/travels/map2"
        />
      </v-col>
    </v-row>
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
