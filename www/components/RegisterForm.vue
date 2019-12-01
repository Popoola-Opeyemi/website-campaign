<template>
  <div class="modal-card">
    <header class="modal-card-head">
      <p class="modal-card-title">Job Contact Form</p>
    </header>
    <section class="modal-card-body">
      <div class="forms-content">
        <b-field
          label="Full Name"
          :type="errors.has('Full Name') ? 'is-danger' : ''"
          :message="errors.has('Full Name') ? errors.first('Full Name') : ''"
        >
          <b-input
            autocomplete="name"
            name="Full Name"
            v-validate="'required'"
            placeholder="Full Name"
            v-model="form.fullname"
          />
        </b-field>
        <b-field
          label="Mobile Number"
          :type="errors.has('Mobile Number') ? 'is-danger' : ''"
          :message="errors.has('Mobile Number') ? errors.first('Mobile Number') : ''"
        >
          <b-input
            placeholder="Mobile Number"
            v-model="form.mobile_number"
            name="Mobile Number"
            v-validate="'required|numeric'"
          ></b-input>
        </b-field>
        <b-field
          label="Job Category"
          :type="errors.has('Job Category') ? 'is-danger': ''"
          :message="errors.first('Job Category')"
        >
          <b-select
            v-model="form.order_type"
            expanded
            name="Job Category"
            placeholder="select a category"
            v-validate="'required'"
          >
            <option
              :value="option.id"
              v-for="(option, index) in RegisterSelect"
              :key="index"
            >{{ option.name }}</option>
          </b-select>
        </b-field>

        <b-field
          v-if="form.category_id == 3"
          label="Others Description"
          :type="errors.has('Description') ? 'is-danger' : ''"
          :message="errors.has('Description') ? errors.first('Description') : ''"
        >
          <b-input
            autocomplete="name"
            name="Description"
            v-validate="'required'"
            placeholder="Description"
            v-model="form.description"
          />
        </b-field>
        <b-field
          label="Email"
          :type="errors.has('Email') ? 'is-danger' : ''"
          :message="errors.has('Email') ? errors.first('Email') : ''"
        >
          <b-input
            name="Email"
            v-validate="'required|email'"
            placeholder="Email"
            v-model="form.email"
          />
        </b-field>

        <div class="register-button">
          <button class="button is-danger" @click="cancel">Cancel</button>
          <button class="button reg-button" @click="submit">Submit</button>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { clone } from "@/utils/helpers";
export default {
  data() {
    return {
      form: {},
      RegisterSelect: [
        { name: "Web Application Development", id: 1 },
        { name: "Mobile Application Development", id: 2 },
        { name: "Others", id: 3 }
      ]
    };
  },
  methods: {
    async submit() {
      let form = clone(this.form);
      let result = await this.$validator.validateAll();
      if (result) {
        this.$emit("save", form);
        this.$parent.close();
      }
    },
    cancel() {
      this.$parent.close();
    }
  }
};
</script>

<style>
</style>