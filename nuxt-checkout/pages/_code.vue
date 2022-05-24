<template>
  <div class="container">
    <main>
      <div class="py-5 text-center">
        <h2>Welcome</h2>
        <p class="lead">
          {{ user.first_name }} {{ user.last_name }}has invited you to buy these
          products
        </p>
      </div>

      <div class="row g-5">
        <div class="col-md-5 col-lg-4 order-md-last">
          <h4 class="d-flex justify-content-between align-items-center mb-3">
            <span class="text-primary">Products</span>
            <span class="badge bg-primary rounded-pill">3</span>
          </h4>
          <ul class="list-group mb-3">
            <div v-for="(product, i) in products" :key="i">
              <li class="list-group-item d-flex justify-content-between lh-sm">
                <div>
                  <h6 class="my-0">{{ product.title }}</h6>
                  <small class="text-muted">{{ product.description }}</small>
                </div>
                <span class="text-muted">${{ product.price }}</span>
              </li>
              <li class="list-group-item d-flex justify-content-between lh-sm">
                <div>
                  <h6 class="my-0">Quantity</h6>
                </div>
                <input
                  v-model="quantites[product.id]"
                  class="text-muted form-control quantity"
                  type="number"
                  min="0"
                />
              </li>
            </div>

            <li class="list-group-item d-flex justify-content-between">
              <span>Total (USD)</span>
              <strong>${{ total }}</strong>
            </li>
          </ul>
        </div>
        <div class="col-md-7 col-lg-8">
          <h4 class="mb-3">Personal Info</h4>
          <form class="needs-validation" novalidate @submit.prevent="submit">
            <div class="row g-3">
              <div class="col-sm-6">
                <label for="firstName" class="form-label">First name</label>
                <input
                  id="firstName"
                  v-model="first_name"
                  type="text"
                  class="form-control"
                  placeholder="First Name"
                  required
                />
              </div>

              <div class="col-sm-6">
                <label for="lastName" class="form-label">Last name</label>
                <input
                  id="lastName"
                  v-model="last_name"
                  type="text"
                  class="form-control"
                  placeholder="Last Name"
                  required
                />
              </div>

              <div class="col-12">
                <label for="email" class="form-label">Email </label>
                <input
                  id="email"
                  v-model="email"
                  type="email"
                  class="form-control"
                  placeholder="you@example.com"
                />
              </div>

              <div class="col-12">
                <label for="address" class="form-label">Address</label>
                <input
                  id="address"
                  v-model="address"
                  type="text"
                  class="form-control"
                  placeholder="1234 Main St"
                  required
                />
              </div>

              <div class="col-md-5">
                <label for="country" class="form-label">Country</label>
                <input
                  id="country"
                  v-model="country"
                  type="text"
                  class="form-control"
                  placeholder="Country"
                />
              </div>

              <div class="col-md-4">
                <label for="state" class="form-label">City</label>
                <input
                  id="city"
                  v-model="city"
                  type="text"
                  class="form-control"
                  placeholder="City"
                />
              </div>
              <div class="col-md-3">
                <label for="zip" class="form-label">Zip</label>
                <input
                  id="zip"
                  v-model="zip"
                  type="text"
                  class="form-control"
                  placeholder="Zip"
                  required
                />
              </div>
            </div>
            <hr class="my-4" />
            <button class="w-100 btn btn-primary btn-lg" type="submit">
              Checkout
            </button>
          </form>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import Vue from 'vue'

export default Vue.extend({
  name: 'IndexPage',
  async asyncData(ctx) {
    const { data } = await ctx.$axios.get(`links/${ctx.params.code}`)
    const user = data.user
    const products = data.products
    const quantites = {}
    products.forEach((p) => {
      quantites[p.id] = 0
    })

    return {
      user,
      quantites,
      products,
    }
  },
  data() {
    return {
      first_name: '',
      last_name: '',
      email: '',
      address: '',
      country: '',
      city: '',
      zip: '',
    }
  },
  computed: {
    total() {
      return this.products.reduce((s, p) => {
        return (s = s + p.price * this.quantites[p.id])
      }, 0)
    },
  },
  methods: {
    async submit() {
      const { data } = await this.$axios.post('orders', {
        first_name: this.first_name,
        last_name: this.last_name,
        email: this.email,
        address: this.address,
        country: this.country,
        city: this.city,
        zip: this.zip,
        code: this.$route.params.code,
        products: this.products.map((p) => ({
          product: p.id,
          quantity: parseInt(this.quantites[p.id]),
        })),
      })

      console.log(data)

      this.$stripe.redirectToCheckout({
        sessionId: data.transaction_id,
      })
    },
  },
})
</script>

<style scoped>
.quantity {
  width: 65px;
}
</style>
