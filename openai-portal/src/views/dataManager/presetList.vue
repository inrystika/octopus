<template>
  <div>
    <div class="searchForm">
      <searchForm 
        :searchForm=searchForm 
        @searchData="getSearchData" 
        :blurName="'数据集名称 搜索'">
      </searchForm>
    </div>
    <div class="index">
      <el-table :data="datasetList" style="width: 100%;font-size: 15px;"
        :header-cell-style="{'text-align':'left','color':'black'}" :cell-style="{'text-align':'left'}"
        v-loading="loading">
        <el-table-column label="名称">
          <template slot-scope="scope">
            <span>{{ scope.row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column label="类型">
          <template slot-scope="scope">
            <span>{{ scope.row.type }}</span>
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
            <span>{{ parseTime(scope.row.createdAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button type="text" @click="getVersionList(scope.$index, scope.row)" style="padding-right:10px">版本列表
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="block">
      <el-pagination 
        @size-change="handleSizeChange" 
        @current-change="handleCurrentChange" 
        :current-page="searchData.pageIndex"
        :page-sizes="[10, 20, 50, 80]" 
        :page-size="searchData.pageSize" 
        layout="total, sizes, prev, pager, next, jumper"
        :total="total">
      </el-pagination>
    </div>

    <versionList v-if="versionListVisible" @cancel="cancel" @confirm="confirm"
      @close="close" :data="this.row" :typeChange="this.typeChange">
    </versionList>
  </div>
</template>

<script>
  import versionList from "./components/versionList.vue";
  import searchForm from '@/components/search/index.vue'
  import { getPresetDatasetList } from "@/api/datasetManager";
  import { parseTime } from '@/utils/index'
  import { getErrorMsg } from '@/error/index'
  export default {
    name: "presetList",
    components: {
      versionList,
      searchForm
    },
    props: {
      Type: { type: Number }
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
        ],
        searchData: {
          pageIndex: 1,
          pageSize: 10,
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
        this.typeChange = this.Type
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
        this.searchData={pageIndex:1,pageSize:this.searchData.pageSize}
        this.searchData = Object.assign(val, this.searchData)
        if (this.searchData.time) {
          this.searchData.createdAtGte = this.searchData.time[0]/1000
          this.searchData.createdAtLt = this.searchData.time[1]/1000
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
      //时间戳转换日期
      parseTime(val) {
        return parseTime(val)
      }
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