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
          { name: 'James Davies', role: 'Best Man',  bio: 'James and Ryan have been close friends since primary school, despite 3 interstate moves between them. Over the years they have built great memories driving paddock bashers on the farm, talking astronomy and philosophy while staring at the stars as kids, or playing live music alongside one another through high school.'},
          { name: 'Brendan Flett', role: 'Groomsman',  bio: 'Brendan has been by Ryan\'s side ever since he was 3 years old and will be again for this special day. You\'d often find Brendan pushing Ryan around on his tricycle or trialing some dangerous activity \'just to be sure\'.'},
          { name: 'James Flett', role: 'Groomsman', bio: 'James was always that pesky 3rd child so it makes sense for him to be 3rd groomsman too. Look out ladies because this one is single as well - just like his golf handicap!'},
          { name: 'Luke Meeking', role: 'Groomsman & MC', bio: 'Luke and Ryan have been great friends from high school and have become closer over the years through high school productions and living in Brisbane together. Kara and Luke have been Rogers Street neighbours since they were born and completed their Deb together in 2010.'},
        ],
        girls: [
          { name: 'Laura Sepe', role: 'Maid of Honour', bio: 'Kara’s friend and housemate from uni, Laura was always known as the ‘mother hen’ while living on terraces, and certainly tried her best to keep Kara in line. She is the second brain of Kara, who can be terribly indecisive most of the time, so Laura may have helped make more decisions about this wedding than Ryan!'},
          { name: 'Laura Thomson', role: 'Bridesmaid', bio: 'Kara’s older sister, it’s impossible not to have a close bond when the two shared a bedroom for over 15 years. Despite the battle of the teenage years, the two are only getting closer in their adult lives and conversations can go on for hours whether in person or over the phone.'},
          { name: 'Danica Lang', role: 'Bridesmaid', bio: 'Kara’s friend and fellow Physio student at uni, Dern was prone to staying on Kara’s couch after a night out together. The two girls love to share an adventure with each-other whether that be Hawaii or the snow, or drinking too many gins. Ryan and Ash also love a cheeky motorbike ride or a craft beer.'},
          { name: 'Kasey Hazelman', role: 'Bridesmaid', bio: 'Kara’s sister-in-law, but really another big sister to love and share special family moments and a couple of drinks together. Kasey and Kara always remain close despite how many kilometres are between them.'},
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
