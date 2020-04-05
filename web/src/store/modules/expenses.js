import { createExpense, listExpenses, deleteExpense } from "@/api/expenses";

const state = {
  expense: {},
  expenses: [],
};

const getters = {}

const actions = {
  createExpense({ commit }, payload) {
    return new Promise(resolve => {
      createExpense(payload).then(
        item => {
          commit('setExpense', item);
          resolve(item);
        });
    });
  },

  listExpenses({ commit }) {
    listExpenses().then(items => commit('setExpenses', items));
  },

  deleteExpense(context, id) {
    return deleteExpense(id);
  }
}

const mutations = {
  setExpense(state, payload) {
    state.expense = payload;
  },
  setExpenses(state, payload) {
    state.expenses = payload;
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
