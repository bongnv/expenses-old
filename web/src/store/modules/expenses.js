import { createExpense } from "@/api/expenses";

const state = {
  expense: {},
};

const getters = {}

const actions = {
  createExpense({ commit }, payload) {
    createExpense(payload).then(item => commit('setExpense', item));
  }
}

const mutations = {
  setExpense(state, payload) {
    console.log("received payload", payload);
    state.expense = payload
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
