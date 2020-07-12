<template>
  <section class="hero is-fullheight">

    <div class="hero-body">
      <div class="date-text">
        <h1 class="save-the-date">Save the date</h1>
        <h1 class="date">26 . 03 . 21</h1>
      </div>
    </div>

    <div class="hero-footer">
      <div class="columns">

        <div class="column location-tag">
          <i class="fa fa-map-marker-alt" aria-hidden="true"></i>
          <a href="https://goo.gl/maps/TfczQik3ymnXD9vN8" target="_blank">COROWA | NSW</a>

          <p></p>
          
          <i class="fab fa-instagram"/>
          <a href="https://www.instagram.com/explore/tags/flettsgetmarried/" target="_blank">#FLETTSGETMARRIED</a>
        </div>

        <div class="column more-info">
          <h2>More info including RSVPs coming soon...</h2>
        </div>

        <div class="column nicole">
          <a href="https://www.instagram.com/njd.photos/" target="_blank">
            <i class="fab fa-instagram"/>
            njd.photos
          </a>
        </div>

      </div>
    </div>
  </section>
</template>

<script lang="ts">
  export default {
    name: 'SaveTheDate',

    components: {
    }

  }
</script>

<style lang="scss" scoped>
  @import "~assets/all";

  h1, h2 {
    color: white;
    font-weight: 400;
  }

  i {
    color: white;
  }

  a {
    color: white;
    text-align: right;
  }

  a:hover {
    color: #5fd2ff;
    text-decoration: underline;
  }

  .hero {
    background-position: center;
    background-size: cover;
    background-repeat: no-repeat;
    position: relative;
  }

  .hero-body {
    padding: 0px;
  }

  .date-text {
    width: 100%;
    position: absolute;
    text-align: center;
    letter-spacing: 8px;
  }

  .save-the-date {
    font-size: 5em;
    font-family: 'Scholastica', cursive;
    height: 100px;
  }

  .date {
    font-size: 2em;
    font-weight: 300;
  }

  .more-info {
    font-size: 18px;
    font-weight: 300;
    text-align: center;
  }

  .location-tag {
    font-size: 24px;
  }

  @include col-sm() {
    .hero {
      background-image: url("~assets/img/desktop.jpg");
    }
    .hero-footer {
      text-align: left;
      padding: 30px;
    }
    .date-text {
      bottom: 160px;
    }
    .save-the-date {
      font-size: 8em;
      height: 140px;
    }
    .date {
      font-size: 3em;
    }
    .columns {
      width: 100%;
    }
    .nicole {
      text-align: right;
      float: right;
      padding-top: 40px;
    }
  }

  @include sm-down() {
    .hero {
      background-image: url("~assets/img/mobile.jpg");
    }
    .date-text {
      bottom: 200px;
    }
    .hero-footer {
      text-align: center;
    }
    .columns, .columns:last-child {
      margin: 0px;
    }
    .nicole {
      text-align: right;
    }
  }
  
</style>
