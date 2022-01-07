<template>
  <div>
    <el-dialog title="创建预置数据集" width="650px" :visible.sync="CreateFormVisible" :before-close="handleDialogClose"
      :close-on-click-modal="false" :show-close="close">
      <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px">
        <el-form-item label="数据集名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="ruleForm.name" :disabled="disabled" placeholder="请输入数据集名称" />
        </el-form-item>
        <el-form-item label="数据类型" :label-width="formLabelWidth" prop="typeId">
          <el-select v-model="ruleForm.typeId" :disabled="disabled" placeholder="请选择数据类型">
            <el-option v-for="item in typeOptions" :key="item.id" :label="item.lableDesc" :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="标注类型" :label-width="formLabelWidth">
          <el-select v-model="ruleForm.applyIds" :disabled="disabled" placeholder="请选择标注类型" multiple>
            <el-option v-for="item in useOptions" :key="item.id" :label="item.lableDesc" :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="数据集描述" :label-width="formLabelWidth" prop="desc">
          <el-input v-model="ruleForm.desc" :disabled="disabled" :autosize="{ minRows: 2, maxRows: 4}"
            placeholder="请输入数据集描述" maxlength="300" show-word-limit />
        </el-form-item>
        <el-form-item :label-width="formLabelWidth">
          <el-button v-show="!showUpload" type="text" @click="nextStep('ruleForm')">下一步</el-button>
        </el-form-item>
        <el-form-item v-if="showUpload" label="数据集上传" :label-width="formLabelWidth" prop="path">
          <upload v-model="ruleForm.path" :upload-data="uploadData" @confirm="confirm" @cancel="cancel"
            @upload="isCloseX" />
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
  import upload from '@/components/upload/index.vue'
  import { createPreDataset, datasetType,datasetUse } from "@/api/dataManager"
  import { getErrorMsg } from '@/error/index'
  export default {
    name: "PreDatasetCreation",
    components: {
      upload
    },
    props: {},
    props: {
      row: {
        type: Object,
        default: () => { }
      }
    },
    created() {
      this.datasetType(),
      this.datasetUse()
    },
    data() {
      return {
        showUpload: false,
        disabled: false,
        uploadData: { data: {}, type: undefined },
        ruleForm: {
        },
        rules: {
          name: [
            { required: true, message: '请输入数据集名称', trigger: 'blur' },
            { min: 4, max: 30, message: '长度在 4 到 30 个字符', trigger: 'blur' }
          ],
          typeId: [
            { required: true, message: '请选择数据类型', trigger: 'change' }
          ],
          path: [
            { required: true, message: '请上传数据集', trigger: 'change' }
          ]
        },
        CreateFormVisible: true,
        formLabelWidth: '120px',
        close: true,
        typeOptions:[],
        useOptions:[]
      }
    },
    methods: {
      getErrorMsg(code) {
        return getErrorMsg(code)
      },
      nextStep(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            createPreDataset(this.ruleForm).then(response => {
              if (response.success) {
                this.showUpload = true
                this.disabled = true
                this.uploadData.id = response.data.id
                this.uploadData.type = "preDatasetCreation"
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
        this.$emit('close', false)
      },
      cancel() {
        this.$emit('cancel', false)
      },
      confirm(val) {
        this.$emit('confirm', val)
      },
      isCloseX(val) {
        this.close = val
      },
      // 获取数据集类型
      datasetType() {
        datasetType({ pageIndex: 1, pageSize: 20 }).then(response => {
          if (response.success) {
            this.typeOptions = response.data.lables
          } else {
            // this.showUpload = false
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      },
       // 获取数据集用途
       datasetUse() {
        datasetUse({ pageIndex: 1, pageSize: 20 }).then(response => {
          if (response.success) {
            this.useOptions = response.data.lables
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
  }
</script>