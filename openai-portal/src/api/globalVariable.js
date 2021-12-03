
import { themeChange } from '@/api/themeChange.js'

let DOMAIN
if (process.env.NODE_ENV === 'development') {
  DOMAIN = process.env.VUE_APP_BASE_DOMAIN || 'http://192.168.202.73'
} else {
  // eslint-disable-next-line eqeqeq
  if (!window.location.port || window.location.port == '') { DOMAIN = window.location.protocol + '//' + document.domain } else { DOMAIN = window.location.protocol + '//' + document.domain + ':' + window.location.port }
}

let obj = themeChange()
console.log("obj:",obj)
let THEMECOLOR = obj.data && obj.data.themeColor ? obj.data.themeColor : '#1a1a23'
let THEMETITLEZH = obj.data && obj.data.systemNameZh ? obj.data.systemNameZh : '启智章鱼'
let THEMETITLEEN = obj.data && obj.data.systemNameEn ? obj.data.systemNameEn : 'Octopus'
let THEMELOGOADDR = obj.data && obj.data.logoAddr ? obj.data.logoAddr : './'

export default {
  DOMAIN,
  THEMECOLOR,
  THEMETITLEZH,
  THEMETITLEEN,
  THEMELOGOADDR
}
// 本地调试执行npm run dev指令需要修改DOMAIN地址，将http://192.168.202.73替换成服务器地址
