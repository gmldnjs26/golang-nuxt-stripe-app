import { createStore, Commit } from "vuex";
import { User } from "@/types/user";

export default createStore({
  state: {
    user: null,
  },
  mutations: {
    SET_USER(state, payload) {
      state.user = payload;
    },
  },
  actions: {
    set_user({ commit }: { commit: Commit }, payload: User) {
      commit("SET_USER", payload);
    },
  },
  modules: {},
});
