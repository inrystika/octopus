<template>
  <div>
    <el-tabs v-model="activeName" class="Wrapper" style="margin:20px 0px 0px 20px" @tab-click="handleClick">
      <el-tab-pane label="用户数据集" name="userDataset">
        <userList v-if="tabRefresh.userMenu" :data-tab-type="1" />
      </el-tab-pane>
      <el-tab-pane label="预置数据集" name="preDataset">
        <templateList v-if="tabRefresh.templateMenu" :data-tab-type="2" />
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
          templateMenu: false
        }
      };
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
      }
    }
  };
</script>

<style lang="scss" scoped>
  .Wrapper {
    margin: 15px !important;
    background-color: #fff;
    padding: 20px;
    min-height: 800px;
  }
</style>