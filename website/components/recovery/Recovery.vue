<template>
  <section class="section is-small">
    <h1 class="title" id="recovery">The Day After</h1>
    <p>
      The families and bridal party will be staying at an air bnb in Yarrawonga for the weekend, and we invite you all to call in from 10am onwards on Saturday the 26th.
      This gives the bride and groom another chance to catch up once the whirlwind of a day has passed, and enjoy brunch/lunch together before they head back to QLD.
      If you would like to continue the celebrations, please BYO drinks.
      Please advise Kara or Ryan if you wish to join us on the Saturday, and a rough estimate of time or if you would like to stay for lunch.
    </p>
    <br />
    <br />
    <p>Kara 0400 964 349 Ryan 0429 235 371</p>
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
