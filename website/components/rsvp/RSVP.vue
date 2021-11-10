<template>
    <section>
      <div class="container">
        <b-steps
          v-model="activeStep"
          :animated="true"
          :rounded="true"
          :has-navigation="true"
          icon-prev="chevron-left"
          icon-next="chevron-right"
          label-position="bottom"
          mobile-mode="compact">

          <b-step-item step="1" label="Search" icon="magnify">
            <h1 class="title has-text-centered">Find your invite</h1>
            <Search @inviteSelected="inviteSelected($event)"/>
          </b-step-item>

          <b-step-item step="2" label="RSVP" icon="ticket">
            <h1 class="title has-text-centered">Attendance</h1>
            Lorem ipsum dolor sit amet.
            <p v-for="guest in guests">{{ guest.firstName }}</p>
          </b-step-item>

          <b-step-item step="3" label="Diet" icon="food">
            <h1 class="title has-text-centered">Dietary Requirements</h1>
            Lorem ipsum dolor sit amet.
          </b-step-item>

          <b-step-item step="4" label="Music" icon="music">
            <h1 class="title has-text-centered">Dietary Requirements</h1>
            Lorem ipsum dolor sit amet.
          </b-step-item>

          <b-step-item step="5" label="COVID" icon="needle">
            <h1 class="title has-text-centered">Attendance</h1>
            Lorem ipsum dolor sit amet.
          </b-step-item>
        </b-steps>
      </div>
    </section>
</template>

<script lang="ts">
  import  Vue from "vue"
  import Buefy from 'buefy'
  import { Guest } from "~/models/invite";

  Vue.use(Buefy)

  export default Vue.extend({
    name: 'RSVP',

    components: {
    },

    data() {
      return {
        activeStep: 0,
        guests: [],
      }
    },

    methods: {
      inviteSelected(guests: Guest[]) {
        this.guests = guests;
        this.activeStep = 1;
      }
    }
  })
</script>

<style lang="scss" scoped>
  @import "~assets/all";

  .container {
    padding-top: 20px;
  }

</style>
