<template>
  <div>
    <el-tabs v-model="activeName" class="Wrapper">
      <el-tab-pane label="NoteBook" name="myNoteBook">
        <notebookList :notebook="notebook"></notebookList>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import notebookList from "./components/notebook/notebookList.vue";
export default {
  components: {
    notebookList
  },
  data() {
    return {
      activeName: "myNoteBook",
      notebook: false
    }
  },
  created() {
    if (this.$route.query.data === undefined) {
      this.notebook = false
    } else if (this.$route.query.data.notebook) {
      this.notebook = true
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