<template>
  <el-tabs v-model="activeName" @tab-click="handleClick" class="Wrapper">
    <el-tab-pane label="部署列表" name="first">
      <deploy></deploy>
    </el-tab-pane>
  </el-tabs>
</template>
<script>
  import deploy from "./deploy.vue";
  export default {
    components: {
      deploy
    },
    data() {
      return {
        activeName: 'first'
      }
    },
    created() {

    },
    methods: {
      handleClick() { }
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