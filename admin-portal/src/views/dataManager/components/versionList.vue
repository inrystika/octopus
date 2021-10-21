<template>
  <div>
    <el-dialog :title="title" width="70%" :visible.sync="versionListVisible" :before-close="handleDialogClose"
      :close-on-click-modal="false">
      <el-table v-loading.fullscreen.lock="loading" label-width="100px" :data="versionList" style="width: 100%"
        height="350">
        <el-table-column label="版本号" props="version">
          <template slot-scope="scope">
            <span>{{ scope.row.version }}</span>
          </template>
        </el-table-column>
        <el-table-column label="版本描述" props="desc" :show-overflow-tooltip="true">
          <template slot-scope="scope">
            <span>{{ scope.row.desc }}</span>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" props="desc">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.createdAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="数据集状态" props="status">
          <template slot-scope="scope">
            <span>{{ getDatasetStatus(scope.row.status) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="上传进度" v-if="versionListType===2">
          <template slot-scope="scope">
            <span v-if="scope.row.progress&&scope.row.progress!=0" style="color:#409EFF">{{
              scope.row.progress+'%' }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" props="action">
          <template slot-scope="scope">
            <el-button v-if="(scope.row.status === 1 ) || (scope.row.status === 4 ) ? true : false"
              v-show="versionListType === 1 ? false : true" type="text" @click="reupload(scope.row)">
              重新上传
            </el-button>
            <el-button type="text" style="padding-right:10px" :disabled="scope.row.status === 3 ? false : true"
              @click="handlePreview(scope.row)">
              预览
            </el-button>
            <el-button v-if="versionListType === 1 ? false : true" slot="reference" type="text"
              @click="confirmDelete(scope.row)">
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
    <preview v-if="preVisible" :row="versionData" @close="close" />
    <reuploadDataset v-if="myDatasetVisible" :data="data" :version-data="versionData" @close="close" @cancel="cancel"
      @confirm="confirm" />
  </div>
</template>

<script>
  import { getVersionList, deleteDatasetVersion } from "@/api/dataManager"
  import { parseTime } from '@/utils/index'
  import preview from './preview.vue'
  import reuploadDataset from "./reuploadDataset.vue"
  import { getErrorMsg } from '@/error/index'
  import store from '@/store'
  export default {
    name: "VersionList",
    components: {
      preview,
      reuploadDataset
    },
    props: {
      versionListType: { type: Number, default: undefined },
      data: {
        type: Object,
        default: () => { }
      }
    },
    data() {
      return {
        title: '版本列表/' + this.data.name,
        versionListVisible: true,
        myDatasetVisible: false,
        preVisible: false,
        versionData: undefined,
        loading: false,
        pageIndex: 1,
        pageSize: 20,
        total: undefined,
        versionList: [],
        timer: null
      }
    },
    created() {
      this.timer = setInterval(() => { this.getVersionList() }, 1000)

    },
    destroyed() {
      clearInterval(this.timer)
      this.timer = null
    },
    destroyed() {
      clearInterval(this.timer)
    },
    methods: {
      getErrorMsg(code) {
        return getErrorMsg(code)
      },
      reupload(row) {
        this.myDatasetVisible = true
        this.versionData = row,
          store.commit('user/SET_PROGRESSID', row.datasetId + row.version)
      },
      handlePreview(row) {
        this.preVisible = true
        this.versionData = row
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
        if (!param) {
          param = { pageIndex: this.pageIndex, pageSize: this.pageSize }
        }
        param.id = this.data.id
        getVersionList(param).then(response => {
          if (response.success) {
            this.versionList = response.data.versions;
            this.versionList.forEach(item => {
              if (sessionStorage.getItem(JSON.stringify(item.datasetId + item.version))) {
                item.progress = sessionStorage.getItem(JSON.stringify(item.datasetId + item.version))
              }

            })
            this.total = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        });
      },
      confirmDelete(row) {
        this.$confirm('是否删除此版本数据集？', '提示', {
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
        deleteDatasetVersion(row).then(response => {
          if (response.success) {
            this.$message.success("删除成功");
            this.loading = false
            this.getVersionList()
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
            this.loading = false
          }
        })
      },
      getDatasetStatus(value) {
        switch (value) {
          case 1:
            return "等待解压中"
          case 2:
            return "解压中"
          case 3:
            return "解压完成"
          case 4:
            return "解压失败"
        }
      },
      handleDialogClose() {
        this.$emit("close", false);
      },
      close(val) {
        this.preVisible = val
        this.myDatasetVisible = val
      },
      cancel(val) {
        this.myDatasetVisible = val
      },
      confirm(val) {
        this.myDatasetVisible = val
      },
      // 时间戳转换日期
      parseTime(val) {
        return parseTime(val)
      }
    }
  }
</script>