import request from '@/utils/request'

export function judgeParam(params) {
  let conditions = []
  conditions.push(`pageSize=`+params.pageSize);
  conditions.push(`pageIndex=`+params.pageIndex);
  params.orderBy?conditions.push(`orderBy=`+params.orderBy):null;
  params.sortBy?conditions.push(`sortBy=`+params.sortBy):null;
  params.searchKey?conditions.push(`searchKey=`+params.searchKey):null;
  return conditions
}

export async function getUserHour() {
  const res = await request({
    url: "/v1/billingmanage/user",
    method: "get"
  })
  return res
}

export async function getGroupHour() {
  const res = await request({
    url: "/v1/billingmanage/workspace",
    method: "get"
  })
  return res
}

export async function getUserConsumptionRecord(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/billingmanage/user/payrecord?" + conditions.join("&"),
    method: "get"
  })
  return res
}

export async function getGroupConsumptionRecord(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/billingmanage/workspace/payrecord?" + conditions.join("&"),
    method: "get"
  })
  return res
}

export async function getUserRechargeRecord(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/billingmanage/user/rechargerecord?" + conditions.join("&"),
    method: "get"
  })
  return res
}

export async function getGroupRechargeRecord(payload) {
  const conditions = judgeParam(payload)
  const res = await request({
    url: "/v1/billingmanage/workspace/rechargerecord?" + conditions.join("&"),
    method: "get"
  })
  return res
}
