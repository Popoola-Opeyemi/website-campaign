<template>
  <div class="zero">
    <nav-bar />
    <nuxt />
  </div>
</template>

<script>
import NavBar from "@/components/navbar";
export default {
  components: {
    NavBar
  }
};
</script>

<style lang="scss">
@import "~assets/style.scss";
</style>
