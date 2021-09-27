<template>
  <div>
    <div class="searchForm">
      <searchForm :searchForm="searchForm" :blurName="'NoteBook名称 搜索'" @searchData="getSearchData" />
    </div>
    <el-button type="primary" size="medium" class="create" @click="create">
      创建
    </el-button>
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
      <el-table-column label="规格">
        <template slot-scope="scope">
          <span>{{ scope.row.resourceSpecName+" "+scope.row.resourceSpecPrice+"机时/h" }}</span>
        </template>
      </el-table-column>
      <el-table-column label="描述" :show-overflow-tooltip="true">
        <template slot-scope="scope">
          <span>{{ scope.row.desc }}</span>
        </template>
      </el-table-column>
      <el-table-column label="算法">
        <template slot-scope="scope">
          <span>{{ scope.row.algorithmName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="镜像">
        <template slot-scope="scope">
          <span>{{ scope.row.imageName }}</span>
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
          <!-- <div v-if="scope.row.status === '停止' ? true : false"> -->
          <div v-if="({'stopped':true,'succeeded':true,'failed':true})[scope.row.status] || false">
            <el-button slot="reference" type="text" style="padding-right:10px" @click="confirmStart(scope.row)">
              启动
            </el-button>
            <el-button slot="reference" type="text" @click="confirmDelete(scope.row)">删除</el-button>
            <!-- <el-button type="text" @click="saveAlgorithm(scope.row)">保存算法</el-button> -->
          </div>
          <el-button
            v-if="({'running':true})[scope.row.status] || false"
            type="text"
            style="padding-right:10px"
            @click="jumpUrl(scope.row.url)"
          >打开</el-button>
          <el-button
            v-if="({'preparing':true,'pending':true,'running':true})[scope.row.status] || false"
            slot="reference"
            type="text"
            @click="confirmStop(scope.row)"
          >
            停止
          </el-button>
          <el-button slot="reference" type="text" @click="getNotebookInfo(scope.row)">
            信息
          </el-button>
        </template>
      </el-table-column>
    </el-table>
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

    <notebookInfo
      v-if="notebookInfoVisible"
      :initInfo="initInfo"
      @confirm="confirm"
      @cancel="cancel"
      @close="close"
    />
    <notebookCreation v-if="notebookVisible" @cancel="cancel" @confirm="confirm" @close="close" />
  </div>
</template>

<script>
  import notebookCreation from "./notebookCreation.vue"
  import notebookInfo from "./notebookInfo.vue"
  import searchForm from '@/components/search/index.vue'
  import { getNotebookInfo, getNotebookList, stopNotebook, deleteNotebook, startNotebook } from "@/api/modelDev";
  import { parseTime } from '@/utils/index'
  import { getResourceList } from "@/api/trainingManager"
  import { getErrorMsg } from '@/error/index'
  export default {
    name: "NotebookList",
    components: {
      notebookCreation,
      notebookInfo,
      searchForm
    },
    props: {
      notebook: {
        type: Boolean,
        default: false
      }
    },
    data() {
      return {
        row: {},
        initInfo: "",
        notebookVisible: false,
        notebookInfoVisible: false,
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
        searchData: {
          pageIndex: 1,
          pageSize: 10
        },
        resourceList: [],
        title: "是否启动NoteBook？",
        statusText: {
          'preparing': ['status-ready', '初始中'],
          'pending': ['status-agent', '等待中'],
          'running': ['status-running', '运行中'],
          'failed': ['status-danger', '失败'],
          'succeeded': ['status-success', '成功'],
          'stopped': ['status-stopping', '已停止']
        }
      }
    },
    created() {
      this.getNotebookList(this.searchData);
      this.getResource()
      if (this.notebook) {
        this.notebookVisible = true
      }
    },
    methods: {
      getErrorMsg(code) {
        return getErrorMsg(code)
      },
      jumpUrl(url) {
        const jumpUrl = this.GLOBAL.DOMAIN + url
        window.open(jumpUrl)
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
            })
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
      getResource() {
        getResourceList().then(response => {
          if (response.success) {
            this.resourceList = response.data.mapResourceSpecIdList.debug.resourceSpecs
          }
        })
      },
      confirmStart(row) {
        this.$confirm(this.title, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          center: true
        }).then(() => {
          this.handleStart(row)
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消'
          });
        });
      },
      handleStart(row) {
        // this.resourceList.forEach(
        //   item => {
        //     if (item.id === row.id) {
        //       if (item.price !== row.resourceSpecPrice) {
        //         this.title = `<div>资源价格有变动,
        //                     <strong>旧价格:${row.resourceSpecPrice}机时/h,</strong>
        //                     <strong>新价格:${item.price}机时/h,</strong>
        //                     是否启动Notebook？
        //                     </div>`
        //       }
        //     }
        //   }
        // )
        startNotebook(row.id).then(response => {
          if (response.success) {
            this.$message.success("已启动");
            this.getNotebookList(this.searchData);
          } else {
            if (response.error.subcode === 11014) {
              this.$message({
                message: '资源规格已被删除，请重新提交NoteBook',
                type: 'warning'
              })
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              })
            }
          }
        })
      },
      confirmStop(row) {
        this.$confirm('是否停止NoteBook？', '提示', {
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
      getNotebookInfo(row) {
        getNotebookInfo(row.id).then(response => {
          if (response.success) {
            this.notebookInfoVisible = true
            const notebookDialogString = response.payload.notebook.initInfo ? response.payload.notebook.initInfo.replace(/\n/g, "<br>") : ''
            const notebookDialogData = JSON.parse(notebookDialogString)
            for (const pid in notebookDialogData['podEvents']) {
              const eventList = notebookDialogData['podEvents'][pid]
              const roleName = notebookDialogData['podRoleName'][pid]
              if (roleName == "") {
                continue
              }
              let message = ""
              for (const key in eventList) {
                const event = eventList[key]
                if (event['reason'] == "" && event['message'] == "") {
                  continue
                }
                message += "[" + event['reason'] + "]" + "<br>"
                message += event['message'] + "<br><br>"
              }
              for (const key in notebookDialogData['extras']) {
                const event = notebookDialogData['extras'][key]
                if (event['reason'] == "" && event['message'] == "") {
                  continue
                }
                message += "[" + event['reason'] + "]" + "<br>"
                message += event['message'] + "<br><br>"
              }
              message += "<br>"
              this.initInfo = message
            }
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            })
          }
        })
      },
      handleStop(row) {
        stopNotebook(row.id).then(response => {
          if (response.success) {
            this.$message.success("已停止");
            this.getNotebookList(this.searchData);
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            })
          }
        })
      },
      confirmDelete(row) {
        this.$confirm('是否删除NoteBook？', '提示', {
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
        deleteNotebook(row.id).then(response => {
          if (response.success) {
            this.$message.success("已删除");
            this.getNotebookList(this.searchData);
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
      },
      create() {
        this.notebookVisible = true;
      },
      close(val) {
        this.notebookVisible = val;
        this.notebookInfoVisible = val;
        this.getNotebookList(this.searchData);
      },
      cancel(val) {
        this.notebookVisible = val;
        this.notebookInfoVisible = val;
        this.getNotebookList(this.searchData);
      },
      confirm(val) {
        this.notebookVisible = val
        this.notebookInfoVisible = val;
        this.getNotebookList(this.searchData);
      }
    }
  }
</script>

<style lang="scss" scoped>
  .Wrapper {
    margin: 20px !important;
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