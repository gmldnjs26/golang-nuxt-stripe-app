<template>
  <div class="col-md-12 mb-4 input-group">
    <input class="form-control" placeholder="Search" @input="search" />
    <div class="input-group-append">
      <select class="form-select" @change="sort($event.target.value)">
        <option>Select</option>
        <option value="asc">Price Ascending</option>
        <option value="desc">Price Descending</option>
      </select>
    </div>
  </div>
  <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 g-3">
    <div class="col" v-for="product in products" :key="product.id">
      <div class="card shadow-sm">
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

    return {
      search,
      sort,
      loadMore,
    };
  },
};
</script>
