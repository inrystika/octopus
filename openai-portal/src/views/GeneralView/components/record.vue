<template>
  <div>
    <el-dialog
      :title="recordType === 1 ? '机时消费记录' : '机时充值记录'"
      width="70%"
      :visible.sync="createFormVisible"
      :before-close="handleDialogClose"
    >
      <el-popover style="float: right" placement="left-end" title="计时规则" width="500" trigger="hover">
        <div>
          <div style="padding-bottom:2%">
            任务 <code>i</code> 机时 = 子任务 <code>1</code> 机时 + 子任务 <code>2</code> 机时 + …… + 子任务 <code>n</code> 机时;<br>
          </div>
          <div style="padding-bottom:2%">
            子任务 <code>i</code> 机时 = 副本 <code>1</code> 机时 + 副本 <code>2</code> 机时 + …… + 副本 <code>n</code> 机时;<br>
          </div>
          <div>
            副本 <code>k</code> 机时 = 资源规格价格 * 运行时间;
          </div>
        </div>
        <el-button v-if="recordType === 1 ? true : false" slot="reference" type="text">计时规则</el-button>
      </el-popover>
      <el-table :data="recordList" :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
        <div v-if="recordType === 1 ? true : false">
          <el-table-column label="类型" align="center" prop="type">
            <template slot-scope="scope">
              <span>{{ changeType(scope.row.bizType) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="名称" align="center" prop="name">
            <template slot-scope="scope">
              <span>{{ scope.row.Title }}</span>
            </template>
          </el-table-column>
          <el-table-column label="消费机时(h)" align="center" prop="consumption">
            <template slot-scope="scope">
              <span>{{ scope.row.amount }}</span>
            </template>
          </el-table-column>
          <el-table-column label="开始时间" align="center" prop="startTime">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.startedAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="结束时间" align="center" prop="endTime">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.endedAt) }}</span>
            </template>
          </el-table-column>
        </div>
        <div v-if="recordType === 2 ? true : false">
          <el-table-column label="充值时间" align="center" prop="rechargeTime">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.updatedAt) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="充值机时(h)" align="center" prop="rechargeHour">
            <template slot-scope="scope">
              <span>{{ scope.row.amount }}</span>
            </template>
          </el-table-column>
        </div>
      </el-table>
      <div class="pagination">
        <el-pagination
          :current-page="pageIndex"
          :page-sizes="[10, 20, 50]"
          :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
      <div slot="footer" />
    </el-dialog>
  </div>
</template>

<script>
import { parseTime } from '@/utils/index'
import { getUserConsumptionRecord, getGroupConsumptionRecord, getUserRechargeRecord, getGroupRechargeRecord } from "@/api/generalView";
import { getErrorMsg } from '@/error/index'
export default {
  name: "Record",
  props: {
    groupName: {
      type: String,
      default: " "
    },
    recordType: {
      type: Number,
      default: 1
    }
  },
  data() {
    return {
      createFormVisible: true,
      recordList: {},
      pageIndex: 1,
      pageSize: 20,
      total: undefined
    }
  },
  created() {
    if (this.recordType === 1) {
      this.getConsumptionRecord()
    } else {
      this.getRechargeRecord()
    }
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    changeType(value) {
      switch (value) {
        case 1:
          return '训练'
        default:
          return 'notebook'
      }
    },
    handleSizeChange(val) {
      this.pageSize = val
      this.getConsumptionRecord()
    },
    handleCurrentChange(val) {
      this.pageIndex = val
      this.getConsumptionRecord()
    },
    getConsumptionRecord() {
      const param = {
        pageIndex: this.pageIndex,
        pageSize: this.pageSize
      }
      if (this.groupName === "default-workspace") {
        getUserConsumptionRecord(param).then(response => {
          if (response.success) {
            this.recordList = response.data.records
            this.total = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            })
          }
        })
      } else {
        getGroupConsumptionRecord(param).then(response => {
          if (response.success) {
            this.recordList = response.data.records
            this.total = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            })
          }
        })
      }
    },
    getRechargeRecord() {
      const param = {
        pageIndex: this.pageIndex,
        pageSize: this.pageSize
      }
      if (this.groupName === "default-workspace") {
        getUserRechargeRecord(param).then(response => {
          if (response.success) {
            this.recordList = response.data.records
            this.total = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            })
          }
        })
      } else {
        getGroupRechargeRecord(param).then(response => {
          if (response.success) {
            this.recordList = response.data.records
            this.total = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            })
          }
        })
      }
    },
    handleDialogClose() {
      this.$emit("close", false);
    },
    // 时间戳转换日期
    parseTime(val) {
      return parseTime(val)
    }
  }
};
</script>

<style lang="scss" scoped>
  .pagination {
    float: right;
    margin: 20px;
  }
</style>