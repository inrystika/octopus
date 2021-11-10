<template>
  <div>
    <el-dialog
      title="创建训练任务"
      width="55%"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
      <el-form
        :model="ruleForm"
        :rules="rules"
        ref="ruleForm"
        :label-width="formLabelWidth"
        class="demo-ruleForm"
      >
        <el-form-item label="任务名称" :label-width="formLabelWidth" prop="taskName">
          <el-input v-model="ruleForm.taskName" placeholder="请输入任务名称"></el-input>
        </el-form-item>
        <el-form-item label="描述" :label-width="formLabelWidth" prop="remark">
          <el-input
            :autosize="{ minRows: 2, maxRows: 4}"
            placeholder="请输入notebook描述"
            maxlength="300"
            show-word-limit
            v-model="ruleForm.remark"
          />
        </el-form-item>
        <!-- <el-form-item label="训练引擎框架" prop="frameworkLanguage" style="display:inline-block;">
          <el-select v-model="ruleForm.frameworkLanguage" placeholder="请选择训练引擎" v-loadmore='loadAlgorithmName'
            @change="changeAlgorithmName">
            <el-option v-for="item in algorithmNameOption" :key="item.algorithmId+item.algorithmName"
              :label="item.algorithmName" :value='item.algorithmId'>
            </el-option>
          </el-select>
        </el-form-item> -->
        <!-- 数据集 -->
        <div>
          <el-form-item label="训练数据集" prop="dataSet">
            <el-select
              value-key="dataSetCode"
              v-model="ruleForm.dataSet"
              placeholder="请选择训练数据集"
              v-loadmore='loadDataSetName'
              @change="changeDataSetName"
            >
              <el-option
                v-for="item in dataSetNameOption"
                :key="item.dataSetCode"
                :label="item.name"
                :value='item'
              >
              </el-option>
            </el-select>
          </el-form-item>
          <!-- <el-form-item label="数据集版本" prop="dataSetVersion" v-if="dataSetVersion"
            style="display: inline-block;">
            <el-select v-model="ruleForm.dataSetVersion" placeholder="请选择数据集版本"
              v-loadmore='loadDataSetVersion'>
              <el-option v-for="item in dataSetVersionOption" :key="item.datasetId+item.version"
                :label="item.version" :value='item.version'>
              </el-option>
            </el-select>
          </el-form-item> -->
        </div>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="submit('ruleForm')">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import { getErrorMsg } from '@/error/index'
import { 
  createCloudTrainJob,
  getCloudDatasetList,
  getCloudDatasetVersionList,
  getCloudFrameworkList,
  getCloudFrameworkVersionList,
  getCloudInterpreterList,
  getCloudInterpreterVersionList 
} from "@/api/cloudInterconnection"
export default {
  name: "trainingTaskCreate",
  data() {
    return {
      CreateFormVisible: true,
      formLabelWidth: "120px",
      ruleForm: {
        taskName: '',
        remark: '',
        frameworkLanguage: '',
        frameworkVersion: '',
        interpreterLanguage: '',
        interpreterVersion: '',
        dataSet: "",
        dataSetVersion: '',
        dataSetPath: '',
        execCommand: '',
        outputPath: "",
        cpuSize: "",
        cpuMemory: "",
        gpuSize: "",
        gpuMemory: "",
        parameters: [{
          key: "",
          value: ""
        }],
      },
      rules: {
        taskName: [
          { required: true, message: '请输入任务名称', trigger: 'blur' },
        ],
        frameworkLanguage: [
          { required: true, message: '请选择引擎框架', trigger: 'blur' },
        ],
        frameworkVersion: [
          { required: true, message: '请选择引擎版本', trigger: 'blur' },
        ],
        interpreterLanguage: [
          { required: true, message: '请选择引擎解释器', trigger: 'blur' },
        ],
        interpreterVersion: [
          { required: true, message: '请选择引擎解释器版本', trigger: 'blur' },
        ],
        dataSet: [
          { required: true, message: '请选择训练数据集', trigger: 'blur' },
        ],
        dataSetVersion: [
          { required: true, message: '请选择训练数据集版本', trigger: 'blur' },
        ],
        dataSetPath: [
          { required: true, message: '请选择挂载目录', trigger: 'blur' },
        ],
        execCommand: [
          { required: true, message: '请输入运行命令', trigger: 'blur' },
        ],
        outputPath: [
          { required: true, message: '请填写模型输出路径', trigger: 'blur' },
        ],
        cpuMemory: [
          { required: true, message: '请选择CPU大小', trigger: 'blur' },
        ],
        cpuSize: [
          { required: true, message: '请选择CPU个数', trigger: 'blur' },
        ],
        gpuMemory: [
          { required: true, message: '请选择GPU大小', trigger: 'blur' },
        ],
        gpuSize: [
          { required: true, message: '请选择GPU个数', trigger: 'blur' },
        ],
      },
      // 数据集
      dataSetNameOption: [],
      dataSetNameCount: 1,
      dataSetNameTotal: undefined,
      dataSetVersion: false,
      dataSetVersionCount: 1,
      dataSetVersionOption: [],
      dataSetVersionTotal: undefined,
    }
  },
  created(){
    this.getCloudDatasetList()
    // this.getCloudDatasetVersionList()
    // this.getCloudFrameworkList()
    // this.getCloudFrameworkVersionList()
    // this.getCloudInterpreterList()
    // this.getCloudInterpreterVersionList()
  },
  directives: {
    loadmore: {
      inserted: function (el, binding) {
        const SELECTWRAP_DOM = el.querySelector('.el-select-dropdown .el-select-dropdown__wrap');
        SELECTWRAP_DOM.addEventListener('scroll', function () {
          const CONDITION = this.scrollHeight - this.scrollTop <= this.clientHeight;
          if (CONDITION) {
            binding.value();
          }
        })
      }
    }
  },
  methods: {
    // 数据集
    getCloudDatasetList(){
      const param = {
        pageIndex: this.dataSetNameCount, 
        pageSize: 10
      }
      getCloudDatasetList(param).then(response => {
        if(response.data !== null) {
          this.dataSetNameOption = this.dataSetNameOption.concat(response.data.dataSets)
        } else {
          this.$message({
            message: '获取训练数据集失败',
            type: 'warning'
          })
        }
      })
    },
    loadDataSetName(){
      this.dataSetNameCount = this.dataSetNameCount + 1
      if (this.dataSetNameOption.length < this.dataSetNameTotal) {
        this.getCloudDatasetList()
      }
    },
    changeDataSetName() {
      this.dataSetVersion = true
      this.dataSetVersionCount = 1
      this.dataSetVersionOption = [],
      this.ruleForm.dataSetVersion = '',
      this.getDataSetVersionList()
    },
    getDataSetVersionList() {
      let data = {}
      data.dataSetCode = this.ruleForm.dataSet.dataSetCode
      data.pageIndex = this.dataSetVersionCount
      data.pageSize = 10
      getCloudDatasetVersionList(data).then(response => {
        if (response.data.versions !== null) { 
          this.dataSetVersionOption = this.dataSetVersionOption.concat(response.data.versions) 
          this.dataSetVersionTotal = response.data.totalSize
        }
      })
    },
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    handleDialogClose() {
      this.$emit("close", false);
    },
    submit(formName) {
      console.log("this.ruleForm",this.ruleForm)
      this.$refs[formName].validate((valid) => {
        if (valid) {

        }
      })
    },
    cancel() {
      this.$confirm('此操作将被取消，是否继续?','提示',{
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$emit('cancel', false)
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已中断取消操作'
        });
      })
    },
  }
}
</script>
<style lang="scss" scoped>
 .line {
    text-align: center;
  }

  .inline {
    display: inline-block !important;
  }

  .block {
    display: block !important;
  }
</style>