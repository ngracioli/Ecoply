import { createRouter, createWebHistory } from "vue-router";
// Imports diretos para páginas principais
import LandingPage from "../views/LandingPage.vue";
import Login from "../views/auth/Login.vue";
import Register from "../views/auth/Register.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: LandingPage,
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
  {
    path: "/register",
    name: "Register",
    component: Register,
  },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: () => import("../views/dashboard/Dashboard.vue"),
  },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: () => import("../views/dashboard/Dashboard.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,

  scrollBehavior(to, from, savedPosition) {
    if (to.hash) {
      return new Promise((resolve) => {
        setTimeout(() => {
          resolve({
            el: to.hash,
            behavior: "smooth",
            top: 80,
          });
        }, 100);
      });
    }

    if (savedPosition) {
      return savedPosition;
    }

    return { top: 0 };
  },
});

export default router;
