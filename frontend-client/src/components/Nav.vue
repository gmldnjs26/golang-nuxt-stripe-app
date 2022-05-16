<template>
  <div class="container">
    <header
      class="d-flex flex-wrap align-items-center justify-content-center justify-content-md-between py-3 mb-4 border-bottom"
    >
      <ul class="nav col-12 col-md-auto mb-2 justify-content-center mb-md-0">
        <li><a href="#" class="nav-link px-2 link-secondary">Frontend</a></li>
        <li><a href="#" class="nav-link px-2 link-dark">Backend</a></li>
      </ul>

      <div class="col-md-3 text-end" v-if="!user">
        <button type="button" class="btn btn-outline-primary me-2">
          Login
        </button>
        <button type="button" class="btn btn-primary">Sign-up</button>
      </div>
      <div v-else>
        <button type="button" class="btn btn-outline-primary me-2">
          Logout
        </button>
        <button type="button" class="btn btn-primary">
          {{ user.first_name }} {{ user.last_name }}
        </button>
      </div>
    </header>
  </div>
</template>

<script>
import { computed } from "@vue/runtime-core";
import { useStore } from "vuex";
export default {
  setup() {
    const store = useStore();

    const user = computed(() => {
      return store.state.user;
    });

    return {
      user,
    };
  },
};
</script>

<style></style>
