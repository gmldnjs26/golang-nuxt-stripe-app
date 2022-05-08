<template>
  <header
    class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow"
  >
    <a class="navbar-brand col-md-3 col-lg-2 me-0 px-3" href="#"
      >Company name</a
    >
    <div class="navbar-nav">
      <div class="nav-item text-nowrap">
        <router-link class="nav-link px-3" to="/profile"
          >{{ user.first_name }} {{ user.last_name }}
        </router-link>
      </div>
      <div class="nav-item text-nowrap">
        <a class="nav-link px-3" @click="logout" href="#">Sign out</a>
      </div>
    </div>
  </header>
</template>

<script lang="ts">
import Vue from "vue";
import { User } from "@/types/user";
import axios from "axios";

export default Vue.extend({
  name: "Nav",
  props: {
    user: {
      type: Object as () => User,
      default: () => {},
    },
  },
  methods: {
    async logout() {
      await axios.post("logout");
      this.$router.push("/login");
    },
  },
});
</script>

<style scoped>
.navbar-nav {
  flex-direction: row !important;
}
</style>
