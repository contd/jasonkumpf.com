<template>
  <v-layout
    class="d-flex flex-row flex-wrap justify-space-around"
  >
    <MainCard
      title="Monterrey, Mexico"
      text="August 2014"
      imgheight="200px"
      imgsrc="/img/monterrey-8-2014.jpg"
      @click="showMap1Click"
      v-if="!showMap2 && !showMap2"
    />
    <MainCard
      title="Monterrey, Mexico"
      text="July 2015"
      imgheight="200px"
      imgsrc="/img/monterrey-7-2015.jpg"
      @click="showMap2Click"
      v-if="!showMap2 && !showMap2"
    />
    <div v-if="showMap1">
      <iframe
        src="https://www.google.com/maps/d/embed?mid=13PagFei3Yz0lHY1qpvbN6TCFGnA"
        style="position:fixed; top:0px; left:0px; bottom:0px; right:0px; width:100%; height:100%; border:none; margin:0; padding:0; overflow:hidden; z-index:999999;"
      />
    </div>
    <div v-if="showMap2">
      <iframe
        src="https://www.google.com/maps/d/embed?mid=1j0XbxJDNzlkKL5xJ5jahy_qmXJQ"
        style="position:fixed; top:0px; left:0px; bottom:0px; right:0px; width:100%; height:100%; border:none; margin:0; padding:0; overflow:hidden; z-index:999999;"
      />
    </div>
  </v-layout>
</template>

<script>
import MainCard from '@/components/MainCard'

export default {
  name: 'Travels',
  components: {
    MainCard
  },
  data () {
    return {
      showMap1: false,
      showMap2: false
    }
  },
  methoods: {
    showMap1Click () {
    },
    showMap2Click () {
    }
  }
}
</script>
