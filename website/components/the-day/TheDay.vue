<template>
  <section class="section is-small">
    <div class="columns">
      <div class="column">
        <h1 class="title" id="ceremony">Ceremony</h1>
        <h1 class="title is-6">3:00PM</h1>
        <p>On the bank of the Murray by the Corowa and Wahgunyah Rowing Clubs.</p>
        <p>Enter via Bridge Road, on the NSW side of the river.</p>
        <p>Stay after the ceremony while photos are being taken for lawn games and light refreshments.</p>
        <br />
        <br />
        <figure class="image is-1by1">
          <img src="~assets/img/mudmap.PNG">
        </figure>
      </div>
      <div class="column">
        <h1 class="title" id="reception">Reception</h1>
        <h1 class="title is-6">6:00PM</h1>
        <p>Corowa Whisky and Chocolate</p>
        <a href="https://goo.gl/maps/3isV4jjQxFgdQ6XZA" target="_blank">20-24 Steel Street, Corowa</a>
        <p>The Courtyard and Café will be available from 5pm for drinks and canapes.</p>
        <p>Doors to reception will open at 6pm, please be seated for the bridal party to arrive at 6:15pm.</p>
        <br />
        <figure class="image is-1by1">
          <img src="~assets/img/whisky.PNG">
        </figure>
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

  section {
    background-color: #ebdacc;
  }

  p {
    padding-top: 10px;
    padding-bottom: 10px;
  }
</style>
