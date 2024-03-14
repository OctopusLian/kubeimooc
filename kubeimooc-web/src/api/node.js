import request from '@/utils/request'

export function updateNodeLabel(data) {
  return request({
    url: '/k8s/node/label',
    method: 'put',
    data
  })
}

export function updateNodeTaint(data) {
  return request({
    url: '/k8s/node/taint',
    method: 'put',
    data
  })
}


export function getNodeItemOrList(params) {
  let name = params.name
  let keyword = params.keyword
  let url = "/k8s/node"
  let query = []
  if (name !== undefined && name !== "") {
    query.push("nodeName=" + name)
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
