import request from '@/utils/request'
import requestMinIO from '@/utils/requestMinIO'
export function judgeParam(params) {
  const conditions = []
  params.imageType ? conditions.push(`imageType=` + params.imageType) : null;
  params.imageStatus ? conditions.push(`imageStatus=` + params.imageStatus) : null;
  params.orderBy ? conditions.push(`orderBy=` + params.orderBy) : null;
  params.sortBy ? conditions.push(`sortBy=` + params.sortBy) : null;
  conditions.push(`pageSize=`+params.pageSize);
  conditions.push(`pageIndex=`+params.pageIndex);  params.imageAddrLike ? conditions.push(`imageAddrLike=` + params.imageAddrLike) : null;
  params.imageNameLike ? conditions.push(`imageNameLike=` + params.imageNameLike) : null;
  params.userId ? conditions.push(`userId=` + params.userId) : null;
  params.spaceId ? conditions.push(`spaceId=` + params.spaceId) : null;
  params.sourceType ? conditions.push(`sourceType=` + params.sourceType) : null;
  params.imageVersion ? conditions.push(`imageVersion=` + params.imageVersion) : null;
  params.searchKey ? conditions.push(`searchKey=` + params.searchKey) : null;
  params.nameVerLike ? conditions.push(`nameVerLike=` + params.nameVerLike) : null;
  return conditions
}
// 我的镜像列表
export function getMyImage(params) {
  const conditions = judgeParam(params)
  return request({
    url: '/v1/imagemanage/userimage?' + conditions.join("&"),
    method: 'get'
  })
}
// 公共镜像列表
export function getPublicImage(params) {
  const conditions = judgeParam(params)
  return request({
    url: '/v1/imagemanage/commimage?' + conditions.join("&"),
    method: 'get'
  })
}
// 预置镜像列表
export function getPreImage(params) {
  const conditions = judgeParam(params)
  return request({
    url: '/v1/imagemanage/preimage?' + conditions.join("&"),
    method: 'get'
  })
}

// 创建镜像
export function createImage(data) {
  return request({
    url: '/v1/imagemanage/image',
    method: 'post',
    data
  })
}
// 上传我的镜像
export function uploadImage(data) {
  return request({
    url: `/v1/imagemanage/image/${data.id}/upload`,
    method: 'post',
    data: { id: data.id, fileName: data.fileName, domain: data.domain }
  })
}
// 删除我的镜像
export function deleteImage(params) {
  return request({
    url: `/v1/imagemanage/image/${params}`,
    method: 'delete'
  })
}
// 编辑我的镜像
export function editImage(data) {
  return request({
    url: `/v1/imagemanage/image/${data.id}`,
    method: 'put',
    data
  })
}
// 分享我的镜像
export function shareImage(params) {
  return request({
    url: `/v1/imagemanage/image/${params}/share`,
    method: 'post'
  })
}
// 取消分享我的镜像
export function cancelImage(params) {
  return request({
    url: `/v1/imagemanage/image/${params}/closeshare`,
    method: 'post'
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
    url: `/v1/imagemanage/image/${params}/uploadconfirm`,
    method: 'put'
  })
}

