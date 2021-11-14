<template>
  <section>

    <div class="columns is-centered">
      <div class="column is-one-third">
        <div class="block">
          <p>
            Which of the following people <b>have</b> received 2 doses of a COVID-19 vaccine,
            or <b>will</b> have received 2 doses by the 11th March 2022.
          </p>
        </div>
        <div class="block">
          <p>You may be asked to provide proof of vaccination on the day.</p>
        </div>
      </div>
    </div>

    <div class="columns is-centered">
      <div class="column is-one-quarter">
        <div class="block">
          <b-field v-for="guest in guests" v-bind:data="guest" v-bind:key="guest.id">
            <b-checkbox v-model="guest.vaccinated" v-if="guest.rsvp">
              {{ guest.firstName | titlecase }} {{ guest.lastName | titlecase }}
            </b-checkbox>
          </b-field>
        </div>
        <div class="block">
          <b-field>
            <b-button type="is-light"
                      icon-left="magnify"
                      label="Next"
                      @click="next"
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
    name: 'Covid',

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

</style>
