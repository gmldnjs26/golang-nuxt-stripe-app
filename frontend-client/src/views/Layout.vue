<template>
  <div>
    <Nav />
    <main>
      <Header v-show="showHeader" />
      <div class="album py-5 bg-light">
        <div class="container">
          <router-view></router-view>
        </div>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import Nav from "@/components/Nav.vue";
import Header from "@/components/Header.vue";
import { computed, onMounted } from "vue";
import axios from "axios";
import { useStore } from "vuex";
import { useRoute } from "vue-router";

export default {
  components: { Nav, Header },
  setup() {
    const store = useStore();
    const route = useRoute();

    const showHeader = computed(() => {
      return route.path === "/" || route.path === "/backend";
    });

    onMounted(async () => {
      try {
        const { data } = await axios.get("user");
        store.dispatch("set_user", data);
      } catch (err) {
        console.log(err);
        store.dispatch("set_user", null);
      }
    });

    return {
      showHeader,
    };
  },
};
</script>

<style></style>
