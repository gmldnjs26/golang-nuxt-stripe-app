<template>
  <div class="container">
    <main>
      <div class="py-5 text-center">
        <h2>Success!</h2>
        <p class="lead">Your purchase has been completed!</p>
        <button @click="onClick">asdasd</button>
      </div>
    </main>
  </div>
</template>

<script>
export default {
  data() {
    return {
      text: 's',
    }
  },
  async fetch() {
    try {
      const { data } = await this.$axios.post('/api/admin/product', {
        source: this.$route.query.source,
      })
      this.text = data
    } catch (err) {
      console.error(
        `
          ${new Date().toLocaleString()}
          Stauts: ${err.response.status}
          StatusText: ${err.response.statusText}
          Message: ${err.response.data && err.response.data.message}
        `
      )
    }
  },
  methods: {
    async onClick() {
      console.log(this.$axios.defaults.baseURL)
      const { data } = await this.$axios.post('/api/checkout/orders/confirm', {
        source: this.$route.query.source,
      })
      console.log(data)
    },
  },
}
</script>

<style></style>
