<template>
  <main class="form-signin">
    <form @submit.prevent="onSumbit">
      <h1 class="h3 mb-3 fw-normal">Please sign up</h1>

      <div class="form-floating">
        <input
          v-model="first_name"
          class="form-control"
          placeholder="First Name"
        />
        <label>First Name</label>
      </div>
      <div class="form-floating">
        <input
          v-model="last_name"
          class="form-control"
          placeholder="Last Name"
        />
        <label>Last Name</label>
      </div>
      <div class="form-floating">
        <input
          v-model="email"
          type="email"
          class="form-control"
          placeholder="name@example.com"
        />
        <label>Email address</label>
      </div>
      <div class="form-floating">
        <input
          v-model="password"
          type="password"
          class="form-control"
          placeholder="Password"
        />
        <label>Password</label>
      </div>
      <div class="form-floating">
        <input
          v-model="password_confirm"
          type="password"
          class="form-control"
          placeholder="Password Confirm"
        />
        <label>Password Confirm</label>
      </div>
      <button class="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
    </form>
  </main>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      first_name: "",
      last_name: "",
      email: "",
      password: "",
      password_confirm: "",
    };
  },
  methods: {
    async onSumbit() {
      const result = await axios.post("register", {
        first_name: this.first_name,
        last_name: this.last_name,
        email: this.email,
        password: this.password,
        password_confirm: this.password_confirm,
      });
      if (result) this.$router.push("/login");
      return result;
    },
  },
};
</script>

<style scoped>
html,
body {
  height: 100%;
}

body {
  display: flex;
  align-items: center;
  padding-top: 40px;
  padding-bottom: 40px;
  background-color: #f5f5f5;
}

.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: auto;
}

.form-signin .checkbox {
  font-weight: 400;
}

.form-signin .form-floating:focus-within {
  z-index: 2;
}

.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}

.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>
