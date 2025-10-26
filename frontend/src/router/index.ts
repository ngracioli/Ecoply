import { createRouter, createWebHistory } from "vue-router";
// Imports diretos para páginas principais
import LandingPage from "../views/LandingPage.vue";
import Login from "../views/auth/Login.vue";
import Register from "../views/auth/Register.vue";
import EcoplyInfo from "../views/EcoplyInfo/EcoplyInfo.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: LandingPage,
  },
  {
      path: "/info",
    name: "Info",
    component: EcoplyInfo,
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
    meta: { requiresAuth: true },
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

router.beforeEach((to, _from, next) => {
  const token = localStorage.getItem("token");
  const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);

  if (requiresAuth && !token) {
    next({ name: "Login" });
  } else if (to.name === "Login" && token) {
    next({ name: "Dashboard" });
  } else {
    next();
  }
});

export default router;
