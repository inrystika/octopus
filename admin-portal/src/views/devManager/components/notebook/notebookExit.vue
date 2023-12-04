<template>
  <div>
    <div>
      <el-row>
        <el-col :span="12">
          <div>任务名称:<span>{{ notebookInfo.name }}</span></div>
        </el-col>
        <el-col :span="12">
          <div>是否分布式:<span>{{ this.notebookInfo.tasks.length > 1 ? '是' : '否'}}</span></div>
        </el-col>

      </el-row>
      <el-input v-model="exitMsg" type="textarea" :readonly="true" :rows="20" />
    </div>

    <div slot="footer" class="dialog-footer" />
  </div>
</template>

<script>
  export default {
    name: "NotebookExit",
    props: {
      notebookData: {
        type: Object,
        default: () => { }
      }
    },
    data() {
      return {
        notebookInfo: {},
        exitMsg: '暂无相关信息'
      }
    },
    created() {
      this.notebookInfo = this.notebookData
      if(this.notebookInfo.exitMsg) {
        this.exitMsg = this.notebookInfo.exitMsg
      }
    }
  }
</script>

<style lang="scss" scoped>
  .el-col {
    margin: 10px 0 20px 0;
    font-size: 15px;
    font-weight: 800;

    span {
      font-weight: 400;
      margin-left: 20px
    }
  }


  .block {
    float: right;
    margin: 20px;
  }
</style>