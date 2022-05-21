<template>
  <Products
    :products="filteredProducts"
    :filters="filters"
    @on-filtering="setFilter"
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
    const allProducts = ref<Product[]>([]);
    const filteredProducts = ref<Product[]>([]);
    const filters = reactive<Filter>({
      s: "",
      sort: "",
      page: 1,
    });
    const lastPage = ref(0);
    const perPage = 9;

    onMounted(async () => {
      const { data } = await axios.get("products/frontend");
      allProducts.value = data;
      filteredProducts.value = data.slice(0, filters.page * perPage);
      lastPage.value = Math.ceil(data.length / perPage);
    });

    const setFilter = (value: Filter) => {
      filters.s = value.s;
      filters.sort = value.sort;
      filters.page = value.page;

      let products = allProducts.value.filter((p) => {
        return (
          p.title.toLowerCase().indexOf(filters.s.toLowerCase()) >= 0 ||
          p.description.toLowerCase().indexOf(filters.s.toLowerCase()) >= 0
        );
      });

      if (filters.sort === "asc" || filters.sort === "desc") {
        products.sort((a, b) => {
          const diff = a.price - b.price;

          if (diff === 0) return 0;

          const sign = Math.abs(diff) / diff;

          return filters.sort === "asc" ? sign : -sign;
        });
      }

      lastPage.value = Math.ceil(products.length / perPage);
      filteredProducts.value = products.slice(0, filters.page * perPage);

      // if (filters.sort === "asc") {
      //   filteredProducts.value.sort((a, b) => a.price - b.price);
      // } else if (filters.sort === "desc") {
      //   filteredProducts.value.sort((a, b) => b.price - a.price);
      // }
    };

    return {
      filteredProducts,
      filters,
      setFilter,
      lastPage,
    };
  },
};
</script>

<style></style>
