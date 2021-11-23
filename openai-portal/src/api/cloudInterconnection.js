import request from '@/utils/request'

export function createCloudTrainJob(params) {
  return request({
    url: `/v1/jointcloudmanage/jointcloudtrainjob`,
    method: 'post',
    data: params
  })
}

export function stopCloudTrainJob(id) {
  return request({
    url: `/v1/jointcloudmanage/trainjob/${id}/stop`,
    method: 'post',
  })
}

export function getCloudTrainJobList(params) {
  let conditions = []
  conditions.push(`pageIndex=` + params.pageIndex);
  conditions.push(`pageSize=` + params.pageSize);
  params.ids?conditions.push(`ids=` + params.ids):null;
  return request({
    url: `/v1/jointcloudmanage/job?` + conditions.join("&"),
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

export function getCloudFrameworkList() {
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
    url: `/v1/jointcloudmanage/interpreter/${key}/version`,
    method: 'get',
  })
}