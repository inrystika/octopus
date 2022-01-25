<template>
  <div>
    <div class="searchForm">
      <searchForm
        :search-form="searchForm"
        :blur-name="'算法名称/描述 搜索'"
        @searchData="getSearchData"
      />
    </div>
    <el-table
      :data="algorithmList"
      style="width: 100%;font-size: 15px;"
      :header-cell-style="{'text-align':'left','color':'black'}"
      :cell-style="{'text-align':'left'}"
    >
      <el-table-column label="算法名称">
        <template slot-scope="scope">
          <span>{{ scope.row.algorithmName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="模型名称">
        <template slot-scope="scope">
          <span>{{ scope.row.modelName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="当前版本号">
        <template slot-scope="scope">
          <span>{{ scope.row.algorithmVersion }}</span>
        </template>
      </el-table-column>
      <el-table-column label="模型类别">
        <template slot-scope="scope">
          <span>{{ scope.row.applyName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="算法框架">
        <template slot-scope="scope">
          <span>{{ scope.row.frameworkName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="算法描述" :show-overflow-tooltip="true">
        <template slot-scope="scope">
          <span>{{ scope.row.algorithmDescript }}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间">
        <template slot-scope="scope">
          <span>{{ scope.row.createdAt | parseTime }}</span>
        </template>
      </el-table-column>
      <el-table-column label="群组">
        <template slot-scope="scope">
          <span>{{ scope.row.spaceId === 'default-workspace' ? '默认群组' : scope.row.spaceName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="提供者">
        <template slot-scope="scope">
          <span>{{ scope.row.userName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button type="text" @click="getAlgorithmVersionList(scope.row)">版本列表</el-button>
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

    <versionList
      v-if="versionListVisible"
      :row="row"
      :algorithm-type="typeChange"
      @close="close"
    />
  </div>
</template>

<script>
import { getUserAlgorithmList } from "@/api/modelDev"
import versionList from "./versionList.vue"
import searchForm from '@/components/search/index.vue'
export default {
  name: "UserList",
  components: {
    versionList,
    searchForm
  },
  props: {
    algorithmTabType: { type: Number, default: undefined }
  },
  data() {
    return {
      row: {},
      total: undefined,
      versionListVisible: false,
      algorithmName: "",
      typeChange: undefined,
      algorithmList: [],
      searchForm: [
         { type: 'Time', label: '创建时间', prop: 'time', placeholder: '请选择创建时间' }
      ],
      searchData: {
        pageIndex: 1,
        pageSize: 10
      }
    }
  },
  created() {
    this.getAlgorithmList(this.searchData);
  },
  methods: {
    handleSizeChange(val) {
      this.searchData.pageSize = val
      this.getAlgorithmList(this.searchData)
    },
    handleCurrentChange(val) {
      this.searchData.pageIndex = val
      this.getAlgorithmList(this.searchData)
    },
    getAlgorithmList(param) {
      this.typeChange = this.algorithmTabType
      getUserAlgorithmList(param).then(response => {
        if (response.success) {
          this.algorithmList = response.data.algorithms;
          this.total = response.data.totalSize
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          });
        }
      })
    },
    getSearchData(val) {
      this.searchData = { pageIndex: 1, pageSize: this.searchData.pageSize }
      this.searchData = Object.assign(val, this.searchData)
      if (this.searchData.time) {
        this.searchData.createdAtGte = this.searchData.time[0] / 1000
        this.searchData.createdAtLt = this.searchData.time[1] / 1000
        delete this.searchData.time
      }
      this.getAlgorithmList(this.searchData)
    },
    getAlgorithmVersionList(row) {
      this.versionListVisible = true;
      this.row = row
    },
    close(val) {
      this.versionListVisible = val
      this.getAlgorithmList(this.searchData);
    }
  }
}
</script>

<style lang="scss" scoped>
  .Wrapper {
    margin: 15px!important;
  }
  .block {
    float: right;
    margin: 20px;
  }
  .searchForm {
    display: inline-block;
  }
</style>