import Vue from 'vue'
import Vuex from 'vuex'
import * as auth from './auth'

Vue.use(Vuex)

export default new Vuex.Store({
  actions: {
    async init({ dispatch }) {
      await dispatch('auth/init')
    },
  },
  modules: { auth },
})
