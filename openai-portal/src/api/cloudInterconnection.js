import request from '@/utils/request'

export function createCloudTrainJob(params) {
  return request({
    url: `/v1/jointcloudmanage/jointcloudtrainjob`,
    method: 'post',
    params
  })
}

export function getCloudTrainJobList(params) {
  return request({
    url: `/v1/jointcloudmanage/job?pageIndex=${params.pageIndex}&pageSize=${params.pageSize}&ids=${params.ids}`,
    method: 'get',
  })
}

export function getCloudDatasetList(params) {
  return request({
    url: `/v1/jointcloudmanage/dataset?pageIndex=${params.pageIndex}&pageSize=${params.pageSize}`,
    method: 'get',
  })
}

export function getCloudDatasetVersionList(params) {
  return request({
    url: `/v1/jointcloudmanage/dataset/${params.dataSetCode}/version?pageIndex=${params.pageIndex}&pageSize=${params.pageSize}`,
    method: 'get',
  })
}

export function getCloudFrameworkList(params) {
  return request({
    url: `/v1/jointcloudmanage/framework`,
    method: 'get',
  })
}

export function getCloudFrameworkVersionList(key) {
  return request({
    url: `/v1/jointcloudmanage/framework/${key}/version`,
    method: 'get',
  })
}

export function getCloudInterpreterList() {
  return request({
    url: `/v1/jointcloudmanage/interpreter`,
    method: 'get',
  })
}

export function getCloudInterpreterVersionList(key) {
  return request({
    url: `/v1/jointcloudmanage/framework/${key}/version`,
    method: 'get',
  })
}