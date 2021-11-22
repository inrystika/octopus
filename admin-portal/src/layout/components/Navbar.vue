<template>
  <div class="navbar">
    <hamburger :is-active="sidebar.opened" class="hamburger-container" @toggleClick="toggleSideBar" />
    <breadcrumb class="breadcrumb-container" />
    <div class="right-menu">
      <el-row class="demo-avatar demo-basic">
        <el-dropdown>
          <i class="el-icon-document" style="color:#666699;"></i>
          <a href="https://octopus.openi.org.cn/docs/management/intro" target="_blank" class="manual">管理手册</a>
          <i class="el-icon-service" style="color:#666699 ;"></i>
          <a href="https://git.openi.org.cn/OpenI/octopus/issues" target="_blank" class="manual">问题意见</a>
          <el-dropdown-menu slot="dropdown" />
        </el-dropdown>
        <el-avatar :src="circleUrl" :size="size" />
        <el-dropdown>
          <span class="el-dropdown-link">
            管理员<i class="el-icon-arrow-down el-icon--right"></i>
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item divided @click.native="logout">
              <span style="display:block;">退出登录</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </el-row>
    </div>
  </div>
</template>

<script>
  import { mapGetters } from 'vuex'
  import Breadcrumb from '@/components/Breadcrumb'
  import Hamburger from '@/components/Hamburger'
  import { removeToken } from '@/utils/auth'
  export default {
    components: {
      Breadcrumb,
      Hamburger
    },
    data() {
      return {
        circleUrl: "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png",
        size: 'small'
      }
    },
    computed: {
      ...mapGetters([
        'sidebar',
        'avatar',
        'name'
      ])
    },
    methods: {
      toggleSideBar() {
        this.$store.dispatch('app/toggleSideBar')
      },
      async logout() {
        removeToken()
        this.$store.commit('user/SET_TOKEN', '')
        this.$router.push(`/?redirect=${this.$route.fullPath}`)
      }
    }
  }
</script>

<style lang="scss" scoped>
  .navbar {
    height: 60px;
    width: 100%;
    overflow: hidden;
    position: relative;
    box-shadow: 0 1px 4px rgba(0, 21, 41, .08);
    background-color: #1a1a23;

    .hamburger-container {
      line-height: 46px;
      height: 100%;
      float: left;
      cursor: pointer;
      transition: background .3s;
      -webkit-tap-highlight-color: transparent;

      &:hover {
        background: rgba(0, 0, 0, .025)
      }
    }

    .breadcrumb-container {
      float: left;
    }

    .right-menu {
      font-size: 14px;
      margin: 20px 30px 0 20px;
      float: right;
      height: 100%;

      .avatar-container {
        margin-right: 30px;
      }

      .el-dropdown {
        position: relative;
        top: -8px;
        display: inline-block;
        font-size: 15px;
        margin-right: 5px;
        margin-left: 5px;
        color: #ffffff;
      }

      .manual {
        color: #666699;
        margin: 0 50px 0 10px;

      }
    }
  }
</style>