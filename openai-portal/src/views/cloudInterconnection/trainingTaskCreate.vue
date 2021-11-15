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
        <!-- 框架 -->
        <div>
          <el-form-item label="学习框架" prop="framework" :class="{inline:frameworkVersionVisible}">
            <el-select
              v-model="ruleForm.framework"
              placeholder="请选择学习框架"
              @change="changeFramework"
            >
              <el-option
                v-for="item in frameworkOption"
                :key="item.key"
                :label="item.value"
                :value='item.key'
              >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item
            label="框架版本"
            prop="frameworkVersion"
            v-if="frameworkVersionVisible"
            style="display: inline-block;"
          >
            <el-select
              v-model="ruleForm.frameworkVersion"
              placeholder="请选择框架版本"
            >
              <el-option
                v-for="item in frameworkVersionOption"
                :key="item.key"
                :label="item.value"
                :value='item.key'
              >
              </el-option>
            </el-select>
          </el-form-item>
        </div>
        <!-- 解释器 -->
        <div>
          <el-form-item label="解释器" prop="interpreter" :class="{inline:interpreterVersionVisible}">
            <el-select
              v-model="ruleForm.interpreter"
              placeholder="请选择解释器"
              @change="changeInterpreter"
            >
              <el-option
                v-for="item in interpreterOption"
                :key="item.key"
                :label="item.value"
                :value='item.key'
              >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item
            label="解释器版本"
            prop="interpreterVersion"
            v-if="interpreterVersionVisible"
            style="display: inline-block;"
          >
            <el-select
              v-model="ruleForm.interpreterVersion"
              placeholder="请选择解释器版本"
            >
              <el-option
                v-for="item in interpreterVersionOption"
                :key="item.key"
                :label="item.value"
                :value='item.key'
              >
              </el-option>
            </el-select>
          </el-form-item>
        </div>
        <!-- 数据集 -->
        <div>
          <el-form-item label="数据集" prop="dataSet" :class="{inline:dataSetVersionVisible}">
            <el-select
              value-key="dataSetCode"
              v-model="ruleForm.dataSet"
              placeholder="请选择数据集"
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
          <el-form-item
            label="数据集版本"
            prop="dataSetVersion"
            v-if="dataSetVersionVisible"
            style="display: inline-block;"
          >
            <el-select
              v-model="ruleForm.dataSetVersion"
              placeholder="请选择数据集版本"
              v-loadmore="loadDataSetVersion"
              @change="changeMountPath"
            >
              <el-option
                v-for="item in dataSetVersionOption"
                :key="item.version"
                :label="item.version"
                :value='item.version'
              >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item
            v-if="mountPathVisible"
            label="挂载目录"
            prop="mountPath"
            style="display: inline-block;"
          >
            <el-input v-model="ruleForm.mountPath" placeholder="请输入挂载目录"></el-input>
          </el-form-item>
        </div>
        <div>
          <el-form-item label="运行命令" prop="execCommand">
            <el-input type="textarea" v-model="ruleForm.execCommand"></el-input>
          </el-form-item>
        </div>
        <div>
          <el-form-item
            label="模型输出位置"
            prop="outputPath"
            style="display: inline-block;"
          >
            <el-input v-model="ruleForm.outputPath" placeholder="请输入模型输出目录"></el-input>
          </el-form-item>
        </div>

        <!-- <div>
          <el-row >
            <el-col :span="6">
              <el-form-item label="内存" prop="memorySize">
                <el-input v-model="ruleForm.memorySize" placeholder="请填写内存大小">
                </el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="内存"  prop="memoryUnits">
                <el-select v-model="ruleForm.memoryUnits">
                  <el-option label="Gi" value="Gi"></el-option>
                  <el-option label="Mi" value="Mi"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6"></el-col>
            <el-col :span="6"></el-col>
          </el-row>
        </div> -->


        <div>
          <el-form-item label="内存大小" prop="memorySize" style="display: inline-block;">
            <el-input v-model.number="ruleForm.memorySize" placeholder="请填写内存大小">
            </el-input>
          </el-form-item>
          <el-form-item label="内存单位" prop="memoryUnits" style="display: inline-block;">
            <el-select v-model="ruleForm.memoryUnits">
              <el-option label="Gi" value="Gi"></el-option>
              <el-option label="Mi" value="Mi"></el-option>
            </el-select>
          </el-form-item>
        </div>
        <div>
          <el-form-item label="GPU个数" prop="gpuSize" style="display: inline-block;">
            <el-input v-model.number="ruleForm.gpuSize" placeholder="请填写GPU个数">
            </el-input>
          </el-form-item>
          <el-form-item label="GPU类型" prop="gpuUnits" style="display: inline-block;">
            <el-select v-model="ruleForm.gpuUnits">
              <el-option label="nvidia.com/gpu" value="nvidia.com/gpu"></el-option>
              <el-option label="npu.huawei.com/NPU" value="npu.huawei.com/NPU"></el-option>
              <el-option label="cambricon.com/mlu" value="cambricon.com/mlu"></el-option>
              <el-option label="不限" value="default"></el-option>
            </el-select>
          </el-form-item>
        </div>
        <div>
          <el-form-item label="CPU个数" prop="cpuSize" style="display: inline-block;">
            <el-input v-model.number="ruleForm.cpuSize" placeholder="请填写CPU个数">
            </el-input>
          </el-form-item>
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
      formLabelWidth: "130px",
      ruleForm: {
        taskName: '',
        remark: '',
        framework: '',
        frameworkVersion: '',
        interpreter: '',
        interpreterVersion: '',
        dataSet: "",
        dataSetVersion: '',
        mountPath: '/dataset ',
        execCommand: '',
        outputPath: "/model",
        memorySize: 1,
        memoryUnits: "Gi",
        cpuSize: 1,
        gpuSize: 1,
        gpuUnits: "nvidia.com/gpu",
        parameters: [{
          key: "",
          value: ""
        }],
      },
      rules: {
        taskName: [
          { required: true, message: '请输入任务名称', trigger: 'blur' },
        ],
        framework: [
          { required: true, message: '请选择引擎框架', trigger: 'blur' },
        ],
        frameworkVersion: [
          { required: true, message: '请选择引擎版本', trigger: 'blur' },
        ],
        interpreter: [
          { required: true, message: '请选择引擎解释器', trigger: 'blur' },
        ],
        interpreterVersion: [
          { required: true, message: '请选择引擎解释器版本', trigger: 'blur' },
        ],
        dataSet: [
          { required: true, message: '请选择数据集', trigger: 'blur' },
        ],
        dataSetVersion: [
          { required: true, message: '请选择数据集版本', trigger: 'blur' },
        ],
        mountPath: [
          { required: true, message: '请选择挂载目录', trigger: 'blur' },
        ],
        execCommand: [
          { required: true, message: '请输入运行命令', trigger: 'blur' },
        ],
        outputPath: [
          { required: true, message: '请填写模型输出路径', trigger: 'blur' },
        ],
        memorySize:[
          { required: true, message: '请填写内存大小', trigger: 'blur' }
        ],
        memoryUnits: [
          { required: true, message: '请选择内存类型', trigger: 'blur' },
        ],
        cpuSize: [
          { required: true, message: '请填写CPU个数', trigger: 'blur' },
        ],
        gpuUnits: [
          { required: true, message: '请选择GPU类型', trigger: 'blur' },
        ],
        gpuSize: [
          { required: true, message: '请填写GPU个数', trigger: 'blur' },
        ],               
      },
      // 引擎框架
      frameworkOption: [],
      frameworkVersionVisible: false,
      frameworkVersionOption: [],
      // 解释器
      interpreterOption: [],
      interpreterVersionVisible: false,
      interpreterVersionOption: [],
      // 数据集
      dataSetNameOption: [],
      dataSetNameCount: 1,
      dataSetNameTotal: undefined,
      dataSetVersionVisible: false,
      dataSetVersionCount: 1,
      dataSetVersionOption: [],
      dataSetVersionTotal: undefined,
      mountPathVisible: false
    }
  },
  created(){
    this.getCloudFrameworkList()
    this.getCloudDatasetList()
    this.getCloudInterpreterList()
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
    // 训练框架
    getCloudFrameworkList(){
      getCloudFrameworkList().then(response => {
        if(response.data !== null && response.data.frameworks !== null) {
          this.frameworkOption = response.data.frameworks
        } else {
          this.$message({
            message: '获取训练学习框架失败',
            type: 'warning'
          })
        }
      })
    },
    changeFramework() {
      this.frameworkVersionVisible = true
      this.frameworkVersionOption = [],
      this.ruleForm.frameworkVersion = '',
      this.getCloudFrameworkVersionList()
    },
    getCloudFrameworkVersionList() {
      const key = this.ruleForm.framework
      getCloudFrameworkVersionList(key).then(response => {
        if (response.data !== null && response.data.versions !== null) { 
          this.frameworkVersionOption = response.data.versions
        } else {
          this.$message({
            message: '获取训练学习框架版本失败',
            type: 'warning'
          })
        }
      })
    },
    // 解释器
    getCloudInterpreterList(){
      getCloudInterpreterList().then(response => {
        if(response.data !== null && response.data.interpreters !== null) {
          this.interpreterOption = response.data.interpreters
        } else {
          this.$message({
            message: '获取训练训练解释器失败',
            type: 'warning'
          })
        }
      })
    },
    changeInterpreter() {
      this.interpreterVersionVisible = true
      this.frameworkVersionOption = [],
      this.ruleForm.interpreterVersion = '',
      this.getCloudInterpreterVersionList()
    },
    getCloudInterpreterVersionList() {
      const key = this.ruleForm.interpreter
      getCloudInterpreterVersionList(key).then(response => {
        if (response.data !== null && response.data.versions !== null) { 
          this.interpreterVersionOption = response.data.versions
        } else {
          this.$message({
            message: '获取训练学习框架版本失败',
            type: 'warning'
          })
        }
      })
    },
    // 数据集
    getCloudDatasetList(){
      const param = {
        pageIndex: this.dataSetNameCount, 
        pageSize: 10
      }
      getCloudDatasetList(param).then(response => {
        if(response.data !== null && response.data.dataSets !== null) {
          this.dataSetNameOption = this.dataSetNameOption.concat(response.data.dataSets)
          this.dataSetNameTotal = response.data.totalSize
        } else {
          this.$message({
            message: '获取数据集失败',
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
      this.dataSetVersionVisible = true
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
        if (response.data !== null && response.data.versions !== null) { 
          this.dataSetVersionOption = this.dataSetVersionOption.concat(response.data.versions) 
          this.dataSetVersionTotal = response.data.totalSize
        } else {
          this.$message({
            message: '获取数据集版本失败',
            type: 'warning'
          })
        }
      })
    },
    loadDataSetVersion() {
      this.dataSetVersionCount = this.dataSetVersionCount + 1
      if (this.dataSetVersionOption.length < this.dataSetVersionTotal) {
        this.getDataSetVersionList()
      }
    },
    changeMountPath(){
      this.mountPathVisible = true
    },
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    handleDialogClose() {
      this.$emit("close", false);
    },
    submit(formName) {
      console.log("this.ruleForm",this.ruleForm)
          console.log("this.ruleForm.cpuSize",typeof(this.ruleForm.cpuSize))
    console.log("this.ruleForm.memorySize",typeof(this.ruleForm.memorySize))
    console.log("this.ruleForm.gpuSize",typeof(this.ruleForm.gpuSize))
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let resourceList = []
          if(this.ruleForm.memorySize != 0){
            let memoryParam = {}
            memoryParam.name = "memory"
            memoryParam.type = "memory"
            memoryParam.unit = this.ruleForm.memoryUnits
            memoryParam.size = this.ruleForm.memorySize
            resourceList.push(memoryParam)
          }
          if(this.ruleForm.gpuSize != 0){
            let gpuParam = {}
            gpuParam.name = this.ruleForm.gpuUnits
            gpuParam.type = "gpu"
            gpuParam.unit = "个"
            gpuParam.size = this.ruleForm.gpuSize
            resourceList.push(gpuParam)
          }
          if(this.ruleForm.cpuSize != 0){
            let cpuParam = {}
            cpuParam.name = "cpu"
            cpuParam.type = "cpu"
            cpuParam.unit = "个"
            cpuParam.size = this.ruleForm.cpuSize
            resourceList.push(cpuParam)
          }
          const params = {
            taskName: this.ruleForm.taskName,
            remark: this.ruleForm.remark,
            framework: this.ruleForm.framework + this.ruleForm.frameworkVersion,
            interpreter: this.ruleForm.interpreter + this.ruleForm.interpreterVersion,
            dataSetVersionVoList: [{
              dataSetCode:this.ruleForm.dataSet.dataSetCode,
              dataSetName:this.ruleForm.dataSet.name,
              path:this.ruleForm.mountPath,
              version:this.ruleForm.dataSetVersion,
            }],
            execCommand: this.ruleForm.execCommand,
            outputPath: this.ruleForm.outputPath,
            resourceParams: resourceList
          }
          console.log("params",params)
          createCloudTrainJob(params).then(response => {
            if(response.success) {            
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
          return false;
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