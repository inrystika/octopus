<template>
  <div>
    <el-dialog
      title="存储配置列表"
      width="70%"
      :visible.sync="createFormVisible"
      :before-close="handleDialogClose" 
      :close-on-click-modal="false"
    >
      <el-button type="primary" size="medium" @click="create" class="create">
        创建
      </el-button>
      <el-table
        :data="storageList"
        style="width: 100%;font-size: 15px;"
        :header-cell-style="{'text-align':'left','color':'black'}"
        :cell-style="{'text-align':'left'}"
      >
        <el-table-column label="平台名称">
            <span>{{ this.platformDetail.name }}</span>
        </el-table-column>
        <el-table-column label="配置名称">
          <template slot-scope="scope">
            <span>{{ scope.row.name }}</span>
          </template>
        </el-table-column>
        <el-table-column label="创建时间">
          <template slot-scope="scope">
            <span>{{ parseTime(scope.row.createdAt)  }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button type="text" @click="confirmDeletion(scope.row)">删除</el-button>
            <el-button type="text" @click="showStorageConfigDetailVisible(scope.row)">详情</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="block">
        <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="searchData.pageIndex"
          :page-sizes="[10, 20, 50, 80]"
          :page-size="searchData.pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
        >
        </el-pagination>
      </div>

      <storageConfigCreation
        v-if="storageConfigCreationVisible"
        :platform-detail="platformDetail"
        @cancel="cancel"
        @confirm="confirm"
        @close="close"
      >
      </storageConfigCreation>

      <storageConfigDetails
        v-if="storageConfigDetailsVisible"
        :storage-config-detail="storageConfigDetail"
        @close="close"
      >
      </storageConfigDetails>
    </el-dialog>
  </div>
</template>
<script>
import storageConfigCreation from "./storageConfigCreation.vue"
import storageConfigDetails from "./storageConfigDetails.vue"
import { getErrorMsg } from '@/error/index'
import { parseTime } from '@/utils/index'
import { getStorageConfigList, deleteStorageConfig } from "@/api/platformManager"
export default {
  name: "storageConfigList",
  components: {
    storageConfigCreation,
    storageConfigDetails
  },
  props: {
    platformDetail: {
      type: Object,
      default: () => {}
    },
  },
  data() {
    return {
      createFormVisible: true,
      storageConfigCreationVisible: false,
      storageConfigDetailsVisible: false,
      storageConfigDetail: [],
      storageList: [],
      searchData: {
        pageIndex: 1,
        pageSize: 10,
      },
      total: 0
    }
  },
  created() {
    this.getStorageConfigList()
  },
  methods: {
    close(val) {
      this.storageConfigCreationVisible = val;
      this.storageConfigDetailsVisible = val;
      this.getStorageConfigList();
    },
    cancel(val) {
      this.storageConfigCreationVisible = val;
      this.getStorageConfigList();
    },
    confirm(val) {
      this.storageConfigCreationVisible = val
      this.getStorageConfigList();
    },
    handleSizeChange(val) {
      this.searchData.pageSize = val
      this.getStorageConfigList()
    },
    handleCurrentChange(val) {
      this.searchData.pageIndex = val
      this.getStorageConfigList()
    },
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    handleDialogClose() {
      this.$emit('close', false)
    },
    create() {
      this.storageConfigCreationVisible = true;
    },
    showStorageConfigDetailVisible(row){
      this.storageConfigDetailsVisible = true
      let list = []
      let obj = row.options
      for (let item in obj) {
        let option = obj[item]
        for(let param in option) {
            list.push({
              key: param,
              value: option[param]
            })
        }
        this.storageConfigDetail = list
      }
    },
    getStorageConfigList() {
      const params = {
        id: this.platformDetail.id,
        pageIndex: this.searchData.pageIndex,
        pageSize: this.searchData.pageSize
      }
      getStorageConfigList(params).then(response => {
        if(response.success){
            this.storageList = response.data.platformStorageConfigs;
            this.total = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          } 
      })
    },
    confirmDeletion(row) {
      this.$confirm('是否删除存储配置？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      }).then(() => {
        this.deleteStorageConfig(row)
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消'
        });
      });
    },
    deleteStorageConfig(row) {
      const params = {
        platformId: this.platformDetail.id,
        name: row.name
      }
      deleteStorageConfig(params).then(response => {
        if (response.success) {
            this.$message.success("已删除");
            this.$emit('cancel', false);
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            })
          }
      })
    },
    parseTime(val) {
      return parseTime(val)
    },
  }
}
</script>
<style lang="scss" scoped>
  .Wrapper {
    margin: 20px !important;
  }
  .block {
    float: right;
  }
  .create {
    float: right;
  }
</style>