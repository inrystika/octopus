<template>
  <div>
    <el-tabs class="Wrapper">
      <el-tab-pane label="训练任务">
        <el-table
          :data="trainJobList"
          style="width: 100%;font-size: 15px;"
          :header-cell-style="{'text-align':'left','color':'black'}"
          :cell-style="{'text-align':'left'}"
        >
          <el-table-column label="任务名称">
            <template slot-scope="scope">
              <span>{{ scope.row.taskName }}</span>
            </template>
          </el-table-column>
          <el-table-column label="学习框架">
            <template slot-scope="scope">
              <span>{{ scope.row.framework }}</span>
            </template>
          </el-table-column>
          <el-table-column label="解释器">
            <template slot-scope="scope">
              <span>{{ scope.row.interpreter }}</span>
            </template>
          </el-table-column>
          <el-table-column label="状态">
            <template slot-scope="scope">
              <span :class="statusOption[scope.row.status][0]"></span>
              <span>{{ statusOption[scope.row.status][1] }}</span>
            </template>
          </el-table-column>
          <el-table-column label="创建时间">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.createTime) }}</span>
            </template>
          </el-table-column>
        </el-table>

          <div class="block">
            <el-pagination
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
              :current-page="searchData.pageIndex"
              :page-sizes="[10, 20, 50, 80]"
              :page-size="searchData.pageSize"
              layout="total, sizes, prev, pager, next, jumper"
              :total="total"
            >
            </el-pagination>
          </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import { parseTime } from '@/utils/index'
import { getErrorMsg } from '@/error/index'
import { getCloudTrainJobList } from "@/api/cloudInterconnection"
export default {
  name: "cloudInterconnection",
  data() {
    return {
      searchData: {
        pageIndex: 1,
        pageSize: 10,
      },
      trainJobList: [],
      total: 0,
      statusOption: {
        '0': ['status-ready', '初始中'],
        '1': ['status-agent', '已分派'],
        '2': ['status-running', '分中心处理中'],
        '-1': ['status-danger', '失败'],
        '3': ['status-success', '成功'],
        '4': ['status-stopping', '停止'],
        '7': ['status-reassign', '重分派'],
      }
    }
  },
  created(){
    this.getCloudTrainJobList(this.searchData)
  },
  methods: {
    getCloudTrainJobList(params){
      getCloudTrainJobList(params).then(response => {
        if (response.success) {
            this.trainJobList = response.data.list;
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
      this.searchData.pageSize = val
      this.getCloudTrainJobList(this.searchData)
    },
    handleCurrentChange(val) {
      this.searchData.pageIndex = val
      this.getCloudTrainJobList(this.searchData)
    },
    parseTime(val) {
      return parseTime(val)
    },
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
  }
}
</script>
<style lang="scss" scoped>
  .Wrapper {
    margin: 15px!important;
    background-color:#fff;
    padding: 20px;
    min-height: 900px
  }
  .create {
    float: right;
  }
  .block {
    float: right;
    margin: 20px;
  }
</style>