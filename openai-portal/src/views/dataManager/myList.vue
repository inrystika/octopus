<template>
  <div>
    <div class="searchForm">
      <searchForm
        :searchForm="searchForm"
        :blurName="'数据集名称 搜索'"
        @searchData="getSearchData"
      >
      </searchForm>
    </div>
    <el-button type="primary" size="medium" class="create" @click="create">
      创建
    </el-button>
    <div class="index">
      <el-table
        v-loading="loading"
        :data="datasetList"
        style="width: 100%;font-size: 15px;"
        :header-cell-style="{'text-align':'left','color':'black'}"
        :cell-style="{'text-align':'left'}"
      >
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
            <!-- <el-button type="text">预览</el-button> -->
            <el-button type="text" @click="createNewVersion(scope.row)">创建新版本
            </el-button>
            <el-button type="text" style="padding-right:10px" @click="getVersionList(scope.$index, scope.row)">版本列表
            </el-button>
            <el-button slot="reference" type="text" @click="confirmDelete(scope.row)">删除</el-button>
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
      >
      </el-pagination>
    </div>

    <myDatasetCreation v-if="myDatasetVisible" @confirm="confirm" @cancel="cancel" @close="close">
    </myDatasetCreation>
    <newVersionCreation
      v-if="newVersionCreationVisible"
      :row="this.data"
      @cancel="cancel"
      @confirm="confirm"
      @close="close"
    ></newVersionCreation>
    <versionList
      v-if="versionListVisible"
      :data="this.data"
      :typeChange="this.typeChange"
      @cancel="cancel"
      @confirm="confirm"
      @close="close"
    >
    </versionList>

  </div>
</template>

<script>
  import newVersionCreation from "./components/newVersionCreation.vue";
  import myDatasetCreation from "./components/myDatasetCreation.vue"
  import versionList from "./components/versionList.vue";
  import searchForm from '@/components/search/index.vue'
  import { deleteDataset, getMyDatasetList } from "@/api/datasetManager";
  import { parseTime } from '@/utils/index'
  import { getErrorMsg } from '@/error/index'
  export default {
    name: "myList",
    components: {
      newVersionCreation,
      versionList,
      myDatasetCreation,
      searchForm
    },
    props: {
      Type: { type: Number },
      dataset: {
        type: Boolean,
        default: false
      }
    },

    data() {
      return {
        input: "",
        data: undefined,
        newVersionCreationVisible: false,
        formLabelWidth: "120px",
        versionListVisible: false,
        myDatasetVisible: false,
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
      if (this.dataset) {
        this.myDatasetVisible = true
      }
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
          getMyDatasetList(param).then(response => {
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
      createNewVersion(row) {
        this.data = row;
        this.newVersionCreationVisible = true;
      },
      getVersionList(index, row) {
        this.data = row;
        this.versionListVisible = true;
      },
      create() {
        this.myDatasetVisible = true;
      },
      cancel(val) {
        this.newVersionCreationVisible = val;
        this.versionListVisible = val;
        this.myDatasetVisible = val
      },
      confirm(val) {
        this.newVersionCreationVisible = val;
        this.versionListVisible = val;
        this.myDatasetVisible = val
        this.getDataList(this.searchData)
      },
      close(val) {
        this.newVersionCreationVisible = val;
        this.versionListVisible = val;
        this.myDatasetVisible = val
        this.getDataList(this.searchData)
      },
      confirmDelete(row) {
        this.$confirm('此操作将永久删除此数据集（如该数据集已分享，则分享的数据集也会被删除)，是否继续?', '提示', {
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
      // 时间戳转换日期
      parseTime(val) {
        return parseTime(val)
      }
    }
  };
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