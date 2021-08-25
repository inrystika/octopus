import request from '@/utils/request'
// 群组列表
export function groupList(params) {
  return request({
    url: '/v1/billingmanage/workspace',
    method: 'get',
    params
  })
}
// 用户列表
export function userList(params) {
  return request({
    url: '/v1/billingmanage/user',
    method: 'get',
    params
  })
}
// 用户充值
export function userRecharge(data) {
  return request({
    url: `/v1/billingmanage/user/${data.userId}/recharge`,
    method: 'post',
    data: { amount: data.amount }
  })
}
// 群组充值
export function groupRecharge(data) {
  return request({
    url: `/v1/billingmanage/workspace/${data.spaceId}/recharge`,
    method: 'post',
    data: { amount: data.amount }
  })
}

