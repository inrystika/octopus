<template>
  <div>
    <el-dialog title="创建我的算法" width="650px" :visible.sync="CreateFormVisible" :before-close="handleDialogClose"
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
        <el-form-item label="算法框架" :label-width="formLabelWidth" >
          <el-select v-model="ruleForm.frameworkId" placeholder="请选择算法框架">
            <el-option v-for="item in optionFrame" :key="item.id" :label="item.lableDesc" :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="算法描述" :label-width="formLabelWidth" prop="desc">
          <el-input v-model="ruleForm.desc" :disabled="disabled" :autosize="{ minRows: 2, maxRows: 4}"
            placeholder="请输入算法描述" maxlength="300" show-word-limit />
        </el-form-item>
        <el-form-item label="模型名称" :label-width="formLabelWidth" prop="modelName">
          <el-input v-model="ruleForm.modelName" :disabled="disabled" placeholder="请输入模型名称" />
        </el-form-item>
        <el-form-item :label-width="formLabelWidth">
          <div v-show="show">
            <span>是否上传代码？</span>
            <br>
            <el-button type="primary" @click="nextStep('ruleForm')">是</el-button>
            <el-button @click="noUpload">否</el-button>
          </div>
        </el-form-item>
        <el-form-item v-if="showUpload" label="上传代码包" :label-width="formLabelWidth" prop="path">
          <upload v-model="ruleForm.path" :upload-data="uploadData" @confirm="confirm" @cancel="cancel"
            @upload="isCloseX" />
        </el-form-item>
      </el-form>
      <span v-show="showConfirm" slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="submit('ruleForm')">创 建</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
  import upload from '@/components/upload/index.vue'
  import { addMyAlgorithm, algorithmType,algorithmFrame } from "@/api/modelDev";
  export default {
    name: "MyAlgorithmCreation",
    components: {
      upload
    },
    data() {
      return {
        isEmpty: false,
        disabled: false,
        showUpload: false,
        show: true,
        showConfirm: false,
        ruleForm: {
          algorithmName: "",
          modelName: '',
          desc: "",
          path: "",
          applyId: '',
          frameworkId:''
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
            },
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
          ] 
        },
        CreateFormVisible: true,
        formLabelWidth: "120px",
        close: true,
        optionType: [],
        optionFrame:[]
      };
    },
    created() {
      this.algorithmType()
      this.algorithmFrame()
    },
    methods: {
      handleDialogClose() {
        this.$emit("close", false);
      },
      noUpload() {
        this.show = false;
        this.showConfirm = true;
        this.isEmpty = true
      },
      nextStep(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            const param = {      
              algorithmName: this.ruleForm.algorithmName,
              algorithmDescript: this.ruleForm.desc,
              modelName: this.ruleForm.modelName,
              isEmpty: this.isEmpty,
              applyId:this.ruleForm.applyId,
              frameworkId:this.ruleForm.frameworkId
            }
            addMyAlgorithm(param).then(response => {
              if (response.success) {
                this.show = false
                this.showUpload = true;
                this.disabled = true
                this.uploadData.AlgorithmId = response.data.algorithmId
                this.uploadData.Version = response.data.version
                this.uploadData.type = 'myAlgorithmCreation'
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
      },
      submit(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            const param = {
              spaceId: '',
              userId: null,
              IsPrefab: true,
              AlgorithmName: this.ruleForm.algorithmName,
              AlgorithmDescript: this.ruleForm.desc,
              modelname: this.ruleForm.modelName,
              isEmpty: this.isEmpty,
              applyId:this.ruleForm.applyId,
              frameworkId:this.ruleForm.frameworkId
            }
            addMyAlgorithm(param).then(response => {
              if (response.success) {
                this.disabled = true
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
            this.$message.error("请填写数据");
          }
        })
      },
      isCloseX(val) {
        this.close = val
      },
      // 获取算法类型
      algorithmType() {
        algorithmType({ pageIndex: 1, pageSize: 50 }).then(response => {
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
        algorithmFrame({ pageIndex: 1, pageSize: 50 }).then(response => {
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