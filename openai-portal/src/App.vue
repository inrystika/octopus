<template>
  <div id="app">
    <router-view v-if="isShow" />
  </div>
</template>

<script>
  import Vue from 'vue'
  import { themeChange } from "@/api/themeChange.js"
  import { changeThemeColor, curColor } from '@/utils/themeColorClient'
  import store from '@/store'
  import { GetUrlParam } from '@/utils/index.js'
  import { setToken } from '@/utils/auth'
  export default {
    name: 'App',
    data() {
      return {
        isShow: false,
        mainColor: curColor
      }
    },
    mounted() {
      var url = window.location.href
      if (url.indexOf('token') !== -1) {
        setToken(GetUrlParam('token'))
        this.$router.push({ path: '/index', })
      }
      // window.location.href = url
      this.themeChange()
    },
    methods: {
      themeChange() {
        themeChange().then(response => {
          const reg = /^#([0-9a-fA-F]{6}|[0-9a-fA-F]{3})$/  // 传入颜色格式须为16进制
          if (response.success) {
            if (response.data) {
              if (response.data.thirdPlatform) {
                sessionStorage.setItem('platform', response.data.thirdPlatform)
              }
              Vue.prototype.GLOBAL.THEME_TITLE_ZH = response.data.systemNameZh
              Vue.prototype.GLOBAL.THEME_TITLE_EN = response.data.systemNameEn
              Vue.prototype.GLOBAL.THEME_LOGO_ADDR = response.data.logoAddr
              Vue.prototype.GLOBAL.THEME_ORG_NAME = response.data.organization
              Vue.prototype.GLOBAL.THEME_MANUAL_INVISIBLE = response.data.manualInvisible
              document.title = response.data.systemNameZh // 修改网页标签的title
              const link = document.querySelector('link[rel="icon"]')
              link.href = response.data.logoAddr // 修改网页标签icon
              if (response.data.themeColor && reg.test(response.data.themeColor)) {
                Vue.prototype.GLOBAL.THEME_COLOR = response.data.themeColor
                this.mainColor = response.data.themeColor
                this.changeColor(this.mainColor)
              } else {
                Vue.prototype.GLOBAL.THEME_COLOR = ''
              }
              // debugger
            } else {
              Vue.prototype.GLOBAL.THEME_COLOR = ''
              Vue.prototype.GLOBAL.THEME_TITLE_ZH = ''
              Vue.prototype.GLOBAL.THEME_TITLE_EN = ''
              Vue.prototype.GLOBAL.THEME_LOGO_ADDR = ''
              Vue.prototype.GLOBAL.THEME_ORG_NAME = ''
              Vue.prototype.GLOBAL.THEME_MANUAL_INVISIBLE = false
            }
            this.isShow = true
          }
        }).catch(err => {
          Vue.prototype.GLOBAL.THEME_COLOR = ''
          Vue.prototype.GLOBAL.THEME_TITLE_ZH = ''
          Vue.prototype.GLOBAL.THEME_TITLE_EN = ''
          Vue.prototype.GLOBAL.THEME_LOGO_ADDR = ''
          Vue.prototype.GLOBAL.THEME_ORG_NAME = ''
          Vue.prototype.GLOBAL.THEME_MANUAL_INVISIBLE = false
          this.isShow = true
        })
      },
      changeColor(newColor) {
        changeThemeColor(newColor).then(() => {
          // this.$message.success('主题色切换成功')
        })
      },
    }
  }
</script>