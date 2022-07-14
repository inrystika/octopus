<template>
  <div>
    <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleClick">
      <el-tab-pane label="用户算法" name="userAlgorithm">
        <userList v-if="tabRefresh.userMenu" :algorithm-tab-type="1" />
      </el-tab-pane>
      <el-tab-pane label="预置算法" name="preAlgorithm">
        <templateList v-if="tabRefresh.templateMenu" :algorithm-tab-type="2" />
      </el-tab-pane>
      <el-tab-pane label="算法配置" name="algorithmConfig">
        <algorithmConfig v-if="tabRefresh.algorithmConfig" :data-tab-type="3" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
  import userList from "./components/algorithm/userList.vue";
  import templateList from "./components/algorithm/templateList.vue";
  import algorithmConfig from "./components/algorithm/algorithmConfig.vue";
  export default {
    components: {
      userList,
      templateList,
      algorithmConfig
    },
    data() {
      return {
        activeName: "userAlgorithm",
        tabRefresh: {
          userMenu: true,
          templateMenu: false,
          algorithmConfig: false
        }
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
        this.activeName = tab.name
        switch (this.activeName) {
          case 'userAlgorithm':
            this.switchTab('userMenu')
            break
          case 'preAlgorithm':
            this.switchTab('templateMenu')
            break
          default:
            this.switchTab('algorithmConfig')
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
  }
</script>

<style scoped>
  .Wrapper {
    margin: 15px !important;
    background-color: #fff;
    padding: 20px;
    min-height: 800px;
  }
</style>