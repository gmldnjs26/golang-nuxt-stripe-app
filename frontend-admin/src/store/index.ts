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
      state.user.first_name = payload.first_name;
      state.user.last_name = payload.last_name;
      state.user.email = payload.email;
    },
  },
  actions: {
    setUser({ commit }: { commit: Commit }, payload: User) {
      commit("SET_USER", payload);
    },
  },
  modules: {},
});
