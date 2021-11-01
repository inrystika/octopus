<template>
  <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleTabClick">
    <el-tab-pane label="部署列表" name="menu1">
      <deployList />
    </el-tab-pane>


  </el-tabs>
</template>
<script>
  import deployList from "./deployList.vue";
  export default {
    components: {
      deployList
    },
    data() {
      return {
        activeName: 'menu1',

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
      handleTabClick(tab) {

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