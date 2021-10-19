<template>
  <div>
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      background-color="#409EFF"
      text-color="#fff"
      active-text-color="#ffd04b"
      @select="handleSelect"
    >
      <el-menu-item index="1" class="index1">鹏城云脑开源平台</el-menu-item>
      <el-menu-item index="2"><a href="" target="_blank">首页</a></el-menu-item>
      <el-menu-item index="3"><a href="" target="_blank">使用文档</a></el-menu-item>
      <el-menu-item index="4"><a href="" target="_blank">社区</a></el-menu-item>
      <el-menu-item index="5"><a href="" target="_blank">FAQ</a></el-menu-item>
      <el-menu-item index="6"><a href="" target="_blank">代码仓库</a></el-menu-item>
    </el-menu>
  </div>
</template>
<script>
  export default {
    data() {
      return {
        activeIndex: '2'
      };
    },
    methods: {
      handleSelect(key, keyPath) {
        console.log(key, keyPath);
      }
    }
  }
</script>
<style>
</style>