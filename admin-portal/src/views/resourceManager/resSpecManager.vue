<template>
  <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleClick">
    <el-tab-pane label="资源规格列表" name="first">
      <ResourceSpec />
    </el-tab-pane>
  </el-tabs>
</template>
<script>
  import ResourceSpec from "./components/ResourceSpec.vue";

  export default {
    components: {
      ResourceSpec

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
    margin: 20px !important;
    background-color: #fff;
    padding: 20px;
    min-height: 900px;
  }
</style>