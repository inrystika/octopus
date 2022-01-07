<template>
  <div>
    <div class="searchForm">
      <searchForm :search-form="searchForm" :blur-name="'数据集名称 搜索'" @searchData="getSearchData" />
    </div>
    <div class="index">
      <el-table v-loading="loading" :data="datasetList" style="width: 100%;font-size: 15px;"
        :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}">
        <el-table-column label="数据集名称">
          <template slot-scope="scope">
            <span>{{ scope.row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column label="数据类型">
          <template slot-scope="scope">
            <span>{{ scope.row.typeDesc }}</span>
          </template>
        </el-table-column>
        <el-table-column label="标注类型">
          <template slot-scope="scope">
            <span>{{ scope.row.applyDesc }}</span>
          </template>
        </el-table-column>
        <el-table-column label="最新版本号">
          <template slot-scope="scope">
            <span>{{ scope.row.latestVersion }}</span>
          </template>
        </el-table-column>
        <el-table-column label="数据集描述" :show-overflow-tooltip="true">
          <template slot-scope="scope">
            <span>{{ scope.row.desc }}</span>
          </template>
        </el-table-column>
        <el-table-column label="所属群组">
          <template slot-scope="scope">
            <span>{{ scope.row.spaceId === 'default-workspace' ? '默认群组' : scope.row.spaceName }}</span>
          </template>
        </el-table-column>
        <el-table-column label="提供者">
          <template slot-scope="scope">
            <span>{{ scope.row.userName }}</span>
          </template>
        </el-table-column>
        <el-table-column label="创建时间">
          <template slot-scope="scope">
            <span>{{ scope.row.createdAt | parseTime }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button type="text" @click="getVersionList(scope.$index, scope.row)">版本列表</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="block">
      <el-pagination :current-page="searchData.pageIndex" :page-sizes="[10, 20, 50, 80]"
        :page-size="searchData.pageSize" :total="total" layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange" @current-change="handleCurrentChange" />
    </div>

    <versionList v-if="versionListVisible" :data="row" :version-list-type="versionListType" @close="close" />
  </div>
</template>

<script>
  import versionList from "./components/versionList.vue"
  import searchForm from '@/components/search/index.vue'
  import { getUserDatasetList } from "@/api/dataManager"
  export default {
    name: "UserList",
    components: {
      versionList,
      searchForm
    },
    props: {
      payload: { type: Array, default: () => [] },
      dataTabType: { type: Number, default: undefined }
    },
    data() {
      return {
        input: "",
        row: {},
        versionListVisible: false,
        versionListType: 1,
        total: undefined,
        loading: false,
        datasetList: [],
        typeChange: undefined,
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
      this.getDataList(this.searchData);
    },
    methods: {
      handleSizeChange(val) {
        this.searchData.pageSize = val
        this.getDataList(this.searchData)
      },
      handleCurrentChange(val) {
        this.searchData.pageIndex = val
        this.getDataList(this.searchData)
      },
      getDataList(param) {
        this.typeChange = this.dataTabType
        getUserDatasetList(param).then(response => {
          if (response.success) {
            this.datasetList = response.data.datasets;
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
        this.getDataList(this.searchData)
      },
      close(val) {
        this.editDataSet = val;
        this.versionListVisible = val;
      },
      getVersionList(index, row) {
        this.row = row;
        this.versionListVisible = true;
        this.versionListType = this.typeChange
      }  
    }
  }
</script>

<style lang="scss" scoped>
  .block {
    float: right;
    margin: 20px;
  }

  .searchForm {
    display: inline-block;
  }
</style>