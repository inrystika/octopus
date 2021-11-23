<template>
  <div>
    <el-table 
      :data="taskList" 
      style="width: 100%;font-size: 15px;"
      :header-cell-style="{'text-align':'left','color':'black'}" 
      :cell-style="{'text-align':'left'}"
    >
      <el-table-column label="名称" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="数据集" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.datasets }}</span>
        </template>
      </el-table-column>
      <el-table-column label="联系方式" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.contactInfo }}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createdAt) }}</span>
        </template>
      </el-table-column>
    </el-table>
    <div class="block">
      <el-pagination
        :current-page="searchData.pageIndex" 
        :page-sizes="[10, 20, 50, 80]" 
        :page-size="searchData.pageSize"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper" 
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>


  </div>
</template>
<script>
import searchForm from '@/components/search/index.vue'
import { getPlatformTrainingTaskList } from "@/api/platformManager"
import { parseTime } from '@/utils/index'
import { getErrorMsg } from '@/error/index'
export default {
  name: "platformTrainingTaskList",
  components: {
    searchForm
  },
  data() {
    return {
      taskList: [],
      total: 0,
      searchData: {
        pageIndex: 1,
        pageSize: 10,
      }
    }
  },
  created() {
    this.getPlatformTrainingTaskList(this.searchData);
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    getPlatformTrainingTaskList(param){
      getPlatformTrainingTaskList(param).then(response => {
        console.log("res:",response.data.trainJobs)
        if(response.success){
          this.taskList = response.data.trainJobs;
          this.total = response.data.totalSize
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          });
        }    
      })
    },
    handleSizeChange(val){
      this.searchData.pageSize = val
      this.getPlatformTrainingTaskList(this.searchData)
    },
    handleCurrentChange(val) {
      this.searchData.pageIndex = val
      this.getPlatformTrainingTaskList(this.searchData)
    },
    //时间戳转换日期
    parseTime(val) {
      return parseTime(val)
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