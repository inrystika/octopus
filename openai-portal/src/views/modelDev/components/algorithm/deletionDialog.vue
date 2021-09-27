<template>
  <div>
    <el-dialog
      title="提示"
      :close-on-click-modal="false"
      :visible.sync="centerDialogVisible"
      width="30%"
      center
      :before-close="handleDialogClose"
      :top="'40vh'"
    >
      <span>是否删除？</span>
      <span slot="footer" class="dialog-footer">
          <el-button @click="cancel">否</el-button>
          <el-button type="primary" @click="confirm">是</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "DeletionDialog",
  data() {
    return {
      centerDialogVisible: true
    }
  },
  methods: {
    cancel() {
      this.$emit('cancel', false)
    },
    confirm() {
      this.$emit('confirm', false)
    },
    handleDialogClose() {
      this.$emit('close', false)
    }
  }
}
</script>

<style lang="scss" scoped>
  .el-dialog--center .el-dialog__body {
    text-align: center;
  }
</style>