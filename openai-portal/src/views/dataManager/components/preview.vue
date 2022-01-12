<template>
  <div>
    <el-dialog
      title="预览"
      :visible.sync="dialogTableVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
      <el-table :data="preList" height="300">
        <el-table-column property="name" label="名称" />
        <el-table-column property="type" label="类型" />
      </el-table>
    </el-dialog>
  </div>
</template>

<script>
import { previewDataset } from '@/api/datasetManager.js'
export default {
  name: "Preview",
  props: {
    row: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      dialogTableVisible: true,
      data: undefined,
      preList: []
    }
  },
  created() {
    this.data = this.row
    this.getPreList()
  },
  methods: {
    getPreList() {
      const param = {
        datasetId: this.data.datasetId,
        version: this.data.version
      }
      previewDataset(param).then(response => {
        if (response.success) {
          this.preList = response.data.files
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          });
        }
      })
    },
    handleDialogClose() {
      this.$emit('close', false)
    }
  }
}
</script>