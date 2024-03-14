import {getSecretItemOrList, deleteSecret, createSecret} from "@/api/secret"
import de from "element-ui/src/locale/lang/de";

const state = {
  podName: ""
}

const mutations = {
  SET_POD_NAME: (state, podName) => {
    state.podName = podName
  }
}

const actions = {
  //删除
  deleteSecret({commit, state}, params) {
    return new Promise((resolve, reject) =>
      deleteSecret(params).then(res => {
        resolve(res)
      }).catch(err => {
        reject(err)
      })
    )
  },
  //创建
  createSecret({commit, state}, params) {
    return new Promise((resolve, reject) =>
      createSecret(params).then(res => {
        resolve(res)
      }).catch(err => {
        reject(err)
      })
    )
  },
  //查看Pod详情或列表
  getSecretItemOrList({commit, state}, params) {
    return new Promise((resolve, reject) =>
      getSecretItemOrList(params).then(res => {
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

