<template>
  <section>

    <div class="columns is-centered">
      <p>If you have a recommendation for a song to play then let us know!</p>
    </div>

    <div class="columns is-centered">
      <div class="column is-one-third">
        <div class="block" v-for="guest in guests" v-bind:data="guest" v-bind:key="guest.id">
          <p v-if="guest.rsvp"><b>{{ guest.firstName | titlecase }} {{ guest.lastName | titlecase }}</b></p>
          <b-field v-if="guest.rsvp">
              <b-input v-model="guest.music"></b-input>
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
    name: 'Music',

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

  .block {
    p {
      padding-bottom: 5px;
    }
  }
</style>
