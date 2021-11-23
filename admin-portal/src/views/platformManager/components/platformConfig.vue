<template>
  <div>
    <el-dialog
      title="平台配置"
      :visible.sync="createFormVisible"
      :before-close="handleDialogClose" 
      :close-on-click-modal="false"
    >
      <el-form 
        :model="ruleForm"
        :rules="rules"
        ref="ruleForm"        
      >
        <el-form-item>
          <div v-for="(item, index) in ruleForm.platformConfig" :key="index">
            <el-form-item>
              <strong>{{item.key+": "}}</strong>
              <el-popover
                placement="top"
                width="400"
                trigger="hover"
                :content="item.desc"
                style="margin-right:2%"
              >
                <i v-if="item.desc?true:false" style="color:orange" class="el-icon-question" slot="reference"></i>
              </el-popover>             
              <el-input v-if="item.type === 'input'" v-model="item.value" style="width: 40%;"></el-input>
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
import { getPlatformConfigKey, getPlatformConfigValue, updatePlatformConfig } from "@/api/platformManager"
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
      createFormVisible: true,
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
            this.ruleForm.platformConfig.push({
              key: item.key,
              value: obj[item.key],
              type: item.type,
              options: item.options,
              title: item.title
            })
          }
        })
      } else {
        this.platformConfigKeyList.map(item => {
          this.ruleForm.platformConfig.push({
              key: item.key,
              value: "",
              type: item.type,
              options: item.options,
              title: item.title,
              desc: item.desc
            })
        })
      }
      // this.ruleForm.platformConfig.push({
      //   key: "test1",
      //   value: "",
      //   type: "radio",
      //   options: "yes",
      //   title: "test title",
      //   desc: "this is desc"
      // })
    },
    update(formName) {
      const params = {}
      this.ruleForm.platformConfig.map(item => {
        params[item.key] = item.value?item.value:item.options
      })
      updatePlatformConfig(this.platformDetail.id,params).then(response => {
        if(response.success) {            
          this.$message.success("平台配置更新成功");
          this.$emit('confirm', false)
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          });
        }
      })
    },
    cancel() {
      this.$emit('cancel', false)
    },
  }
}
</script>
<style lang="scss" scoped>
</style>