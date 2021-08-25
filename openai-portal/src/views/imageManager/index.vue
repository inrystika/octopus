<template>
  <el-tabs v-model="activeName" @tab-click="handleTabClick" class="Wrapper">
    <el-tab-pane label="我的镜像" name="menu1">
      <mirror :image="image" :Type=1 v-if="tabRefresh.menu1"></mirror>
    </el-tab-pane>
    <el-tab-pane label="公共镜像" name="menu2">
      <mirror :Type=3 v-if="tabRefresh.menu2"></mirror>
    </el-tab-pane>
    <el-tab-pane label="预置镜像" name="menu3">
      <mirror :Type=2 v-if="tabRefresh.menu3"></mirror>
    </el-tab-pane>
  </el-tabs>
</template>
<script>
  import mirror from "./Image.vue";
  export default {
    components: {
      mirror
    },
    data() {
      return {
        activeName: 'menu1',
        tabRefresh: {
          menu1: true,
          menu2: false,
          menu3: false
        },
        image: false
      }
    },
    created(){
      if (this.$route.params.data === undefined) {
        this.dataset = false
      } else if (this.$route.params.data.image) {
        this.image = true
      }
    },
    methods: {
      handleTabClick(tab) {
        this.activeName = tab.name
        switch (this.activeName) {
          case 'menu1':
            this.switchTab('menu1')
            break
          case 'menu2':
            this.switchTab('menu2')
            this.image = false
            break
          case 'menu3':
            this.switchTab('menu3')
            this.image = false
            break
        }
      },
      switchTab(tab) {
        for (let key in this.tabRefresh) {
          if (key === tab) {
            this.tabRefresh[key] = true
          } else {
            this.tabRefresh[key] = false
          }
        }
      },
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