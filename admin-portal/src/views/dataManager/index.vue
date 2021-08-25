<template>
  <div>
    <el-tabs class="Wrapper" v-model="activeName" style="margin:20px 0px 0px 20px" @tab-click="handleClick" >
      <el-tab-pane label="用户数据集" name="userDataset">
        <userList v-if="tabRefresh.userMenu" :Type=1></userList>
      </el-tab-pane>
      <el-tab-pane label="预置数据集" name="preDataset">
        <templateList v-if="tabRefresh.templateMenu" :Type=2></templateList>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import userList from "./userList.vue";
import templateList from "./templateList.vue";
export default {
  components: {
    userList,
    templateList
  },
  data() {
    return {
      activeName: 'userDataset',
      tabRefresh: {
        userMenu: true,
        templateMenu: false,
      }
    };
  },
  methods: {
    handleClick(tab, event) {
      this.activeName = tab.name
      switch (this.activeName) {
        case 'userDataset':
          this.switchTab('userMenu')
          break
        default:
          this.switchTab('templateMenu')
      }
    },
    switchTab(tab) {
      for (const key in this.tabRefresh) {
        if (key === tab) {
          this.tabRefresh[key] = true
        } else {
          this.tabRefresh[key] = false
        }
      }
    },
  }
};
</script>

<style lang="scss" scoped>
.Wrapper {
  margin: 15px!important;
  background-color:#fff;
  padding: 20px;
  min-height: 800px;
}
</style>