<template>
  <section>

    <div class="columns is-centered">
      <p>If you're responding for you and a guest (or your family), you'll be able to RSVP for your entire group.</p>
    </div>

    <div class="columns is-centered">
      <div class="column is-one-quarter">
        <div class="block">
          <b-field label="Full Name">
            <b-input v-model="searchVal"
                     v-on:keyup.native.enter="search"
                     type="text"
                     placeholder="Ryan Flett"
                     required
                     validation-message="First and last name separated by space"
                     pattern="[a-zA-Z]{1,} [a-zA-Z]{1,}"
            ></b-input>
          </b-field>
        </div>
        <b-field>
          <b-button type="is-light"
                    icon-left="magnify"
                    label="Search"
                    @click="search"
                    :disabled="!nameIsValid"
                    expanded />
        </b-field>
      </div>
    </div>

    <div class="block"></div>

    <div class="columns is-centered" v-if="searchCommenced">
        <p v-if="!foundInvites">Oops! We’re having trouble finding your invite. Please try another spelling of your name or contact the couple</p>

      <div v-if="foundInvites">
          <div class="block" v-for="match in searchResults.matches">
            <b-radio v-model="selectedInvite" name="name" :native-value="match.inviteId">
              <p class="invites" v-for="guest in match.guests">
                {{ guest.firstName | titlecase }} {{ guest.lastName | titlecase }}
              </p>
            </b-radio>
          </div>

          <b-field>
            <b-button type="is-primary"
                    label="Begin RSVP"
                    @click="selectInvite"
                    :disabled="selectedInvite === ''"
                      expanded />
          </b-field>
        </div>
    </div>

    </section>
</template>

<script lang="ts">
  import  Vue from "vue"
  import Buefy from 'buefy'
  import { SearchResult } from "@/models/invite";
  import {ISearchMatch, ISearchResult} from "~/models/interfaces";

  Vue.use(Buefy)

  export default Vue.extend({
    name: 'Search',

    components: {},

    data() {
      return {
        searchVal: "",
        selectedInvite: "",
        foundInvites: false,
        searchCommenced: false,
        searchResults: new SearchResult(),
      }
    },

    methods: {
      search() {
        this.selectedInvite = "";
        const firstName = this.searchVal.split(" ")[0];
        const lastName = this.searchVal.split(" ")[1];
        fetch(
          `https://api.flettsgetmarried.com/search?firstName=${firstName}&lastName=${lastName}`
        )
        .then(res => res.json())
        .then(data => this.handleSearchResponse(data));
      },
      handleSearchResponse(data: ISearchResult) {
        this.searchResults = new SearchResult(data);
        this.searchCommenced = true;
        this.foundInvites = this.searchResults.matches.length >= 1;
      },
      selectInvite() {
        const invite = this.searchResults.matches.find((m: ISearchMatch) => m.inviteId == this.selectedInvite);
        if (invite) {
          this.$emit('inviteSelected', invite);
        }
      }
    },

    filters: {
      titlecase(value: string){
        return value.replace(/(?:^|\s|-)\S/g, x => x.toUpperCase());
      }
    },

    computed: {
      nameIsValid(): boolean {
        const re = new RegExp("[a-zA-Z]+ [a-zA-Z]+");
        return re.test(this.searchVal);
      }
    }

  })
</script>

<style lang="scss" scoped>
  @import "~assets/all";

  .invites {
    padding: 0 0 5px 5px;
  }
</style>
