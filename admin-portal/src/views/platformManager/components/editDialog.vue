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
        :rules="rules"
        ref="ruleForm"
        label-width="100px"
        class="demo-ruleForm"
      >
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
        <el-button type="primary" @click="update('ruleForm')">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { updatePlatform } from "@/api/platformManager"
import { getResourcePool } from '@/api/resourceManager.js'
export default {
  name: "editDialog",
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
      resourcePools: [],
      ruleForm: {
        contactName: "",
        contactInfo: undefined,
        resourcePool: ""
      },
      rules: {
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
    let { contactName, contactInfo } = this.platformDetail
    this.ruleForm = { contactName, contactInfo }
    this.getResourcePool();
  },
  methods: {
    getResourcePool() {
      getResourcePool().then(response => {
        if (response.success) {
          if (response.data !== null && response.data.resourcePools !== null) {
            this.resourcePools = response.data.resourcePools
            // 判断平台资源是否存在于获取的资源池列表中；若存在，则设为默认展示；若不存在，默认展示为空。
            let resourcePool = this.resourcePools.find(item => item.name == this.platformDetail.resourcePool)
            if(resourcePool) {
              this.ruleForm.resourcePool = resourcePool.name
            }
          }
          else {
            this.resourcePools = []
          }
          this.ruleForm = {...this.ruleForm}
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          });
        }
      })
    },
    update(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          const param = this.ruleForm
          param.id = this.platformDetail.id
          updatePlatform(param).then(response => {
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
    },
    handleDialogClose() {
      this.$emit('close', false)
    },
    cancel() {
      this.$emit('cancel', false)
    }
  }
}
</script>
<style lang="scss" scoped>
</style>