import axios from "axios";
import { API_BASE_URL } from "./api/config";
import {
  requestErrorInterceptor,
  requestInterceptor,
} from "./api/interceptors/request";
import {
  responseErrorInterceptor,
  responseInterceptor,
} from "./api/interceptors/response";

const api = axios.create({
  baseURL: API_BASE_URL,
});

api.interceptors.request.use(requestInterceptor, requestErrorInterceptor);
api.interceptors.response.use(responseInterceptor, responseErrorInterceptor);

export default api;
