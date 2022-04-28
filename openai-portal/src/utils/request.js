import axios from 'axios'
import { Message } from 'element-ui'
import { getToken, removeToken } from '@/utils/auth'
import router from '../router'
// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 60000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent
    if (getToken()) {
      // let each request carry token
      // ['X-Token'] is a custom headers key
      // please modify it according to the actual situation
      config.headers['Authorization'] = 'Bearer ' + getToken()
    }
    // eslint-disable-next-line no-undef
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
    const res = response.data
    if (response.data.payload) {
      res.data = response.data.payload
    }
    if (response.status === 200 && response.data === '' && response.headers.url) {
      window.open(response.headers.url, '_blank')
    } else if (!response.data.success && (response.data.error.subcode === 16004 || response.data.error.subcode === 16010 || response.data.error.subcode === 16007)) {

      setTimeout(function() {
        removeToken()
        router.replace({ path: '/' })
      }, 1000)
    } else {
      return res
    }
  },
  error => {
    console.log('err' + error) // for debug
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
