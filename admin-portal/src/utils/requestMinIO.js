import axios from 'axios'
import { Message } from 'element-ui'
// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_MINIO
})

// request interceptor
service.interceptors.request.use(
  config => {
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
  */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */

  response => {
    // res = response.data
    // res.data = response.data.payload
    if (response.status === 200) {
      return { success: true }
    } else { return { success: false } }

  },
  error => {
    console.log('err' + error) // for debug
    Message({
      message: '上传失败',
      type: 'error',
      duration: 5 * 1000
    })
    setTimeout(function() { location.reload() }, 1000)
    // location.reload();
    return Promise.reject(error)
  }
)
export default service
