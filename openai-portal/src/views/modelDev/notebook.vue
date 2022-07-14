<template>
  <div>
    <el-tabs v-model="activeName" class="Wrapper">
      <el-tab-pane label="NoteBook" name="myNoteBook">
        <notebookList :notebook="notebook" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
  import notebookList from "./components/notebook/notebookList.vue";
  import { clearProgress } from '@/utils/index.js'
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
    },
    mounted() {
      window.addEventListener("beforeunload", (e) => {
        clearProgress()
      });
    },
    destroyed() {
      window.removeEventListener("beforeunload", (e) => {
        clearProgress()
      });
    },
  }
</script>

<style lang="scss" scoped>
  .Wrapper {
    margin: 15px !important;
    background-color: #fff;
    padding: 20px;
    min-height: 900px;
  }
</style>