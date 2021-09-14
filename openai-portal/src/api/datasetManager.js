import request from '@/utils/request'

export function judgeParam(params) {
  let conditions = []
  params.pageSize?conditions.push(`pageSize=`+params.pageSize):null;
  params.pageIndex?conditions.push(`pageIndex=`+params.pageIndex):null;
  params.orderBy?conditions.push(`orderBy=`+params.orderBy):null;
  params.sortBy?conditions.push(`sortBy=`+params.sortBy):null;
  params.searchKey?conditions.push(`searchKey=`+params.searchKey):null;
  params.createdAtGte?conditions.push(`createdAtGte=`+params.createdAtGte):null;
  params.createdAtLt?conditions.push(`createdAtLt=`+params.createdAtLt):null;
  params.shared?conditions.push(`shared=`+params.shared):null;
  params.path?conditions.push(`path=`+params.path):null;
  params.status?conditions.push(`status=`+params.status):null;
  params.nameLike?conditions.push(`nameLike=`+params.nameLike):null;
  return conditions
}

export async function getMyDatasetList(payload){
  let conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/datasetmanage/mydataset?" + conditions.join("&"),
    method: "get",
  })
  return res
}

export async function getPublicDatasetList(payload){
  let conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/datasetmanage/commdataset?" + conditions.join("&"),
    method: "get",
  })
  return res
}

export async function getPresetDatasetList(payload){
  let conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/datasetmanage/predataset?" + conditions.join("&"),
    method: "get",
  })
  return res
}

export async function getVersionList(data){
  let conditions = judgeParam(data)
  const res = await request({
    url: `/v1/datasetmanage/dataset/${data.datasetId}/version?` + conditions.join("&"),
    method: "get",
  })
  return res
}

export async function shareDatasetVersion(payload){
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.datasetId}/version/${payload.version}/share`,
    method: "post",
  })
  return res
}

export async function cancelShareDatasetVersion(payload){
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.datasetId}/version/${payload.version}/closeshare`,
    method: "post",
  })
  return res
}

export async function deleteDatasetVersion(payload){
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.datasetId}/version/${payload.version}`,
    method: "delete",
  })
  return res
}

export async function deleteDataset(id){
  const res = await request({
    url: `/v1/datasetmanage/dataset/${id}`,
    method: "delete",
  })
  return res
}

export async function createMyDataset(payload){
  const res = await request({
    url: `/v1/datasetmanage/dataset`,
    method: "post",
    data: payload,
  })
  return res
}

export async function uploadMyDataset(payload){
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

export async function myDatasetFinishUpload(payload){
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.id}/version/${payload.version}/uploadconfirm`,
    method: "put",
    data: {
      fileName: payload.fileName
    }
  })
  return res
}

export async function createNewVersion(payload){
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.datasetId}/version`,
    method: "post",
    data:{
      desc: payload.desc
    }
  })
  return res
}

export async function uploadNewVersion(payload){
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

export async function newVersionFinishUpload(payload){
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.id}/version/${payload.version}/uploadconfirm`,
    method: "put",
    data:{
      fileName: payload.fileName
    }
  })
  return res
}

export async function previewDataset(payload){
  let conditions = judgeParam(payload)
  const res = await request({
    url: `/v1/datasetmanage/dataset/${payload.datasetId}/version/${payload.version}/file?`+ conditions.join("&"),
    method: "get",
  })
  return res
}
