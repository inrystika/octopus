import request from '@/utils/request'
export function getMyModel(params) {
  return request({
    url: '/v1/modelmanage/mymodel',
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
export function getPublicModel(params) {
  return request({
    url: '/v1/modelmanage/commmodel',
    method: 'get',
    params
  })
}
// 查询公共模型版本列表
export function getPublicList(params) {
  return request({
    url: `/v1/modelmanage/commmodel/${params.modelId}`,
    method: 'get',
    params: {
      pageIndex: params.pageIndex,
      pageSize: params.pageSize
    }
  })
}
// 查询非公共版本列表
export function getNoPublicList(params) {
  return request({
    url: `/v1/modelmanage/model/${params.modelId}`,
    method: 'get',
    params: {
      pageIndex: params.pageIndex,
      pageSize: params.pageSize
    }
  })
}
// 分享模型
export function shareModel(data) {
  return request({
    url: `/v1/modelmanage/model/${data.modelId}/version/${data.version}/share`,
    method: 'post'
  })
}
// 取消分享
export function cancelShareModel(data) {
  return request({
    url: `/v1/modelmanage/model/${data.modelId}/version/${data.version}/closeshare`,
    method: 'post'
  })
}
// 删除我的模型
export function deleteMyModel(params) {
  return request({
    url: `/v1/modelmanage/mymodel/${params.modelId}`,
    method: 'delete'
  })
}
// 删除模型版本
export function deleteModelVersion(params) {
  return request({
    url: `/v1/modelmanage/mymodel/${params.modelId}/version/${params.version}`,
    method: 'delete'
  })
}
// 模型下载
export function downloadModel(params) {
  return request({
    url: `/v1/modelmanage/model/${params.modelId}/version/${params.version}/download?domain=${params.domian}`,
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
