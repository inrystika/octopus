<template>
  <div>
    <el-dialog title="创建预置算法" width="650px" :visible.sync="CreateFormVisible" :before-close="handleDialogClose"
      :close-on-click-modal="false" :show-close="close">
      <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px">
        <el-form-item label="算法名称" :label-width="formLabelWidth" prop="algorithmName">
          <el-input v-model="ruleForm.algorithmName" :disabled="disabled" placeholder="请输入算法名称" />
        </el-form-item>
        <el-form-item label="模型类别" :label-width="formLabelWidth">
          <el-select v-model="ruleForm.applyId" placeholder="请选择模型类别">
            <el-option v-for="item in optionType" :key="item.id" :label="item.lableDesc" :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="框架类型" :label-width="formLabelWidth" >
          <el-select v-model="ruleForm.frameworkId" placeholder="请选择框架类型">
            <el-option v-for="item in optionFrame" :key="item.id" :label="item.lableDesc" :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="描述" :label-width="formLabelWidth" prop="desc">
          <el-input v-model="ruleForm.desc" :autosize="{ minRows: 2, maxRows: 4}" :disabled="disabled"
            placeholder="请输入算法描述" maxlength="300" show-word-limit />
        </el-form-item>
        <el-form-item label="模型名称" :label-width="formLabelWidth" prop="modelName">
          <el-input v-model="ruleForm.modelName" :disabled="disabled" placeholder="请输入模型名称" />
        </el-form-item>
        <el-form-item :label-width="formLabelWidth">
          <el-button v-show="!showUpload" type="text" @click="nextStep('ruleForm')">下一步</el-button>
        </el-form-item>
        <el-form-item v-if="showUpload" label="上传代码包" :label-width="formLabelWidth" prop="path">
          <upload v-model="ruleForm.path" :upload-data="uploadData" @confirm="confirm" @cancel="cancel"
            @upload="isCloseX" />
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
  import upload from '@/components/upload/index.vue'
  import { addPreAlgorithm, algorithmType, frameType } from "@/api/modelDev";
  import { getErrorMsg } from '@/error/index'
  export default {
    name: "PreAlgorithmCreation",
    components: {
      upload
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
        ruleForm: {
          algorithmName: '',
          path: '',
          modelName: '',
          desc: '',
          applyId: '',
          frameworkId: ''
        },
        uploadData: { data: {}, type: undefined },
        rules: {
          algorithmName: [
            {
              required: true,
              message: "请输入算法名称",
              trigger: "blur"
            },
            {
              min: 4,
              max: 30,
              message: "长度在 4 到 30 个字符",
              trigger: "blur"
            }
          ],
          path: [
            {
              required: true,
              message: "请上传数据集",
              trigger: "change"
            }
          ],
          modelName: [
            {
              required: true,
              message: "请输入模型名称",
              trigger: "blur"
            },
            {
              min: 4,
              max: 30,
              message: "长度在 4 到 30 个字符",
              trigger: "blur"
            }
          ],
        },
        CreateFormVisible: true,
        formLabelWidth: "120px",
        close: true,
        optionType: [],
        optionFrame: []
      };
    },
    created() {
      this.algorithmType()
      this.algorithmFrame()
    },
    methods: {
      getErrorMsg(code) {
        return getErrorMsg(code)
      },
      handleDialogClose() {
        this.$emit('close', false)
      },
      nextStep(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            const param = {
              spaceId: '',
              userId: null,
              IsPrefab: true,
              algorithmName: this.ruleForm.algorithmName,
              algorithmDescript: this.ruleForm.desc,
              modelname: this.ruleForm.modelName,
              isEmpty: false,
              frameworkId: this.ruleForm.frameworkId,
              applyId: this.ruleForm.applyId
            }
            addPreAlgorithm(param).then(response => {
              if (response.success) {
                this.showUpload = true;
                this.disabled = true
                this.uploadData.algorithmId = response.data.algorithmId
                this.uploadData.version = response.data.version
                this.uploadData.type = 'newPreAlgorithm'
              } else {
                this.showUpload = false
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
      },
      isCloseX(val) {
        this.close = val
      },
      // 获取算法类型
      algorithmType() {
        algorithmType({ pageIndex: 1, pageSize: 20 }).then(response => {
          if (response.success) {
            this.optionType = response.data.lables
          } else {
            // this.showUpload = false
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      },
      // 获取算法框架
      algorithmFrame() {
        frameType({ pageIndex: 1, pageSize: 20 }).then(response => {
          if (response.success) {
            this.optionFrame = response.data.lables
          } else {
            // this.showUpload = false
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      }
    }
  };
</script>