<template>
  <section>

    <div class="columns is-centered">
      <div class="column is-one-third">
        <div class="block">
          <p v-if="attending">Thanks for your RSVP - we can't wait to celebrate with you on the 25th of March!</p>
          <p v-if="!attending">It's a shame we won't be able to celebrate with you!</p>
        </div>
        <div class="block">
          <p>You can come back any time to update this invitation.</p>
        </div>
      </div>
    </div>

    <div class="columns is-centered">
      <div class="column is-one-quarter">
        <div class="block">
          <b-field>
            <b-button type="is-primary"
                      label="Close"
                      @click="finishClicked"
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
    name: 'Finished',

    props: {
      guests: Array,
      attending: Boolean,
    },

    components: {},

    data() {
      return {}
    },

    methods: {
      finishClicked() {
        this.$emit('finishClicked');
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
