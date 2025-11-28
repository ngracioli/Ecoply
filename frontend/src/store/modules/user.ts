import type { Module } from "vuex";
import api from "../../axios";
import { AUTH_ENDPOINTS } from "../../api/endpoints";
import type { User } from "../../types/user";
import type { UserMeResponse } from "../../types/responses/auth";

const TOKEN_STORAGE_KEY = "token";

export interface UserState {
  token: string | null;
  currentUser: User | null;
}

const user: Module<UserState, unknown> = {
  namespaced: true,

  state: {
    token: localStorage.getItem(TOKEN_STORAGE_KEY) ?? null,
    currentUser: null,
  },

  mutations: {
    setToken(state, token: string) {
      state.token = token;
      localStorage.setItem(TOKEN_STORAGE_KEY, token);
    },

    setUser(state, user: User) {
      state.currentUser = user;
    },

    clearUser(state) {
      state.currentUser = null;
      state.token = null;
      localStorage.removeItem(TOKEN_STORAGE_KEY);
    },
  },

  actions: {
    async fetchCurrentUser({ commit, state }) {
      if (!state.token) {
        throw new Error("Token não encontrado");
      }

      try {
        const response = await api.get<UserMeResponse>(AUTH_ENDPOINTS.ME);
        commit("setUser", response.data.data);
      } catch (error) {
        throw error;
      }
    },

    updateUser({ commit }, user: User) {
      commit("setUser", user);
    },

    clearUser({ commit }) {
      commit("clearUser");
    },
  },

  getters: {
    currentUser(state): User | null {
      return state.currentUser;
    },

    userName(state): string | null {
      return state.currentUser?.name ?? null;
    },

    userType(state): string | null {
      return state.currentUser?.user_type ?? null;
    },

    isBuyer(state): boolean {
      return state.currentUser?.user_type === "buyer";
    },

    isSupplier(state): boolean {
      return state.currentUser?.user_type === "supplier";
    },

    isAuthenticated(state): boolean {
      return !!state.token;
    },

    userToken(state): string | null {
      return state.token;
    },

    userSubmarket(state): string | null {
      return state.currentUser?.agent?.submarket_name ?? null;
    },
  },
};

export default user;
