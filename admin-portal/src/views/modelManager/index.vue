<template>
  <el-tabs v-model="activeName" @tab-click="handleClick" class="Wrapper">
    <el-tab-pane label="用户模型" name="first">
      <Model :Type=1></Model>
    </el-tab-pane>
    <el-tab-pane label="预置模型" name="third">
      <Model :Type=3></Model>
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