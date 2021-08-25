<template>
  <div>
    <el-dialog
      title="创建预置算法"
      width="35%"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px">
        <el-form-item label="算法名称" :label-width="formLabelWidth" prop="algorithmName">
          <el-input v-model="ruleForm.algorithmName" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="描述" :label-width="formLabelWidth" prop="desc">
          <el-input v-model="ruleForm.algorithmDescript" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="模型名称" :label-width="formLabelWidth" prop="modelName">
          <el-input v-model="ruleForm.modelName" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="上传代码包" :label-width="formLabelWidth" prop="path">
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
export default {
  name: "reuploadAlgorithm",
  components: {
    upload,
  },
  props: {
  data: {
    type: Object,
    default: {}
  }
  },
  data() {
    return {
      showUpload: false,
      ruleForm: {
        path: '',
      },
      uploadData: { data: {}, type: undefined },
      rules: {
        path: [
          {
            required: true,
            message: "请上传数据集",
            trigger: "change"
          }
        ],
      },
      CreateFormVisible: true,
      formLabelWidth: "120px"
    };
  },
  created(){
    let {algorithmName,algorithmDescript,modelName} = this.data
    this.ruleForm = {algorithmName,algorithmDescript,modelName}
    this.uploadData.algorithmId = this.data.algorithmId
    this.uploadData.version = this.data.algorithmVersion
    this.uploadData.type = "newPreAlgorithm"
  },
  methods: {
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