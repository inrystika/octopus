import request from '@/utils/request'
// 获取已配置的第三方登录平台接口
export function themeChange() {
  return request({
    url: `/v1/systemmanage/webconfig`,
    method: 'get'
  })
}
// 2、跳转到第三方登录界面的接口
export function getInterface(params) {
  return request({
    url: `/v1/oauth2/${params}/authorize`,
    method: 'get'
  })
}
// 4、注册并绑定接口
export function register(data) {
  return request({
    url: `/v1/authmanage/registerandbind`,
    method: 'post',
    data
  })
}
