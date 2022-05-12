import Vue from "vue";
import Vuex, { Commit } from "vuex";

Vue.use(Vuex);

import { User } from "@/types/user";

export default new Vuex.Store({
  state: {
    user: {} as User,
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
