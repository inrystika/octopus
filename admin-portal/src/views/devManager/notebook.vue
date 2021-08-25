<template>
  <div>
    <el-tabs v-model="activeName" class="Wrapper">
      <el-tab-pane label="NoteBook列表" name="notebookList">
        <notebookList></notebookList>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import notebookList from "./components/notebook/notebookList.vue";
export default {
  components: {
    notebookList,  
  },
  data() {
    return {
      activeName: "notebookList",
    }
  },
  methods: {

  }}
</script>

<style scoped>
  .Wrapper {
    margin: 15px!important;
    background-color:#fff;
    padding: 20px;
    min-height: 800px;
  }
</style>