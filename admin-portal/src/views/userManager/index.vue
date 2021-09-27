<template>
  <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleTabClick">
    <el-tab-pane label="用户列表" name="menu1">
      <user v-if="tabRefresh.menu1" :user-tab-type="1" />
    </el-tab-pane>
    <el-tab-pane label="群组列表" name="menu2">
      <user v-if="tabRefresh.menu2" :user-tab-type="2" />
    </el-tab-pane>
  </el-tabs>
</template>
<script>
  import user from "./user.vue";
  export default {
    components: {
      user
    },
    data() {
      return {
        activeName: "menu1",
        tabRefresh: {
          menu1: true,
          menu2: false,
          menu3: false

        }
      }
    },
    methods: {
      handleTabClick(tab) {
        this.activeName = tab.name
        switch (this.activeName) {
          case "menu1":
            this.switchTab("menu1")
            break
          case "menu2":
            this.switchTab("menu2")
            break
          case "menu3":
            this.switchTab("menu3")
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
    margin: 15px!important;
    background-color:#fff;
    padding: 20px;
    min-height: 900px
  }
</style>