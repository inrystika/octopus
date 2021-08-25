import request from '@/utils/request'
// 获取资源池列表
export function getResourcePool() {
    return request({
        url: '/v1/resourcemanage/resourcepool',
        method: 'get',

    })
}
// 删除资源池
export function deleteResourcePool(data) {
    return request({
        url: `/v1/resourcemanage/resourcepool/${data}`,
        method: 'delete'

    })
}
// 创建资源规格
export function createResource(data) {
    return request({
        url: '/v1/resourcemanage/resourcespec',
        method: 'post',
        data
    })
}
// 获取资源规格列表
export function getResource(params) {
    return request({
        url: '/v1/resourcemanage/resourcespec',
        method: 'get',
        params

    })
}
// 创建资源池
export function createResourcePool(data) {
    return request({
        url: '/v1/resourcemanage/resourcepool',
        method: 'post',
        data

    })
}
// 更新资源池
export function updateResourcePool(data) {
    return request({
        url: `/v1/resourcemanage/resourcepool/${data.id}`,
        method: 'put',
        data: { desc:data.desc, bindingNodes:data.bindingNodes, mapResourceSpecIdList:data.mapResourceSpecIdList }


    })
}
// 获取节点列表
export function getNodeList() {
    return request({
        url: '/v1/resourcemanage/node',
        method: 'get',
    })
}
// 创建自定义资源
export function createCustomizeResource(data) {
    return request({
        url: '/v1/resourcemanage/resource',
        method: 'post',
        data
    })
}
// 更新资源
export function updateResource(data) {
    return request({
        url: `/v1/resourcemanage/resource/${data.id}`,
        method: 'put',
        data:{desc:data.desc,resourceRef:data.resourceRef,bindingNodes:data.bindingNodes}

    })
}
// 删除资源规格
export function deleteSpecification(params) {
    return request({
        url: `/v1/resourcemanage/resourcespec/${params}`,
        method: 'delete',

    })
}
// 获取资源列表
export function getResourceList() {
    return request({
        url: '/v1/resourcemanage/resource',
        method: 'get',

    })
}
// 删除资源
export function deleteResource(params) {
    return request({
        url: `/v1/resourcemanage/resource/${params}`,
        method: 'delete',

    })
}