<template>
  <div>
    <el-dialog
      title="创建我的数据集"
      width="35%"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px">
        <el-form-item label="数据集名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="ruleForm.name" :disabled="disabled" placeholder="请输入数据集名称，长度在 4 到 30 个字符"></el-input>
        </el-form-item>
        <el-form-item label="数据类型" :label-width="formLabelWidth" prop="type">
          <el-select v-model="ruleForm.type" :disabled="disabled" placeholder="请选择数据集类型">
            <el-option label="图片" value="picture"></el-option>
            <el-option label="视频" value="video"></el-option>
            <el-option label="文字" value="text"></el-option>
            <el-option label="语音" value="voice"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="数据集描述" :label-width="formLabelWidth" prop="desc">
          <el-input
            :disabled="disabled"
            :autosize="{ minRows: 2, maxRows: 4}"
            placeholder="请输入数据集描述"
            maxlength="300"
            show-word-limit
            v-model="ruleForm.desc"
          ></el-input>
        </el-form-item>
        <el-form-item :label-width="formLabelWidth">
          <el-button type="text" @click="nextStep('ruleForm')" v-show="!showUpload">下一步</el-button>
        </el-form-item>
        <el-form-item label="数据集上传" :label-width="formLabelWidth" prop="path" v-if="showUpload">
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
import { createMyDataset } from "@/api/datasetManager.js"
import { getErrorMsg } from '@/error/index'
export default {
  name: "myDatasetCreation",
  components: {
    upload,
  },
  props: {
    // row: {
    //   type: Object,
    //   default: {}
    // }
  },
  data() {
    return {
      showUpload: false,
      disabled: false,
      uploadData: { data: {}, type: undefined },
      ruleForm: {
        name: '',
        desc: '',
        type: '',
        path: ''
      },
      rules: {
        name: [
          { required: true, message: "请输入数据集名称", trigger: "blur"},
          { min: 4, max: 30, message: "长度在 4 到 30 个字符", trigger: "blur"}
        ],
        type: [
          {
            required: true,
            message: "请选择数据集类型",
            trigger: "change"
          }
        ],
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
    };
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    nextStep(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          delete this.ruleForm.path
          const param = {
            name: this.ruleForm.name,
            type: this.ruleForm.type,
            desc: this.ruleForm.desc
          }
          createMyDataset(param).then(response => {
            if(response.success) {
              this.showUpload = true
              this.disabled = true
              this.uploadData.id = response.data.id
              this.uploadData.type = "myDatasetCreation"
              this.uploadData.version = response.data.version
            } else {
              // this.showUpload = false
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
};
</script>