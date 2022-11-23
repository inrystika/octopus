<template>
  <div>
    <el-dialog 
      title="minio用户信息" 
      width="35%" 
      :visible.sync="CreateFormVisible"           
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm">
        <el-form-item label="minio用户名称" prop="minioUserName" :label-width="formLabelWidth">
          <el-input v-model="ruleForm.minioUserName" placeholder="请填写账号" :disabled="isShow" minlength="3" maxlength="20" show-word-limit/>
        </el-form-item>
        <el-form-item label="minio账户密码" prop="minioPassword" :label-width="formLabelWidth">
          <el-input v-model="ruleForm.minioPassword" show-password placeholder="请输入密码" minlength="8" maxlength="40" show-word-limit />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="submit">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { createMinIOAccount } from "@/api/userManager.js"
export default {
  name: "minioAccount",
  props: {
    row: {
      type: Object,
      default: () => { }
    }
  },
  data() {
    // var checkName = (rule, value, callback) => {
    //   const regName = /^[a-zA-Z][0-9a-zA-Z_]{3,30}$/;
    //   if (regName.test(value)) {
    //     return callback();
    //   }
    //   callback(new Error("账号由字母开头,长度4-30个字符,允许字母数字下划线"));
    // };
    return {
      isShow: false,
      CreateFormVisible: true,
      formLabelWidth: '120px',
      ruleForm: {
        minioUserName: "",
        minioPassword: ""
      },
      rules: {
        minioUserName: [
          { required: true, message: "请输入账号", trigger: "blur" },
          { min: 3, max: 20, message: '名称长度在3-20个字符', trigger: 'blur' }

        ], minioPassword: [
          { required: true, message: '请输入密码！', trigger: 'blur' },
          { min: 8, max: 40, message: '密码长度在8-40位之间', trigger: 'blur' }
        ]
      },
    }
  },
  created () {
    if (this.row.minioUserName) {
      this.ruleForm.minioUserName = this.row.minioUserName
      this.isShow = true
    }
  },
  methods: {
      handleDialogClose() {
        this.$emit('close', false)
      },
      cancel() {
        this.$emit('cancel', false)
      },
      submit() {
        this.$refs.ruleForm.validate(valid => {
        if (valid) {
          const params = {
            minioUserName: this.ruleForm.minioUserName,
            minioPassword: this.ruleForm.minioPassword
          }
          const userId = this.row.id
          createMinIOAccount(userId,params).then((response) => {
            if (response.success) {
              this.$message.success("创建成功")
              this.ruleForm.minioPassword = ''
              this.isShow = true
              this.$emit('confirm', false)
            } else {
              this.$message({
                message: this.getErrorMsg(
                  response.error.subcode
                ),
                type: "warning"
              })
            }
          })
        } else {
          console.log('error')
          }
        })
      }
  }
}
</script>