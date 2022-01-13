import request from '@/utils/request'

export function judgeParam(params) {
  let conditions = []
  conditions.push(`pageSize=` + params.pageSize);
  conditions.push(`pageIndex=` + params.pageIndex);
  params.orderBy ? conditions.push(`orderBy=` + params.orderBy) : null;
  params.sortBy ? conditions.push(`sortBy=` + params.sortBy) : null;
  params.searchKey ? conditions.push(`searchKey=` + params.searchKey) : null;
  params.createdAtGte ? conditions.push(`createdAtGte=` + params.createdAtGte) : null;
  params.createdAtLt ? conditions.push(`createdAtLt=` + params.createdAtLt) : null;
  params.shared ? conditions.push(`shared=` + params.shared) : null;
  params.path ? conditions.push(`path=` + params.path) : null;
  params.status ? conditions.push(`status=` + params.status) : null;
  return conditions
}

export async function getMyDatasetList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/datasetmanage/mydataset?" + conditions.join("&"),
    method: "get"
  })
  return res
}

export async function getPublicDatasetList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/datasetmanage/commdataset?" + conditions.join("&"),
    method: "get"
  })
  return res
}

export async function getPresetDatasetList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/datasetmanage/predataset?" + conditions.join("&"),
    method: "get"
  })
  return res
}

export async function getVersionList(data) {
  const conditions = judgeParam(data)
  const res = await request({
    url: `/v1/datasetmanage/dataset/${data.datasetId}/version?` + conditions.join("&"),
    method: "get"
  })
  return res
}

export async function shareDatasetVersion(payload) {
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.datasetId}/version/${payload.version}/share`,
    method: "post"
  })
  return res
}

export async function cancelShareDatasetVersion(payload) {
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.datasetId}/version/${payload.version}/closeshare`,
    method: "post"
  })
  return res
}

export async function deleteDatasetVersion(payload) {
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.datasetId}/version/${payload.version}`,
    method: "delete"
  })
  return res
}

export async function deleteDataset(id) {
  const res = await request({
    url: `/v1/datasetmanage/dataset/${id}`,
    method: "delete"
  })
  return res
}

export async function createMyDataset(payload) {
  const res = await request({
    url: `/v1/datasetmanage/dataset`,
    method: "post",
    data: payload
  })
  return res
}

export async function uploadMyDataset(payload) {
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.id}/version/${payload.version}/upload`,
    method: "post",
    data: {
      fileName: payload.fileName,
      domain: payload.domain
    }
  })
  return res
}

export async function myDatasetFinishUpload(payload) {
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.id}/version/${payload.version}/uploadconfirm`,
    method: "put",
    data: {
      fileName: payload.fileName
    }
  })
  return res
}

export async function createNewVersion(payload) {
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.datasetId}/version`,
    method: "post",
    data: {
      desc: payload.desc
    }
  })
  return res
}

export async function uploadNewVersion(payload) {
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.id}/version/${payload.version}/upload`,
    method: "post",
    data: {
      fileName: payload.fileName,
      domain: payload.domain
    }
  })
  return res
}

export async function newVersionFinishUpload(payload) {
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.id}/version/${payload.version}/uploadconfirm`,
    method: "put",
    data: {
      fileName: payload.fileName
    }
  })
  return res
}

export async function previewDataset(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.datasetId}/version/${payload.version}/file?` + conditions.join("&"),
    method: "get"
  })
  return res
}

export async function datasetType(params) {
  const res = await request({
    url: `/v1/datasetmanage/datasettype`,
    method: "get",
    params: params
  })
  return res
}
export async function datasetUse(params) {
  const res = await request({
    url: `/v1/datasetmanage/datasetapply`,
    method: "get",
    params: params
  })
  return res
}
// 修改数据集
export async function editDataSet(data) {
  const res = await request({
    url: `/v1/datasetmanage/mydataset/${data.datasetId}`,
    method: "put",
    data: { typeId: data.typeId, applyIds: data.applyIds, desc: data.desc }
  })
  return res
}
