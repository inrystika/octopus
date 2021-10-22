<template>
  <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleClick">
    <el-tab-pane label="训练任务列表" name="first">
      <traningTask />
    </el-tab-pane>
  </el-tabs>
</template>
<script>
  import traningTask from "./traningTaskList.vue";
  export default {
    components: {
      traningTask

    },
    data() {
      return {
        activeName: 'first'
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
      handleClick(tab, event) {
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