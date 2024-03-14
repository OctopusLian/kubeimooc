import request from '@/utils/request'

export function examplePing() {
  return request({
    url: '/example/ping',
    method: 'get'
  })
}
