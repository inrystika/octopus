import request from '@/utils/request'

export function getPlatformList(params) {
  return request({
    url: '/v1/platformmanage/platform',
    method: 'get',
    params
  })
}

export function createPlatform(params) {
  return request({
    url: '/v1/platformmanage/platform',
    method: 'post',
    params
  })
}

export function updatePlatform(params) {
  return request({
    url: '/v1/platformmanage/platform',
    method: 'put',
    params
  })
}

export function getPlatformConfigKey() {
  return request({
    url: `/v1/platformmanage/platformconfigkey`,
    method: 'get'
  })
}

export function getPlatformConfigValue(platformId) {
  return request({
    url: `/v1/platformmanage/platform/${platformId}/config`,
    method: 'get'
  })
}