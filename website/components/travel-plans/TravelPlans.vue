<template>
  <section class="section is-medium">
    <div class="columns">
      <div class="column">
        <h1 class="title is-4" id="travel">Travel</h1>
        <p>We recommend you find a place to stay in Corowa on the Friday night after the wedding, with a few good options within close proximity to Corowa Whisky and Chocolate. Being a small town, there will be limited taxi services available, and don’t even bother with Uber!</p>
      </div>
      <div class="column">
        <h1 class="title is-4" id="accommodation">Accommodation</h1>
        <p>We recommend you find a place to stay in Corowa on the Friday night after the wedding, with a few good options within close proximity to Corowa Whisky and Chocolate. Being a small town, there will be limited taxi services available, and don’t even bother with Uber!</p>
      </div>
    </div>
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
