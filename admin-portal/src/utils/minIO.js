import requestMinIO from '@/utils/requestMinIO'
export async function minIO(payload) {
  const res = await requestMinIO({
    url: payload.uploadUrl,
    method: "put",
    data: payload.file,
    onUploadProgress: function(progress) {
      sessionStorage.setItem(JSON.stringify(payload.id), JSON.stringify(parseInt(((progress.loaded / progress.total) * 100))));
      console.log(progress.loaded / progress.total)
    }
  })
  if (res && res.success) {
    return res
  } else {
    return {
      success: false
    }
  }
}