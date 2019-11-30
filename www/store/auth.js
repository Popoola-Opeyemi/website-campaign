import { saveSession, getSession, removesession } from "@/utils/auth";

export const state = () => ({
  admin: null,
  isAuthenticated: getSession("isAuthenticated") || ""
});

export const mutations = {
  saveUser(state, data) {
    saveSession("isAuthenticated", true);
  }
};

export const getters = {
  isAuthenticated(state) {
    return state.isAuthenticated;
  }
};

export const actions = {
  async logout() {
    try {
      let req = await this.$axios.$get(`api/auth/logout`);
      if (req.done) {
        removesession("isAuthenticated");
        return { status: true };
      }
      return { status: false };
    } catch (error) {
      console.log("error");
      throw error;
    }
  },

  async login({ commit }, data) {
    try {
      let req = await this.$axios.$post("api/auth/login", data);
      if (req.done) {
        commit("saveUser", req.store);
        return { status: true };
      }
      return { status: false };
    } catch (error) {
      console.log("error");
      throw error;
    }
  }
};
