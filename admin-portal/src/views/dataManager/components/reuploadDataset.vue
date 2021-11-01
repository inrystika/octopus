<template>
  <div>
    <el-dialog
      title="创建新版本"
      width="650px"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
      :show-close="close"
    >
      <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px">
        <el-form-item label="数据集名称" :label-width="formLabelWidth" prop="name">
            <el-input v-model="ruleForm.name" :disabled="true" />
        </el-form-item>
        <el-form-item label="数据类型" :label-width="formLabelWidth" prop="type">
            <el-input v-model="ruleForm.type" :disabled="true" />
        </el-form-item>
        <el-form-item label="版本描述" :label-width="formLabelWidth" prop="desc">
          <el-input
            v-model="ruleForm.desc"
            :disabled="true"
            :autosize="{ minRows: 2, maxRows: 4}"
            placeholder="请输入数据集描述"
            maxlength="300"
            show-word-limit
          />
        </el-form-item>
        <el-form-item label="数据集上传" :label-width="formLabelWidth" prop="path">
          <upload
            v-model="ruleForm.path"
            :upload-data="uploadData"
            @confirm="confirm"
            @cancel="cancel"
            @upload="isCloseX"
          />
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import upload from '@/components/upload/index.vue'
export default {
  name: "ReuploadDataset",
  components: {
    upload
  },
  props: {
    versionData: {
      type: Object,
      default: () => { }
    },
    data: {
      type: Object,
      default: () => { }
    }
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
      formLabelWidth: "120px",
      close:true
    }
  },
  created() {
    const { desc } = this.versionData
    const { name, type } = this.data
    this.ruleForm = { name, type, desc }
    this.uploadData.id = this.versionData.datasetId
    this.uploadData.type = "preDatasetCreation"
    this.uploadData.version = this.versionData.version
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
    },
    isCloseX(val) {
                this.close = val
            }
  }
}
</script>