<template>
  <nuxt-link :to="`/photos?path=${folder.name}`">
    {{ folder.name }}
  </nuxt-link>
</template>

<script>
export default {
  props: {
    folder: {
      // eslint-disable-next-line vue/require-default-prop
      type: Object
    }
  }
}
</script>
