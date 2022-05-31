# golang-nuxt-stripe-app

VueJS with Typescript, Nuxt.js, Vuetify, Golang, Docker, Redis, Stripe, Gorutines, Channels

### Directory 

```bash
├── backend
│ ├── src
│ │ ├── commands // batch
│ │ ├── controllers 
│ │ ├── database
│ │ │ ├── db.go // mysql
│ │ │ ├── redis.go // redis for Revenue Rank and product list
│ │ ├── middlewares
│ │ │ ├── auth.go // jwt token, Login
│ │ ├── models
│ │ ├── routes
│ └── 이하생략..
├── frontend-admin // create product check user and order
├── frontend-client // create link
├── nuxt-checkout // for payment and checkout

```

- [ ] create GCS
- [ ] create Clound Functions(image resizing?)
- [ ] create prod env
- [ ] create stg env? 
- [ ] Fix BUG
