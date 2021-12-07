<template>
  <section class="hero">
    <div class="hero-head">
      <b-navbar :spaced="true">
        <template #start>
          <b-navbar-item  @click="goto('home')">
            the wedding
          </b-navbar-item>
          <b-navbar-item href="#" @click="goto('ourStory')">our story</b-navbar-item>
          <b-navbar-item href="#" @click="goto('bridalParty')">bridal party</b-navbar-item>
          <b-navbar-item href="#" @click="goto('travelPlans')">travel plans</b-navbar-item>
          <b-navbar-item href="#" @click="goto('accom')">accommodation options</b-navbar-item>
          <b-navbar-item href="#" @click="goto('gift')">gifts</b-navbar-item>
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

      <section class="section is-medium">
        <h1 class="title is-4" id="rsvp">RSVP</h1>
        <p>
          Please click on the link below to register your attendance, including those who are sadly unable to attend. You are able to respond for all guests on the same invitation.
        </p>
      </section>

      <RSVP v-if="activePage === 'rsvp'" />
    </div>
  </section>
</template>

<script>

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
        this.activePage = page;
    }
  }
}
</script>

<style lang="scss" scoped>
  .hero-body {
    padding: 0 20px 10px 20px;
  }
</style>
