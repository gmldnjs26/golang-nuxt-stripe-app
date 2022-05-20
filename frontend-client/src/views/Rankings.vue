<template>
  <div class="table-responsive">
    <table class="table table-striped table-sm">
      <thead>
        <tr>
          <th>#</th>
          <th>Name</th>
          <th>Revenue</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(value, key, index) in rankings" :key="index">
          <td>{{ index + 1 }}</td>
          <td>{{ key }}</td>
          <td>{{ value }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import { Ranking } from "@/types/ranking";
import { ref } from "@vue/reactivity";
import { onMounted } from "@vue/runtime-core";
import axios from "axios";
export default {
  setup() {
    const rankings = ref<Ranking>({});

    onMounted(async () => {
      const { data } = await axios.get("rankings");
      rankings.value = data;
    });

    return {
      rankings,
    };
  },
};
</script>

<style></style>
