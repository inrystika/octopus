<template>
  <div>
    <el-dialog title="详情" width="80%" :visible.sync="CreateFormVisible" :before-close="handleDialogClose"
      :append-to-body="true" custom-class="dialog" :close-on-click-modal="false">
      <el-tabs>
        <el-tab-pane label="任务简况">
          <notebookProfile :notebook-data="notebookData" />
        </el-tab-pane>
        <el-tab-pane label="运行信息">
          <notebookInfo :notebook-data="notebookData" />
        </el-tab-pane>
        <el-tab-pane label="事件记录">
          <notebookEventRecord :notebook-data="notebookData" />
        </el-tab-pane>
        <el-tab-pane label="退出信息">
          <notebookExit :notebook-data="notebookData" />
        </el-tab-pane>
      </el-tabs>
    </el-dialog>
  </div>
</template>
<script>
  import notebookInfo from "./notebookInfo.vue"
  import notebookProfile from "./notebookProfile.vue"
  import notebookEventRecord from "./notebookEventRecord.vue"
  import notebookExit from "./notebookExit.vue"
  export default {
    name: "DetailDialog",
    components: {
      notebookInfo,
      notebookProfile,
      notebookEventRecord,
      notebookExit
    },
    props: {
      detailData: {
        type: Object,
        default: () => { }
      }
    },
    data() {
      return {
        CreateFormVisible: true,
        notebookData: {}
      }
    },
    created() {
      this.notebookData = this.detailData
    },
    methods: {
      handleDialogClose() {
        this.$emit('close', false)
      }
    }
  }
</script>
<style lang="scss">
  .dialog {
    min-height: 800px
  }
</style>