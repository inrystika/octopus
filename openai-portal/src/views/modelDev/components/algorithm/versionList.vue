<template>
  <div>
    <el-dialog :close-on-click-modal="false" :title="title" width="70%" :visible.sync="versionListVisible"
      :before-close="handleDialogClose">
      <el-table :data="versionList" style="width: 100%" height="350" v-loading="loading">
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
        <el-table-column label="提供者" v-if="Type === 1 ? false :true">
          <template slot-scope="scope">
            <span>{{ scope.row.userName }}</span>
          </template>
        </el-table-column>
        <el-table-column label="算法状态" props="status">
          <template slot-scope="scope">
            <span>{{ getAlgorithmStatus(scope.row.fileStatus) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <!-- <el-button type="text">预览</el-button> -->
            <el-button 
              @click="reupload(scope.row)" 
              type="text" 
              v-show="Type === 1 ? true :false"
              v-if="(scope.row.fileStatus === 1 ) || (scope.row.fileStatus === 4 ) ? true : false">重新上传
            </el-button>
            <el-button 
              type="text" 
              style="padding-right:10px" 
              @click="createTask(scope.row)"
              :disabled="(scope.row.fileStatus === 3)? false : true"
              >
              创建训练任务
            </el-button>
            <el-button 
              slot="reference" 
              type="text" 
              @click="confirmDownload(scope.row)" 
              :disabled="scope.row.fileStatus === 3 ? false : true"
            >
              下载
            </el-button>
            <el-button 
              style="padding-right:10px" 
              slot="reference" 
              @click="confirmShare(scope.row)" 
              type="text"
              :disabled="scope.row.fileStatus === 3 ? false : true"
              v-if="Type === 1 ? true :false"
            >
              {{scope.row.isShared?"取消分享":"分享"}}
            </el-button>
            <el-button 
              @click="confirmDelete(scope.row)"
              type="text" 
              slot="reference" 
              v-if="Type === 1 ? true :false"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="block">
        <el-pagination 
          @size-change="handleSizeChange" 
          @current-change="handleCurrentChange" 
          :current-page="pageIndex"
          :page-sizes="[10, 20, 50, 80]" 
          :page-size="pageSize" 
          layout="total, sizes, prev, pager, next, jumper"
          :total="total">
        </el-pagination>
    </div>
    <div slot="footer">
    </div>
    </el-dialog>
    <reuploadAlgorithm
      v-if="myAlgorithmVisible"
      :data="this.data"
      @close="close" 
      @cancel="cancel" 
      @confirm="confirm">
    </reuploadAlgorithm>
  </div>
</template>

<script>
  import { getPubAlgorithmVersionList, getAlgorithmVersionList, queryAlgorithmVersion, compressAlgorithm, downloadAlgorithm, shareAlgorithmVersion, cancelShareAlgorithmVersion, deleteAlgorithmVersion } from "@/api/modelDev"
  import { parseTime } from '@/utils/index'
  import reuploadAlgorithm from './reuploadAlgorithm.vue'
  import { getErrorMsg } from '@/error/index'
  export default {
    name: "versionList",
    components: {
      reuploadAlgorithm
    },
    props: {
      Type: { type: Number },
      data: {
        type: Object,
        default: {}
      },
    },
    data() {
      return {
        title:'版本列表/'+this.data.algorithmName,
        versionListVisible: true,
        myAlgorithmVisible: false,
        loading: false,
        pageIndex:1,
        pageSize:20,
        total: undefined,
        typeChange: undefined,
        versionList: [],
        shareTitle: "是否分享至本群组，分享后群内所有人员可见"
      }
    },
    created() {
      this.getVersionList();
    },
    methods: {
      getErrorMsg(code) {
        return getErrorMsg(code)
      },
      reupload(row){
        this.myAlgorithmVisible = true
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
        this.typeChange = this.Type
        if (!param) { 
          param = { pageIndex: this.pageIndex, pageSize: this.pageSize } 
        }
        param.algorithmId = this.data.algorithmId
        if (this.typeChange === 2){
        getPubAlgorithmVersionList(param).then(response => {
          if(response.success) {
            this.versionList = response.data.algorithms
            this.total = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      } else {
        getAlgorithmVersionList(param).then(response => {
          if(response.success) {
            let newArr = []
            response.data.algorithms.filter(function(item,index) {
              let obj = item.algorithmDetail
              obj.isShared = item.isShared
              newArr.push(obj)
            })
            this.versionList = newArr
            this.total = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      }
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
      confirmDownload(row){
        this.$confirm('是否下载此版本？','提示',{
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      }).then(() =>{
        this.handleDownload(row)
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消'
        });
      });
      },
      handleDownload(row) {
        let that = this
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
            let interval = setInterval(function() {
              queryAlgorithmVersion(param).then(response => {
                if(response.success) {
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
              if ( param.compressAt <= latestCompressed) {
                that.loading = false
                clearInterval(interval)
                downloadAlgorithm(param).then(response => {
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
            },3000)
          } else {
            that.loading = false
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
      confirmShare(row){
        if (row.isShared > 0) {
          this.shareTitle = "是否取消本群组分享？"
        } else {
          this.shareTitle = "是否分享至本群组，分享后群内所有人员可见"
        }
        this.$confirm(this.shareTitle,'提示',{
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          center: true
        }).then(() =>{
          this.handleShare(row)
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消'
          });
        });
      },
      handleShare(row) {
        const param = {
          algorithmId: row.algorithmId,
          version: row.algorithmVersion
        }
        if (row.isShared) {
          cancelShareAlgorithmVersion(param).then(response => {
            if (response.success) {
              this.$message.success("已取消本群组分享");
              this.getVersionList();
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
        } else {
          shareAlgorithmVersion(param).then(response => {
            if (response.success) {
              this.$message.success("已分享至群组");
              this.getVersionList();
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
        }
      },
      getAlgorithmStatus(value){
        switch (value) {
          case 1:
            return "等待上传中"
          case 2:
            return "上传中"
          case 3: 
            return "上传完成"
          case 4:
            return "上传失败"
        }
      },
      cancel(val) {
        this.algorithmVersionDeleteDialogVisible = val
        this.myAlgorithmVisible = val
      },
      close(val) {
        this.algorithmVersionDeleteDialogVisible = val
        this.myAlgorithmVisible = val
      },
      confirm(val) {
        this.myAlgorithmVisible = val
      },
      confirmDelete(row){
        this.$confirm('是否删除此版本？','提示',{
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          center: true
        }).then(() =>{
          this.handleDelete(row)
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消'
          });
        });
      },
      handleDelete(row) {
        deleteAlgorithmVersion(row).then(response => {
          if (response.success) {
            this.$message.success("删除成功");
            this.getVersionList();
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      },
      //时间戳转换日期
      parseTime(val) {
        return parseTime(val)
      },
      //创建训练任务
      createTask(row) {
        let data = row
        data.trainingTask = true
        // data.open = true
        switch (this.Type) {
          case 1:
            data.type = '我的算法'
            break;
          case 2:
            data.type = '公共算法'
            break;
          default:
            data.type = '预置算法'
        }
        this.$router.push({
          name: 'trainingManager',
          params: { data: data }
        })
      }
    }
  }
</script>