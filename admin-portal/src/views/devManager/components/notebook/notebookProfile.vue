<template>
  <div>
    <el-row>
      <el-col :span="12">
        <div>
          任务名称:
          <span>{{ profileInfo.name }}</span>
        </div>
      </el-col>
      <el-col :span="12">
        <div>
          描述:
          <span>{{ profileInfo.desc }}</span>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <div>
          选用算法:
          <span>{{ profileInfo.algorithmName + ":" + profileInfo.algorithmVersion }}</span>
        </div>
      </el-col>
      <el-col :span="12">
        <div>
          镜像选择:
          <span>{{ profileInfo.imageName + ":" + profileInfo.imageVersion }}</span>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <div>
          选用数据集:
          <span>{{ profileInfo.datasetShow }}</span>
        </div>
      </el-col>
      <el-col :span="12">
        <div>
          是否分布式:
          <span>{{ this.profileInfo.tasks.length > 1 ? '是' : '否' }}</span>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <div>
          资源规格:
          <span>{{ profileInfo.resourceSpecName }}</span>
        </div>
      </el-col>
      <el-col :span="12">
        <div>
          任务状态:
          <span>{{ statusText[profileInfo.status][1] }}</span>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12">
        <div>
          自定义启动命令:
          <span>{{ profileInfo.command }}</span>
        </div>
      </el-col>
    </el-row>
  </div>
</template>
<script>
export default {
  name: "NotebookProfile",
  props: {
    notebookData: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      profileInfo: {},
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
    this.profileInfo = this.notebookData
    if(!this.profileInfo.datasetName) {
      this.profileInfo.datasetShow = ''
      return
    }
    this.profileInfo.datasetShow = this.profileInfo.datasetName + ":" + this.profileInfo.datasetVersion
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

  // .taskList {
  //   font-weight: 800;
  // }

  .block {
    float: right;
    margin: 20px;
  }
</style>