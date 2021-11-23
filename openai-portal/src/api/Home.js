import request from '@/utils/request'
export function login(data) {
  return request({
    url: '/v1/authmanage/token',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/v1/usermanage/user',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/v1/authmanage/token',
    method: 'delete'
  })
}
export function getSpace(params) {
  return request({
    url: `/v1/usermanage/user/${params}/workspace`,
    method: 'get'
  })
}
export function changeSpace(data) {
  return request({
    url: `/v1/usermanage/user/${data.userId}/workspace`,
    method: 'put',
    data: { workspaceId: data.workspaceId }
  })
}
// 查询用户配置
export function getUserConfig() {
  return request({
    url: `/v1/usermanage/config`,
    method: 'get',
  })
}

