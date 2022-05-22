<template>
  <div class="col-md-12 mb-4" v-if="link">
    <div class="alert alert-info" role="alert">
      {{ link }}
    </div>
  </div>
  <div class="col-md-12 mb-4" v-if="error">
    <div class="alert alert-danger" role="alert">
      {{ error }}
    </div>
  </div>
  <div class="col-md-12 mb-4 input-group">
    <input class="form-control" placeholder="Search" @input="search" />
    <div class="input-group-append" v-if="selectedItem.length !== 0">
      <button class="btn btn-info" @click="generate">Generate Link</button>
    </div>
    <div class="input-group-append">
      <select class="form-select" @change="sort($event.target.value)">
        <option>Select</option>
        <option value="asc">Price Ascending</option>
        <option value="desc">Price Descending</option>
      </select>
    </div>
  </div>
  <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 g-3">
    <div
      class="col"
      v-for="product in products"
      :key="product.id"
      @click="select(product.id)"
    >
      <div
        class="card shadow-sm"
        :class="selectedItem.some((s) => s === product.id) ? 'selected' : ''"
      >
        <img :src="product.image" height="200" />
        <div class="card-body">
          <p class="card-text">
            {{ product.title }}
          </p>
          <div class="d-flex justify-content-between align-items-center">
            <small class="text-muted">{{ product.price }}</small>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div
    class="d-flex justify-content-center mt-4"
    v-if="filters.page < lastPage"
  >
    <button class="btn btn-primary" @click="loadMore">Load More</button>
  </div>
</template>
<script>
import { ref } from "vue";
import axios from "axios";
export default {
  props: {
    products: {
      type: Array,
      required: true,
    },
    filters: {
      type: Object,
      default: () => ({}),
    },
    lastPage: {
      type: Number,
      required: true,
    },
  },
  emits: ["onFiltering"],
  setup(props, { emit }) {
    const selectedItem = ref([]);
    const link = ref("");
    const error = ref("");

    const search = (e) => {
      emit("onFiltering", {
        ...props.filters,
        s: e.target.value,
        page: 1,
      });
    };

    const sort = (value) => {
      emit("onFiltering", {
        ...props.filters,
        sort: value,
        page: 1,
      });
    };

    const loadMore = () => {
      emit("onFiltering", {
        ...props.filters,
        page: props.filters.page + 1,
      });
    };

    const select = (id) => {
      if (selectedItem.value.some((s) => s === id)) {
        selectedItem.value = selectedItem.value.filter((s) => s !== id);
      } else {
        selectedItem.value.push(id);
      }
    };

    const generate = async () => {
      try {
        const { data } = await axios.post("links", {
          products: selectedItem.value,
        });
        link.value = `Link generate: ${process.env.VUE_APP_CHECKOUT_URL}/${data.code}`;
      } catch (err) {
        error.value = "You should be logged in to generate a link!";
      } finally {
        setTimeout(() => {
          link.value = "";
          error.value = "";
          selectedItem.value = [];
        }, 5000);
      }
    };

    return {
      search,
      sort,
      loadMore,
      select,
      selectedItem,
      generate,
      link,
      error,
    };
  },
};
</script>

<style scoped>
.card {
  cursor: pointer;
}
.card.selected {
  border: 4px solid darkcyan;
}
</style>
