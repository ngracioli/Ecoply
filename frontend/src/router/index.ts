import { createRouter, createWebHistory, type Router } from "vue-router";
import { routes } from "./routes";
import { scrollBehavior } from "./scrollBehavior";
import { authGuard } from "./guards";

const router: Router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  scrollBehavior,
});

router.beforeEach(authGuard);

export default router;
export { RouteNames, RoutePaths } from "./types";
export type { RouteMeta, AppRoute } from "./types";
