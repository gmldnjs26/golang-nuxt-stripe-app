// import { logger2 } from '../server/'

export default ({ $axios }) => {
  $axios.onResponse((response) => {
    console.log(`[${response.status}] ${response.request.path}`)
  })

  $axios.onError(async (err) => {
    if (!process.browser) {
      $axios.defaults.baseURL = 'http://localhost:5000'
      await $axios.post('/api/error', {
        message: `code: [${err.response && err.response.status}] path: ${
          err.response && err.response.request.path
        } message: ${
          err.response && err.response.data && err.response.data.message
        }`,
      })
    }
  })
}
