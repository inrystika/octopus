<template>
  <div>
    <el-dialog
      title="存储配置创建"
      width="35%"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose" 
      :close-on-click-modal="false"
      :append-to-body="true"
    >
      <el-form 
        :model="ruleForm"
        :rules="rules"
        ref="ruleForm"        
      >
        <el-form-item label="平台" :label-width="formLabelWidth" prop="platform">
          <el-input v-model="ruleForm.platform" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="名称" :label-width="formLabelWidth" prop="storageConfigName">
          <el-input v-model="ruleForm.storageConfigName" placeholder="请输入名称"></el-input>
        </el-form-item>
        <el-form-item label="存储类型" :label-width="formLabelWidth" prop="type">
          <el-select v-model="ruleForm.type" placeholder="请选择类型">
            <el-option label="juicefs" value="juicefs"></el-option>
          </el-select>
        </el-form-item>
        <el-divider></el-divider>
        <el-form-item label="name" :label-width="formLabelWidth">
          <el-input v-model="ruleForm.name" placeholder="请输入名称"></el-input>
        </el-form-item>
        <el-form-item label="Meta Url" :label-width="formLabelWidth">
          <el-input v-model="ruleForm.metaUrl" placeholder="请输入名称"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
          <el-button @click="cancel">取 消</el-button>
          <el-button type="primary" @click="create('ruleForm')">创建</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import { createStorageConfig } from "@/api/platformManager"
export default {
  name: "storageConfig",
  props: {
    platformDetail: {
      type: Object,
      default: () => {}
    },
  },
  data() {
    return {
      formLabelWidth: "120px",
      CreateFormVisible: true,
      ruleForm: {
        platform: this.platformDetail.name,
        storageConfigName: "",
        type: "juicefs",
        metaUrl: "",
        name: "",
      },
      rules: {
        storageConfigName: [
          { required: true, message: '请填写配置名称', trigger: 'blur' }
        ],
        type: [
          { required: true, message: '请选择存储配置类型', trigger: 'change' }
        ],
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
    create() {
      this.$refs['ruleForm'].validate((valid) => {
        if (valid) {
          const params = {
            name: this.ruleForm.storageConfigName,
            type: this.ruleForm.type,
            options: {
              juicefs: {
                metaUrl: this.ruleForm.metaUrl,
                name: this.ruleForm.name
              }
            }
          }
          createStorageConfig(this.platformDetail.id,params).then(response => {
            if (response.success) {
              this.$message({
                  message: '创建存储配置成功',
                  type: 'success'
              });
              this.$emit('confirm', false)
            }
            else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
        }
      })
    }
  }
}
</script>
<style  lang="scss" scoped>

</style>