<template>
  <Products :products="products" :filters="filters" @on-filtering="setFilter" />
</template>

<script lang="ts">
import Products from "./Products.vue";
import { Product } from "@/types/products";
import { onMounted, reactive, ref } from "vue";
import axios from "axios";
import { Filter } from "@/types/filter";
export default {
  components: { Products },
  setup() {
    const products = ref<Product[]>([]);
    const filters = reactive({
      s: "",
    });

    const setFilter = (value: Filter) => {
      filters.s = value.s;
      load(filters);
    };

    const load = async (f: Filter) => {
      filters.s;
      const arr = [];

      if (filters.s) {
        arr.push(`s=${filters.s}`);
      }

      const { data } = await axios.get(`products/backend?${arr.join("&")}`);
      products.value = data.data;
    };

    onMounted(() => {
      load(filters);
    });

    return {
      products,
      filters,
      setFilter,
    };
  },
};
</script>

<style></style>
