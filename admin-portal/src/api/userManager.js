import request from '@/utils/request'

// 获取用户列表
export function getUserList(params) {
  return request({
    url: '/v1/usermanage/user',
    method: 'get',
    params
  })
}
// 获取用户群组
export function getUserGroup(params) {
  return request({
    url: `/v1/usermanage/user/${params}/workspace`,
    method: 'get'
  })
}
// 新建用户
export function createUser(data) {
  return request({
    url: '/v1/usermanage/user',
    method: 'post',
    data
  })
}
// 修改用户信息
export function editeUser(data) {
  return request({
    url: `/v1/usermanage/user/${data.id}`,
    method: 'put',
    data: { fullname: data.fullname, password: data.password }
  })
}
// 冻结账号
export function freeze(params) {
  return request({
    url: `/v1/usermanage/user/${params}/freeze`,
    method: 'post'
  })
}
// 激活账号
export function activation(params) {
  return request({
    url: `/v1/usermanage/user/${params}/thaw`,
    method: 'post'
  })
}
// 用户详细
export function userDetail(params) {
  return request({
    url: `/v1/usermanage/user/${params}`,
    method: 'get'
  })
}
// 创建群组
export function createGroup(data) {
  return request({
    url: `/v1/usermanage/workspace`,
    method: 'post',
    data
  })
}
// 编辑群组
export function editeGroup(data) {
  return request({
    url: `/v1/usermanage/workspace/${data.id}`,
    method: 'put',
    data: { name: data.name, userIds: data.userIds, resourcePoolId: data.resourcePoolId }
  })
}
// 群组列表
export function groupList(params) {
  return request({
    url: '/v1/usermanage/workspace',
    method: 'get',
    params
  })
}
// 群组详细
export function groupDetail(params) {
  return request({
    url: `/v1/usermanage/workspace/${params}`,
    method: 'get'
  })
}
// 删除群组
export function deleteGroup(params) {
  return request({
    url: `/v1/usermanage/workspace/${params}`,
    method: 'delete'
  })
}
// 查询配置key列表
export function getUserConfigKey() {
  return request({
    url: `/v1/usermanage/userconfigkey`,
    method: 'get'
  })
}
// 查询用户配置
export function getUserConfig(userId) {
  return request({
    url: `/v1/usermanage/user/${userId}/config`,
    method: 'get'
  })
}
// 更新用户配置
export function updateUserConfig(userId,params) {
  return request({
    url: `/v1/usermanage/user/${userId}/config`,
    method: 'put',
    params
  })
}

