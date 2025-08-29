import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("../views/LandingPage.vue"),
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/auth/Login.vue"),
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
