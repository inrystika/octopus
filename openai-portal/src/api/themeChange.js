import request from '@/utils/request'

export function themeChange() {
  return request({
    url: `/v1/systemmanage/webconfig`,
    method: 'get',
  })
}