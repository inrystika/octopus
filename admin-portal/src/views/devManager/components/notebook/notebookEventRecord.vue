<template>
  <div>
    <el-table :data="recordList" style="width: 100%;font-size: 15px;"
      :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
      <el-table-column label="时间">
        <template slot-scope="scope">
          <span>{{ scope.row.time | parseTime }}</span>
        </template>
      </el-table-column>
      <el-table-column label="事件类型">
        <template slot-scope="scope">
          <span>{{ typeText[scope.row.type] }}</span>
        </template>
      </el-table-column>
      <el-table-column label="备注">
        <template slot-scope="scope">
          <div v-if="scope.row.remark.state">{{ "状态 : "+scope.row.remark.state}}</div>
          <div v-if="scope.row.remark.reason">{{ "原因 : "+scope.row.remark.reason}}</div>
        </template>
      </el-table-column>
    </el-table>

    <div class="block">
      <el-pagination :current-page="pageIndex" :page-sizes="[10, 20, 50, 80]" :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper" :total="total" @size-change="handleSizeChange"
        @current-change="handleCurrentChange" />
    </div>
  </div>
</template>
<script>
  import { getNotebookEventRecord } from "@/api/modelDev";
  export default {
    name: 'notebookEventRecord',
    props: {
      notebookData: {
        type: Object,
        default: () => { }
      }
    },
    data() {
      return {
        recordList: [],
        total: 0,
        pageIndex: 1,
        pageSize: 10,
        typeText: {
          1: "创建",
          2: "重新启动",
          3: "开始运行",
          4: "停止",
          5: '保存环境'
        },
      }
    },
    created() {
      this.getNotebookEventRecord()
    },
    methods: {
      getNotebookEventRecord() {
        const params = {
          id: this.notebookData.id,
          pageIndex: this.pageIndex,
          pageSize: this.pageSize
        }
        getNotebookEventRecord(params).then(response => {
          if (response.success) {
            if (response.data.records == null) { response.data.records = [] }
            response.data.records.forEach(
              item => {
                if (item.remark !== '') {
                  item.remark = JSON.parse(item.remark)
                }

              }
            )
            this.recordList = response.data.records
            this.total = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            })
          }
        })
      },
      handleSizeChange(val) {
        this.pageSize = val
        this.pageIndex = 1
        this.getNotebookEventRecord()
      },
      handleCurrentChange(val) {
        this.pageIndex = val
        this.getNotebookEventRecord()
      }
    }
  }
</script>
<style lang="scss" scoped>
  .block {
    float: right;
    margin: 20px;
  }
</style>