<template>
  <div>
    <el-dialog title="创建新版本" width="650px" :visible.sync="CreateFormVisible" :before-close="handleDialogClose"
      :close-on-click-modal="false">
      <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px">
        <el-form-item label="名称：" :label-width="formLabelWidth" prop="name">
          <el-input v-model="ruleForm.name" :disabled="true" />
        </el-form-item>
        <el-form-item label="描述：" :label-width="formLabelWidth" prop="desc">
          <el-input v-model="ruleForm.desc" :autosize="{ minRows: 2, maxRows: 4}" placeholder="请输入算法描述" maxlength="300"
            show-word-limit />
        </el-form-item>
        <el-form-item :label-width="formLabelWidth">
          <el-button v-show="!showUpload" type="text" @click="nextStep('ruleForm')">下一步</el-button>
        </el-form-item>
        <el-form-item v-if="showUpload" label="代码包上传" :label-width="formLabelWidth" prop="path">
          <upload v-model="ruleForm.path" :upload-data="uploadData" @confirm="confirm" @cancel="cancel" />
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
  import upload from '@/components/upload/index.vue'
  import { addPreAlgorithmVersion } from "@/api/modelDev.js";
  import { getErrorMsg } from '@/error/index'
  export default {
    name: "PreAlgorithmVersionCreation",
    components: {
      upload
    },
    props: {
      row: {
        type: Object,
        default: () => { }
      },
      dialogType: {
        type: Boolean,
        default: ""
      }
    },
    data() {
      return {
        showUpload: false,
        uploadData: { data: {}, type: undefined },
        ruleForm: {
          desc: "",
          path: ""
        },
        rules: {
          path: [
            {
              required: true,
              message: "请选择基础版本",
              trigger: "blur"
            }
          ]
        },
        CreateFormVisible: true,
        pageIndex: 1,
        pageSize: 20,
        formLabelWidth: "120px",
        algorithmList: []
      }
    },
    created() {
      this.ruleForm.name = this.row.algorithmName

    },
    methods: {
      getErrorMsg(code) {
        return getErrorMsg(code)
      },
      handleDialogClose() {
        this.$emit("close", false);
      },
      nextStep(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.showUpload = true
            const param = {
              algorithmDescript: this.ruleForm.desc,
              algorithmId: this.row.algorithmId,
              oriVersion: this.row.algorithmVersion
            }
            addPreAlgorithmVersion(param).then(response => {
              if (response.success) {
                this.uploadData.type = "newPreAlgorithmVersion"
                this.uploadData.algorithmId = response.data.algorithmId
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
      cancel() {
        this.$emit("cancel", false);
      },
      confirm(val) {
        this.$emit("confirm", val);
      }
    }
  }
</script>