<template>
  <div>
    <el-tabs class="Wrapper">
      <el-tab-pane label="训练任务">
        <div class="create">
            <el-button type="primary" @click="create">创建训练任务</el-button>
        </div>
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
          <el-table-column label="训练引擎框架">
            <template slot-scope="scope">
              <span>{{ scope.row.framework }}</span>
            </template>
          </el-table-column>
          <el-table-column label="训练引擎解释器">
            <template slot-scope="scope">
              <span>{{ scope.row.interpreter }}</span>
            </template>
          </el-table-column>
          <el-table-column label="状态">
            <template slot-scope="scope">
              <span>{{ statusText[scope.row.status] }}</span>
            </template>
          </el-table-column>
          <el-table-column label="创建时间">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.createTime) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作">
            <el-button>停止</el-button>
          </el-table-column>

        </el-table>
      </el-tab-pane>
    </el-tabs>

    <trainingTaskCreate 
      v-if="createVisible"
      @cancel="cancel"
      @confirm="confirm"
      @close="close"
    />
  </div>
</template>
<script>
import { parseTime } from '@/utils/index'
import { getErrorMsg } from '@/error/index'
import trainingTaskCreate from './trainingTaskCreate.vue'
import { getCloudTrainJobList } from "@/api/cloudInterconnection"
export default {
  name: "cloudInterconnection",
  components: {
    trainingTaskCreate
  },
  data() {
    return {
      createVisible: false,
      searchData: {
        pageIndex: 1,
        pageSize: 10,
      },
      trainJobList: [],
      total: undefined,
      statusText: {
        0: '初始化',
        1: '已分派',
        2: '分中心处理中',
        7: '重新分派',
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
    create() {
      this.createVisible = true
    },
    close(val) {
      this.createVisible = val;
      this.getCloudTrainJobList(this.searchData)
    },
    cancel(val) {
      this.createVisible = val;
      this.getCloudTrainJobList(this.searchData)
    },
    confirm(val) {
      this.createVisible = val
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
</style>