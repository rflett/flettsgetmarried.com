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

          <b-step-item step="1" label="Search" icon="magnify">
            <h1 class="title has-text-centered">Find your invite</h1>
            <Search @inviteSelected="inviteSelected($event)"/>
          </b-step-item>

          <b-step-item step="2" label="RSVP" icon="ticket">
            <h1 class="title has-text-centered">Attendance</h1>
            <Attendance :guests="$data.selectedInvite.guests" @nextClicked="guestsAttending()" @backClicked="backClicked()"/>
          </b-step-item>

          <b-step-item step="3" label="Diet" icon="food" :clickable="$data.allAttending">
            <h1 class="title has-text-centered">Dietary Requirements</h1>
            <Diet :guests="$data.selectedInvite.guests" @nextClicked="guestDietFinished()" @backClicked="backClicked()"/>
          </b-step-item>

          <b-step-item step="4" label="Music" icon="music" :clickable="$data.allAttending">
            <h1 class="title has-text-centered">Pick a song to play</h1>
            <Music :guests="$data.selectedInvite.guests" @nextClicked="musicFinished()" @backClicked="backClicked()"/>
          </b-step-item>

          <b-step-item step="5" label="COVID" icon="needle" :clickable="$data.allAttending">
            <h1 class="title has-text-centered">COVID-19 Vaccination Status</h1>
            <COVID :guests="$data.selectedInvite.guests" @nextClicked="covidFinished()" @backClicked="backClicked()"/>
          </b-step-item>

          <b-step-item step="6" label="Finish" icon="glass-flute">
            <h1 class="title has-text-centered">Thank you!</h1>
            <Finished :guests="$data.selectedInvite.guests" :attending="$data.allAttending" @finishClicked="finishClicked()"/>
          </b-step-item>
        </b-steps>
      </div>

      <div class="unvaccinatedWarning" v-if="$data.unvaccinatedAttendees">
        <b-notification
          type="is-danger"
          aria-close-label="Close notification"
          role="alert">
          Due to government restrictions we can't have any unvaccinated guests attend the wedding.
        </b-notification>
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
        allAttending: boolean,
        unvaccinatedAttendees: boolean,
        selectedInvite: SearchMatch
      } = {
        activeStep: 0,
        allAttending: false,
        unvaccinatedAttendees: false,
        selectedInvite: new SearchMatch()
      };
      return data;
    },

    methods: {
      backClicked() {
        this.activeStep--;
      },
      inviteSelected(invite: SearchMatch) {
        this.activeStep = 1;
        this.selectedInvite = invite;
      },
      guestsAttending() {
        this.allAttending = this.selectedInvite.guests.filter((g: Guest) => g.rsvp).length > 0;
        if (this.allAttending) {
          this.activeStep = 2;
        } else {
          this.sumbit();
        }
      },
      guestDietFinished() {
        this.activeStep = 3;
      },
      musicFinished() {
        this.activeStep = 4;
      },
      covidFinished() {
        this.sumbit();
        this.unvaccinatedAttendees = this.selectedInvite.guests.filter((g: Guest) => g.rsvp && !g.vaccinated).length > 0;
      },
      sumbit() {
        this.sendData();
        this.activeStep = 5;
      },
      finishClicked() {
        window.location.href = "https://flettsgetmarried.com/";
      },
      sendData() {
        fetch("https://api.flettsgetmarried.com/rsvp", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(this.selectedInvite.toDto())
        });
      }
    }
  })
</script>

<style lang="scss" scoped>
  @import "~assets/all";

  .container {
    padding: 20px 10px 0px 10px;
  }

  .unvaccinatedWarning {
    position: absolute;
    bottom: 5%;
    right:0;
    left:0;
    padding: 0 20px 0 20px;
  }

</style>
