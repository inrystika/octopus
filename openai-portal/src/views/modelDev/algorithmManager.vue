<template>
    <div>
      <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleClick">
        <el-tab-pane label="我的算法" name="myAlgorithm">
          <myList :algorithm="algorithm" :Type=1 v-if="tabRefresh.myMenu"></myList>
        </el-tab-pane>
        <el-tab-pane label="公共算法" name="publicAlgorithm">
          <publicList :Type=2 v-if="tabRefresh.pubMenu"></publicList>
        </el-tab-pane>
        <el-tab-pane label="预置算法" name="preAlgorithm">
          <presetList :Type=3 v-if="tabRefresh.preMenu"></presetList>
        </el-tab-pane>
      </el-tabs>
    </div>
</template>

<script>
import myList from "./components/algorithm/myList.vue";
import publicList from "./components/algorithm/publicList.vue";
import presetList from "./components/algorithm/presetList.vue";
export default { 
  components: {
    myList,
    publicList,
    presetList
  },
  data() {
    return {
      activeName: "myAlgorithm",
      tabRefresh: {
        myMenu: true,
        pubMenu: false,
        preMenu: false,
      },
      algorithm: false
    }
  },
  created() {
    if (this.$route.query.data === undefined) {
      this.algorithm = false
    } else if (this.$route.query.data.algorithm) {
      this.algorithm = true
    }
  },
  methods: {
    handleClick(tab, event) {
      this.activeName = tab.name
      switch (this.activeName) {
        case 'myAlgorithm':
          this.switchTab('myMenu')
          break
        case 'publicAlgorithm':
          this.switchTab('pubMenu')
          this.algorithm = false
          break
        default:
          this.switchTab('preMenu')
          this.algorithm = false
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
    margin: 20px!important;
    background-color:#fff;
    padding: 20px;
    min-height: 800px;
  }
</style>