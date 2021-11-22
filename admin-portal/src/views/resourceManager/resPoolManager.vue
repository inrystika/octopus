<template>
  <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleClick">
    <el-tab-pane label="资源池列表" name="first">
      <poolManager />
    </el-tab-pane>

  </el-tabs>
</template>
<script>
  import poolManager from "./components/poolManager.vue";

  export default {
    components: {
      poolManager

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