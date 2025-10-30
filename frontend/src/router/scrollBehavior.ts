import type { RouteLocationNormalized, RouterScrollBehavior } from "vue-router";

const SCROLL_OFFSET = 80;
const HASH_SCROLL_DELAY = 100;

export const scrollBehavior: RouterScrollBehavior = (
  to: RouteLocationNormalized,
  _from: RouteLocationNormalized,
  savedPosition,
) => {
  if (to.hash) {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve({
          el: to.hash,
          behavior: "smooth",
          top: SCROLL_OFFSET,
        });
      }, HASH_SCROLL_DELAY);
    });
  }

  if (savedPosition) {
    return savedPosition;
  }

  return { top: 0 };
};
