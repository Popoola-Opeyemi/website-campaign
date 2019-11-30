<template>
  <div>
    <form class="__loginform" @submit.prevent="userLogin">
      <b-notification
        :type="authstatus.class"
        has-icon
        :active.sync="authstatus.toggle"
        >{{ authmsg }}</b-notification
      >

      <p class="__loginheader">Admin Login</p>
      <b-field
        label="Email"
        :type="errors.has('email') ? 'is-danger' : ''"
        :message="errors.has('email') ? errors.first('email') : ''"
        class="label"
      >
        <b-input
          name="email"
          v-validate="'required|email'"
          v-model="form.email"
          placeholder="Enter email address"
        />
      </b-field>

      <b-field
        label="Password"
        :type="errors.has('Password') ? 'is-danger' : ''"
        :message="errors.has('Password') ? errors.first('Password') : ''"
        class="label"
      >
        <b-input
          type="password"
          name="Password"
          v-validate="'required'"
          password-reveal
          v-model="form.password"
          placeholder="Enter Password"
        />
      </b-field>

      <div class="__loginbutton">
        <div class="columns">
          <div class="column">
            <button class="button is-outlined is-fullwidth is-primary">
              Login
            </button>
          </div>
          <div class="column" />
        </div>
      </div>
    </form>
  </div>
</template>

<script>
import { clone } from "@/utils/helpers";
import { Notification } from "@/utils/helpers";
export default {
  layout: "login",
  data() {
    return {
      checked: true,
      loading: null,
      authmsg: null,
      authstatus: {
        class: "",
        toggle: false
      },

      form: {
        email: null,
        password: null
      }
    };
  },
  methods: {
    async userLogin() {
      let result = await this.$validator.validateAll();

      if (!result) {
        return;
      }

      try {
        let req = await this.$store.dispatch("auth/login", clone(this.form));
        if (req.status) {
          Notification(this, "User Authenticated", "is-success");

          setTimeout(() => (window.location = "/admin/dashboard"), 1500);
        } else {
          Notification(this, "incorrect Username or Password", "is-danger");
        }
      } catch (error) {
        Notification(this, "An error occured", "is-danger");
      }
    }
  }
};
</script>

<style></style>
