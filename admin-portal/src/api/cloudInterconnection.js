import request from '@/utils/request'

export function getCloudTrainJobList(params) {
  let conditions = []
  conditions.push(`pageIndex=` + params.pageIndex);
  conditions.push(`pageSize=` + params.pageSize);
  params.ids?conditions.push(`ids=` + params.ids):null;
  return request({
    url: `/v1/jointcloudmanage/job?` + conditions.join("&"),
    method: 'get',
  })
}