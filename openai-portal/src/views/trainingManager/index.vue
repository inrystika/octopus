<template>
  <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleTabClick">
    <el-tab-pane label="训练任务" name="menu1">
      <traningTask v-if="tabRefresh.menu1" :training-task="trainingTask" />
    </el-tab-pane>
    <el-tab-pane label="任务模板" name="menu2">
      <taskTemplate v-if="tabRefresh.menu2" :training-template="trainingTemplate" @createTraning="createTraning" />
    </el-tab-pane>
  </el-tabs>
</template>
<script>
  import traningTask from "./traningTask.vue";
  import taskTemplate from "./taskTemplate.vue";
  import { clearProgress } from '@/utils/index.js'
  export default {
    components: {
      traningTask,
      taskTemplate

    },
    data() {
      return {
        activeName: undefined,
        tabRefresh: {
          menu1: true,
          menu2: false
        },
        trainingTemplate: false,
        trainingTask: false
      }
    },
    created() {
      if (this.$route.params.data === undefined) {
        this.activeName = 'menu1'
        this.switchTab('menu1')
      } else if (this.$route.params.data.trainingTemplate) {
        this.activeName = 'menu2'
        this.switchTab('menu2')
        this.trainingTemplate = true
      } else if (this.$route.params.data.trainingTask) {
        this.activeName = 'menu1'
        this.switchTab('menu1')
        this.trainingTask = true
      }
    },
    mounted() {
      window.addEventListener("beforeunload", (e) => {
        clearProgress()
      });
    },
    destroyed() {
      window.removeEventListener("beforeunload", (e) => {
        clearProgress()
      });
    },
    methods: {
      handleTabClick(tab) {
        this.activeName = tab.name
        this.$route.params.data = null
        switch (this.activeName) {
          case 'menu1':
            this.switchTab('menu1')
            this.trainingTemplate = false
            break
          case 'menu2':
            this.switchTab('menu2')
            this.trainingTask = false
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
      },
      createTraning() {
        this.activeName = "menu1"
        this.tabRefresh.menu1 = true
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