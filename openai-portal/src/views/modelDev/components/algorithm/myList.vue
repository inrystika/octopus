<template>
  <div>
    <div class="searchForm">
      <searchForm
        :search-form="searchForm"
        :blur-name="'算法名称/描述 搜索'"
        @searchData="getSearchData"
      />
    </div>
    <el-button
      v-if="algorithmTabType === 1 ? true : false"
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
          <span>{{ parseTime(scope.row.createdAt) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button type="text" @click="getAlgorithmVersionList(scope.row)">版本列表</el-button>
          <el-button type="text" @click="copyAlgorithm(scope.row)">复制算法</el-button>
          <el-button type="text" style="padding-right:10px" @click="createNewVersion(scope.row)">创建新版本</el-button>
          <el-button slot="reference" type="text" @click="confirmDelete(scope.row)">删除</el-button>
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

    <myAlgorithmCreation
      v-if="creationVisible"
      @cancel="cancel"
      @close="close"
      @confirm="confirm"
    />
    <newVersionCreation
      v-if="newVersionVisible"
      :row="row"
      @close="close"
      @cancel="cancel"
      @confirm="confirm"
    />
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
import newVersionCreation from "./newVersionCreation.vue";
import algorithmCopy from "./algorithmCopy.vue";
import versionList from "./versionList.vue";
import myAlgorithmCreation from "./myAlgorithmCreation.vue"
import searchForm from '@/components/search/index.vue'
import { getMyAlgorithmList, deleteMyAlgorithm } from "@/api/modelDev"
import { parseTime } from '@/utils/index'
import { getErrorMsg } from '@/error/index'
export default {
  name: "MyList",
  components: {
    newVersionCreation,
    algorithmCopy,
    versionList,
    myAlgorithmCreation,
    searchForm
  },
  props: {
    algorithmTabType: {
      type: Number,
      default: undefined
    },
    algorithm: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      row: {},
      newVersionVisible: false,
      algorithmCopyVisible: false,
      versionListVisible: false,
      creationVisible: false,
      newVersionName: "",
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
    if (this.algorithm) {
      this.creationVisible = true
    }
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
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
        getMyAlgorithmList(param).then(response => {
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
    createNewVersion(row) {
      this.newVersionName = row.algorithmName
      this.newVersionVisible = true;
      this.row = row
    },
    create() {
      this.creationVisible = true;
    },
    close(val) {
      this.newVersionVisible = val;
      this.algorithmCopyVisible = val;
      this.versionListVisible = val;
      this.creationVisible = val
      this.getAlgorithmList(this.searchData)
    },
    cancel(val) {
      this.newVersionVisible = val;
      this.algorithmCopyVisible = val;
      this.creationVisible = val;
      this.getAlgorithmList(this.searchData)
    },
    copyAlgorithm(row) {
      this.algorithmCopyVisible = true;
      this.row = row
    },
    confirm(val) {
      this.algorithmCopyVisible = val
      this.newVersionVisible = val
      this.creationVisible = val
      this.getAlgorithmList(this.searchData)
    },
    confirmDelete(row) {
      this.$confirm('此操作将永久删除此算法(如该算法已分享，则分享算法也将被删除)，是否继续', '提示', {
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
      deleteMyAlgorithm(row.algorithmId).then(response => {
        if (response.success) {
          this.$message.success("删除成功");
          this.getAlgorithmList(this.searchData)
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
  .pagination {
    float: right;
    margin: 20px;
  }
  .create {
    float: right;
  }
  .searchForm {
    display: inline-block;
  }
</style>