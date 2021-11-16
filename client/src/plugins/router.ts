import {
  createWebHistory,
  createRouter,
  RouteLocationNormalized,
  NavigationGuardNext,
  RouteRecordNormalized,
} from "vue-router";
import { RouteRecordRaw } from "vue-router";
import decode from "jwt-decode";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    alias: "/home",
    name: "Home",
    component: () => import("../pages/Home.vue"),
    meta: { requiresAuth: true },
    children: [
      {
        path: "",
        alias: "/chat",
        name: "Chatbox",
        component: () => import("../components/ChatBox.vue"),
        meta: { requiresAuth: true },
        props: true,
      },
      {
        path: "/search/:searchString",
        name: "SearchResult",
        component: () => import("../components/SearchResult.vue"),
        meta: { requiresAuth: true },
        props: true,
      },
    ],
  },
  {
    path: "/login",
    name: "Login",
    meta: { requiresAuth: false },
    component: () => import("../pages/Login.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(
  (
    to: RouteLocationNormalized,
    from: RouteLocationNormalized,
    next: NavigationGuardNext
  ) => {
    if (
      to.matched.some(
        (record: RouteRecordNormalized) => record.meta.requiresAuth
      )
    ) {
      const mainToken = localStorage.getItem("main_token");
      if (mainToken === null) {
        next({
          path: "/login",
        });
      } else {
        const data = decode(mainToken) as any;
        const now = Date.now();
        const exp = data?.exp * 1000;
        if (now >= exp) {
          localStorage.removeItem("main_token");
          next({
            path: "/login",
          });
        } else {
          console.log("here");
          next();
        }
      }
    } else {
      next();
    }
  }
);

export default router;
