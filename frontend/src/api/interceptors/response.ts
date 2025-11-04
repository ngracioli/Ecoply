import type { AxiosError, AxiosResponse } from "axios";
import router from "../../router";
import { AUTH_TOKEN_KEY } from "../config";

const HTTP_UNAUTHORIZED = 401;

export const responseInterceptor = (response: AxiosResponse): AxiosResponse => {
  return response;
};

export const responseErrorInterceptor = (error: AxiosError): Promise<never> => {
  if (error.response?.status === HTTP_UNAUTHORIZED) {
    localStorage.removeItem(AUTH_TOKEN_KEY);
    router.push({ name: "Login" });
  }

  return Promise.reject(error);
};
