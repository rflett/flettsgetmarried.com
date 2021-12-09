<template>
  <section class="section is-small">
    <h1 class="title" id="ourStory">Our Story</h1>

    <p>
      Kara and Ryan have ‘officially’ been together since the morning of Wednesday, July 28th, 2010. However those closest to them at the time will attest that their love was blossoming much earlier than that.
      <br /><br />
      Ryan and Kara both attended Kyabram Secondary College, sharing the odd elective together since Year 7, but not properly meeting until Year 9 when they were in the same form and Ryan would adore Kara from across the room.
      <br /><br />
      It wasn’t until the Flett family returned from their brief hiatus in Warren NSW, that Kara paid closer attention, and their friendship developed further once lunchtime friend circles merged later that year.
      <br /><br />
      The next 18 months saw a lot of back and forth messaging, racing home from school to chat on MSN, and heavy flirting at weekend group gatherings. The school library at study breaks in Year 11, the green/blue area at lunchtimes, and near the music studio after school before Ryan’s guitar lesson, always presented an opportunity for the two love birds to exchange looks and a cheeky handhold. It wasn’t until one Wednesday morning on a misty day that the two met up at Edis Park and finally sealed the deal, excited to head back to tell their friends that morning at recess.
    </p>

    <div class="columns">
      <div class="column is-6 is-offset-3">
        <figure class="image">
          <img src="~assets/img/young.jpg" />
        </figure>
      </div>
    </div>

    <p>
      The next 11 years that followed have seen growth together as individuals and as a couple, achieving many important milestones by each others side: graduation from year 12; moving to Bendigo for university at LaTrobe; securing a job each at Jojoes pizza and pasta restaurant – Ryan the pizza maker, Kara the waitress; graduating from uni, obtaining their first professional jobs of their careers; moving interstate to Brisbane; travelling to New Zealand where they got engaged; and most recently, buying their first home together in Brisbane.
      <br /><br />
      New Zealand will always hold a special place in Kara and Ryan’s heart, as Ryan popped the question by the turquoise waters of Lake Pukaki on October 13th, 2019. It was a freezing evening, but Ryan was very insistent that they go for a walk along the lake at sunset. It made for a perfect start to their two-week holiday together.
      <br /><br />
      Planning a wedding from interstate would be challenging at the best of times, so the couple are very excited to finally share their special day with those that mean the most to them, by the Murray River which has always felt like home.
    </p>

    <div class="columns">
      <div class="column is-6 is-offset-3">
        <figure class="image">
          <img src="~assets/img/engaged.jpg" />
        </figure>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
  import  Vue from "vue"
  import Buefy from 'buefy'

  Vue.use(Buefy)

  export default Vue.extend({
    name: 'OurStory',

    components: {
    },

    data(){
      return {
        carousels: [
          { text: 'Our story!', color: 'primary' },
          { text: 'Our story!', color: 'info' },
        ]
      }
    },

    methods: {
    }
  })
</script>

<style lang="scss" scoped>
  @import "~assets/all";

  .image {
    margin-right: auto;
    margin-left: auto;
    padding-top: 20px;
  }

  @include col-lg() {
    .image {
      max-width: 768px;
    }
  }
</style>
