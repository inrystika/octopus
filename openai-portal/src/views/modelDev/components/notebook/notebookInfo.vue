<template>
  <div>
    <el-dialog
      title="启动信息"
      width="55%"
      :visible.sync="infoVisible"
    >
      <el-card>
      </el-card>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "notebookInfo",
  props: {

  },
  data() {
    return {
      infoVisible: true,
    }
  }
}
</script>