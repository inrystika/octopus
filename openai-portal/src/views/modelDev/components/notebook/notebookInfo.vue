<template>
  <div>
    <el-dialog
      title="启动信息"
      width="55%"
      :visible.sync="infoVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
      <div v-html="this.initInfo"></div>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="confirm">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "notebookInfo",
  props: {
    initInfo: {
      type: String,
      default: ""
    }
  },
  data() {
    return {
      infoVisible: true,
    }
  },
  methods: {
    handleDialogClose() {
      this.$emit('close', false)
    },
    cancel() {
      this.$emit('cancel', false)
    },
    confirm() {
      this.$emit('confirm', false)
    }
  }
}
</script>