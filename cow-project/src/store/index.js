import { createStore } from "vuex";
export default createStore({
  state: {
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

