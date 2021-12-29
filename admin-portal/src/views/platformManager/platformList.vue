<template>
  <div>
    <div class="create">
        <el-button type="primary" @click="create">创建</el-button>
    </div>
    <el-table 
      :data="platformList" 
      style="width: 100%;font-size: 15px;"
      :header-cell-style="{'text-align':'left','color':'black'}" 
      :cell-style="{'text-align':'left'}"
    >
      <el-table-column label="名称" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="联系人" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.contactName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="联系方式" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.contactInfo }}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.createdAt | parseTime }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <el-button type="text" @click="detail(scope.row)">详情</el-button>
          <el-button type="text" @click="edit(scope.row)">编辑</el-button>
          <el-button type="text" @click="getPlatformConfig(scope.row)">平台配置</el-button>
          <el-button type="text" @click="showStorageConfigList(scope.row)">存储配置</el-button>
        </template>
      </el-table-column>
    </el-table>
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

    <createDialog
      v-if="createVisible"
      @cancel="cancel" 
      @confirm="confirm" 
      @close="close"
    />

    <detailsDialog
      v-if="detailsVisible"
      :platform-detail="platformDetail"
      @close="close"
    />

    <editDialog
      v-if="editVisible"
      :platform-detail="platformDetail"
      @cancel="cancel" 
      @confirm="confirm" 
      @close="close"
    />

    <platformConfig
      v-if="platformConfigVisible"
      :platform-detail="platformDetail"
      @cancel="cancel" 
      @confirm="confirm" 
      @close="close"
    />
    
    <storageConfigList
      v-if="storageConfigListVisible"
      :platform-detail="platformDetail"
      @cancel="cancel" 
      @confirm="confirm" 
      @close="close"
    />
  </div>
</template>
<script>
import searchForm from '@/components/search/index.vue'
import createDialog from "./components/createDialog.vue"
import detailsDialog from "./components/detailsDialog.vue"
import editDialog from "./components/editDialog.vue"
import platformConfig from "./components/platformConfig.vue"
import storageConfigList from "./components/storageConfigList.vue"
import { getPlatformList } from "@/api/platformManager"
import { getErrorMsg } from '@/error/index'
export default {
  name: "platformList",
  components: {
    createDialog,
    detailsDialog,
    editDialog,
    platformConfig,
    storageConfigList
  },
  data() {
    return {
      createVisible: false,
      detailsVisible: false,
      editVisible: false,
      platformConfigVisible: false,
      platformList: [],
      platformDetail: {},
      storageConfigListVisible: false,
      total: 0,
      searchData: {
        pageIndex: 1,
        pageSize: 10,
      }
    }
  },
  created() {
    this.getPlatformList(this.searchData);
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    getPlatformList(param){
      getPlatformList(param).then(response => {
          if(response.success){
            this.platformList = response.data.platforms;
            this.total = response.data.totalSize
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }    
        })
    },
    handleSizeChange(val){
      this.searchData.pageSize = val
      this.getPlatformList(this.searchData)
    },
    handleCurrentChange(val) {
      this.searchData.pageIndex = val
      this.getPlatformList(this.searchData)
    },
    cancel(val) {
      this.createVisible = val
      this.editVisible = val
      this.platformConfigVisible = val
      this.storageConfigListVisible = val
      this.getPlatformList(this.searchData);
    },
    confirm(val) {
      this.createVisible = val
      this.editVisible = val
      this.platformConfigVisible = val
      this.storageConfigListVisible = val
      this.getPlatformList(this.searchData);
    },
    close(val) {
      this.createVisible = val
      this.detailsVisible = val
      this.editVisible = val
      this.platformConfigVisible = val
      this.storageConfigListVisible = val
      this.getPlatformList(this.searchData);
    },
    create() {
      this.createVisible = true
    },
    detail(row) {
      this.detailsVisible = true
      this.platformDetail = row
    },
    edit(row) {
      this.editVisible = true
      this.platformDetail = row
    },
    getPlatformConfig(row) {
      this.platformConfigVisible = true
      this.platformDetail = row
    },
    showStorageConfigList(row) {
      this.platformDetail = row
      this.storageConfigListVisible = true
    }
  }
}
</script>
<style lang="scss" scoped>
  .create {
    float: right;
  }
  .block {
    float: right;
    margin: 20px;
  }
</style>