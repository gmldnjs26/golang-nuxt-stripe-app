<template>
  <v-app>
    <Nav :user="user" />
    <div class="container-fluid">
      <div class="row">
        <Menu />
        <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
          <div class="table-responsive">
            <router-view></router-view>
          </div>
        </main>
      </div>
    </div>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import Menu from "@/components/Menu.vue";
import Nav from "@/components/Nav.vue";
import axios from "axios";
import { User } from "@/types/user";

export default Vue.extend({
  components: { Nav, Menu },
  name: "Layout",
  async mounted() {
    try {
      const { data } = await axios.get("user");
      this.user = data;
    } catch (err) {
      this.$router.push("/login");
      console.log(err);
    }
  },

  data: () => ({
    user: {} as User,
  }),
});
</script>

<style></style>
