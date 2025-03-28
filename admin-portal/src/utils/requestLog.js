import axios from 'axios'
import { Message } from 'element-ui'

// create an axios instance
const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API2, // url = base url + request url
    // withCredentials: true, // send cookies when cross-domain requests
    timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
    config => {
        // do something before request is sent

        // if (store.getters.token) {
        //     config.headers['Authorization'] = 'Bearer ' + getToken()
        // }
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
        return response
        // if the custom code is not 20000, it is judged as an error.
        // if (res.error) {
        //   Message({
        //     message: res.message || 'Error',
        //     type: 'error',
        //     duration: 5 * 1000
        //   })

        //   // if (res.code === 50008 || res.code === 50012 || res.code === 50014) {
        //   //   // to re-login
        //   //   MessageBox.confirm('You have been logged out, you can cancel to stay on this page, or log in again', 'Confirm logout', {
        //   //     confirmButtonText: 'Re-Login',
        //   //     cancelButtonText: 'Cancel',
        //   //     type: 'warning'
        //   //   }).then(() => {
        //   //     store.dispatch('user/resetToken').then(() => {
        //   //       location.reload()
        //   //     })
        //   //   })
        //   // }
        //   return Promise.reject(new Error(res.message || 'Error'))
        // } else {
        //   return res.payload
        // }
    },
    error => {
        // Message({
        //     message: '没有找到日志',
        //     type: 'warning',
        //     duration: 5 * 1000
        // })
        console.log("error:", error)
        // return Promise.reject(error)
    }
)

export default service
