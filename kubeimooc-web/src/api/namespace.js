import request from '@/utils/request'

export function getNamespaceList() {
  return request({
    url: '/k8s/namespace',
    method: 'get'
  })
}
