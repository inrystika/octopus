import Vue from 'vue'
import 'normalize.css/normalize.css' // A modern alternative to CSS resets
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
// import locale from 'element-ui/lib/locale/lang/en' // lang i18n
import zhLocale from 'element-ui/lib/locale/lang/zh-CN'
import '@/styles/index.scss' // global css
import App from './App'
import store from './store'
import router from './router'
import '@/icons' // icon
import '@/permission' // permission control
import VueAwesomeSwiper from 'vue-awesome-swiper'
import '../theme/index.css'
import '@/styles/dot.scss'
import globalVariable from '@/api/globalVariable.js'
import install from './preventReClick'
import { parseTime } from "@/utils/index"
import { mixin } from '@/error/index'
import directives from './directives'
//设置为全局过滤器，避免冗余代码
Vue.filter('parseTime', parseTime)
// 将多个子组件公用的getErrorMsg函数，提取成一个全局混入对象
Vue.mixin(mixin)

Vue.prototype.GLOBAL = globalVariable
/**
 * If you don't want to use mock-server
 * you want to use MockJs for mock api
 * you can execute: mockXHR()
 *
 * Currently MockJs will be used in the production environment,
 * please remove it before going online ! ! !
 */
// if (process.env.NODE_ENV === 'production') {
//   const { mockXHR } = require('../mock')
//   mockXHR()
// }

// set ElementUI lang to EN
Vue.use(ElementUI, { zhLocale })
// 如果想要中文版 element-ui，按如下方式声明
// Vue.use(ElementUI)
Vue.use(VueAwesomeSwiper /* { default global options } */)
Vue.use(directives)
Vue.use(install)
Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
