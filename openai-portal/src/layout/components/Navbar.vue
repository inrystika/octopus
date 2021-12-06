<template>
  <div class="navbar" :style="{'background-color':this.GLOBAL.THEME_COLOR?this.GLOBAL.THEME_COLOR:''}">
    <hamburger :is-active="sidebar.opened" class="hamburger-container" @toggleClick="toggleSideBar" />
    <breadcrumb class="breadcrumb-container" />
    <div class="right-menu">
      <el-row class="demo-avatar demo-basic">
        <el-dropdown>
          <i class="el-icon-document" :style="{'color':fontColor?fontColor:'#666699'}" />
          <a href="https://octopus.openi.org.cn/docs/manual/intro" target="_blank" class="manual" :style="{'color':fontColor}">使用手册</a>
          <i class="el-icon-service" :style="{'color':fontColor?fontColor:'#666699'}" />
          <a href="https://git.openi.org.cn/OpenI/octopus/issues" target="_blank" class="manual" :style="{'color':fontColor}">问题意见</a>
          <el-dropdown-menu slot="dropdown" />
        </el-dropdown>
        <el-avatar :src="circleUrl" :size="size" />
        <el-dropdown>
          <span class="el-dropdown-link">
            {{ name }}<i class="el-icon-arrow-down el-icon--right" />
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item divided @click.native="logout">
              <span style="display:block;">退出登录</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <el-dropdown @command="handleCommand" @visible-change="change">
          <span class="el-dropdown-link">
            {{ current }}<i class="el-icon-arrow-down el-icon--right" />
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item v-for="(item) in options" :key="item.index" :command="item">{{ item.name }}
            </el-dropdown-item>
          </el-dropdown-menu>
          <!-- <el-color-picker size="mini" @change="changeColor" v-model="mainColor">主题切换</el-color-picker> -->
        </el-dropdown>
      </el-row>
    </div>
  </div>
</template>

<script>
  import { changeSpace } from '@/api/Home'
  import { mapGetters } from 'vuex'
  import Breadcrumb from '@/components/Breadcrumb'
  import Hamburger from '@/components/Hamburger'
  import { getSpace } from '@/api/Home'
  // import { changeThemeColor, curColor } from '@/utils/themeColorClient'
  export default {
    components: {
      Breadcrumb,
      Hamburger
    },
    data() {
      return {
        options: [],
        circleUrl: "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png",
        size: 'small',
        current: '默认群组',
        userMsg: undefined,
        fontColor: this.GLOBAL.THEME_COLOR ? '#FFFFFF' : ''
        // mainColor: curColor
      }
    },
    computed: {
      ...mapGetters([
        'sidebar',
        'avatar',
        'name',
        'workspaces',
        'id',
        'workspaceId'
      ])
    },
    created() {
      this.getSpace()
    },
    methods: {
      change(val) {
        if (val) {
          this.getSpace()
        }
      },
      getSpace() {
        getSpace(this.id).then(response => {
          let data = []
          if (response.payload != null) {
            data = response.payload.workspaces; data.forEach(item => {
              if (item.id === 'default-workspace') {
                item.name = '默认群组'
              }
            })
          }
          this.options = data
          this.options.forEach(
            item => {
              if (item.id === this.workspaceId) {
                this.current = item.name
              }
            }
          )
        })
      },
      toggleSideBar() {
        this.$store.dispatch('app/toggleSideBar')
      },
      async logout() {
        await this.$store.dispatch('user/logout')
        this.$router.push(`/?redirect=${this.$route.fullPath}`)
      },
      handleCommand(command) {
        // 切换群组页面刷新但是保留页面当前群组状态
        const data = { userId: this.id, workspaceId: command.id }
        this.current = command.name
        changeSpace(data).then(response => {
          this.$message({
            message: '切换成功',
            type: 'success'
          });
          location.reload()
        })
      },
      // changeColor(newColor) {
      //   changeThemeColor(newColor).then(() => {
      //     this.$message.success('主题色切换成功')
      //   })
      // }
    },
  }
</script>

<style lang="scss" scoped>
  .navbar {
    height: 60px;
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
      margin: 20px 30px 0 20px;
      float: right;
      height: 100%;
      color: #409EFF;
      font-size: 20px;

      .avatar-container {
        margin-right: 30px;
      }

      .el-dropdown {
        position: relative;
        top: -8px;
        display: inline-block;
        color: #fff;
        font-size: 15px;
        margin-right: 5px;
        margin-left: 5px
      }

      .manual {
        color: #666699;
        margin: 0 50px 0 10px;
      }
    }
  }
  .pkuNavbar {
    height: 60px;
    overflow: hidden;
    position: relative;
    // box-shadow: 0 1px 4px rgba(0, 21, 41, .08);
    background-color: #94070A;

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
      margin: 20px 30px 0 20px;
      float: right;
      height: 100%;
      color: #409EFF;
      font-size: 20px;

      .avatar-container {
        margin-right: 30px;
      }

      .el-dropdown {
        position: relative;
        top: -8px;
        display: inline-block;
        color: #fff;
        font-size: 15px;
        margin-right: 5px;
        margin-left: 5px
      }

      .manual {
        color: #666699;
        margin: 0 50px 0 10px;
      }
    }
  }
</style>