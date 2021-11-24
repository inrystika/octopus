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
    data: { amount: data.amount, title: data.title }
  })
}
// 群组充值
export function groupRecharge(data) {
  return request({
    url: `/v1/billingmanage/workspace/${data.spaceId}/recharge`,
    method: 'post',
    data: { amount: data.amount, title: data.title }
  })
}
// 查询用户充值记录列表
export function getUserRecharge(params) {
  return request({
    url: `/v1/billingmanage/userrechargerecord`,
    method: 'get',
    params
  })
}
// 查询群组充值记录
export function getGroupRecharge(params) {
  return request({
    url: `/v1/billingmanage/workspacerechargerecord`,
    method: 'get',
    params
  })
}
// 查询用户消费记录
export function getUserPay(params) {
  return request({
    url: `/v1/billingmanage/userpayrecord`,
    method: 'get',
    params
  })
}
// 查询群组消费记录
export function getGroupPay(params) {
  return request({
    url: `/v1/billingmanage/workspacepayrecord`,
    method: 'get',
    params
  })
}

