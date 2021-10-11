import request from '@/utils/request'
export function getMyModel(params) {
  return request({
    url: '/v1/modelmanage/usermodel',
    method: 'get',
    params
  })
}
export function getPreModel(params) {
  return request({
    url: '/v1/modelmanage/premodel',
    method: 'get',
    params
  })
}

// 查询模型版本列表
export function getModelList(params) {
  return request({
    url: `/v1/modelmanage/model/${params.modelId}`,
    method: 'get',
    params: { pageIndex: params.pageIndex, pageSize: params.pageSize }
  })
}
// 删除预置模型
export function deletePreModel(params) {
  return request({
    url: `/v1/modelmanage/premodel/${params.modelId}`,
    method: 'delete'
  })
}
// 删除预置模型版本
export function deletePreModelVersion(params) {
  return request({
    url: `/v1/modelmanage/premodel/${params.modelId}/version/${params.version}`,
    method: 'delete'
  })
}
// 模型版本下载
export function downloadModel(params) {
  return request({
    url: `/v1/modelmanage/model/${params.modelId}/version/${params.version}/download?domain=${params.domain}`,
    method: 'get'
  })
}
// 模型管理预览
export function preview(data) {
  return request({
    url: `/v1/modelmanage/model/${data.modelId}/version/${data.version}/file`,
    method: 'get'
  })
}
// 新增预置模型
export function addPreModel(data) {
  return request({
    url: '/v1/modelmanage/premodel',
    method: 'post',
    data
  })
}
// 新增预置模型版本列表
export function addPreList(params) {
  return request({
    url: `/v1/modelmanage/premodel/${params.modelId}`,
    method: 'post',
    data: { modelId: params.modelId, descript: params.descript }
  })
}
// 上传
export function uploadModel(data) {
  return request({
    url: `/v1/modelmanage/premodel/${data.modelId}/version/${data.version}/upload`,
    method: 'post',
    data: { fileName: data.fileName, domain: data.domain }

  })
}
// 上传确认
export function modelFinishUpload(data) {
  return request({
    url: `/v1/modelmanage/premodel/${data.modelId}/version/${data.version}/uploadconfirm`,
    method: 'put',
    data: { fileName: data.fileName }
  })
}

