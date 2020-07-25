import { auth } from '@/firebase'
import { SignInError } from '@/errors'

export const namespaced = true

export const state = () => ({
  user: null,
})

export const getters = {
  loggedIn: ({ user }) => !!user,
}

export const mutations = {
  setUser(state, user) {
    state.user = user
  },
}

export const actions = {
  async signInAnonymously({ commit }) {
    const user = await auth.signInAnonymously().catch((err) => {
      throw new SignInError(err, 'Internal auth error')
    })
    commit('setUser', user)
  },
}
