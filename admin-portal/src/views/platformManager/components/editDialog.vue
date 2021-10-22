<template>
  <div>
    <el-dialog
      title="编辑"
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
        <el-form-item label="联系人" :label-width="formLabelWidth" prop="person">
          <el-input v-model="ruleForm.person"></el-input>
        </el-form-item>
        <el-form-item label="联系方式" :label-width="formLabelWidth" prop="connect">
          <el-input v-model="ruleForm.connect"></el-input>
        </el-form-item>
        <el-form-item label="资源池" :label-width="formLabelWidth" prop="pool">
          <el-input v-model="ruleForm.pool"></el-input>
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
  name: "editDialog",
  data() {
    return {
      formLabelWidth: '120px',
      CreateFormVisible: true,
      ruleForm: {
        name: "",
        person: "",
        connect: undefined,
        pool: ""
      }
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
      
    }
  }
}
</script>
<style lang="scss" scoped>
</style>