<template>
    <el-tabs v-model="activeName" @tab-click="handleClick" class="Wrapper">
      <el-tab-pane label="资源池列表" name="first">
        <poolManager></poolManager>
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