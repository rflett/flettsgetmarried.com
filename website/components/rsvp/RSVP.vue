<template>
    <section>
        <b-field label="Attending">
            <b-taginput
                v-model="attendees"
                icon="account"
                placeholder="Add a name"
                type="is-success"
                size="is-medium"
                aria-close-label="Remove this person">
            </b-taginput>
        </b-field>
        
        <b-field label="Not Attending">
            <b-taginput
                v-model="absentees"
                icon="account"
                placeholder="Add a name"
                type="is-danger"
                size="is-medium"
                aria-close-label="Remove this person">
            </b-taginput>
        </b-field>

        <b-field label="Dietry Requirements">
            <b-input maxlength="200" type="textarea"></b-input>
        </b-field>

        <div class="notification is-primary">
            <div class="buttons">
                <b-button type="is-primary" inverted outlined>Cancel</b-button>
                <b-button type="is-primary" inverted @click="submit">Submit RSVP</b-button>
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
        absentees: []
      }
    },

    methods: {
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
