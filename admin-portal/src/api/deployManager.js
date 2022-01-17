import request from '@/utils/request'
// import requestLog from '@/utils/requestLog'
// 获取模型部署列表
export function getDeployList(params) {
    return request({
        url: '/v1/deploymanage/modeldeploy',
        method: 'get',
        params
    })
}

// 获取模型部署详情
export function deployDetail(params) {
    return request({
        url: `/v1/deploymanage/modeldeploy/${params}`,
        method: 'get'
    })
}

// 获取模型部署事件列表
export function deployEvent(params) {
    return request({
        url: `/v1/deploymanage/modeldeployevent`,
        method: 'get',
        params
    })
}
