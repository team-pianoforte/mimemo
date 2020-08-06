import ncmb from '@/plugins/ncmb'


export class Board extends ncmb.DataStore('Board') {}
export class UserBoard extends ncmb.DataStore('UserBoard') {}

export class Memo extends ncmb.DataStore('Memo') {}


export const state = () => ({
  current: null,
  list: [],
})

export const mutations = {
  setList(state, list) {
    state.list = list
  },
  setCurrent(state, v) {
    state.current = v
  },
  append(state, v) {
    state.list.push(v)
  },
  removeIndex(state, i) {
    state.splice(i, 1)
  },
}

export const actions = {
  async fetchAll({ commit, rootState }) {
    const userId = rootState.users.currentUser.objectId
    const boards = await Board.equalTo('userId', userId).fetchAll()
    commit('setList', boards)
  },
  async setCurrent({ commit, state }, id) {
    commit('setCurrent', state.list.find((v) => v.objectId === id))
  },
  async create({ dispatch, rootState }, { name }) {
    const userId = rootState.users.currentUser.objectId
    await new Board()
      .set('name', name)
      .set('userId', userId)
      .save()
    await dispatch('fetchAll')
  },
  async delete({ dispatch }, obj) {
    await obj.delete()
    await dispatch('fetchAll')
  },
}
