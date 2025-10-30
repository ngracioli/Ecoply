import type { RouteRecordRaw } from "vue-router";

export interface RouteMeta {
  requiresAuth?: boolean;
}

export type AppRoute = RouteRecordRaw & {
  meta?: RouteMeta;
};

export const enum RouteNames {
  HOME = "Home",
  LOGIN = "Login",
  REGISTER = "Register",
  DASHBOARD = "Dashboard",
  OFFER_DETAIL = "OfferDetail",
}

export const enum RoutePaths {
  HOME = "/",
  LOGIN = "/login",
  REGISTER = "/register",
  DASHBOARD = "/dashboard",
  OFFER_DETAIL = "/offer/:id",
}
