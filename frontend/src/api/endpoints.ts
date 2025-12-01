export const AUTH_ENDPOINTS = {
  LOGIN: "/api/v1/auth/login",
  SIGNUP: "/api/v1/auth/signup",
  AVAILABLE: "/api/v1/auth/available",
  ME: "/api/v1/me",
  MY_OFFERS: "/api/v1/me/offers",
  ANALYTICS: "/api/v1/me/analytics",
} as const;

export const ANALYTICS_ENDPOINTS = {
  PLATFORM: "/api/v1/analytics/platform",
} as const;

export const OFFER_ENDPOINTS = {
  LIST: "/api/v1/offers",
  DETAIL: (uuid: string) => `/api/v1/offers/${uuid}`,
  DELETE: (uuid: string) => `/api/v1/offers/${uuid}`,
  UPDATE: (uuid: string) => `/api/v1/offers/${uuid}`,
  CREATE: "/api/v1/offers",
  MY_OFFERS: "/api/v1/me/offers",
  PURCHASE: (uuid: string) => `/api/v1/offers/${uuid}/purchases`,
} as const;

export const PURCHASE_ENDPOINTS = {
  LIST: "/api/v1/purchases",
  SALES: "/api/v1/sales",
  CONTRACT: (purchaseUuid: string) =>
    `/api/v1/purchases/${purchaseUuid}/contract`,
} as const;

export const CCEE_ENDPOINTS = {
  AGENTS: (cnpj: string) => `/api/v1/ccee/agents/${cnpj}`,
} as const;
