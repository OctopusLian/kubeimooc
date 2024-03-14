import {createConfigMap, getCmItemOrList,deleteCm} from "@/api/configmap"

const state = {
  podName: ""
}

const mutations = {
  SET_POD_NAME: (state, podName) => {
    state.podName = podName
  }
}

const actions = {
  createConfigMap({commit, state}, params) {
    return new Promise((resolve, reject) =>
      createConfigMap(params).then(res => {
        resolve(res)
      }).catch(err => {
        reject(err)
      })
    )
  },
  //删除cm
  deleteCm({commit, state}, params) {
    return new Promise((resolve, reject) =>
      deleteCm(params).then(res => {
        resolve(res)
      }).catch(err => {
        reject(err)
      })
    )
  },
  //查看Pod详情或列表
  getCmItemOrList({commit, state}, params) {
    return new Promise((resolve, reject) =>
      getCmItemOrList(params).then(res => {
        resolve(res)
      }).catch(err => {
        reject(err)
      })
    )
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

