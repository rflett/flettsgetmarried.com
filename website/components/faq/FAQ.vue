<template>
  <section class="section is-small">
    <h1 class="title" id="faq">FAQ</h1>

    <div v-for="item in faq" v-bind:data="item" v-bind:key="item.question">
      <h1 class="title is-6">{{ item.question }}</h1>
      <p>{{ item.answer }}</p>
      <br />
    </div>
  </section>
</template>

<script lang="ts">
  import  Vue from "vue"
  import Buefy from 'buefy'

  Vue.use(Buefy)

  export default Vue.extend({
    name: 'Accom',

    components: {
    },

    data(){
      return {
        faq: [
          { question: 'What is the dress code?', answer: 'The dress code is formal. We encourage everyone to wear something with a little pop of colour! Ladies, please keep in mind the ceremony will take place on grass, so if you’re opting for heels, a block heel, wedge or flats may be most comfortable – we also want you to be able to hit up the dance floor later in the night! For the gents, a suit and tie with an optional jacket.'},
          { question: 'Are kids invited to the wedding?', answer: 'Unfortunately, the only children that will be attending the wedding are the nieces and nephews of the couple.' },
          { question: 'Will there be a bus?', answer: 'We will not be providing transport for our guests to return to their accommodation from the Corowa Whisky Chocolate Factory, hence why we encourage you find somewhere local to stay the night. The distance between the ceremony and reception venues is approximately 2km so please ensure you have a designated driver to get there, cars can be left at the rowing club or reception carparks.'},
          { question: 'Do I need to be vaccinated?', answer: 'All guests are required to be fully vaccinated against COVID as per government restrictions.'},
        ]
      }
    },

    methods: {
    }
  })
</script>

<style lang="scss" scoped>
  @import "~assets/all";

  section {
    background-color: #ebdacc;
  }

</style>
