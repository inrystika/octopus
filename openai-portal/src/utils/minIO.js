import requestMinIO from '@/utils/requestMinIO'
export async function minIO(payload) {
  const res = await requestMinIO({
    url: payload.uploadUrl,
    method: "put",
    data: payload.file
  })
  if (res && res.success) {
    return res
  } else {
    return {
      success: false
    }
  }
}