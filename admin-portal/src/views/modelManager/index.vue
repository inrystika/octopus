<template>
  <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleClick">
    <el-tab-pane label="用户模型" name="first">
      <Model :model-tab-type="1" />
    </el-tab-pane>
    <el-tab-pane label="预置模型" name="third">
      <Model :model-tab-type="3" />
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