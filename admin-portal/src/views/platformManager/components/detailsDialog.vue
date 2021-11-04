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
        :model="ruleForm"  
        ref="ruleForm" 
        label-width="100px" 
        class="demo-ruleForm"
      >
        <el-form-item label="平台名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="ruleForm.name" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="联系人" :label-width="formLabelWidth" prop="contactName">
          <el-input v-model="ruleForm.contactName" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="联系方式" :label-width="formLabelWidth" prop="contactInfo">
          <el-input v-model="ruleForm.contactInfo" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="资源池" :label-width="formLabelWidth" prop="resourcePool">
          <el-input v-model="ruleForm.resourcePool" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="客户端ID" :label-width="formLabelWidth" prop="id">
          <el-input v-model="ruleForm.id" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="客户端Secret" :label-width="formLabelWidth" prop="clientSecret">
          <el-input v-model="ruleForm.clientSecret" :disabled="true"></el-input>
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="confirm">确 定</el-button>
      </div>
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
      ruleForm: {
        name: "",
        contactName: "",
        contactInfo: undefined,
        resourcePool: "",
        id: "",
        clientSecret: ""
      }
    }
  },
  created() {
    let { name, contactName, contactInfo, resourcePool, id, clientSecret } = this.platformDetail
    this.ruleForm = { name, contactName, contactInfo, resourcePool, id, clientSecret }
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
<style lang="scss" scoped>
</style>