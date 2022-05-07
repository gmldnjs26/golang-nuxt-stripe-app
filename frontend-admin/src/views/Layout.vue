<template>
  <v-app>
    <Nav />
    <div class="container-fluid">
      <div class="row">
        <Menu />
        <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
          <div class="table-responsive">
            <table class="table table-striped table-sm">
              <thead>
                <tr>
                  <th scope="col">#</th>
                  <th scope="col">Header</th>
                  <th scope="col">Header</th>
                  <th scope="col">Header</th>
                  <th scope="col">Header</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>1,001</td>
                  <td>random</td>
                  <td>data</td>
                  <td>placeholder</td>
                  <td>text</td>
                </tr>
              </tbody>
            </table>
          </div>
        </main>
      </div>
    </div>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import Menu from "@/components/Menu.vue";
import Nav from "@/components/Nav.vue";

export default Vue.extend({
  components: { Nav, Menu },
  name: "Layout",

  data: () => ({
    //
  }),
});
</script>

<style></style>
