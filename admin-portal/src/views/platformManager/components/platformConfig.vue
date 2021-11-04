<template>
  <div>
    <el-dialog
      title="平台配置"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose" 
      :close-on-click-modal="false"
    >
      <el-form 
        :model="ruleForm"
        :rules="rules"
        ref="ruleForm"        
      >
        <el-form-item prop="platformConfig">
          <div v-for="(item, index) in ruleForm.platformConfig" :key="index">
            <el-form-item :label="item.title">
              <el-popover
                placement="top"
                width="400"
                trigger="hover"
                :content="item.title">
                <i class="el-icon-question" slot="reference"></i>
              </el-popover>
              <el-input v-if="item.type === 'input'" v-model="item[item.key]" style="width: 40%;"></el-input>
              <el-radio-group v-if="item.type === 'radio'" v-model="item.options">
                <el-radio :label="'yes'"></el-radio>
                <el-radio :label="'no'"></el-radio>
              </el-radio-group>
            </el-form-item>
          </div>
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="update('ruleForm')">确 定</el-button>
      </div>

    </el-dialog>
  </div>
</template>
<script>
import { getPlatformConfigKey, getPlatformConfigValue } from "@/api/platformManager"
import { getErrorMsg } from '@/error/index'
export default {
  name: "platformConfig",
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
      platformConfigKeyList: [],
      ruleForm: {
        platformConfig: []
      },
      rules: {
        platformConfig: {
          required: true, message: '请选择配置信息', trigger: ['change', 'blur']
        },
      }
    }
  },
  created() {
    this.getPlatformConfigKey()
    this.getPlatformConfigValue()
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    handleDialogClose() {
      this.$emit('close', false)
    },
    getPlatformConfigKey() {
      getPlatformConfigKey().then(response => {
        if (response.success) {
          this.platformConfigKeyList = response.data.configKeys
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          });
        }
      })
    },
    getPlatformConfigValue() {
      const platformId = this.platformDetail.id
      getPlatformConfigValue(platformId).then(response => {
        if (response.success) {
          let configValue = response.data.config
          this.judgeObjectEmpty(configValue)
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          });
        }
      })
    },
    judgeObjectEmpty(obj){
      if(obj && Object.getOwnPropertyNames(obj).length) {
        this.platformConfigKeyList.map(item => {
          if (obj[item.key]) {
            item[item.key] = obj[item.key]
          }
        })
      } else {
        this.platformConfigKeyList.map(item => {
          item[item.key] = ""
        })
      }
      this.ruleForm.platformConfig = this.platformConfigKeyList
      // this.ruleForm.platformConfig.push({
      //   desc:"this is desc",key:"this.is key",options:"yes", title:"this is title",type:"radio"
      // })
      // console.log("platformConfigKeyList",this.platformConfigKeyList)
    },
    cancel() {
      this.$emit('cancel', false)
    },
  }
}
</script>
<style lang="scss" scoped>
</style>