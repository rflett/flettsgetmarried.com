<template>
  <section class="section is-medium">
    <h1 class="title is-4" id="gifts">Gifts</h1>
    <p>We are thrilled to have you all present for our special day, however if you would like to honour us with a gift, we will have a wishing well on the day. We have just purchased our new home together and plan to have a honeymoon somewhere warm and sunny, so any contributions would help us fund these exciting adventures.</p>
  </section>
</template>

<script lang="ts">
import  Vue from "vue"
import Buefy from 'buefy'

Vue.use(Buefy)

export default Vue.extend({
  name: 'TravelPlans',

  components: {
  },

  data(){
    return {
      carousels: [
        { text: 'Coming Soon!', color: 'primary' },
        { text: 'Coming Soon!', color: 'info' },
      ]
    }
  },

  methods: {
  }
})
</script>

<style lang="scss" scoped>
@import "~assets/all";
</style>
