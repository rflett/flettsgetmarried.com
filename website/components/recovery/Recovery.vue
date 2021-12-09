<template>
  <section class="section is-small">
    <h1 class="title" id="recovery">The Day After</h1>
    <p>
      The families and bridal party will be staying at an Airbnb in Yarrawonga for the weekend,
      and we invite you all to call in from 10am on Saturday the 26th.
      If you would like to continue the celebrations please advise Kara or Ryan and BYO drinks.
    </p>
    <br />
    <ul>
      <li>Kara | 0400 964 349</li>
      <li>Ryan | 0429 235 371</li>
      <li><a href="https://goo.gl/maps/Muct4XfonHetpd1t5" target="_blank">4 Waterview Pl, Yarrawonga</a></li>
    </ul>
  </section>
</template>

<script lang="ts">
  import  Vue from "vue"
  import Buefy from 'buefy'

  Vue.use(Buefy)

  export default Vue.extend({
    name: 'Recovery',

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
