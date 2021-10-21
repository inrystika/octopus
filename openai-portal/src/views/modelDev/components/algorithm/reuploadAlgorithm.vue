<template>
  <div>
    <el-dialog
      title="创建我的算法"
      width="35%"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
      <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px">
        <el-form-item label="算法名称" :label-width="formLabelWidth" prop="algorithmName">
          <el-input v-model="ruleForm.algorithmName" :disabled="true" />
        </el-form-item>
        <el-form-item label="算法描述" :label-width="formLabelWidth" prop="desc">
          <el-input v-model="ruleForm.algorithmDescript" :disabled="true" />
        </el-form-item>
        <el-form-item label="模型名称" :label-width="formLabelWidth" prop="modelName">
          <el-input v-model="ruleForm.modelName" :disabled="true" />
        </el-form-item>
        <el-form-item label="上传代码包" :label-width="formLabelWidth" prop="path">
          <upload
            v-model="ruleForm.path"
            :upload-data="uploadData"
            @confirm="confirm"
            @cancel="cancel"
          />
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import upload from '@/components/upload/index.vue'
export default {
  name: "ReuploadAlgorithm",
  components: {
    upload
  },
  props: {
    reuploadData: {
      type: Object,
      default: () => { }
    }
  },
  data() {
    return {
      isEmpty: false,
      showUpload: false,
      show: true,
      showConfirm: false,
      ruleForm: {
        path: ''
      },
      uploadData: { data: {}, type: undefined },
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
    };
  },
  created() {
    const { algorithmName, algorithmDescript, modelName } = this.reuploadData
    this.ruleForm = { algorithmName, algorithmDescript, modelName }
    this.uploadData.AlgorithmId = this.reuploadData.algorithmId
    this.uploadData.Version = this.reuploadData.algorithmVersion
    this.uploadData.type = "myAlgorithmCreation"
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