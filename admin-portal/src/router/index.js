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
    component: () => import('@/views/Home/login')
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
      component: () => import('@/views/clusterMonitor/index'),
      meta: { title: '集群监控', icon: 'monitor' }
    }]
  },
  {
    path: '/userManager',
    component: Layout,
    children: [
      {
        path: 'index',
        name: '/',
        component: () => import('@/views/userManager/index'),
        meta: { title: '用户管理', icon: 'user' }
      }
    ]
  },
  {
    path: '/resourceManager',
    component: Layout,
    redirect: '/nodeManager',
    name: 'resourceManager',
    meta: {
      title: '资源管理',
      icon: 'resource'
    },
    children: [
      {
        path: 'nodeManager',
        component: () => import('@/views/resourceManager/nodeManager'),
        name: 'nodeManager',
        meta: { title: '节点', icon: 'dot' }
      },
      {
        path: 'resourceMsg',
        component: () => import('@/views/resourceManager/resourceMsg'),
        name: 'resourceMsg',
        meta: { title: '资源', icon: 'dot' }
      },
      {
        path: 'resSpecManager',
        component: () => import('@/views/resourceManager/resSpecManager'),
        name: 'resSpecManager',
        meta: { title: '资源规格', icon: 'dot' }
      },
      {
        path: 'resPoolManager',
        component: () => import('@/views/resourceManager/resPoolManager'),
        name: 'resPoolManager',
        meta: { title: '资源池', icon: 'dot' }
      }
    ]
  },
  {
    path: '/timeManager',
    component: Layout,
    meta: {
      title: '机时管理',
      icon: 'time'
    },
    children: [
      {
        path: 'machine',
        name: 'machine',
        component: () => import('@/views/timeManager/machine'),
        meta: { title: '机时列表', icon: 'dot' }
      },
      {
        path: 'recharge',
        name: 'recharge',
        component: () => import('@/views/timeManager/recharge'),
        meta: { title: '充值记录', icon: 'dot' }
      },
      {
        path: 'consumption',
        name: 'consumption',
        component: () => import('@/views/timeManager/consumption'),
        meta: { title: '消费记录', icon: 'dot' }
      }
    ]
  },
  {
    path: '/dataManager',
    component: Layout,
    redirect: '/dataManager',
    name: '/',
    // meta: { title: '数据管理', icon: '' },
    children: [
      {
        path: 'index',
        name: 'dataManager',
        component: () => import('@/views/dataManager/index'),
        meta: { title: '数据管理', icon: 'data' }
      }
    ]
  },
  {
    path: '/imageManager',
    component: Layout,
    children: [
      {
        path: 'index',
        name: '/',
        component: () => import('@/views/imageManager/index'),
        meta: { title: '镜像管理', icon: 'image' }
      }
    ]
  },
  {
    path: '/devManager',
    component: Layout,
    meta: { title: '模型开发', icon: 'model' },
    alwaysShow: true,
    children: [
      {
        path: 'notebook',
        name: '/',
        component: () => import('@/views/devManager/notebook'),
        meta: { title: 'NoteBook', icon: 'dot' }
      },
      {
        path: 'algorithmManager',
        name: '/',
        component: () => import('@/views/devManager/algorithmManager'),
        meta: { title: '算法管理', icon: 'dot' }
      }
    ]
  },
  {
    path: '/traningManager',
    component: Layout,
    children: [
      {
        path: 'index',
        name: '/',
        component: () => import('@/views/traningManager/index'),
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
  {
    path: '/platformManager',
    component: Layout,
    redirect: '/platform',
    name: 'platformManager',
    meta: {
      title: '平台管理',
      icon: 'example'
    },
    alwaysShow: true,
    children: [
      {
        path: 'platform',
        component: () => import('@/views/platformManager/platform'),
        name: 'platform',
        meta: { title: '平台', icon: 'dot' }
      },
    ]
  },

  // {
  //   path: 'external-link',
  //   component: Layout,
  //   children: [
  //     {
  //       path: 'https://panjiachen.github.io/vue-element-admin-site/#/',
  //       meta: { title: '管理手册', icon: '' }
  //     }
  //   ]
  // },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  base: '/adminportal/',
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
