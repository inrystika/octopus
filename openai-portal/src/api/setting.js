import request from '@/utils/request'
export function updateUserFtpAccount(data) {
  return request({
    url: '/v1/usermanage/user/ftpaccount',
    method: 'put',
    data
  })
}