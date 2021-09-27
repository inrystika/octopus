<template>
  <div>
    <div class="searchForm">
      <searchForm
        :searchForm="searchForm"
        :blurName="'算法名称/描述 搜索'"
        @searchData="getSearchData"
      />
    </div>
      <el-button
        type="primary"
        size="medium"
        class="create"
        @click="create"
      >
        创建
      </el-button>
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
      <el-table-column label="算法描述" :show-overflow-tooltip="true">
        <template slot-scope="scope">
          <span>{{ scope.row.algorithmDescript }}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createdAt) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button type="text" @click="getAlgorithmVersionList(scope.row)">版本列表</el-button>
          <el-button type="text" style="padding-right:10px" @click="createNewVersion(scope.row)">创建新版本</el-button>
          <!-- <el-button type="text" @click="editAlgorithm(scope.row)" v-if="Type === 1 ? false : true">编辑</el-button> -->
          <el-button slot="reference" type="text" @click="confirmDelete(scope.row)">删除</el-button>
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
      :Type="this.typeChange"
      @close="close"
    />
    <preAlgorithmVersionCreation
      v-if="standardDialogVisible"
      :row="this.row"
      :dialogType="dialogType"
      @close="close"
      @cancel="cancel"
      @confirm="confirm"
    />
    <preAlgorithmCreation
      v-if="creationVisible"
      @cancel="cancel"
      @close="close"
      @confirm="confirm"
    />
  </div>
</template>

<script>
import { getPresetAlgorithmList, getAlgorithmVersionList, deletePreAlgorithm } from "@/api/modelDev"
import versionList from "./versionList.vue"
import preAlgorithmVersionCreation from "./preAlgorithmVersionCreation.vue"
import preAlgorithmCreation from './preAlgorithmCreation.vue'
import searchForm from '@/components/search/index.vue'
import { parseTime } from '@/utils/index'
import { getErrorMsg } from '@/error/index'
export default {
  name: "templateList",
  components: {
    versionList,
    preAlgorithmVersionCreation,
    preAlgorithmCreation,
    searchForm
  },
  props: {
    Type: { type: Number, default: undefined }
  },
  data() {
    return {
      row: {},
      total: undefined,
      versionListVisible: false,
      standardDialogVisible: false,
      algorithmName: "",
      dialogType: false,
      creationVisible: false,
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
    getErrorMsg(code) {
      return getErrorMsg(code)
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
      this.typeChange = this.Type
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
    create() {
      this.creationVisible = true;
    },
    close(val) {
      this.versionListVisible = val
      this.standardDialogVisible = val
      this.creationVisible = val
      this.getAlgorithmList(this.searchData);
    },
    createNewVersion(row) {
      this.row = row
      this.standardDialogVisible = true;
    },
    editAlgorithm(row) {
      this.algorithmName = row.AlgorithmName
      this.standardDialogVisible = true;
      this.dialogType = true
    },
    cancel(val) {
      this.standardDialogVisible = val
      this.creationVisible = val
    },
    confirm(val) {
      this.standardDialogVisible = val
      this.creationVisible = val
      this.getAlgorithmList(this.searchData);
    },
    confirmDelete(row) {
      this.$confirm('是否删除此算法？', '提示', {
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
      deletePreAlgorithm(row.algorithmId).then(response => {
        if (response.success) {
          this.$message.success("删除成功");
          this.getAlgorithmList(this.searchData);
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          });
        }
      })
    },
    // 时间戳转换日期
    parseTime(val) {
      return parseTime(val)
    }
  }
}
</script>

<style lang="scss" scoped>
  .Wrapper {
    margin: 15px!important;
  }
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