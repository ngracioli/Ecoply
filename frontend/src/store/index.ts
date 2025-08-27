import { createStore } from "vuex";

export interface RootState {
  // count: number;
}

export const store = createStore<RootState>({
  state: {
    // count: 0,
  },
  mutations: {},
  actions: {},
  getters: {},
});
