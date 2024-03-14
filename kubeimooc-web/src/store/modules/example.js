import { examplePing } from '@/api/example'

const state = {}

const mutations = {}

const actions = {
  // get user info
  examplePing({ commit, state }) {
    return new Promise((resolve, reject) => {
      examplePing(state.token).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

