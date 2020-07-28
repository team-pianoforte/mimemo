import { auth, db, refKeys } from '@/firebase'
import { SignInError } from '@/errors'

export const namespaced = true

export const state = () => ({
  authUser: null,
  user: null,
  userUnsubscribe: () => {},
})

export const getters = {
  loggedIn: ({ user }) => !!user,
}

export const mutations = {
  setAuthUser(state, v) {
    state.authUser = v
  },
  setUser(state, v) {
    state.user = v
  },
  setUserUnsubscribe(state, f) {
    state.userUnsubscribe = f
  },
}

const newAuthUser = ({ uid }) => ({
  uid,
  boardIds: [],
})

export const actions = {
  async changeUser({ commit, state }, user) {
    const { uid } = user
    const ref = db.collection(refKeys.users).doc(uid)

    if (!ref.get().exists) {
      await ref.set(newAuthUser({ uid }))
    }

    const unsubscribe = ref.onSnapshot(async (doc) => {
      state.userUnsubscribe()
      commit('setUser', await doc.data())
    })

    commit('setAuthUser', user)
    commit('setUserUnsubscribe', unsubscribe)
  },
  async init({ dispatch, commit }) {
    auth.onAuthStateChanged(async (user) => {
      if (user) {
        dispatch('changeUser', user)
      } else {
        commit('setAuth', null)
        commit('setUser', null)
        commit('setUserUnsubscribe', () => {})
      }
    })
  },
  async signInAnonymously({ commit }) {
    const user = await auth.signInAnonymously().catch((err) => {
      throw new SignInError(err, 'Internal auth error')
    })
    commit('setAuthUser', user)
  },
}
