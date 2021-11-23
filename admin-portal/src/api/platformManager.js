import request from '@/utils/request'

export function judgeParam(params) {
  let conditions = []
  conditions.push(`pageSize=`+params.pageSize);
  conditions.push(`pageIndex=`+params.pageIndex);
  params.orderBy?conditions.push(`orderBy=`+params.orderBy):null;
  params.sortBy?conditions.push(`sortBy=`+params.sortBy):null;
  params.searchKey?conditions.push(`searchKey=`+params.searchKey):null;
  params.createdAtGte?conditions.push(`createdAtGte=`+params.createdAtGte):null;
  params.createdAtLt?conditions.push(`createdAtLt=`+params.createdAtLt):null;
  params.name?conditions.push(`name=`+params.name):null;
  return conditions
}

export function getPlatformList(params) {
  let conditions = judgeParam(params)
  return request({
    url: '/v1/platformmanage/platform?' + conditions.join("&"),
    method: 'get',
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

export function getStorageConfigList(params) {
  let conditions = judgeParam(params)
  return request({
    url: `/v1/platformmanage/platform/${params.id}/storageconfig?` + conditions.join("&"),
    method: 'get',
  })
}

export function createStorageConfig(platformId,params) {
  return request({
    url: `/v1/platformmanage/platform/${platformId}/storageconfig`,
    method: 'post',
    data: params
  })
}

export function deleteStorageConfig(params) {
  return request({
    url: `/v1/platformmanage/platform/${params.platformId}/storageconfig/${params.name}`,
    method: 'delete',
  })
}

export function updatePlatformConfig(platformId,params) {
  return request({
    url: `/v1/platformmanage/platform/${platformId}/config`,
    method: 'put',
    data: {
      config: params
    }
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