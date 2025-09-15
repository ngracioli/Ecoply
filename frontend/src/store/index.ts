import { createStore } from "vuex";

// Importe seus módulos aqui
import auth from "./modules/auth";

export default createStore({
  modules: {
    // Adicione seus módulos importados aqui
    auth,
  },
});
