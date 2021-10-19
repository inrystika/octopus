<template>
    <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleClick">
      <el-tab-pane label="节点管理" name="first">
        <nodeList />
      </el-tab-pane>
    </el-tabs>
  </template>
  <script>
import nodeList from "./components/nodeList.vue";

    export default {
      components: {
        nodeList

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