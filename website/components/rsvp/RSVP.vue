<template>
    <section>
        <div class="container">
          <b-field label="Contact" type="null">
              <b-input placeholder="04"
                  v-model="contact"
                  v-on:blue="validate($event, 'contact')"
                  type="text"
                  icon="phone"
                  maxlength="10"
                  validation-message="Must be a mobile number starting with 04"
                  pattern="04[0-9]{8}"
                  has-counter
                  required>
              </b-input>
          </b-field>

          <b-field label="Attending">
              <b-taginput
                  v-model="attendees"
                  icon="account"
                  placeholder="Add a name"
                  type="is-success"
                  aria-close-label="Remove this person">
              </b-taginput>
          </b-field>
          
          <b-field label="Not Attending">
              <b-taginput
                  v-model="absentees"
                  icon="account"
                  placeholder="Add a name"
                  type="is-danger"
                  aria-close-label="Remove this person">
              </b-taginput>
          </b-field>

          <b-field label="Dietry Requirements">
              <b-input maxlength="200" type="textarea" v-model="diet"></b-input>
          </b-field>

          <b-field label="Song Requests">
              <b-input maxlength="300" type="textarea" v-model="song"></b-input>
          </b-field>

          <div class="notification is-primary">
              <div class="buttons">
                  <b-button type="is-primary" inverted outlined>Cancel</b-button>
                  <b-button type="is-primary" inverted @click="submit">Submit RSVP</b-button>
              </div>
          </div>

          <div class="notification is-warning">
            <p v-for="a in attendees" v-bind:key="a">{{ a }}</p>
            <p v-for="a in absentees" v-bind:key="a">{{ a }}</p>
              {{ contact }}
              {{ diet }}
              {{ song }}
          </div>
        </div>
    </section>
</template>


<script lang="ts">
  import  Vue from "vue"
  import Buefy from 'buefy'

  Vue.use(Buefy)

  export default Vue.extend({
    name: 'RSVP',

    components: {
    },

    data() {
      return {
        attendees: [],
        absentees: [],
        contact: null,
        diet: null,
        song: null
      }
    },

    methods: {
      validate(event, field: string) {
        switch (field) {
          case 'contact': {

            break;
          }
        }
      },
      submit() {
        this.$buefy.toast.open({
            duration: 3000,
            message: `Thank you for submitting your RSVP`,
            position: 'is-bottom'
        })
    }
    }

  })
</script>

<style lang="scss" scoped>
  @import "~assets/all";

  .hero {
    background-position: center;
    background-size: cover;
    background-repeat: no-repeat;
    position: relative;
  }

  .hero-body {
    padding: 0px;
  }

</style>
