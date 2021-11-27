<template>
  <section>

    <div class="columns is-centered">
      <p>Let us know who on the invite is coming!</p>
    </div>

    <div class="columns is-centered">
      <div class="column is-one-quarter">
        <div class="block">
          <b-field class="attendance" v-for="guest in guests" v-bind:data="guest" v-bind:key="guest.id">
            <p><b>{{ guest.firstName | titlecase }} {{ guest.lastName | titlecase }}</b></p>
            <b-switch v-model="guest.rsvp" :rounded="false" passive-type="is-dark" type="is-success">
              {{ guest.rsvp ? "Is attending &#127881;" : "Is not attending &#128577;" }}
            </b-switch>
          </b-field>
        </div>
        <div class="block">
          <b-field>
            <b-button type="is-primary"
                      label="Next"
                      @click="next"
                      expanded />
          </b-field>
          <b-field>
            <b-button label="Back"
                      @click="back"
                      expanded />
          </b-field>
        </div>
      </div>
    </div>

  </section>
</template>

<script lang="ts">
  import  Vue from "vue"
  import Buefy from 'buefy'

  Vue.use(Buefy)

  export default Vue.extend({
    name: 'Attendance',

    props: ['guests'],

    components: {},

    data() {
      return {
      }
    },

    methods: {
      next() {
        this.$emit('nextClicked');
      },
      back() {
        this.$emit('backClicked')
      }
    },

    filters: {
      titlecase(value: string){
        return value.replace(/(?:^|\s|-)\S/g, x => x.toUpperCase());
      }
    },

  })
</script>

<style lang="scss" scoped>
  @import "~assets/all";

  .attendance {
    display: block;

    p {
      padding-bottom: 5px;
    }
  }

</style>
