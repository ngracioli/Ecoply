import { createStore } from "vuex";
import auth from "./modules/auth";
import user from "./modules/user";

export type { RootState } from "./types";

export default createStore({
  modules: {
    auth,
    user,
  },
});
