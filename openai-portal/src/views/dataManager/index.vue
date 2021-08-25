<template>
  <div>
    <el-tabs v-model="activeName" class="Wrapper" @tab-click="handleClick">
      <el-tab-pane label="我的数据集" name="myDataset">
        <myList :dataset="dataset" :Type=1 v-if="tabRefresh.myMenu"></myList>
      </el-tab-pane>
      <el-tab-pane label="公共数据集" name="publicDataset">
        <publicList :Type=2 v-if="tabRefresh.pubMenu"></publicList>
      </el-tab-pane>
      <el-tab-pane label="预置数据集" name="preDataset">
        <presetList :Type=3 v-if="tabRefresh.preMenu"></presetList>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import myList from "./myList.vue";
import publicList from "./publicList.vue"
import presetList from "./presetList.vue"
export default {
  components: {
    myList,
    publicList,
    presetList
  },
  data() {
    return {
      activeName: "myDataset",
      tabRefresh: {
        myMenu: true,
        pubMenu: false,
        preMenu: false,
      },
      dataset: false
    };
  },
  created(){
    if (this.$route.params.data === undefined) {
      this.dataset = false
    } else if (this.$route.params.data.dataset) {
      this.dataset = true
    }
  },
  methods: {
    handleClick(tab, event) {
      this.activeName = tab.name
      switch (this.activeName) {
        case 'myDataset':
          this.switchTab('myMenu')
          break
        case 'publicDataset':
          this.switchTab('pubMenu')
          this.dataset = false
          break
        default:
          this.switchTab('preMenu')
          this.dataset = false
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
  },
}
;
</script>

<style lang="scss" scoped>
.Wrapper {
    margin: 20px!important;
    background-color:#fff;
    padding: 20px;
    min-height: 800px;
  }
</style>