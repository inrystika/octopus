<template>
  <div>
    <el-dialog :title="title" width="80%" :visible.sync="versionListVisible" :before-close="handleDialogClose"
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
            <span>{{ scope.row.createdAt | parseTime }}</span>
          </template>
        </el-table-column>
        <el-table-column label="状态">
          <template slot-scope="scope">
            <span v-if="!(scope.row.progress&&scope.row.progress!=0)">{{ getDatasetStatus(scope.row.status) }}</span>
            <span v-if="scope.row.progress&&scope.row.progress!=0">{{ "上传中" }}</span>
            <el-progress :percentage="parseInt(scope.row.progress-1)" v-if="scope.row.progress&&scope.row.progress!=0">
            </el-progress>
          </template>
        </el-table-column>
        <el-table-column label="操作" props="action">
          <template slot-scope="scope">
            <el-button v-if="(scope.row.status === 1 ) || (scope.row.status === 4 ) ? true : false"
              v-show="versionListType === 1 ? false : true" type="text" @click="reupload(scope.row)"
              :disabled="scope.row.progress&&scope.row.progress!=0">
              重新上传
            </el-button>
            <el-button type="text" :disabled="scope.row.status === 3 ? false : true" @click="handlePreview(scope.row)"
              style="margin-left: 0px;">
              预览
            </el-button>
            <el-button v-if="versionListType === 1 ? false : true" slot="reference" type="text"
              @click="confirmDelete(scope.row)" :disabled="scope.row.progress&&scope.row.progress!=0">
              删除
            </el-button>
            <el-button type="text" v-if="scope.row.status === 3 ? true : false" @click="handleCache(scope.row)">
              加速设置
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
    <el-dialog title="加速设置" :visible.sync="dialogCache">
      <el-form :model="cache">
        <el-form-item label="即时配送">
          <el-switch v-model="cache.open" @change="open"></el-switch>
        </el-form-item>
        <el-form-item label="缓存大小" v-if="show">
          <el-select v-model="cache.quota" placeholder="请选择缓存大小">
            <el-option label="500M" value="500M"></el-option>
            <el-option label="1G" value="1G"></el-option>
            <el-option label="2G" value="2G"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="dialogFormVisible = false">确 定</el-button>
      </div>
    </el-dialog>
  </div>

</template>
<script>
  import { getVersionList, deleteDatasetVersion } from "@/api/dataManager"
  import preview from './preview.vue'
  import reuploadDataset from "./reuploadDataset.vue"
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
        timer: null,
        dialogCache: false,
        cache: { open: false, quota: "1G" },
        show:false
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
            return "未上传"
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
      handleCache(val) {
        this.dialogCache = true
      },
      open() {
        if (this.cache.open) {
          this.show = true
        } else { this.show = false }
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