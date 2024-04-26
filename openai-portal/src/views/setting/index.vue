<template>
  <div class="Wrapper">
    <el-row :gutter="20">
      <div style="float:left;width:126px;height:20px;text-align:left;padding-top:12px;padding-left: 10px">设置</div>
    </el-row>
    <el-divider class="dividerClass"></el-divider>
    <el-row :gutter="20">
      <el-col :span="20" :offset="2">
        <el-card shadow="never">
          <div slot="header">
            <el-row :gutter="20">
              <el-col >
                <span>sftp</span>
              </el-col>
            </el-row>
          </div>
          <el-row :gutter="20">
            <el-col >
              <el-form ref="ftpForm" :rules="ftpRules" :model="ftpForm" >
                <el-form-item label="sftp账号:" prop="ftpUserName">
                  <el-input v-model="ftpForm.ftpUserName" placeholder="请填写账号" :disabled="isShow" minlength="4" maxlength="30" show-word-limit/>
                </el-form-item>
                <el-form-item label="sftp密码:" prop="ftpPassword">
                  <el-input v-model="ftpForm.ftpPassword" show-password placeholder="请输入密码" minlength="8" maxlength="30" show-word-limit>
                  </el-input>
                </el-form-item>
                <el-button style="float: right;" type="primary" @click="submit">保存</el-button>
              </el-form>
            </el-col>
          </el-row>
        </el-card>
        <!-- <el-divider></el-divider> -->
        <el-card shadow="never" style="margin-top: 20px">
          <div slot="header">
            <el-row :gutter="20">
              <el-col >
                <span>邮件通知</span>
              </el-col>
            </el-row>
          </div>
          <el-row :gutter="20">
            <el-col >
              是否邮件通知：
              <el-switch
                @change="notifyEmail"
                v-model="isEmailNotify"
                active-text="是"
                inactive-text="否">
              </el-switch>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>

  </div>
</template>
<script>
import { updateUserFtpAccount, updateEmailNotify } from "@/api/setting"
import { clearProgress } from '@/utils/index.js'
import { getInfo } from '@/api/Home'
import { getToken } from '@/utils/auth'
export default {
  name: 'setting',
  data() {
    var checkName = (rule, value, callback) => {
      const regName = /^[a-zA-Z][0-9a-zA-Z_]{3,30}$/;
      if (regName.test(value)) {
        return callback();
      }
      callback(new Error("账号由字母开头，长度4-30个字符，允许字母数字下划线"));
    };
    return {
      isEmailNotify: false,
      isShow: false,
      ftpRules: {
        ftpUserName: [
          { required: true, message: "请输入账号", trigger: "blur" },
          { validator: checkName, trigger: "blur" }

        ], ftpPassword: [
          { required: true, message: '请输入密码！', trigger: 'blur' },
          { min: 8, max: 30, message: '密码长度在8-30位之间', trigger: 'blur' }
        ]
      },
      ftpForm: {
        ftpUserName: '',
        ftpPassword: ''
      },
    }
  },
  created() {
    getInfo(getToken()).then(response => {
      if (response.data.user.ftpUserName) {
        this.ftpForm.ftpUserName = response.data.user.ftpUserName
        this.isShow = true
        this.isEmailNotify = response.data.user.emailNotify
      }     
    })
  },
  mounted() {
    window.addEventListener("beforeunload", (e) => {
      clearProgress()
    });
  },
  destroyed() {
    window.removeEventListener("beforeunload", (e) => {
      clearProgress()
    });
  },
  methods: {
    notifyEmail(res) {
      console.log({res})
      const params = {
        emailNotify: this.isEmailNotify
      }
      updateEmailNotify(params).then((response) => {
        if (response.success) {
          this.$message.success("邮件通知更新成功")
        } else {
          this.$message({
            message: this.getErrorMsg(
              response.error.subcode
            ),
            type: "warning"
          })
        }
      })
    },
    submit() {
      this.$refs.ftpForm.validate(valid => {
        if (valid) {
          const params = {
            ftpUserName: this.ftpForm.ftpUserName,
            ftpPassword: this.ftpForm.ftpPassword
          }
          updateUserFtpAccount(params).then((response) => {
            if (response.success) {
              this.$message.success("创建成功")
              this.ftpForm.ftpPassword = ''
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
<style lang="scss" scoped>
  .Wrapper {
    margin: 15px !important;
    background-color: #fff;
    padding: 20px;
    min-height: 900px;
  }
  ::v-deep .el-card__header {
    background: #f0f0f0 !important;
  }
</style>