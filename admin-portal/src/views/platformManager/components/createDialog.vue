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
        <el-form-item label="联系人" :label-width="formLabelWidth" prop="contactName">
          <el-input v-model="ruleForm.contactName"></el-input>
        </el-form-item>
        <el-form-item label="联系方式" :label-width="formLabelWidth" prop="contactInfo">
          <el-input v-model="ruleForm.contactInfo"></el-input>
        </el-form-item>
        <el-form-item label="资源池" :label-width="formLabelWidth" prop="resourcePool">
          <el-select v-model="ruleForm.resourcePool" placeholder="请选择">
            <el-option 
              v-for="item in resourcePools" 
              :key="item.id" 
              :label="item.id" 
              :value="item.name"
            >
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="create('ruleForm')">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { createPlatform } from "@/api/platformManager"
import { getResourcePool } from '@/api/resourceManager.js'
export default {
  name: "createDialog",
  data() {
    return {
      formLabelWidth: '120px',
      CreateFormVisible: true,
      resourcePools: [],
      ruleForm: {
        name: "",
        contactName: "",
        contactInfo: undefined,
        resourcePool: ""
      },
      rules: {
        name: [
          {
            required: true,
            message: "请输入名称",
            trigger: "blur"
          }
        ],
        resourcePool: [
          {
            required: true,
            message: "请选择资源",
            trigger: "blur"
          }
        ]
      }
    }
  },
  created() {
    this.getResourcePool();
  },
  methods: {
    handleDialogClose() {
      this.$emit('close', false)
    },
    cancel() {
      this.$emit('cancel', false)
    },
    getResourcePool() {
      getResourcePool().then(response => {
        if (response.success) {
          if (response.data !== null && response.data.resourcePools !== null) {
            this.resourcePools = response.data.resourcePools
          }
          else {
            this.resourcePools = []
          }
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          });
        }
      })
    },
    create(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          createPlatform(this.ruleForm).then(response => {
            if(response.success) {
              this.$message.success("创建成功");
              this.$emit('confirm', false)
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
        } else {
          return false
        }
      })
    }
  }
}
</script>
<style lang="scss" scoped>
</style>