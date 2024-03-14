import {getNodeItemOrList, updateNodeLabel, updateNodeTaint} from "@/api/node"

const state = {
  nodeName: ""
}

const mutations = {
  SET_NODE_NAME: (state, podName) => {
    state.nodeName = podName
  }
}

const actions = {
  setNodeName({commit, state}, name) {
    commit("SET_NODE_NAME", name)
  },
  //查看Pod详情或列表
  getNodeItemOrList({commit, state}, params) {
    return new Promise((resolve, reject) =>
      getNodeItemOrList(params).then(res => {
        resolve(res)
      }).catch(err => {
        reject(err)
      })
    )
  },
  updateNodeLabel({commit, state}, data) {
    return new Promise((resolve, reject) =>
      updateNodeLabel(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    )
  },
  updateNodeTaint({commit, state}, data) {
    return new Promise((resolve, reject) =>
      updateNodeTaint(data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    )
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

