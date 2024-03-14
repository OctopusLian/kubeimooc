import request from '@/utils/request'
import da from "element-ui/src/locale/lang/da";

export function createPod(data) {
  return request({
    url: '/k8s/pod',
    method: 'post',
    data
  })
}

export function getPodItemOrList(params) {
  let namespace = params.namespace
  let name = params.name
  let keyword = params.keyword
  let nodeName = params.nodeName
  let url = "/k8s/pod/" + namespace
  let query = []
  if (nodeName !== undefined && nodeName !== "") {
    query.push("nodeName=" + nodeName)
  }
  if (name !== undefined && name !== "") {
    query.push("name=" + name)
  }
  if (keyword !== undefined && keyword !== "") {
    query.push("keyword=" + keyword)
  }
  if (query.length > 0) {
    url = url + "?" + query.join("&")
  }
  return request({
    url: url,
    method: "get",
  })
}

export function deletePod(params) {
  let namespace = params.namespace
  let name = params.name
  let url = "/k8s/pod/" + namespace + "/" + name
  return request({
    url: url,
    method: "delete",
  })
}
