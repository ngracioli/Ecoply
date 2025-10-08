import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import PrimeVue from "primevue/config";
import Lara from "@primeuix/themes/lara";
import ToastService from "primevue/toastservice";

const app = createApp(App);

app.use(router);

app.use(store);

app.use(PrimeVue, {
  theme: {
    preset: Lara,
    options: {
      darkModeSelector: false,
    },
  },
});

app.use(ToastService);

app.mount("#app");
