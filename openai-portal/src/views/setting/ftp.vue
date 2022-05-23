<template>
  <div class="Wrapper">
    <el-row :gutter="20">
      <div style="float:left;width:126px;height:40px;text-align:center;padding-top:12px;">ftp设置</div>
    </el-row>
    <el-divider class="demo"></el-divider>
    <el-row :gutter="20">
      <el-col :span="12" :offset="6">
        <el-form ref="ftpForm" :rules="ftpRules" :model="ftpForm" style="margin: 0 auto;">
          <el-form-item label="ftp账号:" prop="ftpUserName">
            <!-- <el-tooltip content="ftp账号与启智章鱼账号相互独立" placement="top">
              <i class="el-icon-info" style="color: #f7c324"></i>
            </el-tooltip> -->
            <el-input v-model="ftpForm.ftpUserName" placeholder="请填写账号" minlength="4" maxlength="30" show-word-limit/>
          </el-form-item>
          <el-form-item label="ftp密码:" prop="ftpPassword">
            <el-input v-model="ftpForm.ftpPassword" :type="[passFlag?'text':'password']" placeholder="请输入密码" minlength="8" maxlength="30" show-word-limit>
              <i slot="suffix" :class="[passFlag?'el-icon-minus':'el-icon-view']" style="margin-top:8px;font-size:18px;" autocomplete="auto" @click="passFlag=!passFlag" />
            </el-input>
          </el-form-item>
          <el-button style="float: right;" type="primary" @click="submit">创建</el-button>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import { updateUserFtpAccount } from "@/api/setting"
import { clearProgress } from '@/utils/index.js'
import { getInfo } from '@/api/Home'
import { getToken } from '@/utils/auth'
export default {
  name: 'ftp',
  data() {
    var checkName = (rule, value, callback) => {
      const regName = /^[a-zA-Z][0-9a-zA-Z_]{4,30}$/;
      if (regName.test(value)) {
        return callback();
      }
      callback(new Error("账号由字母开头，长度4-30个字符，允许字母数字下划线"));
    };
    return {
      passFlag: false,
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
      formLabelWidth: "120px",
    }
  },
  created() {
    getInfo(getToken()).then(response => {
      if (!response.data) {
        this.ftpForm.ftpUserName = ''
      }
      this.ftpForm.ftpUserName = response.data.user.ftpUserName
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
  .demo {
    margin:5px 0 5px 0 !important;
  }
</style>