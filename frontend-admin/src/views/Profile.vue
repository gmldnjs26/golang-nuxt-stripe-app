<template>
  <div>
    <h3>Account Information</h3>
    <form @submit.prevent="infoSubmit">
      <div class="mb-3">
        <v-text-field
          label="First Name"
          v-model="user.first_name"
        ></v-text-field>
      </div>
      <div class="mb-3">
        <v-text-field label="Last Name" v-model="user.last_name"></v-text-field>
      </div>
      <div class="mb-3">
        <v-text-field label="Email" v-model="user.email"></v-text-field>
      </div>
      <v-btn color="primary" type="submit">Save</v-btn>
    </form>

    <h3 class="mt-4">Change Password</h3>
    <form @submit.prevent="passwordSubmit">
      <div class="mb-3">
        <v-text-field label="Password" v-model="password"></v-text-field>
      </div>
      <div class="mb-3">
        <v-text-field
          label="Password Confirm"
          v-model="password_confirm"
        ></v-text-field>
      </div>
      <v-btn color="primary" type="submit">Save</v-btn>
    </form>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      password: "",
      password_confirm: "",
    };
  },
  computed: {
    user() {
      return this.$store.state.user;
    },
  },
  methods: {
    async infoSubmit() {
      const result = await axios.put("users/info", {
        first_name: this.user.first_name,
        last_name: this.user.last_name,
        email: this.user.email,
      });
      this.$store.dispatch("set_user", result.data);
    },
    async passwordSubmit() {
      await axios.put("users/password", {
        password: this.password,
        password_confirm: this.password_confirm,
      });
    },
  },
};
</script>

<style></style>
