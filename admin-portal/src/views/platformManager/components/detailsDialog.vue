<template>
  <div>
    <el-dialog
      title="详情"
      width="35%"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose" 
      :close-on-click-modal="false"
    >
      <el-form       
        label-width="100px" 
        class="demo-ruleForm"
      >
        <el-form-item label="平台名称:" :label-width="formLabelWidth">
          {{ this.platformDetail.name }}
        </el-form-item>
        <el-form-item label="联系人:" :label-width="formLabelWidth">
          {{ this.platformDetail.contactName }}
        </el-form-item>
        <el-form-item label="联系方式:" :label-width="formLabelWidth">
          {{ this.platformDetail.contactInfo }}
        </el-form-item>
        <el-form-item label="资源池:" :label-width="formLabelWidth">
          {{ this.platformDetail.resourcePool }}
        </el-form-item>
        <el-form-item label="客户端ID:" :label-width="formLabelWidth">
          {{ this.platformDetail.id }}
        </el-form-item>
        <el-form-item label="客户端Secret:" :label-width="formLabelWidth">
          {{ this.platformDetail.clientSecret }}
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "detailsDialog",
  props: {
    platformDetail: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      formLabelWidth: '120px',
      CreateFormVisible: true,
    }
  },
  methods: {
    handleDialogClose() {
      this.$emit('close', false)
    },
  }
}
</script>
<style lang="scss" scoped>
</style>