<template>
  <section class="hero">

    <div class="hero-header"></div>
    <div class="hero-body"></div>
    <div class="hero-foot">
      <div class="content has-text-centered">
        <h1 class="save-the-date">25 . 03 . 22</h1>
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

  h1 {
    font-weight: 400;
    margin-right: auto;
    margin-left: auto;

    color: whitesmoke;
  }

  .hero {
    background-position: center;
    background-size: cover;
    background-repeat: no-repeat;
    min-height: calc(100vh - 6.25rem);
  }

  .save-the-date {
    font-family: 'Scholastica', cursive;
  }

  @include col-sm() {
    .hero {
      background-image: url("~assets/img/desktop.jpg");
    }

    .save-the-date {
      font-size: 8em;
    }
  }

  @include sm-down() {
    .hero {
      background-image: url("~assets/img/mobile.jpg");

      .save-the-date {
        font-size: 8em;
      }
    }
  }

</style>
