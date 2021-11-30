<template>
  <div>
    <el-table
      :data="recordList"
      style="width: 100%;font-size: 15px;"
      :header-cell-style="{'text-align':'left','color':'black'}"
      :cell-style="{'text-align':'left'}"
    >
      <el-table-column label="时间">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.time) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="事件类型">
        <template slot-scope="scope">
          <span>{{ typeText[scope.row.type] }}</span>
        </template>
      </el-table-column>
      <el-table-column label="备注">
        <template slot-scope="scope">
          <span>{{ scope.row.remark }}</span>
        </template>
      </el-table-column>
    </el-table>

    <div class="block">
      <el-pagination
        :current-page="pageIndex"
        :page-sizes="[10, 20, 50, 80]"
        :page-size="pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>
<script>
import { getNotebookEventRecord } from "@/api/modelDev";
import { getErrorMsg } from '@/error/index'
import { parseTime } from '@/utils/index'
export default {
  name: 'notebookEventRecord',
  props: {
    notebookData: {
      type: Object,
      default: () => {}
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
      },
    }
  },
  created() {
    this.getNotebookEventRecord()
  },
  methods: {
    getNotebookEventRecord() {
      const params = {
        id:this.notebookData.id,
        pageIndex: this.pageIndex,
        pageSize: this.pageSize
      }
      getNotebookEventRecord(params).then(response => {
        if(response.success) {
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
      this.getNotebookEventRecord()
    },
    handleCurrentChange(val) {
      this.pageIndex = val
      this.getNotebookEventRecord()
    },
    parseTime(val) {
      return parseTime(val)
    },
    getErrorMsg(code) {
      return getErrorMsg(code)
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