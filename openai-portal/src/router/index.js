import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/',
    component: () => import('@/views/Home/firstPage'),
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/index',
    component: Layout,
    // redirect: '/GeneralView',
    children: [{
      path: '/index',
      name: 'index',
      component: () => import('@/views/GeneralView/index'),
      meta: { title: '概览', icon: 'overView' }
    }]
  },
  {
    path: '/dataManager',
    component: Layout,
    meta: { title: '数据管理', icon: 'data' },
    alwaysShow: true,
    children: [
      {
        path: 'index',
        name: 'dataManager',
        component: () => import('@/views/dataManager/index'),
        meta: { title: '数据集管理', icon: 'dot' }
      },
    ]
  },
  {
    path: '/imageManager',
    component: Layout,
    children: [
      {
        path: 'index',
        name: 'imageManager',
        component: () => import('@/views/imageManager/index'),
        meta: { title: '镜像管理', icon: 'image' }
      }
    ]
  },
  {
    path: '/modelDev',
    component: Layout,
    meta: { title: '模型开发', icon: 'model' },
    alwaysShow: true,
    children: [
      {
        path: 'notebook',
        name: 'notebook',
        component: () => import('@/views/modelDev/notebook'),
        meta: { title: 'NoteBook', icon: 'dot' }
      },
      {
        path: 'algorithmManager',
        name: 'algorithmManager',
        component: () => import('@/views/modelDev/algorithmManager'),
        meta: { title: '算法管理', icon: 'dot' }
      },
    ]
  },
  {
    path: '/trainingManager',
    component: Layout,
    children: [
      {
        path: 'index',
        name: 'trainingManager',
        component: () => import('@/views/trainingManager/index'),
        meta: { title: '训练管理', icon: 'training' }
      }
    ]
  },
  {
    path: '/modelManager',
    component: Layout,
    children: [
      {
        path: 'index',
        name: 'modelManager',
        component: () => import('@/views/modelManager/index'),
        meta: { title: '模型管理', icon: 'model2' }
      }
    ]
  },
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
