import { createWebHistory, createRouter } from "vue-router";
import { RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    alias: "/home",
    name: "Home",
    component: () => import("../pages/Home.vue"),
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../pages/Login.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;