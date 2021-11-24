<template>
  <div :class="{'has-logo':true}">
    <logo :collapse="isCollapse" />
    <el-scrollbar wrap-class="scrollbar-wrapper">
      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :background-color="variables.menuBg"
        :text-color="variables.menuText"
        :unique-opened="false"
        :active-text-color="variables.menuActiveText"
        :collapse-transition="false"
        mode="vertical"
      >
        <div v-for="route in routes" :key="route.path">
          <sidebar-item
            v-if="route.path === '/cloudInterconnection' && isPermission === 'no' ? false : true"
            :key="route.path"
            :item="route"
            :base-path="route.path"
          />
        </div>
        <!-- <sidebar-item v-for="route in routes" :key="route.path" :item="route" :base-path="route.path" /> -->
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Logo from './Logo'
import SidebarItem from './SidebarItem'
import variables from '@/styles/variables.scss'
import { getUserConfig } from '@/api/Home'
import { getErrorMsg } from '@/error/index'

export default {
  components: { SidebarItem, Logo },
  data(){
    return {
      isPermission: 'no'
    }
  },
  created() {
    this.getUserConfig()
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    getUserConfig() {
      getUserConfig().then(response => {
        if (response.success) {
<<<<<<< HEAD
          if(response.data.config.jointCloudPermission === 'no') {
            this.$router.options.routes.forEach((item,index) => {
              if(item.children[0].name === "cloudInterconnection") {
                delete this.$router.options.routes[item]
              }
            })
          }
=======
          this.isPermission = response.data.config&&response.data.config.jointCloudPermission?response.data.config.jointCloudPermission:'no'
>>>>>>> db391938c7eacdd41c8d4cb56729e1bcb908d39b
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          });
        }
      })
    }
  },
  computed: {
    ...mapGetters([
      'sidebar'
    ]),
    routes() {
      return this.$router.options.routes
    },
    activeMenu() {
      const route = this.$route
      const { meta, path } = route
      // if set path, the sidebar will highlight the path you set
      if (meta.activeMenu) {
        return meta.activeMenu
      }
      return path
    },
    variables() {
      return variables
    },
    isCollapse() {
      return !this.sidebar.opened
    }
  }
}
</script>
