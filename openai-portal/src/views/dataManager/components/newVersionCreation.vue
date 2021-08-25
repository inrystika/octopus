<template>
  <div>
    <el-dialog
      title="创建新版本"
      width="35%"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px">
        <el-form-item label="数据集名称" :label-width="formLabelWidth" prop="name">
            <el-input v-model="ruleForm.name" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="数据类型" :label-width="formLabelWidth" prop="type">
            <el-input v-model="ruleForm.type" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="版本描述" :label-width="formLabelWidth" prop="desc">
          <el-input
            v-model="ruleForm.desc"
            :autosize="{ minRows: 2, maxRows: 4}"
            placeholder="请输入数据集描述"
            maxlength="300"
            show-word-limit
          ></el-input>
        </el-form-item>
        <el-form-item :label-width="formLabelWidth">
          <el-button type="text" @click="nextStep('ruleForm')" v-show="!showUpload">下一步</el-button>
        </el-form-item>
        <el-form-item label='数据集上传' :label-width="formLabelWidth" prop="path" v-if="showUpload">
          <upload        
            :uploadData="uploadData" 
            @confirm="confirm" 
            @cancel="cancel"   
            v-model="ruleForm.path"
          >
          </upload>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import upload from '@/components/upload/index.vue'
import { createNewVersion } from "@/api/datasetManager.js";
import { getErrorMsg } from '@/error/index'
export default {
  name: "newVersionCreation",
  components: {
    upload,
  },
  props: {
    row: {
      type: Object,
      default: () => { }
    },
  },
  data() {
    return {
      showUpload: false,
      uploadData: { data: {}, type: undefined },
      ruleForm: {
        desc: ""
      },
      rules: {
        path: [
          {
            required: true,
            message: "请上传数据集",
            trigger: "change"
          }
        ]
      },
      CreateFormVisible: true,
      formLabelWidth: "120px"
    }
  },
  created(){
    let {name,type} = this.row
    this.ruleForm = {name,type}
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    nextStep(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.showUpload = true
          const param = {
            desc: this.ruleForm.desc,
            datasetId: this.row.id
          }
          createNewVersion(param).then(response => {
            if(response.success) {
              this.uploadData.type = "newDatasetVersionCreation"
              this.uploadData.id = response.data.datasetId
              this.uploadData.version = response.data.version
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
        } else {
          return false;
        }
      });
    },
    handleDialogClose() {
      this.$emit("close", false);
    },
    cancel() {
      this.$emit("cancel", false);
    },
    confirm(val) {
      this.$emit("confirm", val);
    }
  }
}
</script>