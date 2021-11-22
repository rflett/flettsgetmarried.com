<template>
    <section>
      <div class="container">
        <b-steps
          v-model="$data.activeStep"
          :animated="true"
          :rounded="true"
          :has-navigation="false"
          icon-prev="chevron-left"
          icon-next="chevron-right"
          label-position="bottom"
          mobile-mode="compact">

          <b-step-item step="1" label="Find Invite" icon="magnify">
            <h1 class="title has-text-centered">Find your invite</h1>
            <Search @inviteSelected="inviteSelected($event)"/>
          </b-step-item>

          <b-step-item step="2" label="RSVP" icon="ticket">
            <h1 class="title has-text-centered">Attendance</h1>
            <Attendance :guests="$data.selectedInvite.guests" @nextClicked="guestsAttending()"/>
          </b-step-item>

          <b-step-item step="3" label="Diet" icon="food" :clickable="$data.attending">
            <h1 class="title has-text-centered">Dietary Requirements</h1>
            <Diet :guests="$data.selectedInvite.guests" @nextClicked="guestDietFinished()"/>
          </b-step-item>

          <b-step-item step="4" label="Music" icon="music" :clickable="$data.attending">
            <h1 class="title has-text-centered">Pick a song to play</h1>
            <Music :guests="$data.selectedInvite.guests" @nextClicked="musicFinished()"/>
          </b-step-item>

          <b-step-item step="5" label="COVID" icon="needle" :clickable="$data.attending">
            <h1 class="title has-text-centered">COVID-19 Vaccination Status</h1>
            <COVID :guests="$data.selectedInvite.guests" @nextClicked="covidFinished()"/>
          </b-step-item>

          <b-step-item step="6" label="Finish" icon="glass-flute">
            <h1 class="title has-text-centered">Thank you!</h1>
            <Finished :guests="$data.selectedInvite.guests" :attending="$data.attending" @finishClicked="finishClicked()"/>
          </b-step-item>
        </b-steps>
      </div>
    </section>
</template>

<script lang="ts">
  import  Vue from "vue"
  import Buefy from 'buefy'
  import {Guest, SearchMatch} from "~/models/invite";

  Vue.use(Buefy)

  export default Vue.extend({
    name: 'RSVP',

    components: {
    },

    data() {
      const data : {
        activeStep: number,
        attending: boolean,
        selectedInvite: SearchMatch
      } = {
        activeStep: 0,
        attending: false,
        selectedInvite: new SearchMatch()
      };
      return data;
    },

    methods: {
      inviteSelected(invite: SearchMatch) {
        this.activeStep = 1;
        this.selectedInvite = invite;
      },
      guestsAttending() {
        this.attending = this.selectedInvite.guests.filter((g: Guest) => g.rsvp).length > 0;
        if (this.attending) {
          this.activeStep = 2;
        } else {
          this.activeStep = 5;
        }
      },
      guestDietFinished() {
        this.activeStep = 3;
      },
      musicFinished() {
        this.activeStep = 4;
      },
      covidFinished() {
        this.activeStep = 5;
      },
      finishClicked() {
        window.location.href = "https://flettsgetmarried.com/";
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
