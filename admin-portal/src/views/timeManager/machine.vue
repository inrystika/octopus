<template>
  <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleTabClick">
    <el-tab-pane label="用户机时列表" name="menu1">
      <Time v-if="tabRefresh.menu1" :time-tab-type="1" />
    </el-tab-pane>
    <el-tab-pane label="群组机时列表" name="menu2">
      <Time v-if="tabRefresh.menu2" :time-tab-type="2" />
    </el-tab-pane>
  </el-tabs>
</template>
<script>
  import Time from "./component/time.vue";

  export default {
    components: {
      Time

    },
    data() {
      return {
        activeName: 'menu1',
        tabRefresh: {
          menu1: true,
          menu2: false

        }
      }
    },
    mounted() {
      window.addEventListener('beforeunload', e => {
        sessionStorage.clear()
      });

    },
    destroyed() {
      window.removeEventListener('beforeunload', e => {
        sessionStorage.clear()
      })
    },
    methods: {
      handleTabClick(tab) {
        this.activeName = tab.name
        switch (this.activeName) {
          case 'menu1':
            this.switchTab('menu1')
            break
          case 'menu2':
            this.switchTab('menu2')
            break
          case 'menu3':
            this.switchTab('menu3')
            break
        }
      },
      switchTab(tab) {
        for (const key in this.tabRefresh) {
          if (key === tab) {
            this.tabRefresh[key] = true
          } else {
            this.tabRefresh[key] = false
          }
        }
      }
    }
  }
</script>
<style lang="scss" scoped>
  .Wrapper {
    margin: 15px !important;
    background-color: #fff;
    padding: 20px;
    min-height: 900px;
  }
</style>