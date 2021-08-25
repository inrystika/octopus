<template>
  <div>
    <el-dialog
      :title="title"
      width="70%"
      :visible.sync="createListVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
      <el-table
        :data="versionList"
        :model="ruleList"
        props="name"
        ref="ruleList"
        style="width: 100%" height="350"
        v-loading.fullscreen.lock="loading"
      >       
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
        <el-table-column label="创建时间" props="createdAt">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.createdAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="数据集状态" props="status">
          <template slot-scope="scope">
            <span>{{ getDatasetStatus(scope.row.status) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" props="action">
          <template slot-scope="scope">
            <el-button 
              @click="reupload(scope.row)" 
              type="text" 
              v-show="typeChange === 1 ? true : false"
              v-if="(scope.row.status === 1 ) || (scope.row.status === 4 ) ? true : false">重新上传
            </el-button>
            <el-button 
              type="text" 
              style="padding-right:10px"
              @click="handlePreview(scope.row)"
              :disabled="scope.row.status === 3 ? false : true"
            >
              预览
            </el-button>
            <el-button 
              style="padding-right:10px" 
              slot="reference"  
              @click="confirmShare(scope.row)"
              type="text" 
              v-show="typeChange === 1 ? true : false"
              :disabled="scope.row.status === 3 ? false : true"
            >
              {{scope.row.shared ? "取消分享":"分享"}}
            </el-button>
            <el-button 
              @click="confirmDelete(scope.row)"
              slot="reference" 
              type="text" 
              v-show="typeChange === 1 ? true : false"
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
    <preview v-if="preVisible" :row="versionData" @close="close">
    </preview>
    <reuploadDataset 
      v-if="myDatasetVisible" 
      :data="this.data"
      :versionData="this.versionData"
      @close="close" 
      @cancel="cancel" 
      @confirm="confirm">
    </reuploadDataset>
  </div>
</template>

<script>
import { shareDatasetVersion, cancelShareDatasetVersion, deleteDatasetVersion, getVersionList } from "@/api/datasetManager";
import { parseTime } from '@/utils/index'
import preview from './preview.vue'
import reuploadDataset from "./reuploadDataset.vue"
import { getErrorMsg } from '@/error/index'
export default {
  name: "versionList",
  components: {      
    preview,
    reuploadDataset
  },
  props: {
    data: {
      type: Object,
      default: {}
    },
    typeChange: { type: Number }
  },
  data() {
    return {
      ruleList: [
        {
          name: "",
          type: "",
          version: "",
          desc: "",
          createdAt: "",
          provider: ""
        }
      ],
      title:'版本列表/'+this.data.name,
      myDatasetVisible:false,
      createListVisible: true,
      preVisible: false,
      formLabelWidth: "120px",
      shared: false,
      shareTitle: "是否分享至本群组，分享后群内所有人员可见",
      loading: false,
      pageIndex: 1,
      pageSize: 20,
      total: undefined,
      Type: undefined,
      versionData: undefined,
      versionList: []
    };
  },
  created() {
    this.getVersionList();
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    reupload(row) {
      this.myDatasetVisible = true
      this.versionData = row
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
    getVersionList(param){
      this.Type = this.typeChange
      if (!param) { 
        param = { pageIndex: this.pageIndex, pageSize: this.pageSize } 
      }
      param.datasetId = this.data.id
      param.shared = this.Type === 2 ? true :false
      getVersionList(param).then(response => {
        if (response.success) {
          this.versionList = response.data.versions;
          this.total = response.data.totalSize
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          })
        }
      });
    },
    confirmShare(row){
      if(row.shared){
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
      this.versionData = row
      this.loading = true
      if(this.versionData.shared){
        cancelShareDatasetVersion(this.versionData).then(response => {
          if(response.success) {
            this.$message.success("已取消本群组分享");
            this.loading = false
            this.getVersionList()
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            })
            this.loading = false
          }
        })
      } else {
        shareDatasetVersion(this.versionData).then(response => {
          if(response.success) {
            this.$message.success("已分享至群组");
            this.loading = false
            this.getVersionList()
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            })
            this.loading = false
          }
        })
      }
    },
    getDatasetStatus(value){
      switch (value) {
        case 1:
          return "等待上传中"
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
      this.loading = true
      const param = {
        datasetId: row.datasetId,
        version: row.version
      }
      deleteDatasetVersion(param).then(response => {
        if(response.success) {
          this.$message.success("已删除此版本");
          this.loading = false
          this.getVersionList()
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          })
          this.loading = false
        }
      })
    },
    //时间戳转换日期
    parseTime(val) {
      return parseTime(val)
    }
  }
};
</script>

<style lang="scss" scoped>
</style>
