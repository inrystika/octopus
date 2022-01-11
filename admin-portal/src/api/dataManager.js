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
  params.path ? conditions.push(`path=` + params.path) : null;
  return conditions
}

export async function getUserDatasetList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/datasetmanage/userdataset?" + conditions.join("&"),
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

export async function getVersionList(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.id}/version?` + conditions.join("&"),
    method: "get"
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

export async function createPreDataset(payload) {
  const res = await request({
    url: "/v1/datasetmanage/dataset",
    method: "post",
    data: payload
  })
  return res
}

export async function uploadPreDataset(payload) {
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

export async function preDatasetFinishUpload(payload) {
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
    url: `/v1/datasetmanage/dataset/${payload.datasetId}/version/${payload.version}/upload`,
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
    url: `/v1/datasetmanage/dataset/${payload.datasetId}/version/${payload.version}/uploadconfirm`,
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
export async function addDatasetType(data) {
  const res = await request({
    url: `/v1/datasetmanage/datasettype`,
    method: "post",
    data: { lableDesc: data }
  })
  return res
}
export async function deleteDatasetType(params) {
  const res = await request({
    url: `/v1/datasetmanage/datasettype/${params}`,
    method: "delete"
  })
  return res
}
export async function updateDatasetType(data) {
  const res = await request({
    url: `/v1/datasetmanage/datasettype/${data.id}`,
    method: "put",
    data: data
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
export async function addDatasetUse(data) {
  const res = await request({
    url: `/v1/datasetmanage/datasetapply`,
    method: "post",
    data: { lableDesc: data }
  })
  return res
}
export async function deleteDatasetUse(params) {
  const res = await request({
    url: `/v1/datasetmanage/datasetapply/${params}`,
    method: "delete"
  })
  return res
}
export async function updateDatasetUse(data) {
  const res = await request({
    url: `/v1/datasetmanage/datasetapply/${data.id}`,
    method: "put",
    data: data
  })
  return res
}
// 修改数据集
export async function editeDataSet(data) {
  const res = await request({
    url: `/v1/datasetmanage/predataset/${data.datasetId}`,
    method: "put",
    params: { typeId: data.typeId, applyIds: data.applyIds, desc: data.desc }
  })
  return res
}
