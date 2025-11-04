export const AUTH_ENDPOINTS = {
  LOGIN: "/api/v1/auth/login",
  SIGNUP: "/api/v1/auth/signup",
  AVAILABLE: "/api/v1/auth/available",
  ME: "/api/v1/me",
  MY_OFFERS: "/api/v1/me/offers",
} as const;

export const OFFER_ENDPOINTS = {
  LIST: "/api/v1/offers",
  DETAIL: (uuid: string) => `/api/v1/offers/${uuid}`,
  DELETE: (uuid: string) => `/api/v1/offers/${uuid}`,
  CREATE: "/api/v1/offers",
  MY_OFFERS: "/api/v1/me/offers",
} as const;
