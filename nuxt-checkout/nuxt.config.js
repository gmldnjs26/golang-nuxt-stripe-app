export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'nuxt-checkout',
    htmlAttrs: {
      lang: 'en',
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' },
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      {
        rel: 'stylesheet',
        href: 'https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta3/dist/css/bootstrap.min.css',
        integrity:
          'sha384-eOJMYsd53ii+scO/bJGFsiCZc+5NDVN2yr8+0RDqr0Ql0h+rP48ckxlpbzKgwra6',
        crossOrigin: 'anonymous',
      },
    ],
  },
  ssr: true,

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [{ src: '~/plugins/axios-init.js', mode: 'server' }],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    '@nuxt/typescript-build',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/bootstrap
    // 'bootstrap-vue/nuxt',
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    // https://www.npmjs.com/package/nuxt-stripe-module
    [
      'nuxt-stripe-module',
      {
        publishableKey:
          'pk_test_51KvyLfCXj5HN8X9OCn8znS9epfEESd1IYpZcyoTdtoZuNZZDWShPu3SD0AzW4JtWboKEgqnRRH3uRw1MzLhSnHou00ZbxubyJ7',
      },
    ],
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // Workaround to avoid enforcing hard-coded localhost:3000: https://github.com/nuxt-community/axios-module/issues/308
    baseURL: 'http://localhost:7070/',
    // proxy: true,
  },
  // proxy: {
  //   '/api/checkout/': {
  //     target: 'http://localhost:7070/',
  //     path: {
  //       '^/api/checkout/': '/api/checkout/',
  //     },
  //     changeOrigin: true,
  //   },
  // },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},
  server: {
    port: 5000,
  },
}
