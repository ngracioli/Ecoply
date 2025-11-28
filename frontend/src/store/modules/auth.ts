import type { ActionContext, Module } from "vuex";
import api from "../../axios";
import { AUTH_ENDPOINTS } from "../../api/endpoints";
import type { LoginRequest, RegisterRequest } from "../../types/requests/auth";
import type { AuthResponse } from "../../types/responses/auth";

export interface AuthState {
  token: string | null;
}

type AuthActionContext = ActionContext<AuthState, unknown>;

const auth: Module<AuthState, unknown> = {
  namespaced: true,

  actions: {
    async login({ commit }: AuthActionContext, credentials: LoginRequest) {
      try {
        const response = await api.post<AuthResponse>(
          AUTH_ENDPOINTS.LOGIN,
          credentials,
        );
        const { token, user } = response.data.data;

        commit("user/setToken", token, { root: true });
        commit("user/setUser", user, { root: true });
      } catch (error) {
        throw error;
      }
    },

    async register({ commit }: AuthActionContext, formData: RegisterRequest) {
      try {
        const response = await api.post<AuthResponse>(
          AUTH_ENDPOINTS.SIGNUP,
          formData,
        );
        const { token, user } = response.data.data;

        commit("user/setToken", token, { root: true });
        commit("user/setUser", user, { root: true });
      } catch (error) {
        throw error;
      }
    },

    logout({ commit }: AuthActionContext) {
      commit("logout");
      commit("user/clearUser", null, { root: true });
    },
  },
};

export default auth;
