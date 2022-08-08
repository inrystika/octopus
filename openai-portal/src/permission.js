import router from './router'
import store from './store'
import { Message } from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import { getToken, setToken, getCookie } from '@/utils/auth' // get token from cookie
import getPageTitle from '@/utils/get-page-title'
import { GetUrlParam } from '@/utils/index.js'
NProgress.configure({ showSpinner: false }) // NProgress Configuration
const whiteList = ['/', '/register'] // no redirect whitelist
router.beforeEach(async (to, from, next) => {
  // start progress bar
  NProgress.start()

  // set page title
  document.title = getPageTitle(to.meta.title)
  if (!getToken() && getCookie('token')) {
    setToken(getCookie('token'))
  }
  // determine whether the user has logged in
  const hasToken = getToken()
  if (hasToken) {
    try {
      // eslint-disable-next-line eqeqeq
      if (store.getters.name === '') { await store.dispatch('user/getInfo') }
      if (store.getters.workspaces.length === 0) { await store.dispatch('user/getSpace') }
    } catch (error) {
      await store.dispatch('user/resetToken')
      // Message.error('')
      next('/index')
      NProgress.done()
    }

    if (to.path === '/') {
      next('/index')
      NProgress.done()
    } else if (to.path === '/register') {
      next('/index')
      NProgress.done()
    } else {
      next()
    }
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      if (to.path === '/register') {

        if (GetUrlParam('token') && GetUrlParam('token') !== '') {

          setToken(GetUrlParam('token'))
          next('/index')
        } else {

          next()
        }
      } else {

        next()
      }
    } else {
      const url = window.location.href
      if (url.indexOf('register') !== -1) {
        next('/register')
      }

      NProgress.done()
    }
  }
})
router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})
