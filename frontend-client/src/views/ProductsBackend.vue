<template>
  <Products
    :products="products"
    :filters="filters"
    @on-filtering="load"
    :lastPage="lastPage"
  />
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
    const filters = reactive<Filter>({
      s: "",
      sort: "",
      page: 1,
    });
    const lastPage = ref(0);

    const load = async (f: Filter) => {
      filters.s = f.s;
      filters.sort = f.sort;
      filters.page = f.page;
      const arr = [];

      if (filters.s) {
        arr.push(`s=${filters.s}`);
      }
      if (filters.sort) {
        arr.push(`sort=${filters.sort}`);
      }
      if (filters.sort) {
        arr.push(`page=${filters.page}`);
      }

      const { data } = await axios.get(`products/backend?${arr.join("&")}`);
      products.value =
        filters.page === 1 ? data.data : [...products.value, ...data.data];
      lastPage.value = data.last_page;
    };

    onMounted(() => {
      load(filters);
    });

    return {
      products,
      filters,
      load,
      lastPage,
    };
  },
};
</script>

<style></style>
