<template>
  <section>
    <div class="container">
      <div class="columns is-centered">
        <div class="column is-one-quarter">
          <div class="block">
            <b-field label="Admin Password">
              <b-input v-model="password"
                       v-on:keyup.native.enter="getGuests"
                       type="password"
                       required
                       placeholder="********"
                       password-reveal
              ></b-input>
            </b-field>
          </div>
          <b-field>
            <b-button type="is-primary"
                      label="Submit"
                      @click="getGuests"
                      expanded />
          </b-field>
        </div>
      </div>

      <div class="block"></div>

      <b-table
        :data="guests.guests"
        :bordered="false"
        :striped="false"
        :narrowed="true"
        :mobile-cards="true"
        :row-class="(row, index) => row.inviteId % 2 === 0 && 'is-dark'"
      >
        <template slot-scope="props">
          <b-table-column field="inviteId" label="Invite" width="15" centered>
            {{ +props.row.inviteId }}
          </b-table-column>
          <b-table-column field="name" label="Name" width="300">
            {{ props.row.firstName | titlecase }} {{ props.row.lastName | titlecase }}
          </b-table-column>
          <b-table-column field="rsvp" label="RSVP" width="15">
              <span class="tag" :class="props.row.rsvp ? 'is-success' : 'is-danger'">
                {{ props.row.rsvp ? "Yes" : "No" }}
              </span>
          </b-table-column>
          <b-table-column field="vaccinated" label="Vaccinated" width="15">
              <span class="tag" :class="props.row.vaccinated ? 'is-success' : 'is-danger'">
                {{ props.row.vaccinated ? "Yes" : "No" }}
              </span>
          </b-table-column>
          <b-table-column field="updatedAt" label="Updated At" width="20">
              <span class="tag" v-if="props.row.updatedAt">
                {{ new Date(props.row.updatedAt).toLocaleDateString() }}
              </span>
            <span class="tag" v-if="!props.row.updatedAt">Not Updated</span>
          </b-table-column>
          <b-table-column field="diet" label="Diet">
            {{ props.row.diet }}
          </b-table-column>
          <b-table-column field="music" label="Music">
            {{ props.row.music }}
          </b-table-column>
        </template>
      </b-table>

    </div>
  </section>
</template>

<script lang="ts">
  import  Vue from "vue"
  import Buefy from 'buefy'
  import {GuestsResult} from "~/models/invite";
  import {IGuestsResult} from "~/models/interfaces";

  Vue.use(Buefy)

  export default Vue.extend({
    name: 'Admin',

    components: {},

    data() {
      return {
        guests: new GuestsResult(),
        password: "",
      }
    },

    methods: {
      getGuests() {
        fetch(
          `https://api.flettsgetmarried.com/guests?password=${this.password}`
        )
          .then(res => res.json())
          .then(data => this.handleSearchResponse(data));
      },
      handleSearchResponse(data: IGuestsResult) {
        this.guests = new GuestsResult(data);
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

  .container {
    padding: 20px 10px 0px 10px;
  }

  .is-dark  {
    background: lightsalmon;
  }
</style>
