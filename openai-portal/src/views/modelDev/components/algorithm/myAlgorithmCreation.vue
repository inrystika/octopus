<template>
  <div>
    <el-dialog
      title="创建我的算法"
      width="35%"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px">
        <el-form-item label="算法名称" :label-width="formLabelWidth" prop="algorithmName">
          <el-input v-model="ruleForm.algorithmName" :disabled="disabled" placeholder="请输入算法名称"></el-input>
        </el-form-item>
        <el-form-item label="算法描述" :label-width="formLabelWidth" prop="desc">
          <el-input
            :disabled="disabled"
            :autosize="{ minRows: 2, maxRows: 4}"
            placeholder="请输入算法描述"
            maxlength="300"
            show-word-limit
            v-model="ruleForm.desc"
          ></el-input>
        </el-form-item>
        <el-form-item label="模型名称" :label-width="formLabelWidth" prop="modelName">
          <el-input v-model="ruleForm.modelName" :disabled="disabled" placeholder="请输入模型名称"></el-input>
        </el-form-item>
        <el-form-item :label-width="formLabelWidth">
          <div v-show="show">
            <span>是否上传代码？</span>
            <br/>
            <el-button type="primary" @click="nextStep('ruleForm')">是</el-button>
            <el-button @click="noUpload">否</el-button>
          </div>
        </el-form-item>
        <el-form-item label="上传代码包" :label-width="formLabelWidth" prop="path" v-if="showUpload">
          <upload        
            :uploadData="uploadData" 
            @confirm="confirm" 
            @cancel="cancel"   
            v-model="ruleForm.path"
          >
          </upload>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer" v-show="showConfirm">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="submit('ruleForm')">创 建</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import upload from '@/components/upload/index.vue'
import { addMyAlgorithm } from "@/api/modelDev";
import { getErrorMsg } from '@/error/index'
export default {
  name: "myAlgorithmCreation",
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
      isEmpty: false,
      disabled: false,
      showUpload: false,
      show:true,
      showConfirm:false,
      ruleForm: {
        algorithmName: "",
        modelName: '',
        desc: "",
        path: ""
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
      formLabelWidth: "120px"
    };
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    handleDialogClose() {
      this.$emit("close", false);
    },
    noUpload(){
      this.show = false;
      this.showConfirm = true;
      this.isEmpty = true
      
    },
    nextStep(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          const param = {
            spaceId: '',
            userId: null,
            IsPrefab: false,
            AlgorithmName: this.ruleForm.algorithmName,
            AlgorithmDescript: this.ruleForm.desc,
            modelname: this.ruleForm.modelName,
            isEmpty: this.isEmpty
          }
          addMyAlgorithm(param).then(response => {
            if(response.success){
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
        if(valid){
          const param = {
            spaceId: '',
            userId: null,
            IsPrefab: true,
            AlgorithmName: this.ruleForm.algorithmName,
            AlgorithmDescript: this.ruleForm.desc,
            modelname: this.ruleForm.modelName,
            isEmpty: this.isEmpty
          }
          addMyAlgorithm(param).then(response => {
            if(response.success){
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
    }
  }
};
</script>