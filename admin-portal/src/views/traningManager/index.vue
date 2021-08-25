<template>
  <el-tabs v-model="activeName" @tab-click="handleClick" class="Wrapper">
    <el-tab-pane label="训练任务列表" name="first">
      <traningTask></traningTask>
    </el-tab-pane>
  </el-tabs>
</template>
<script>
  import traningTask from "./traningTaskList.vue";
  export default {
    components: {
      traningTask,

    },
    data() {
      return {
        activeName: 'first'
      }
    },
    methods: {
      handleClick(tab, event) {
      }
    }
  }
</script>
<style lang="scss" scoped>
  .Wrapper {
    margin: 15px!important;
    background-color:#fff;
    padding: 20px;
    min-height: 900px;
  }
</style>