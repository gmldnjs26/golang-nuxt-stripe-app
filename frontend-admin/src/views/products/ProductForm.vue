<template>
  <form @submit.prevent="submit">
    <div class="mb-3">
      <v-text-field label="Title" v-model="title"></v-text-field>
    </div>
    <div class="mb-3">
      <v-text-field label="Description" v-model="description"></v-text-field>
    </div>
    <div class="mb-3">
      <v-text-field label="Image" v-model="image"></v-text-field>
    </div>
    <div class="mb-3">
      <v-text-field
        label="Price"
        type="number"
        min="0"
        v-model="price"
      ></v-text-field>
    </div>
    <v-btn color="primary" type="submit">Save</v-btn>
  </form>
</template>

<script lang="ts">
import axios from "axios";
import Vue from "vue";
export default Vue.extend({
  data() {
    return {
      title: "",
      description: "",
      image: "",
      price: "",
    };
  },
  methods: {
    async submit() {
      const result = await axios.post(
        "products",
        JSON.stringify({
          title: this.title,
          description: this.description,
          image: this.image,
          price: this.price,
        })
      );
      console.log(result);

      this.$router.push("/products");
    },
  },
});
</script>

<style></style>
