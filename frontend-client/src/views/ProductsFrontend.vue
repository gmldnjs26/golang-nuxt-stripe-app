<template>
  <Products :products="products" />
</template>

<script lang="ts">
import Products from "./Products.vue";
import { Product } from "@/types/products";
import { onMounted, ref } from "vue";
import axios from "axios";
export default {
  components: { Products },
  setup() {
    const products = ref<Product[]>([]);

    onMounted(async () => {
      const { data } = await axios.get("products/frontend");
      products.value = data;
    });

    return {
      products,
    };
  },
};
</script>

<style></style>
