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
      <el-table-column label="框架类型">
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
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button type="text" @click="getAlgorithmVersionList(scope.row)">版本列表</el-button>
          <el-button type="text" @click="copyAlgorithm(scope.row)">复制算法</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination">
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

    <algorithmCopy
      v-if="algorithmCopyVisible"
      :row="row"
      :algorithm-tab-type="typeChange"
      @close="close"
      @cancel="cancel"
      @confirm="confirm"
    />
    <versionList
      v-if="versionListVisible"
      :algorithm-tab-type="typeChange"
      :data="row"
      @close="close"
    />
  </div>
</template>

<script>
import algorithmCopy from "./algorithmCopy.vue";
import versionList from "./versionList.vue";
import searchForm from '@/components/search/index.vue'
import { getPresetAlgorithmList } from "@/api/modelDev"
export default {
  name: "PresetList",
  components: {
    algorithmCopy,
    versionList,
    searchForm
  },
  props: {
    algorithmTabType: { type: Number, default: undefined }
  },
  data() {
    return {
      row: {},
      algorithmCopyVisible: false,
      versionListVisible: false,
      total: undefined,
      typeChange: undefined,
      algorithmList: [],
      searchForm: [
        { type: 'Time', label: '创建时间', prop: 'time', placeholder: '请选择创建时间' },
        { type: 'Input', label: '算法名称', prop: 'nameLike', placeholder: '请输入算法名称' }
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
      getPresetAlgorithmList(param).then(response => {
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
    getAlgorithmVersionList(row) {
      this.versionListVisible = true;
      this.typeChange = this.algorithmTabType
      this.row = row
    },
    close(val) {
      this.algorithmCopyVisible = val;
      this.versionListVisible = val;
      this.getAlgorithmList(this.searchData)
    },
    cancel(val) {
      this.algorithmCopyVisible = val;
      this.getAlgorithmList(this.searchData)
    },
    copyAlgorithm(row) {
      this.algorithmCopyVisible = true;
      this.row = row
    },
    confirm(val) {
      this.algorithmCopyVisible = val
      this.getAlgorithmList(this.searchData)
    }
  }
}
</script>

<style lang="scss" scoped>
  .pagination {
    float: right;
    margin: 20px;
  }
  .searchForm {
    display: inline-block;
  }
</style>