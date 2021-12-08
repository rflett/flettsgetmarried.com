<template>
  <section class="hero">
    <div class="hero-head">
      <b-navbar :spaced="true" :fixed-top="true" :shadow="true">
        <template #start>
          <b-navbar-item  @click="goto('ceremony')">the wedding</b-navbar-item>
          <b-navbar-item href="#" @click="goto('ourStory')">our story</b-navbar-item>
          <b-navbar-item href="#" @click="goto('bridalParty')">bridal party</b-navbar-item>
          <b-navbar-item href="#" @click="goto('travelPlans')">travel + accommodation</b-navbar-item>
          <b-navbar-item href="#" @click="goto('gifts')">gifts</b-navbar-item>
          <b-navbar-item href="#" @click="goto('recovery')">the day after</b-navbar-item>
          <b-navbar-item href="#" @click="goto('faq')">faq</b-navbar-item>
          <b-navbar-item href="https://photos.app.goo.gl/pczNbTH7Yreju3o16" target="_blank">photos from the day</b-navbar-item>
        </template>

        <template #end>
          <b-navbar-item tag="div">
            <div class="buttons">
              <a class="button is-primary" @click="goto('rsvp')">RSVP</a>
            </div>
          </b-navbar-item>
        </template>
      </b-navbar>
    </div>
    <div class="hero-body">
      <SaveTheDate v-if="activePage === 'home'" />

      <TheDay v-if="activePage === 'home'" />
      <OurStory v-if="activePage === 'home'" />
      <BridalParty v-if="activePage === 'home'" />

      <TravelPlans v-if="activePage === 'home'" />

      <Gifts v-if="activePage === 'home'" />

      <Recovery v-if="activePage === 'home'" />
      <FAQ v-if="activePage === 'home'" />

      <section class="section is-small" v-if="activePage === 'home'">
        <h1 class="title" id="rsvp">RSVP</h1>
        <p>
          Please click on the link below to register your attendance, including those who are sadly unable to attend. You are able to respond for all guests on the same invitation.
        </p>
      </section>

      <RSVP v-if="activePage === 'rsvp'" />
    </div>
  </section>
</template>

<script>

  import {scrollToID} from "@/models/utilities";

  export default {
    name: 'Index',

    components: {
    },

    data() {
      return {
        activePage: 'home'
      };
    },

    methods: {
      goto(page) {
          if (page === "rsvp") {
            this.activePage = "rsvp";
          } else {
            this.activePage = "home";
            scrollToID(page, 84);
          }
      }
    }
  }
</script>

<style lang="scss" scoped>
  .hero-body {
    padding: 0 20px 10px 20px;
  }

  .navbar {
    background-color: white;
  }
</style>
