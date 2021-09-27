<template>
  <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleTabClick">
    <el-tab-pane label="我的模型" name="menu1">
      <Model v-if="tabRefresh.menu1" :model-tab-type="1" />
    </el-tab-pane>
    <el-tab-pane label="公共模型" name="menu2">
      <Model v-if="tabRefresh.menu2" :model-tab-type="2" />
    </el-tab-pane>
    <el-tab-pane label="预置模型" name="menu3">
      <Model v-if="tabRefresh.menu3" :model-tab-type="3" />
    </el-tab-pane>
  </el-tabs>
</template>
<script>
  import Model from "./Model.vue";
  export default {
    components: {
      Model
    },
    data() {
      return {
        activeName: 'menu1',
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
    margin: 20px !important;
    background-color: #fff;
    padding: 20px;
    min-height: 900px;
  }
</style>