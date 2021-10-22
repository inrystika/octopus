<template>
  <div>
    <el-tabs class="Wrapper">
      <el-tab-pane label="平台列表">
        <div class="create">
            <el-button type="primary" @click="create">创建</el-button>
        </div>
        <el-table 
          :data="tableData" 
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
              <span>{{ scope.row.person }}</span>
            </template>
          </el-table-column>
          <el-table-column label="联系方式" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.connect }}</span>
            </template>
          </el-table-column>
          <el-table-column label="创建时间" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.created }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center">
            <template slot-scope="scope">
              <el-button type="text" @click="detail">详情</el-button>
              <el-button type="text" @click="edit">编辑</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>

    <createDialog
      v-if="createVisible"
      @cancel="cancel" 
      @confirm="confirm" 
      @close="close"
    />

    <detailsDialog
      v-if="detailsVisible"
      @cancel="cancel" 
      @confirm="confirm" 
      @close="close"
    />

    <editDialog
      v-if="editVisible"
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
export default {
  name: "platform",
  components: {
    createDialog,
    detailsDialog,
    editDialog
  },
  data() {
    return {
      createVisible: false,
      detailsVisible: false,
      editVisible: false,
      tableData: [
        {
          name: "test1",
          person: "Jack1",
          connect: "123@163.com",
          created: "2021-03-25 12:23:12"
        },
        {
          name: "test2",
          person: "Jack2",
          connect: "123@163.com",
          created: "2021-03-25 12:23:12"
        },
      ]
    }
  },
  methods: {
    cancel(val) {
      this.createVisible = val
      this.detailsVisible = val
      this.editVisible = val
    },
    confirm(val) {
      this.createVisible = val
      this.detailsVisible = val
      this.editVisible = val
    },
    close(val) {
      this.createVisible = val
      this.detailsVisible = val
      this.editVisible = val
    },
    create() {
      this.createVisible = true
    },
    detail() {
      this.detailsVisible = true
    },
    edit() {
      this.editVisible = true
    },
  }
}
</script>
<style lang="scss" scoped>
  .Wrapper {
    margin: 15px!important;
    background-color:#fff;
    padding: 20px;
    min-height: 900px
  }
  .create {
    float: right;
  }
</style>