import type { InternalAxiosRequestConfig } from "axios";
import { AUTH_TOKEN_KEY } from "../config";

export const requestInterceptor = (
  config: InternalAxiosRequestConfig,
): InternalAxiosRequestConfig => {
  const token = localStorage.getItem(AUTH_TOKEN_KEY);

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
};

export const requestErrorInterceptor = (error: unknown): Promise<never> => {
  return Promise.reject(error);
};
