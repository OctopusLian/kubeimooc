import {getNamespaceList} from "@/api/namespace";
import Cookies from 'js-cookie'

const state = {
  nsName: ""
}

const mutations = {
  SET_NS_NAME: (state, nsName) => {
    state.nsName = nsName
  },
}

const actions = {
  updateNsName({commit}, nsName) {
    commit("SET_NS_NAME", nsName)
    //写入到cookie，这样刷新页面也能加载上一个命名空间
    Cookies.set("curNs", nsName)
  },
  getCurNs({commit, state}) {
    let curNs = Cookies.get("curNs")
    if ((curNs !== "" && curNs !== undefined) && state.nsName === "") {
      commit("SET_NS_NAME", curNs)
    }
    if ((curNs === "" || curNs === undefined) && state.nsName === "") {
      commit("SET_NS_NAME", "default")
    }
    return state.nsName
  },
  //查看Pod详情或列表
  getNamespaceList({commit, state}) {
    return new Promise((resolve, reject) =>
      getNamespaceList().then(res => {
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

