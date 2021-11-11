import request from '@/utils/request'
import requestLog from '@/utils/requestLog'
// 训练任务接口
export function getList(params) {
  let conditions = []
  conditions.push(`pageSize=`+params.pageSize);
  conditions.push(`pageIndex=`+params.pageIndex);
  params.orderBy?conditions.push(`orderBy=`+params.orderBy):null;
  params.sortBy?conditions.push(`sortBy=`+params.sortBy):null;
  params.searchKey?conditions.push(`searchKey=`+params.searchKey):null;
  params.createAtGte?conditions.push(`createdAtGte=`+params.createAtGte):null;
  params.createAtLt?conditions.push(`createdAtLt=`+params.createAtLt):null;
  params.status?conditions.push(`status=`+params.status):null;
  return request({
    url: '/v1/trainmanage/trainjob?' + conditions.join("&"),
    method: 'get',
  })
}
export function stop(id) {
  return request({
    url: `/v1/trainmanage/trainjob/${id}/stop`,
    method: 'post',

  })
}
export function Delete(data) {
  return request({
    url: '/v1/trainmanage/trainjob',
    method: 'delete',
    data

  })
}
export function createTask(data) {
  return request({
    url: '/v1/trainmanage/trainjob',
    method: 'post',
    data: data

  })
}
export function saveTemplate(data) {
  return request({
    url: '/v1/trainmanage/trainjobtemplate',
    method: 'post',
    data: data

  })
}
// 训练任务详情
export function getTraningDetail(params){
  return request({
    url: `/v1/trainmanage/trainjob/${params}`,
    method: 'get',

  })
}
// 训任务模板详情
export function getTempalteDetail(params){
  return request({
    url: `/v1/trainmanage/trainjobtemplate/${params}`,
    method: 'get',
  })
}
// 任务模板接口
export function getTemplate(params) {
  let conditions = []
  conditions.push(`pageSize=`+params.pageSize);
  conditions.push(`pageIndex=`+params.pageIndex);
  params.orderBy?conditions.push(`orderBy=`+params.orderBy):null;
  params.sortBy?conditions.push(`sortBy=`+params.sortBy):null;
  params.searchKey?conditions.push(`searchKey=`+params.searchKey):null;
  params.createAtGte?conditions.push(`createdAtGte=`+params.createAtGte):null;
  params.createAtLt?conditions.push(`createdAtLt=`+params.createAtLt):null;
  params.status?conditions.push(`status=`+params.status):null;
  return request({
    url: '/v1/trainmanage/trainjobtemplate?' + conditions.join("&"),
    method: 'get',
  })
}
//删除任务模板
export function deleteTemplate(data) {
  return request({
    url: '/v1/trainmanage/trainjobtemplate',
    method: 'delete',
    data

  })
}
// 获取资源规格列表
export function getResourceList() {
  return request({
    url: '/v1/resourcemanage/resourcespec',
    method: 'get',

  })
}
// 编辑任务模板
export function editeTemplate(params) {
  return request({
    url: `/v1/trainmanage/trainjobtemplate/${params.id}`,
    method: 'put',
    data: params.data

  })
}
// 任务日志访问
export function showLog(params) {
  return requestLog({
    url: `/log/user/trainjob/${params.jobId}/${params.subName}/index.log`,
    method: 'get'
  })
}
// //下载训练任务日志
// export function downloadLog(params) {
//   return requestLog({
//       url: `/log/download/user/trainjob/${params.jobId}/${params.subName}/index.log`,
//       method: 'get'
//   })
// }



