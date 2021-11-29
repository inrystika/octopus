import request from '@/utils/request'

export function judgeParam(params) {
  const conditions = []
  conditions.push(`pageSize=` + params.pageSize);
  conditions.push(`pageIndex=` + params.pageIndex);
  params.orderBy ? conditions.push(`orderBy=` + params.orderBy) : null;
  params.sortBy ? conditions.push(`sortBy=` + params.sortBy) : null;
  params.searchKey ? conditions.push(`searchKey=` + params.searchKey) : null;
  params.createdAtGte ? conditions.push(`createdAtGte=` + params.createdAtGte) : null;
  params.createdAtLt ? conditions.push(`createdAtLt=` + params.createdAtLt) : null;
  params.status ? conditions.push(`status=` + params.status) : null;
  params.fileStatus ? conditions.push(`fileStatus=` + params.fileStatus) : null;
  params.algorithmVersion ? conditions.push(`algorithmVersion=` + params.algorithmVersion) : null;
  return conditions
}

export async function getNotebookList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/developmanage/notebook?" + conditions.join("&"),
    method: "get"
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

export async function stopNotebook(id) {
  const res = await request({
    url: `/v1/developmanage/notebook/${id}/stop`,
    method: "post"
  })
  return res
}

export async function getUserAlgorithmList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: `/v1/algorithmmanage/allalgorithm?` + conditions.join("&"),
    method: "get"
  })
  return res
}

export async function getPresetAlgorithmList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/algorithmmanage/prealgorithm?" + conditions.join("&"),
    method: "get"
  })
  return res
}

export async function getAlgorithmVersionList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}?` + conditions.join("&"),
    method: "get"
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

export async function addPreAlgorithm(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm`,
    method: "post",
    data: payload
  })
  return res
}

export async function addPreAlgorithmVersion(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm/${payload.algorithmId}`,
    method: "post",
    data: {
      oriVersion: payload.oriVersion,
      algorithmDescript: payload.algorithmDescript
    }
  })
  return res
}

export async function uploadPreAlgorithm(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm/${payload.algorithmId}/version/${payload.version}/upload`,
    method: "post",
    data: payload
  })
  return res
}

export async function preAlgorithmFinishUpload(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm/${payload.algorithmId}/version/${payload.version}/uploadconfirm`,
    method: "put",
    data: payload
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

export async function downloadAlgorithmVersion(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithm/${payload.algorithmId}/version/${payload.version}/download?domain=${payload.domain}&compressAt=${payload.compressAt}`,
    method: "get"
  })
  return res
}

export async function deletePreAlgorithmVersion(payload) {
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm/${payload.algorithmId}/version/${payload.version}`,
    method: "delete"
  })
  return res
}

export async function deletePreAlgorithm(algorithmId) {
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm/${algorithmId}`,
    method: "delete"
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
export async function addAlgorithmType(data) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithmapply`,
    method: "post",
    data: { lableDesc: data }
  })
  return res
}
export async function deleteAlgorithmType(params) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithmapply/${params}`,
    method: "delete"
  })
  return res
}
export async function updateAlgorithmType(data) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithmapply/${data.id}`,
    method: "put",
    data: data
  })
  return res
}
export async function frameType(params) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithmframework`,
    method: "get",
    params: params
  })
  return res
}
export async function addFrameType(data) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithmframework`,
    method: "post",
    data: { lableDesc: data }
  })
  return res
}
export async function deleteFrameType(params) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithmframework/${params}`,
    method: "delete"
  })
  return res
}
export async function updateFrameType(data) {
  const res = await request({
    url: `/v1/algorithmmanage/algorithmframework/${data.id}`,
    method: "put",
    data: data
  })
  return res
}
// 修改我的算法
export async function editeAlgorithm(params) {
  const res = await request({
    url: `/v1/algorithmmanage/prealgorithm/${params.algorithmId}`,
    method: "put",
    params: { applyId: params.applyId, frameworkId: params.frameworkId, algorithmDescript: params.algorithmDescript }
  })
  return res
}
