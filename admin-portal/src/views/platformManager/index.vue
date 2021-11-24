<template>
  <div>
    <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleClick">
      <el-tab-pane label="平台列表" name="platformList">
        <platformList></platformList>
      </el-tab-pane>
      <el-tab-pane label="平台训练任务列表" name="trainingTask">
        <platformTrainingTaskList></platformTrainingTaskList>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import platformList from "./platformList.vue"
import platformTrainingTaskList from "./platformTrainingTaskList.vue"
export default {
  name: "index",
  components: {
    platformList,
    platformTrainingTaskList
  },
  data() {
    return {
      activeName: 'platformList',
      tabRefresh: {
        platformMenu: true,
        trainingTaskMenu: false,
      },
    }
  },
  methods: {
    handleClick(tab, event) {
      this.activeName = tab.name
      switch (this.activeName) {
        case 'platformList':
          this.switchTab('platformMenu')
          break
        case 'trainingTask':
          this.switchTab('trainingTaskMenu')
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