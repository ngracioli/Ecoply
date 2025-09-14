import type { Module } from "vuex";
import api from "../../axios";

export interface AuthState {
  token: string | null;
}

const auth: Module<AuthState, any> = {
  namespaced: true,
  state: {
    token: localStorage.getItem("token") ?? null,
  },
  mutations: {
    setToken(state, token: string) {
      state.token = token;
      localStorage.setItem("token", state.token);
      console.log("Token set:", state.token);
    },
    logout(state) {
      state.token = null;
    },
  },
  actions: {
    async login(
      { commit },
      { email, password }: { email: string; password: string },
    ) {
      try {
        const response = await api.post("/api/v1/auth/login", {
          email,
          password,
        });
        const token = response.data.data.token;
        commit("setToken", token);
      } catch (error) {
        console.error("Login failed:", error);
        throw error;
      }
    },
  },
  getters: {
    isAuthenticated(state): boolean {
      return !!state.token;
    },
  },
};

export default auth;
