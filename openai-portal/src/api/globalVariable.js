
let DOMAIN
if (process.env.NODE_ENV === 'development') {
  DOMAIN = 'http://192.168.202.73'
} else {
  // eslint-disable-next-line eqeqeq
  if (!window.location.port || window.location.port == '') { DOMAIN = window.location.protocol + '//' + document.domain } else { DOMAIN = window.location.protocol + '//' + document.domain + ':' + window.location.port }
}
export default {
  DOMAIN
}