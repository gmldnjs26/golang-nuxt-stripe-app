<template>
  <nav
    id="sidebarMenu"
    class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse"
  >
    <div class="position-sticky pt-3">
      <ul class="nav flex-column">
        <li class="nav-item">
          <router-link
            class="nav-link active"
            active-class="active"
            to="/users"
          >
            Users
          </router-link>
        </li>
        <li class="nav-item">
          <router-link
            class="nav-link active"
            active-class="active"
            to="/products"
          >
            Products
          </router-link>
        </li>
        <li class="nav-item">
          <router-link
            class="nav-link active"
            active-class="active"
            to="/orders"
          >
            Orders
          </router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
export default {};
</script>

<style></style>
