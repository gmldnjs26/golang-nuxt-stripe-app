<template>
  <v-expansion-panels>
    <v-expansion-panel v-for="(order, i) in orders" :key="i">
      <v-expansion-panel-header>
        {{ order.name }} ${{ order.total }}</v-expansion-panel-header
      >
      <v-expansion-panel-content>
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-left">#</th>
                <th class="text-left">Product title</th>
                <th class="text-left">Price</th>
                <th class="text-left">Quantity</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in order.order_items" :key="item.id">
                <td>{{ item.id }}</td>
                <td>{{ item.product_title }}</td>
                <td>{{ item.price }}</td>
                <td>{{ item.quantity }}</td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      orders: [],
    };
  },
  async mounted() {
    const { data } = axios.get("orders");

    this.orders = data;
  },
};
</script>

<style></style>
