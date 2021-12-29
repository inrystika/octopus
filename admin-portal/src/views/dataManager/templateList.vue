<template>
  <div>
    <div class="searchForm">
      <searchForm :search-form="searchForm" :blur-name="'数据集名称 搜索'" @searchData="getSearchData" />
    </div>
    <el-button type="primary" size="medium" class="create" @click="create">
      创建
    </el-button>
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
        <el-table-column label="创建时间">
          <template slot-scope="scope">
            <span>{{ scope.row.createdAt | parseTime }}</span>
          </template>
        </el-table-column>
        <el-table-column label="数据集描述" :show-overflow-tooltip="true">
          <template slot-scope="scope">
            <span>{{ scope.row.desc }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button type="text" @click="getVersionList(scope.$index, scope.row)">版本列表</el-button>
            <el-button style="padding-right:10px" type="text" @click="createNewVersion(scope.row)">创建新版本</el-button>
            <el-button type="text" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button slot="reference" type="text" @click="confirmDelete(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="block">
      <el-pagination :current-page="searchData.pageIndex" :page-sizes="[10, 20, 50, 80]"
        :page-size="searchData.pageSize" :total="total" layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange" @current-change="handleCurrentChange" />
    </div>
    <preDatasetCreation v-if="preDatasetVisible" @cancel="cancel" @close="close" @confirm="confirm" />
    <versionList v-if="versionListVisible" :data="data" :version-list-type="versionListType" @close="close" />
    <newVersion v-if="newVersionVisible" :row="data" @cancel="cancel" @confirm="confirm" @close="close" />
    <dataSetEdit v-if="editDataSet" :data="data" @cancel="cancel" @confirm="confirm" @close="close" />
  </div>
</template>

<script>
  import versionList from "./components/versionList.vue"
  import newVersion from "./components/newVersion.vue"
  import preDatasetCreation from "./components/preDatasetCreation.vue";
  import dataSetEdit from "./components/dataSetEdit.vue";
  import searchForm from '@/components/search/index.vue'
  import { deleteDataset, getPresetDatasetList } from "@/api/dataManager"
  import { getErrorMsg } from '@/error/index'
  export default {
    name: "TemplateList",
    components: {
      versionList,
      newVersion,
      preDatasetCreation,
      searchForm,
      dataSetEdit
    },
    props: {
      payload: { type: Array, default: () => [] },
      dataTabType: { type: Number, default: undefined }
    },
    data() {
      return {
        input: "",
        data: undefined,
        versionListVisible: false,
        editDataSet: false,
        versionListType: 1,
        total: undefined,
        newVersionVisible: false,
        preDatasetVisible: false,
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
      getErrorMsg(code) {
        return getErrorMsg(code)
      },
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
        getPresetDatasetList(param).then(response => {
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
      create() {
        this.preDatasetVisible = true;
        // this.row = {};
      },
      createNewVersion(row) {
        this.data = row;
        this.newVersionVisible = true;
      },
      confirmDelete(row) {
        this.$confirm('是否删除此数据集？？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          center: true
        }).then(() => {
          this.handleDelete(row)
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消'
          });
        });
      },
      handleDelete(row) {
        this.loading = true
        deleteDataset(row.id).then(response => {
          if (response.success) {
            this.$message.success("删除成功");
            this.loading = false
            this.getDataList(this.searchData)
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
            this.loading = false
          }
        })
      },
      close(val) {
        this.editDataSet = val;
        this.preDatasetVisible = val
        this.newVersionVisible = val
        this.versionListVisible = val    
      },
      cancel(val) {
        this.editDataSet = val;
        this.newVersionVisible = val;
        this.preDatasetVisible = val;
        this.getDataList(this.searchData)
      },
      confirm(val) {
        this.editDataSet = val;
        this.preDatasetVisible = val
        this.newVersionVisible = val
        this.getDataList(this.searchData)
      },
      getVersionList(index, row) {
        this.data = row;
        this.versionListVisible = true;
        this.versionListType = this.typeChange
      },
      handleEdit(val) {
        this.editDataSet = true
        this.data = val
      }
    }
  }
</script>

<style lang="scss" scoped>
  .create {
    float: right;
  }

  .block {
    float: right;
    margin: 20px;
  }

  .searchForm {
    display: inline-block;
  }
</style>