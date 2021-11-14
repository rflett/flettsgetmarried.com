<template>
  <section>

    <div class="columns is-centered">
      <p>Let us know who on the invite is coming!</p>
    </div>

    <div class="columns is-centered">
      <div class="column is-one-quarter">
        <div class="block">
          <b-field v-for="guest in guests" v-bind:data="guest" v-bind:key="guest.id">
            <b-checkbox v-model="guest.rsvp">
              {{ guest.firstName | titlecase }} {{ guest.lastName | titlecase }}
            </b-checkbox>
          </b-field>
        </div>
        <div class="block">
          <b-field>
            <b-button type="is-light"
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
