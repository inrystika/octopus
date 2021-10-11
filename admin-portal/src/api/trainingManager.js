import request from '@/utils/request'
import requestLog from '@/utils/requestLog'
// 获取训练任务列表
export function getTraining(params) {
    return request({
        url: '/v1/trainmanage/trainjob',
        method: 'get',
        params
    })
}
// 停止训练任务
export function stopTraining(params) {
    return request({
        url: `/v1/trainmanage/trainjob/${params}/stop`,
        method: 'post'
    })
}
// 获取训练任务详情
export function trainingDetail(params) {
    return request({
        url: `/v1/trainmanage/trainjob/${params}`,
        method: 'get'
    })
}
// 任务日志访问
export function showLog(params) {
    return requestLog({
        url: `/log/user/trainjob/${params.jobId}/${params.subName}/index.log`,
        method: 'get'
    })
}
// 下载训练任务日志
export function downloadLog(params) {
    return requestLog({
        url: `/log/download/user/trainjob/${params.jobId}/${params.subName}/index.log`,
        method: 'get'
    })
  }

