
let DOMAIN
if (process.env.NODE_ENV === 'development') {
  DOMAIN = 'http://192.168.202.73'
} else {
  if (window.location.port != 80 || window.location.port != 443 || window.location.port) {
    DOMAIN = window.location.protocol + '//' + document.domain + ':' + window.location.port
  } else { DOMAIN = window.location.protocol + '//' + document.domain }
}
export default {
  DOMAIN
}