<template>
  <section class="hero">
    <div class="hero-head">
      <b-navbar :spaced="true" :fixed-top="true" :shadow="true">
        <template #start>
          <b-navbar-item @click="goto('ceremony')">the wedding</b-navbar-item>
          <b-navbar-item @click="goto('ourStory')">our story</b-navbar-item>
          <b-navbar-item @click="goto('bridalParty')">bridal party</b-navbar-item>
          <b-navbar-item @click="goto('travelPlans')">travel + accommodation</b-navbar-item>
          <b-navbar-item @click="goto('gifts')">gifts</b-navbar-item>
          <b-navbar-item @click="goto('recovery')">the day after</b-navbar-item>
          <b-navbar-item @click="goto('faq')">faq</b-navbar-item>
          <b-navbar-item @click="goto('photos')">photos from the day</b-navbar-item>
        </template>

        <template #end>
          <b-navbar-item>
            <div class="buttons">
              <b-button type="is-primary" expanded @click="goto('rsvp')">RSVP</b-button>
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
        <h1 class="title" id="photos">Photo Gallery</h1>

        <div class="columns">
          <div class="column"><img src="~assets/img/boys.png"></div>
          <div class="column"><img src="~assets/img/boys.png"></div>
          <div class="column"><img src="~assets/img/boys.png"></div>
        </div>

        <p>Please click the link below to share any photos that you take on the day, including any posted to social media as this gives us higher quality copies of guests photos.</p>
        <p>We please ask that you refrain from posting any photos to social media on our day, until the bride or groom have shared their own posts after the weekend.</p>
        <p>If you are posting anything to social media, please use the hashtag #flettsgetmarried.</p>
        <p>We will also have a photobooth at reception which we highly encourage you take advantage of!</p>
        <div class="columns">
          <div class="column is-4 is-offset-4">
            <div class="block"></div>
            <div class="buttons">
              <b-button type="is-primary" @click="openPhotos()" expanded>Upload Photos</b-button>
            </div>
          </div>
        </div>
      </section>

      <section class="section is-small rsvp-section" v-if="activePage === 'home'">
        <h1 class="title" id="rsvp">RSVP</h1>
        <p>
          Please click on the button below to register your attendance, including those who are sadly unable to attend. You are able to respond for all guests on the same invitation.
        </p>
        <div class="columns">
          <div class="column is-4 is-offset-4">
            <div class="block"></div>
            <div class="buttons">
              <b-button type="is-primary" expanded @click="goto('rsvp')">RSVP</b-button>
            </div>
          </div>
        </div>
      </section>

      <RSVP v-if="activePage === 'rsvp'" />
    </div>
  </section>
</template>

<script>
  import Vue from "vue"
  import {scrollToID} from "@/models/utilities";
  import VueSmoothScroll from 'vue2-smooth-scroll'

  Vue.use(VueSmoothScroll)

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
      },
      openPhotos() {
        window.open("https://photos.app.goo.gl/pczNbTH7Yreju3o16", "_blank").focus();
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

  .rsvp-section {
    background-color: #ebdacc;
  }
</style>
