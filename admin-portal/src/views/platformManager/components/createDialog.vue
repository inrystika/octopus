<template>
  <div>
    <el-dialog
      title="创建平台"
      width="35%"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose" 
      :close-on-click-modal="false"
    >
      <el-form
        :model="ruleForm" 
        :rules="rules" 
        ref="ruleForm" 
        label-width="100px" 
        class="demo-ruleForm"
      >
        <el-form-item label="平台名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="ruleForm.name"></el-input>
        </el-form-item>
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
  name: "createDialog",
  data() {
    return {
      formLabelWidth: '120px',
      CreateFormVisible: true,
      ruleForm: {
        name: "",
        person: "",
        connect: undefined,
        pool: ""
      },
      rules: {
        name: [
          {
            required: true,
            message: "请输入名称",
            trigger: "blur"
          }
        ],
        pool: [
          {
            required: true,
            message: "请选择资源",
            trigger: "blur"
          }
        ]
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