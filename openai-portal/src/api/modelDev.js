import request from '@/utils/request'

export function judgeParam(params) {
  const conditions = []
  params.pageSize ? conditions.push(`pageSize=` + params.pageSize) : null;
  params.pageIndex ? conditions.push(`pageIndex=` + params.pageIndex) : null;
  params.orderBy ? conditions.push(`orderBy=` + params.orderBy) : null;
  params.sortBy ? conditions.push(`sortBy=` + params.sortBy) : null;
  params.searchKey ? conditions.push(`searchKey=` + params.searchKey) : null;
  params.createdAtGte ? conditions.push(`createdAtGte=` + params.createdAtGte) : null;
  params.createdAtLt ? conditions.push(`createdAtLt=` + params.createdAtLt) : null;
  params.status ? conditions.push(`status=` + params.status) : null;
  params.fileStatus ? conditions.push(`fileStatus=` + params.fileStatus) : null;
  params.algorithmVersion ? conditions.push(`algorithmVersion=` + params.algorithmVersion) : null;
  params.nameLike ? conditions.push(`nameLike=` + params.nameLike) : null;
  return conditions
}

export async function getNotebookList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/developmanage/notebook?" + conditions.join("&"),
    method: 'get'
  })
  return res
}

export async function getNotebookInfo(params) {
  const res = await request({
    url: `/v1/developmanage/notebookevent`,
    method: 'get',
    params
  })
  return res
}

export async function getNotebookEventRecord(params) {
  const res = await request({
    url: `/v1/developmanage/notebook/${params.id}/eventrecord?pageSize=${params.pageSize}&pageIndex=${params.pageIndex}`,
    method: 'get',
  })
  return res
}

export async function createNotebook(payload) {
  const res = await request({
    url: `/v1/developmanage/notebook`,
    method: "post",
    data: payload
  })
  return res
}

export async function stopNotebook(id) {
  const res = await request({
    url: `/v1/developmanage/notebook/${id}/stop`,
    method: "post"
  })
  return res
}

export async function deleteNotebook(id) {
  const res = await request({
    url: `/v1/developmanage/notebook/${id}`,
    method: "delete"
  })
  return res
}

export async function startNotebook(id) {
  const res = await request({
    url: `/v1/developmanage/notebook/${id}/start`,
    method: "post"
  })
  return res
}

export async function getAlgorithmVersionList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}?` + conditions.join("&"),
    method: 'get'
  })
  return res
}

export async function queryAlgorithmVersion(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}/version/${payload.version}`,
    method: 'get'
  })
  return res
}

export async function getPubAlgorithmVersionList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: `/v1/algorithmmanage/commalgorithm/${payload.algorithmId}?` + conditions.join("&"),
    method: 'get'
  })
  return res
}

export async function getMyAlgorithmList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: `/v1/algorithmmanage/myalgorithm?` + conditions.join("&"),
    method: 'get'
  })
  return res
}

export async function getPublicAlgorithmList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: `/v1/algorithmmanage/commalgorithm?` + conditions.join("&"),
    method: 'get'
  })
  return res
}

export async function getPresetAlgorithmList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm?` + conditions.join("&"),
    method: 'get'
  })
  return res
}

export async function copyAlgorithm(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}/version/${payload.version}/copy`,
    method: "post",
    data: {
      algorithmDescript: payload.algorithmDescript,
      newAlgorithmName: payload.newAlgorithmName,
      modelName: payload.modelName
    }
  })
  return res
}

export async function compressAlgorithm(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}/version/${payload.version}/downloadcompress`,
    method: "post"
  })
  return res
}

export async function downloadAlgorithm(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}/version/${payload.version}/download?compressAt=${payload.compressAt}&domain=${payload.domain}`,
    method: "get"
  })
  return res
}

export async function shareAlgorithmVersion(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}/version/${payload.version}/share`,
    method: "post",
    data: payload
  })
  return res
}

export async function cancelShareAlgorithmVersion(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}/version/${payload.version}/closeshare`,
    method: "post",
    data: payload
  })
  return res
}

export async function deleteAlgorithmVersion(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/myalgorithm/${payload.algorithmId}/version/${payload.algorithmVersion}`,
    method: "delete"
  })
  return res
}

export async function deleteMyAlgorithm(algorithmId) {
  const res = await request({
    url: `/v1/algorithmmanage/myalgorithm/${algorithmId}`,
    method: "delete"
  })
  return res
}

export async function addMyAlgorithm(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/myalgorithm`,
    method: "post",
    data: payload
  })
  return res
}

export async function createNewAlgorithmVersion(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/myalgorithm/${payload.algorithmId}`,
    method: "post",
    data: payload
  })
  return res
}

export async function uploadMyAlgorithm(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}/version/${payload.version}/upload`,
    method: "post",
    data: payload
  })
  return res
}

export async function myAlgorithmFinishUpload(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}/version/${payload.version}/uploadconfirm`,
    method: "put",
    data: payload
  })
  return res
}
export async function algorithmType(params) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithmapply`,
    method: "get",
    params: params
  })
  return res
}
export async function algorithmFrame(params) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithmframework`,
    method: "get",
    params: params
  })
  return res
}
// 修改我的算法
export async function editeAlgorithm(params) {
  const res = await request({
    url: `/v1/algorithmmanage/myalgorithm/${params.algorithmId}`,
    method: "put",
    params: { applyId: params.applyId, frameworkId: params.frameworkId, algorithmDescript: params.algorithmDescript }
  })
  return res
}