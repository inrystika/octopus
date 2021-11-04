import requestMinIO from '@/utils/requestMinIO'
import { Message } from 'element-ui'
export async function minIO(payload) {
  const res = await requestMinIO({
    url: payload.uploadUrl,
    method: "put",
    data: payload.file,
    onUploadProgress: function(progress) {
      sessionStorage.setItem(JSON.stringify(payload.id), JSON.stringify(parseInt(((progress.loaded / progress.total) * 100))));
    }
  })
  if (res && res.success) {
    return res
  } else {
    sessionStorage.setItem(JSON.stringify(payload.id), 0)
    return {
      success: false
    }
  }
}