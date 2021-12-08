<template>
  <section class="section is-small">
    <h1 class="title" id="bridalParty">The Bridal Party</h1>
    <div class="container is-small">

      <!--the main one-->
      <div class="columns is-centered is-multiline">
        <div class="column">
          <div class="card">
            <div class="card-header">
              <div class="container">
                <p class="title is-4">Ryan Flett</p>
                <p class="subtitle is-6">Groom</p>
              </div>
            </div>
            <div class="card-image">
              <figure class="image is-1by1">
                <img src="https://bulma.io/images/placeholders/256x256.png">
              </figure>
            </div>
            <div class="card-content">He's the bish</div>
          </div>
        </div>
      </div>

      <!--the rest-->
      <div class="columns is-centered is-multiline">
        <div class="column is-6" v-for="member in boys" v-bind:data="member" v-bind:key="member.name">
          <div class="card">
            <div class="card-header">
              <div class="container">
                <p class="title is-4">{{ member.name }}</p>
                <p class="subtitle is-6">{{ member.role }}</p>
              </div>
            </div>
            <div class="card-image">
              <figure class="image is-1by1">
                <img :src="member.image">
              </figure>
            </div>
            <div class="card-content">{{ member.bio }}</div>
          </div>
        </div>

      </div>
    </div>
    <hr />
    <div class="container is-small">
      <!--the main one-->
      <div class="columns is-centered is-multiline">
        <div class="column">
          <div class="card">
            <div class="card-header">
              <div class="container">
                <p class="title is-4">Kara Hazelman</p>
                <p class="subtitle is-6">Bride</p>
              </div>
            </div>
            <div class="card-image">
              <figure class="image is-1by1">
                <img src="https://bulma.io/images/placeholders/256x256.png">
              </figure>
            </div>
            <div class="card-content">She's the bish</div>
          </div>
        </div>
      </div>

      <!--the rest-->
      <div class="columns is-centered is-multiline">
        <div class="column is-6" v-for="member in girls" v-bind:data="member" v-bind:key="member.name">
          <div class="card">
            <div class="card-header">
              <div class="container">
                <p class="title is-4">{{ member.name }}</p>
                <p class="subtitle is-6">{{ member.role }}</p>
              </div>
            </div>
            <div class="card-image">
              <figure class="image is-1by1">
                <img :src="member.image">
              </figure>
            </div>
            <div class="card-content">{{ member.bio }}</div>
          </div>
        </div>
      </div>
    </div>
    <hr />
    <div class="container is-small">
      <div class="columns is-centered is-multiline">

        <div class="column is-6" v-for="member in kids" v-bind:data="member" v-bind:key="member.name">
          <div class="card">
            <div class="card-header">
              <div class="container">
                <p class="title is-4">{{ member.name }}</p>
                <p class="subtitle is-6">{{ member.role }}</p>
              </div>
            </div>
            <div class="card-image">
              <figure class="image is-1by1">
                <img :src="member.image">
              </figure>
            </div>
            <div class="card-content">{{ member.bio }}</div>
          </div>
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
    name: 'BridalParty',

    components: {
    },

    data(){
      return {
        boys: [
          { name: 'James Davies', role: 'Best Man', img: 'https://bulma.io/images/placeholders/256x256.png', bio: 'Kara’s friend and housemate from uni, Laura was always known as the ‘mother hen’ while living on terraces, and certainly tried her best to keep Kara in line. She is the second brain of Kara, who can be terribly indecisive most of the time, so Laura may have helped make more decisions about this wedding than Ryan!'},
          { name: 'Luke Meeking', role: 'Groomsman', img: 'https://bulma.io/images/placeholders/256x256.png', bio: 'Kara’s friend and fellow Physio student at uni, prone to staying on Kara’s couch after a night out together, and responsible for introducing Kara to overseas travel when they embarked on Hawaii together.'},
          { name: 'Brendan Flett', role: 'Groomsman', img: 'https://bulma.io/images/placeholders/256x256.png', bio: 'Kara’s older sister, it’s impossible not to have a close bond when the two shared a bedroom for over 15 years. Despite the battle of the teenage years, the two are only getting closer in their adult lives and conversations can go on for hours whether in person or over the phone.'},
          { name: 'James Flett', role: 'Groomsman', img: 'https://bulma.io/images/placeholders/256x256.png', bio: 'Kara’s sister-in-law, but really another big sister to love and share special family moments together, Kasey always strives to stay close despite how many kilometers are between them.'},
        ],
        girls: [
          { name: 'Laura Sepe', role: 'Maid of Honour', img: 'https://bulma.io/images/placeholders/256x256.png', bio: 'Kara’s friend and housemate from uni, Laura was always known as the ‘mother hen’ while living on terraces, and certainly tried her best to keep Kara in line. She is the second brain of Kara, who can be terribly indecisive most of the time, so Laura may have helped make more decisions about this wedding than Ryan!'},
          { name: 'Danica Lang', role: 'Bridesmaid', img: 'https://bulma.io/images/placeholders/256x256.png', bio: 'Kara’s friend and fellow Physio student at uni, prone to staying on Kara’s couch after a night out together, and responsible for introducing Kara to overseas travel when they embarked on Hawaii together.'},
          { name: 'Laura Thomson', role: 'Bridesmaid', img: 'https://bulma.io/images/placeholders/256x256.png', bio: 'Kara’s older sister, it’s impossible not to have a close bond when the two shared a bedroom for over 15 years. Despite the battle of the teenage years, the two are only getting closer in their adult lives and conversations can go on for hours whether in person or over the phone.'},
          { name: 'Kasey Hazelman', role: 'Bridesmaid', img: 'https://bulma.io/images/placeholders/256x256.png', bio: 'Kara’s sister-in-law, but really another big sister to love and share special family moments together, Kasey always strives to stay close despite how many kilometers are between them.'},
        ],
        kids: [
          { name: 'Mia Thomson', role: 'Flower Girl', img: 'https://bulma.io/images/placeholders/256x256.png', bio: 'Kara’s niece – often accidentally called ‘Kara’ by her mother when in trouble.'},
          { name: 'Cooper Thomson', role: 'Page Boy', img: 'https://bulma.io/images/placeholders/256x256.png', bio: 'Kara’s nephew, often taking the blame when ‘Kara’/Mia is up to no good.'},
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

  .container {
    max-width: 1024px;
    margin-left: auto;
    margin-right: auto;
  }

  .card-header {
    padding: 15px;
  }

  .card-image {
    width: 50%;
    margin-left: auto;
    margin-right: auto;
  }

  @include col-xl() {
    .carousel {
      max-width: 75%;
      margin-left: auto;
      margin-right: auto;
    }

    .card-image {
      max-width: 75%;
      margin-left: auto;
      margin-right: auto;
    }
  }
</style>
