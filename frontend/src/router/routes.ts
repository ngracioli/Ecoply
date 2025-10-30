import type { AppRoute } from "./types";
import { RouteNames, RoutePaths } from "./types";
import LandingPage from "../views/LandingPage.vue";
import Login from "../views/auth/Login.vue";
import Register from "../views/auth/Register.vue";

export const routes: AppRoute[] = [
  {
    path: RoutePaths.HOME,
    name: RouteNames.HOME,
    component: LandingPage,
  },
  {
    path: RoutePaths.LOGIN,
    name: RouteNames.LOGIN,
    component: Login,
  },
  {
    path: RoutePaths.REGISTER,
    name: RouteNames.REGISTER,
    component: Register,
  },
  {
    path: RoutePaths.DASHBOARD,
    name: RouteNames.DASHBOARD,
    component: () => import("../views/dashboard/Dashboard.vue"),
    meta: { requiresAuth: true },
  },
  {
    path: RoutePaths.OFFER_DETAIL,
    name: RouteNames.OFFER_DETAIL,
    component: () => import("../views/OfferDetail.vue"),
    meta: { requiresAuth: true },
  },
];
