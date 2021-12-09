<template>
  <section class="section is-small">
      <h1 class="title" id="travelPlans">Travel & Accommodation</h1>
      <p>We recommend you find a place to stay in Corowa on the Friday night after the wedding, with a few good options within close proximity to Corowa Whisky and Chocolate.</p>
      <p>Being a small town there will be limited taxi services available.</p>
      <hr />
      <ul>
        <li>Federation Motel Resort</li>
        <li>Heritage Motor Inn</li>
        <li>Statesman Motor Inn</li>
      </ul>
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
