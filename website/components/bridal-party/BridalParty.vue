<template>
  <section class="section is-small">
    <h1 class="title" id="bridalParty">The Bridal Party</h1>
    <div class="container is-small">

      <!--the boys-->
      <div class="columns is-centered is-multiline">
        <div class="column">
          <div class="card">
            <div class="card-image">
              <figure class="image">
                <img src="~assets/img/boys.png">
                <p>James Davies, Brendan Flett, Ryan Flett, Luke Meeking, James Flett</p>
              </figure>
            </div>
            <div class="card-content">
              <div class="block" v-for="boy in boys">
                <hr />
                <h1 class="title is-4">{{ boy.name }}</h1>
                <h1 class="subtitle is-6">{{  boy.role }}</h1>
                <p>{{ boy.bio }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!--the girls-->
      <div class="columns is-centered is-multiline">
        <div class="column">
          <div class="card">
            <div class="card-image">
              <figure class="image">
                <img src="~assets/img/girls.png">
                <p>Danica Day, Laura Sepe, Kara Hazelman, Kasey Hazelman, Laura Thomson</p>
                <p>Cooper Thomson, Mia Thomson</p>
              </figure>
            </div>
            <div class="card-content">
              <div class="block" v-for="girl in girls">
                <hr />
                <h1 class="title is-4">{{ girl.name }}</h1>
                <h1 class="subtitle is-6">{{  girl.role }}</h1>
                <p>{{ girl.bio }}</p>
              </div>
            </div>
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
          { name: 'James Davies', role: 'Best Man',  bio: 'Kara’s friend and housemate from uni, Laura was always known as the ‘mother hen’ while living on terraces, and certainly tried her best to keep Kara in line. She is the second brain of Kara, who can be terribly indecisive most of the time, so Laura may have helped make more decisions about this wedding than Ryan!'},
          { name: 'Brendan Flett', role: 'Groomsman',  bio: 'Kara’s older sister, it’s impossible not to have a close bond when the two shared a bedroom for over 15 years. Despite the battle of the teenage years, the two are only getting closer in their adult lives and conversations can go on for hours whether in person or over the phone.'},
          { name: 'James Flett', role: 'Groomsman', bio: 'Kara’s sister-in-law, but really another big sister to love and share special family moments together, Kasey always strives to stay close despite how many kilometers are between them.'},
          { name: 'Luke Meeking', role: 'Groomsman', bio: 'Kara’s friend and fellow Physio student at uni, prone to staying on Kara’s couch after a night out together, and responsible for introducing Kara to overseas travel when they embarked on Hawaii together.'},
        ],
        girls: [
          { name: 'Laura Sepe', role: 'Maid of Honour', bio: 'Kara’s friend and housemate from uni, Laura was always known as the ‘mother hen’ while living on terraces, and certainly tried her best to keep Kara in line. She is the second brain of Kara, who can be terribly indecisive most of the time, so Laura may have helped make more decisions about this wedding than Ryan!'},
          { name: 'Laura Thomson', role: 'Bridesmaid', bio: 'Kara’s older sister, it’s impossible not to have a close bond when the two shared a bedroom for over 15 years. Despite the battle of the teenage years, the two are only getting closer in their adult lives and conversations can go on for hours whether in person or over the phone.'},
          { name: 'Danica Lang', role: 'Bridesmaid', bio: 'Kara’s friend and fellow Physio student at uni, prone to staying on Kara’s couch after a night out together, and responsible for introducing Kara to overseas travel when they embarked on Hawaii together.'},
          { name: 'Kasey Hazelman', role: 'Bridesmaid', bio: 'Kara’s sister-in-law, but really another big sister to love and share special family moments together, Kasey always strives to stay close despite how many kilometers are between them.'},
          { name: 'Mia Thomson', role: 'Flower Girl', bio: 'Kara’s niece – often accidentally called ‘Kara’ by her mother when in trouble.'},
          { name: 'Cooper Thomson', role: 'Page Boy', bio: 'Kara’s nephew, often taking the blame when ‘Kara’/Mia is up to no good.'},
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

  .card-image {
    margin-left: auto;
    margin-right: auto;
    padding-top: 20px;
  }

  .image p {
    padding-top: 5px;
    text-align: center;
    font-size: 14px;
  }

  @include col-xl() {
    .card-image {
      max-width: 75%;
      margin-left: auto;
      margin-right: auto;
    }
  }
</style>
