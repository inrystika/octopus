<template>
  <div>
    <div>
      <el-row>
        <el-col :span="12">
          <div>任务名称:<span>{{ notebookInfo.name }}</span></div>
        </el-col>
        <el-col :span="12">
          <div>是否分布式:<span>{{ this.notebookInfo.tasks.length > 1 ? '是' : '否'}}</span></div>
        </el-col>
        <el-col v-if="show" :span="12">
          <div>
            子任务名:
            <el-select v-model="taskIndex" placeholder="请选择" class="select" @change="selectLog">
              <el-option v-for="item in options" :key="item.name" :label="item.name" :value="item.value" />
            </el-select>
          </div>
        </el-col>
      </el-row>
      <el-input v-model="subTaskInfo" type="textarea" :readonly="true" :rows="20" />
    </div>

    <div class="block">
      <el-pagination :current-page="pageIndex" :page-sizes="[10, 20, 50, 80]" :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper" :total="total" @size-change="handleSizeChange"
        @current-change="handleCurrentChange" />
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
        default: () => { }
      }
    },
    data() {
      return {
        notebookInfo: {},
        subTaskInfo: "",
        total: 0,
        pageIndex: 1,
        pageSize: 10,
        taskIndex: '',
        replicaIndex: 1,
        options: [],
      }
    },
    computed: {
      show: function () {
        if (this.notebookInfo.tasks.length > 1) {
          return true
        } else {
          return false
        }
      }
    },
    created() {
      this.notebookInfo = this.notebookData
   
      if (!this.show) {
        this.taskIndex = 1
        this.getNotebookInfo()
      }
      else {
        this.options = []
        let i = 1
        this.notebookInfo.tasks.forEach(
          item => {
            this.options.push({ name: item.name, value: i++ })
          }
        )
      }

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
          if (response.success) {
            this.total = response.payload.totalSize
            let infoMessage = ""
            response.payload.notebookEvents.forEach(function (element) {
              const title = element.reason
              const message = element.message
              infoMessage += "\n" + "[" + title + "]"
              infoMessage += "\n" + "[" + message + "]" + "\n"
            })
            this.subTaskInfo = infoMessage
          } else {
            const data = {
              id: this.notebookInfo.notebookJobId,
              pageIndex: this.pageIndex,
              pageSize: this.pageSize,
              taskIndex: 0,
              replicaIndex: 0
            }
            getNotebookInfo(data).then(response => {
              if (response.success) {
                this.total = response.payload.totalSize
                let vcMessage = ""
                response.payload.notebookEvents.forEach(function (element) {
                  const title = element.reason
                  const message = element.message
                  vcMessage += "\n" + "[" + title + "]"
                  vcMessage += "\n" + "[" + message + "]" + "\n"
                })
                this.subTaskInfo = vcMessage
              } else {
                this.subTaskInfo = "暂无相关运行信息"
              }
            })
          }
        }).catch(err => {
          this.$message({
            message: "未知错误",
            type: 'warning'
          });
        });
      },
      handleSizeChange(val) {
        this.pageSize = val
        this.getNotebookInfo()
      },
      handleCurrentChange(val) {
        this.pageIndex = val
        this.getNotebookInfo()
      },
      selectLog() {
        this.getNotebookInfo()
      }
    },
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


  .block {
    float: right;
    margin: 20px;
  }
</style>