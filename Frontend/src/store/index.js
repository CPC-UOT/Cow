import { createStore } from "vuex";
export default createStore({
  state: {
    backend_url: process.env.BACKEND_URL,
    code : `#include <iostream>

int main() {
    std::cout << "Hello World!";
}
`
  },
  getters: {},
  mutations: {},
  actions: {},
  modules: {},
});

