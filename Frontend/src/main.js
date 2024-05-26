import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import { loader } from "@guolao/vue-monaco-editor";
loader.config({
  paths: {
    vs: "https://cdn.jsdelivr.net/npm/monaco-editor@0.43.0/min/vs",
  },
});

createApp(App).use(store).use(router).mount("#app");
