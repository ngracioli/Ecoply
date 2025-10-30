import type { NavigationGuardNext, RouteLocationNormalized } from "vue-router";
import { RouteNames } from "./types";

const TOKEN_STORAGE_KEY = "token" as const;

const getStoredToken = (): string | null => {
  return localStorage.getItem(TOKEN_STORAGE_KEY);
};

const isAuthenticationRequired = (route: RouteLocationNormalized): boolean => {
  return route.matched.some((record) => record.meta.requiresAuth);
};

export const authGuard = (
  to: RouteLocationNormalized,
  _from: RouteLocationNormalized,
  next: NavigationGuardNext,
): void => {
  const token = getStoredToken();
  const requiresAuth = isAuthenticationRequired(to);

  if (requiresAuth && !token) {
    next({ name: RouteNames.LOGIN });
    return;
  }

  if (to.name === RouteNames.LOGIN && token) {
    next({ name: RouteNames.DASHBOARD });
    return;
  }

  next();
};
