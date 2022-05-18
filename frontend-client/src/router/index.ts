import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import Layout from "@/views/Layout.vue";
import Login from "@/views/Login.vue";
import Register from "@/views/Register.vue";
import Profile from "@/views/Profile.vue";
import ProductsFrontend from "@/views/ProductsFrontend.vue";
import Stats from "@/views/Stats.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/login",
    component: Login,
  },
  {
    path: "/register",
    component: Register,
  },
  {
    path: "",
    component: Layout,
    children: [
      { path: "", component: ProductsFrontend },
      { path: "/profile", component: Profile },
      { path: "/backend", component: Profile },
      { path: "/stats", component: Stats },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
