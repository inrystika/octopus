<template>
  <div>
    <el-dialog title="创建新版本" width="650px" :visible.sync="CreateFormVisible" :before-close="handleDialogClose"
      :close-on-click-modal="false" :show-close="close">
      <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px">
        <el-form-item label="数据集名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="ruleForm.name" :disabled="true" />
        </el-form-item>
        <el-form-item label="数据类型" :label-width="formLabelWidth" prop="type">
          <el-input v-model="ruleForm.type" :disabled="true" />
        </el-form-item>
        <el-form-item label="版本描述" :label-width="formLabelWidth" prop="desc">
          <el-input v-model="ruleForm.desc" :disabled="true" />
        </el-form-item>
        <!-- <el-form-item :label-width="formLabelWidth">
          <el-button type="text" @click="nextStep('ruleForm')" v-show="!showUpload">下一步</el-button>
        </el-form-item> -->
        <el-form-item label="数据集上传" :label-width="formLabelWidth" prop="path">
          <upload v-model="ruleForm.path" :upload-data="uploadData" @confirm="confirm" @cancel="cancel"
            @upload="isCloseX" />
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
        // showUpload: false,
        uploadData: { data: {}, type: undefined },
        ruleForm: {},
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
        close: true
      }
    },
    created() {
      const { desc } = this.versionData
      const { name, type } = this.data
      this.ruleForm = { name, type, desc }
      this.uploadData.id = this.versionData.datasetId
      this.uploadData.type = "myDatasetCreation"
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
      close(val) {
        this.$emit("close", false);
      },
      isCloseX(val) {
        this.close = val
      }
    }
  }
</script>