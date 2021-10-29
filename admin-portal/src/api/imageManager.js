import request from '@/utils/request'
import requestMinIO from '@/utils/requestMinIO'
// 用户镜像列表
export function getUserImage(params) {
  return request({
    url: '/v1/imagemanage/userimage',
    method: 'get',
    params
  })
}
// 预置镜像列表
export function getPreImage(params) {
  return request({
    url: '/v1/imagemanage/preimage',
    method: 'get',
    params
  })
}
// 创建预置镜像
export function createPreImage(data) {
  return request({
    url: '/v1/imagemanage/preimage',
    method: 'post',
    data

  })
}
// 编辑预置镜像
export function editePreImage(data) {
  return request({
    url: `/v1/imagemanage/preimage/${data.id}`,
    method: 'put',
    data
  })
}
// 上传预置镜像
export function uploadPreImage(data) {
  return request({
    url: `/v1/imagemanage/preimage/${data.id}/upload`,
    method: 'post',
    data: { id: data.id, fileName: data.fileName, domain: data.domain }

  })
}
// 上传miniIO
export function uploadMiniIO(params) {
  return requestMinIO({
    url: params.uploadUrl,
    method: 'put',
    data: params.file,
    onUploadProgress: function(progress) {
      sessionStorage.setItem(JSON.stringify(params.id), JSON.stringify(parseInt(((progress.loaded / progress.total) * 100))));
    }
  })
}
// 完成镜像上传
export function finishUpload(params) {
  return request({
    url: `/v1/imagemanage/preimage/${params.id}/uploadconfirm`,
    method: 'put'
  })
}

// 删除预置镜像
export function deletePreImage(params) {
  return request({
    url: `/v1/imagemanage/preimage/${params}`,
    method: 'delete'
  })
}