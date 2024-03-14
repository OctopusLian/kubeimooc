import {createPod, getPodItemOrList, deletePod} from "@/api/pod"

const state = {
  podName: ""
}

const mutations = {
  SET_POD_NAME: (state, podName) => {
    state.podName = podName
  }
}

const actions = {
  createPod({commit, state}, data) {
    return new Promise((resolve, reject) =>
      createPod(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    )
  },
  //查看Pod详情或列表
  getPodItemOrList({commit, state}, params) {
    return new Promise((resolve, reject) =>
      getPodItemOrList(params).then(res => {
        resolve(res)
      }).catch(err => {
        reject(err)
      })
    )
  },
  //删除Pod
  deletePod({commit, state}, params) {
    return new Promise((resolve, reject) =>
      deletePod(params).then(res => {
        resolve(res)
      }).catch(err => {
        reject(err)
      })
    )
  },
  setPodName({commit, state}, name) {
    commit("SET_POD_NAME", name)
  },
  getPodName({commit, state}) {
    return state.podName
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

