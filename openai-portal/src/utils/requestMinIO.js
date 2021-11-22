import axios from 'axios'
import { Message } from 'element-ui'
import store from '@/store'

// create an axios instance
const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API// url = base url + request url
    // withCredentials: true, // send cookies when cross-domain requests
    // timeout: 5000 // request timeout
    // eslint-disable-next-line no-undef

})

// request interceptor
service.interceptors.request.use(
    config => {
        return config
    },
    error => { // for debug
        return Promise.reject(error)
    }
)

service.interceptors.response.use(

    response => {
        return response
    },
    error => {
        // eslint-disable-next-line no-undef
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