import { v4 as uuid } from 'uuid'
import ncmb from '@/plugins/ncmb'

class UserProps extends ncmb.DataStore('UserProps') {}


export const state = () => ({
  currentUser: null,
  currentUserProps: null,
})

export const getters = {
  loggedIn: (state) => !!state.currentUser,
}

export const mutations = {
  setCurrentUser(state, u) {
    state.currentUser = u
  },
  setCurrentUserProps(state, u) {
    state.currentUserProps = u
  },
}


const userUuidKey = 'userUuid'

export const actions = {
  async prepare({ commit }) {
    const user = ncmb.User.getCurrentUser()
    commit('setCurrentUser', user)
    if (user === null) {
      commit('setCurrentUserProps', null)
      return
    }

    const userProps = await UserProps.equalTo('userId', user.objectId).fetch()
    commit('setCurrentUserProps', userProps.objectId
      ? userProps
      : await new UserProps()
        .set('userId', user.objectId)
        .save(),
    )
  },
  async loginAsAnonymous() {
    const id = localStorage.getItem(userUuidKey) || uuid()
    localStorage.setItem(userUuidKey, id)
    await ncmb.User.loginAsAnonymous(id)
  },
}
