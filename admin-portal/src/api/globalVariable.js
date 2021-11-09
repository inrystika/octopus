
let DOMAIN
if (process.env.NODE_ENV === 'development') {
  // eslint-disable-next-line no-undef
  DOMAIN = process.env.VUE_APP_BASE_DOMAIN || 'http://192.168.202.73'
} else {
  if (!window.location.port || window.location.port == '') { DOMAIN = window.location.protocol + '//' + document.domain }
  else { DOMAIN = window.location.protocol + '//' + document.domain + ':' + window.location.port }
}
export default {
  DOMAIN
}
// 本地调试执行npm run dev指令需要修改DOMAIN地址，将http://192.168.202.73替换成服务器地址