<template>
  <div>
    <b-navbar class="nav-barr">
      <template slot="brand">
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          <img src="/img/website-logo.png" alt="Pbg Group " />
        </b-navbar-item>
      </template>
      <template slot="start"></template>

      <template slot="end">
        <b-navbar-item tag="div">
          <div class="buttons">
            <a class="button blue-button" @click="signUp()">
              <strong>Sign up</strong>
            </a>
          </div>
        </b-navbar-item>
      </template>
    </b-navbar>
  </div>
</template>

<script>
export default {
  methods: {
    signUp() {
      this.$emit("register");
    }
  }
};
</script>

<style></style>
