<template>
  <div class="sidebar-logo-container" :class="{'collapse':collapse}" :style="{'background':this.GLOBAL.THEME_COLOR?this.GLOBAL.THEME_COLOR:''}">
    <transition name="sidebarLogoFade">
      <router-link v-if="collapse" key="collapse" class="sidebar-logo-link" to="/">
      <img v-if="logo" :src="this.GLOBAL.THEME_LOGO_ADDR?this.GLOBAL.THEME_LOGO_ADDR:logoCollapse" class="sidebar-logo">         
      </router-link>
      <router-link v-else key="expand" class="sidebar-logo-link" to="/">
        <div v-if="this.GLOBAL.THEME_LOGO_ADDR">
          <img :src="this.GLOBAL.THEME_LOGO_ADDR?GLOBAL.THEME_LOGO_ADDR:logo" class="sidebar-logo">
          <p class="sidebar-title">{{this.GLOBAL.THEME_TITLE_ZH}}</p>
        </div>
        <div v-else><img v-if="logo" :src="this.GLOBAL.THEME_LOGO_ADDR?GLOBAL.THEME_LOGO_ADDR:logo" class="sidebar-logo"></div>
      </router-link>
    </transition>
  </div>
</template>

<script>
  export default {
    name: 'SidebarLogo',
    props: {
      collapse: {
        type: Boolean,
        required: true
      }
    },
    data() {
      return {
        logoCollapse: require('@/assets/logoCollapse.svg'),
        logo: require('@/assets/logo-w.svg')
      }
    }
  }
</script>

<style lang="scss" scoped>
  .sidebarLogoFade-enter-active {
    transition: opacity 1.5s;
  }

  .sidebarLogoFade-enter,
  .sidebarLogoFade-leave-to {
    opacity: 0;
  }

  .sidebar-logo-container {
    position: relative;
    width: 100%;
    height: 60px;
    line-height: 60px;
    background: #000;
    text-align: center;
    overflow: hidden;

    & .sidebar-logo-link {
      height: 100%;
      width: 100%;
      display: flex;
      justify-content: center;

      & .sidebar-logo {
        width: 50px;
        height: 50px;
        vertical-align: middle;
        margin-right: 12px;
      }

      & .sidebar-title {
        display: inline-block;
        margin: 0;
        color: #fff;
        font-weight: 600;
        line-height: 50px;
        font-size: 16px;
        font-family: Avenir, Helvetica Neue, Arial, Helvetica, sans-serif;
        vertical-align: middle;
      }
    }

    &.collapse {
      .sidebar-logo {
        width: 50px;
        height: 50px;
        margin-right: 0px;
      }
    }
  }
</style>