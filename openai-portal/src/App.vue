<template>
  <div id="app">
    <router-view v-if="isShow" />
  </div>
</template>

<script>
import Vue from 'vue'
import { themeChange } from "@/api/themeChange.js"
import { changeThemeColor, curColor } from '@/utils/themeColorClient'
export default {
  name: 'App',
  data() {
    return {
      isShow: false,
      ainColor: curColor
    }
  },
  created() {
    this.themeChange()
  },
  methods: {
    themeChange() {
      const reg = /^#([0-9a-fA-F]{6}|[0-9a-fA-F]{3})$/  // 传入颜色格式须为16进制
      themeChange().then(response => {
        if(response.success) {
          if(response.data && response.data.themeColor && reg.test(response.data.themeColor)){
            Vue.prototype.GLOBAL.THEME_COLOR = response.data.themeColor
            Vue.prototype.GLOBAL.THEME_TITLE_ZH = response.data.systemNameZh
            Vue.prototype.GLOBAL.THEME_TITLE_EN = response.data.systemNameEn
            Vue.prototype.GLOBAL.THEME_LOGO_ADDR = response.data.logoAddr
            Vue.prototype.GLOBAL.THEME_ORG_NAME = response.data.organization
            Vue.prototype.GLOBAL.THEME_MANUAL_INVISIBLE = response.data.manualInvisible
            this.mainColor = response.data.themeColor
            document.title = response.data.systemNameZh // 修改网页标签的title
            const link = document.querySelector('link[rel="icon"]')
            link.href = response.data.logoAddr // 修改网页标签icon
            this.changeColor(this.mainColor)
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
        console.log('err:',err)
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
