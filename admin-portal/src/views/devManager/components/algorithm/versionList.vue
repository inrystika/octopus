<template>
  <div>
    <el-dialog :close-on-click-modal="false" :title="title" width="70%" :visible.sync="versionListVisible"
      :before-close="handleDialogClose">
      <el-table v-loading="loading" :data="versionList" style="width: 100%" height="350">
        <el-table-column label="算法名称">
          <template slot-scope="scope">
            <span>{{ scope.row.algorithmName }}</span>
          </template>
        </el-table-column>
        <el-table-column label="版本号">
          <template slot-scope="scope">
            <span>{{ scope.row.algorithmVersion }}</span>
          </template>
        </el-table-column>
        <el-table-column label="版本描述" :show-overflow-tooltip="true">
          <template slot-scope="scope">
            <span>{{ scope.row.algorithmDescript }}</span>
          </template>
        </el-table-column>
        <el-table-column label="创建时间">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.createdAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="算法状态">
          <template slot-scope="scope">
            <span v-if="!(scope.row.progress&&scope.row.progress!=0)">{{ getAlgorithmStatus(scope.row.fileStatus)
              }}</span>
            <span v-if="scope.row.progress&&scope.row.progress!=0">{{ "上传中" }}</span>
            <el-progress :percentage="parseInt(scope.row.progress-1)" v-if="scope.row.progress&&scope.row.progress!=0">
            </el-progress>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <!-- <el-button type="text" style="padding-right:10px">预览</el-button> -->
            <el-button v-if="(scope.row.fileStatus === 1 ) || (scope.row.fileStatus === 4 ) ? true : false"
              v-show="algorithmType === 2 ? true : false" type="text" @click="reupload(scope.row)" :disabled="scope.row.progress&&scope.row.progress!=0">
              重新上传
            </el-button>
            <el-button slot="reference" type="text" :disabled="scope.row.fileStatus === 3 ? false : true"
              @click="confirmDownload(scope.row)">
              下载
            </el-button>
            <el-button v-if="algorithmType === 2 ? true : false" slot="reference" type="text"
              @click="confirmDelete(scope.row)" :disabled="scope.row.progress&&scope.row.progress!=0">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="block">
        <el-pagination :current-page="pageIndex" :page-sizes="[10, 20, 50, 80]" :page-size="pageSize" :total="total"
          layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange"
          @current-change="handleCurrentChange" />
      </div>
      <div slot="footer">
      </div>
    </el-dialog>
    <reuploadAlgorithm v-if="myAlgorithmVisible" :reupload-data="reuploadData" @close="close" @cancel="cancel"
      @confirm="confirm" />
  </div>
</template>

<script>
  import { getAlgorithmVersionList, queryAlgorithmVersion, compressAlgorithm, downloadAlgorithmVersion, deletePreAlgorithmVersion } from "@/api/modelDev";
  import { parseTime } from '@/utils/index'
  import reuploadAlgorithm from "./reuploadAlgorithm.vue"
  import { getErrorMsg } from '@/error/index'
  import store from '@/store'
  export default {
    name: "VersionList",
    components: {
      reuploadAlgorithm
    },
    props: {
      payload: { type: Object, default: () => { } },
      algorithmType: { type: Number, default: undefined },
      row: {
        type: Object,
        default: () => { }
      }
    },
    data() {
      return {
        title: '版本列表/' + this.row.algorithmName,
        versionListVisible: true,
        myAlgorithmVisible: false,
        loading: false,
        pageIndex: 1,
        pageSize: 20,
        total: undefined,
        versionList: [],
        reuploadData: {}
      }
    },
    created() {
      this.getVersionList()
      this.timer = setInterval(() => { this.getVersionList() }, 2000)

    },
    destroyed() {
      clearInterval(this.timer)
      this.timer = null
    },
    methods: {
      getErrorMsg(code) {
        return getErrorMsg(code)
      },
      reupload(row) {
        store.commit('user/SET_PROGRESSID', row.algorithmId + row.algorithmVersion)
        this.myAlgorithmVisible = true
        this.reuploadData = row
      },
      handleSizeChange(val) {
        this.pageSize = val
        this.getVersionList()
      },
      handleCurrentChange(val) {
        this.pageIndex = val
        this.getVersionList()
      },
      getVersionList(param) {
        this.typeChange = this.algorithmType
        if (!param) {
          param = { pageIndex: this.pageIndex, pageSize: this.pageSize }
        }
        param.algorithmId = this.row.algorithmId
        getAlgorithmVersionList(param).then(response => {
          if (response.success) {
            this.versionList = response.data.algorithms,
              this.versionList.forEach(item => {
                if (sessionStorage.getItem(JSON.stringify(item.algorithmId + item.algorithmVersion))) {
                  item.progress = sessionStorage.getItem(JSON.stringify(item.algorithmId + item.algorithmVersion))
                }

              })
            this.total = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      },
      handleDialogClose() {
        this.$emit("close", false);
      },
      // 接受到url下载
      URLdownload(fileName, url) {
        const link = document.createElement('a')
        link.style = 'display: none'; // 创建一个隐藏的a标签
        link.setAttribute('download', fileName)
        link.setAttribute('href', url)
        link.setAttribute('target', "_blank")
        document.body.appendChild(link);
        link.click(); // 触发a标签的click事件
        document.body.removeChild(link);
      },
      confirmDownload(row) {
        this.$confirm('是否下载此版本算法？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          center: true
        }).then(() => {
          this.downloadAlgorithm(row)
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消'
          });
        });
      },
      downloadAlgorithm(row) {
        const that = this
        this.loading = true
        const param = {
          algorithmId: row.algorithmId,
          version: row.algorithmVersion
        }
        let latestCompressed = row.latestCompressed
        compressAlgorithm(param).then(response => {
          if (response.success) {
            param.compressAt = response.data.compressAt
            param.domain = this.GLOBAL.DOMAIN
            const interval = setInterval(function () {
              queryAlgorithmVersion(param).then(response => {
                if (response.success) {
                  latestCompressed = response.data.algorithm.latestCompressed
                } else {
                  that.loading = false
                  clearInterval(interval)
                  that.$message({
                    message: that.getErrorMsg(response.error.subcode),
                    type: 'warning'
                  });
                }
              })
              if (param.compressAt <= latestCompressed) {
                that.loading = false
                clearInterval(interval)
                downloadAlgorithmVersion(param).then(response => {
                  if (response.success) {
                    that.URLdownload(row.algorithmName, response.data.downloadUrl)
                    that.$message.success("下载成功");
                  } else {
                    that.$message({
                      message: that.getErrorMsg(response.error.subcode),
                      type: 'warning'
                    });
                  }
                })
              }
            }, 3000)
          } else {
            that.loading = false
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      },
      confirmDelete(row) {
        this.$confirm('是否删除此版本算法', '提示', {
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
        const param = {
          algorithmId: row.algorithmId,
          version: row.algorithmVersion
        }
        deletePreAlgorithmVersion(param).then(response => {
          if (response.success) {
            this.$message.success('删除成功')
            this.getVersionList();
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      },
      getAlgorithmStatus(value) {
        switch (value) {
          case 1:
            return "未上传"
          case 2:
            return "制作中"
          case 3:
            return "解压完成"
          default:
            return "解压失败"
        }
      },
      cancel(val) {
        this.myAlgorithmVisible = val
      },
      close(val) {
        this.myAlgorithmVisible = val
      },
      confirm(val) {
        this.myAlgorithmVisible = val
      },
      // 时间戳转换日期
      parseTime(val) {
        return parseTime(val)
      }
    }

  }
</script>