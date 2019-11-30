<template>
  <div class="hero is-fullheight login-page">
    <div class="hero-body">
      <div class="container">
        <div class="columns">
          <div class="column"></div>
          <div class="column loginContent">
            <nuxt />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  middleware: ["anonymous"]
};
</script>

<style lang="scss">
@import "~/assets/login.scss";
</style>
