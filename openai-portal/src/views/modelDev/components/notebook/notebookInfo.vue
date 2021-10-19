<template>
  <div>
    <el-dialog
      title="启动信息"
      width="80%"
      :visible.sync="infoVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
    <div>
      <el-input
        v-if="showInfo"
        v-model="subTaskInfo"
        type="textarea"
        :readonly="true"
        :autosize="true"
      />
    </div>

    <div class="block">
      <el-pagination
        v-if="showInfo"
        :current-page="pageIndex"
        :page-sizes="[10, 20, 50, 80]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>

    <div slot="footer" class="dialog-footer" />
    </el-dialog>
  </div>
</template>

<script>
import { getNotebookInfo } from "@/api/modelDev";
export default {
  name: "NotebookInfo",
  props: {
    notebookData: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      showInfo: false,
      infoVisible: true,
      subTaskInfo: "",
      total: 0,
      pageIndex: 1,
      pageSize: 10,
      taskIndex: 1,
      replicaIndex: 1
    }
  },
  created() {
    this.getNotebookInfo()
  },
  methods: {
    getNotebookInfo(){
      const param = {
        id: this.notebookData.notebookJobId,
        pageIndex: this.pageIndex,
        pageSize: this.pageSize,
        taskIndex: this.taskIndex,
        replicaIndex: this.replicaIndex
      }
      getNotebookInfo(param).then(response => {
        if (!response.success) {
          this.$message({
            message: "暂无相关运行信息",
            type: 'warning'
          });
          return
        }

        this.showInfo = response.payload.notebookEvents.length
        this.total = response.payload.totalSize

        let infoMessage = ""

        response.payload.notebookEvents.forEach(function(element) {
          const title = element.reason
          const message = element.message
          infoMessage += "\n" + "[" + title + "]"
          infoMessage += "\n" + "[" + message + "]" + "\n"
        })

        this.subTaskInfo = infoMessage
      }).catch(err => {
        console.log("err:", err)
        this.$message({
          message: "未知错误",
          type: 'warning'
        });
      });
    },
    handleDialogClose() {
      this.$emit('close', false)
    },
    handleSizeChange(val) {
      this.pageSize = val
      this.getNotebookInfo()
    },
    handleCurrentChange(val) {
      this.pageIndex = val
      this.getNotebookInfo()
    }
  },
}
</script>

<style lang="scss" scoped>
  .block {
    float: right;
    margin: 20px;
  }
</style>