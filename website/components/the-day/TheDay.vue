<template>
  <section class="section is-medium">
    <div class="columns">
      <div class="column">
        <h1 class="title is-4" id="ceremony">Ceremony</h1>
        <h1 class="title is-6">3:00PM</h1>
        <p>On the bank of the Murray by the Corowa and Wahgunyah Rowing Clubs</p>
        <p>Enter via Bridge Road from Corowa, just before crossing the John Foord Bridge</p>
        <p>Stay after the ceremony while family photos are being taken, for lawn games and light refreshments</p>
      </div>
      <div class="column">
        <h1 class="title is-4" id="reception">Reception</h1>
        <h1 class="title is-6">6:00PM</h1>
        <p>COROWA WHISKY AND CHOCOLATE</p>
        <p>20-24 Steel Street, Corowa NSW, 2646</p>
        <p>Courtyard and Café will be available from 5pm for drinks and canapes</p>
        <p>Doors will open at 6pm, please be seated for bridal party to arrive at 6:15pm</p>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
  import  Vue from "vue"
  import Buefy from 'buefy'

  Vue.use(Buefy)

  export default Vue.extend({
    name: 'TheDay',

    components: {
    },

    data(){
      return {}
    },

    methods: {
    }
  })
</script>

<style lang="scss" scoped>
  @import "~assets/all";
</style>
