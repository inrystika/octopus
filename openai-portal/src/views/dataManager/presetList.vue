<template>
  <div>
    <div class="searchForm">
      <searchForm
        :search-form="searchForm"
        :blur-name="'数据集名称 搜索'"
        @searchData="getSearchData"
      />
    </div>
    <div class="index">
      <el-table
        v-loading="loading"
        :data="datasetList"
        style="width: 100%;font-size: 15px;"
        :header-cell-style="{'text-align':'left','color':'black'}"
        :cell-style="{'text-align':'left'}"
      >
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
        <el-table-column label="创建时间">
          <template slot-scope="scope">
            <span>{{ scope.row.createdAt | parseTime }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button type="text" style="padding-right:10px" @click="getVersionList(scope.$index, scope.row)">
              版本列表
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="block">
      <el-pagination
        :current-page="searchData.pageIndex"
        :page-sizes="[10, 20, 50, 80]"
        :page-size="searchData.pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>

    <versionList
      v-if="versionListVisible"
      :data="row"
      :type-change="typeChange"
      @cancel="cancel"
      @confirm="confirm"
      @close="close"
    />
  </div>
</template>

<script>
  import versionList from "./components/versionList.vue";
  import searchForm from '@/components/search/index.vue'
  import { getPresetDatasetList } from "@/api/datasetManager";
  import { getErrorMsg } from '@/error/index'
  export default {
    name: "PresetList",
    components: {
      versionList,
      searchForm
    },
    props: {
        dataType: {
        type: Number,
        default: undefined
      }
    },

    data() {
      return {
        input: "",
        row: {},
        formLabelWidth: "120px",
        versionListVisible: false,
        datasetList: [],
        total: undefined,
        loading: false,
        typeChange: undefined,
        searchForm: [
          { type: 'Time', label: '创建时间', prop: 'time', placeholder: '请选择创建时间' },
          { type: 'Input', label: '数据集名称', prop: 'nameLike', placeholder: '请输入数据集名称' }
        ],
        searchData: {
          pageIndex: 1,
          pageSize: 10
        }
      };
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
        this.typeChange = this.dataType
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
      getVersionList(index, row) {
        this.row = row;
        this.versionListVisible = true;
      },
      cancel(val) {
        this.versionListVisible = val;
      },
      confirm(val) {
        this.versionListVisible = val;
        this.getDataList(this.searchData)
      },
      close(val) {
        this.versionListVisible = val;
        this.getDataList(this.searchData)
      },
    }
  };
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