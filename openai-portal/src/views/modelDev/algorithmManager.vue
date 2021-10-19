<template>
    <div>
      <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleClick">
        <el-tab-pane label="我的算法" name="myAlgorithm">
          <myList v-if="tabRefresh.myMenu" :algorithm="algorithm" :algorithm-tab-type="1" />
        </el-tab-pane>
        <el-tab-pane label="公共算法" name="publicAlgorithm">
          <publicList v-if="tabRefresh.pubMenu" :algorithm-tab-type="2" />
        </el-tab-pane>
        <el-tab-pane label="预置算法" name="preAlgorithm">
          <presetList v-if="tabRefresh.preMenu" :algorithm-tab-type="3" />
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
        preMenu: false
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

<style lang="scss" scoped>
.Wrapper {
    margin: 20px!important;
    background-color:#fff;
    padding: 20px;
    min-height: 900px;
  }
</style>