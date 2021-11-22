<template>
  <div>
    <div>
      <searchForm
        :search-form="searchForm"
        :blur-name="'NoteBook名称 搜索'"
        @searchData="getSearchData"
      />
    </div>
    <el-table
      :data="notebookList"
      style="width: 100%;font-size: 15px;"
      :header-cell-style="{'text-align':'left','color':'black'}"
      :cell-style="{'text-align':'left'}"
    >
      <el-table-column label="名称">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="描述" :show-overflow-tooltip="true">
        <template slot-scope="scope">
          <span>{{ scope.row.desc }}</span>
        </template>
      </el-table-column>
      <el-table-column label="群组">
        <template slot-scope="scope">
          <span>{{ scope.row.workspaceId === 'default-workspace' ? '默认群组' : scope.row.workspaceName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="用户">
        <template slot-scope="scope">
          <span>{{ scope.row.userName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="规格">
        <template slot-scope="scope">
          <span>{{ scope.row.resourceSpecName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态">
        <template slot-scope="scope">
          <!-- <span>{{ scope.row.status }}</span> -->
          <span :class="statusText[scope.row.status][0]"></span>
          <span>{{ statusText[scope.row.status][1] }}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createdAt) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <!-- <div v-if="({'running':true,'pending':true,'preparing':true})[scope.row.status] || false"> -->
            <el-button
              v-if="({'running':true,'pending':true,'preparing':true})[scope.row.status] || false" 
              slot="reference" 
              type="text" 
              @click="confirmStop(scope.row)"
            >
              停止
            </el-button>
          <!-- </div> -->
          <!-- <div v-if="({'stopped':true})[scope.row.status] || false">
          </div> -->
          <el-button slot="reference" type="text" @click="showNotebookInfo(scope.row)">
            详情
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <detailDialog
      v-if="detailVisible"
      :detail-data="detailData"
      @confirm="confirm"
      @cancel="cancel"
      @close="close"
    />

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
import detailDialog from "./detailDialog.vue"
import searchForm from '@/components/search/index.vue'
import { getNotebookList, stopNotebook } from "@/api/modelDev"
import { parseTime } from '@/utils/index'
import { getErrorMsg } from '@/error/index'
export default {
  name: "NotebookList",
  components: {
    searchForm,
    detailDialog
  },
  data() {
    return {
      detailData: {},
      detailVisible: false,
      row: {},
      total: undefined,
      notebookList: [],
      searchForm: [
        { type: 'Time', label: '创建时间', prop: 'time', placeholder: '请选择创建时间' },
        {
          type: 'Select', label: '状态', prop: 'status', placeholder: '请选择状态',
          options: [
            { label: '运行中', value: 'running' },
            { label: '等待中', value: 'pending' },
            { label: '已停止', value: 'stopped' },
            { label: '成功', value: 'succeeded' },
            { label: '失败', value: 'failed' },
            { label: '初始中', value: 'preparing' }
          ]
        }
      ],
      statusText: {
        'preparing': ['status-ready', '初始中'],
        'pending': ['status-agent', '等待中'],
        'running': ['status-running', '运行中'],
        'failed': ['status-danger', '失败'],
        'succeeded': ['status-success', '成功'],
        'stopped': ['status-stopping', '已停止']
      },
      searchData: {
        pageIndex: 1,
        pageSize: 10
      }
    }
  },
  created() {
    this.getNotebookList(this.searchData);
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    handleSizeChange(val) {
      this.searchData.pageSize = val
      this.getNotebookList(this.searchData)
    },
    handleCurrentChange(val) {
      this.searchData.pageIndex = val
      this.getNotebookList(this.searchData)
    },
    getNotebookList(param) {
      getNotebookList(param).then(response => {
         if (response.success) {
          this.notebookList = response.data.notebooks;
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
      this.getNotebookList(this.searchData)
    },
    confirmStop(row) {
      this.$confirm('是否停止Notebook？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      }).then(() => {
        this.handleStop(row)
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消'
        });
      });
    },
    showNotebookInfo(row) {
      this.detailVisible = true
      this.detailData = row
    },
    handleStop(row) {
      stopNotebook(row.id).then(response => {
        if (response.success) {
          this.$message.success("已停止");
          setTimeout(function() { location.reload() }, 1000)
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
      })
    },
    close(val) {
      this.detailVisible = val;
      this.getNotebookList(this.searchData);
    },
    cancel(val) {
      this.detailVisible = val;
      this.getNotebookList(this.searchData);
    },
    confirm(val) {
      this.detailVisible = val;
      this.getNotebookList(this.searchData);
    },
    // 时间戳转换日期
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