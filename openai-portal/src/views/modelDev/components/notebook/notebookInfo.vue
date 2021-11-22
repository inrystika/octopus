<template>
  <div>
    <div>
      <el-row>
        <el-col :span="12">
          <div>任务名称:<span>{{ notebookInfo.name }}</span></div>
        </el-col>
        <el-col :span="12">
          <div>是否分布式:<span>否</span></div>
        </el-col>
      </el-row>
      <el-input
        v-if="showInfo"
        v-model="subTaskInfo"
        type="textarea"
        :readonly="true"
        :rows="20"
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
      notebookInfo: {},
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
    this.notebookInfo = this.notebookData
    this.getNotebookInfo()
  },
  methods: {
    getNotebookInfo() {
      const param = {
        id: this.notebookInfo.notebookJobId,
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
  }
}
</script>

<style lang="scss" scoped>
  .el-col {
    margin: 10px 0 20px 0;
    font-size: 15px;
    font-weight: 800;

    span {
      font-weight: 400;
      margin-left: 20px
    }
  }

  // .select {
  //   margin-left: 5px;
  // }

  .block {
    float: right;
    margin: 20px;
  }
</style>