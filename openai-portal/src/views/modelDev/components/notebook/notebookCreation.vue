<template>
  <div>
    <el-dialog
      title="创建NoteBook"
      width="55%"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" :label-width="formLabelWidth"
      class="demo-ruleForm">
        <el-form-item label="名称" :label-width="formLabelWidth" prop="name">
          <el-input v-model="ruleForm.name" placeholder="请输入notebook名称"></el-input>
        </el-form-item>
        <el-form-item label="描述" :label-width="formLabelWidth" prop="desc">
          <el-input
            :autosize="{ minRows: 2, maxRows: 4}"
            placeholder="请输入notebook描述"
            maxlength="300"
            show-word-limit
            v-model="ruleForm.desc"
          ></el-input>
        </el-form-item>
        <!-- 算法三级框 -->
        <div>  
          <!-- <el-form-item label="算法类型" prop="algorithmSource" style="display:inline-block;">
            <el-select v-model="ruleForm.algorithmSource" placeholder="请选择">
              <el-option label="我的算法" value="my"></el-option>
            </el-select>
          </el-form-item> -->
          <el-form-item label="算法名称" prop="algorithmId" style="display:inline-block;">
            <el-select v-model="ruleForm.algorithmId" placeholder="请选择算法名称" v-loadmore='loadAlgorithmName'
              @change="changeAlgorithmName">
              <el-option v-for="item in algorithmNameOption" :key="item.algorithmId+item.algorithmName"
                :label="item.algorithmName" :value='item.algorithmId'>
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="算法版本" prop="algorithmVersion" v-if="algorithmVersion" style="display:inline-block;">
            <el-select v-model="ruleForm.algorithmVersion" placeholder="请选择算法版本"
              v-loadmore='loadAlgorithmVersion'>
              <el-option v-for="item in algorithmVersionOption"
                :key="item.algorithmDetail.algorithmId+item.algorithmDetail.algorithmVersion"
                :label="item.algorithmDetail.algorithmVersion"
                :value='item.algorithmDetail.algorithmVersion'>
              </el-option>
            </el-select>
          </el-form-item>
        </div>
        <!-- 镜像三级框 -->
        <div>
          <el-form-item label="镜像类型" prop="imageSource" :class="{inline:imageName}">
            <el-select v-model="ruleForm.imageSource" @change="changeimageSource" placeholder="请选择">
              <el-option label="我的镜像" value="my"></el-option>
              <el-option label="预置镜像" value="pre"></el-option>
              <el-option label="公共镜像" value="common"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="镜像名称" prop="imageItem" v-if="imageName" style="display: inline-block;">
            <el-select value-key="id" v-model="ruleForm.imageItem" placeholder="请选择镜像" v-loadmore='loadImageName'>
              <el-option v-for="item in imageNameOption" :key="item.id" :label="item.imageName+':'+item.imageVersion"
                  :value="item">
              </el-option>
            </el-select>
          </el-form-item>
        </div>
        <!-- 数据集三级框 -->
        <div>
          <el-form-item label="数据集类型" prop="dataSetSource" :class="{inline:dataSetName}">
            <el-select v-model="ruleForm.dataSetSource" @change="changedataSetSource" clearable placeholder="请选择">
              <el-option label="我的数据集" value="my"></el-option>
              <el-option label="预置数据集" value="pre"></el-option>
              <el-option label="公共数据集" value="common"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="数据集名称" prop="dataSetId" v-if="dataSetName" style="display: inline-block;">
            <el-select v-model="ruleForm.dataSetId" placeholder="请选择数据集名称" v-loadmore='loadDataSetName'
              @change="changeDataSetName">
              <el-option v-for="item in dataSetNameOption" :key="item.id+item.name" :label="item.name"
                :value='item.id'>
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="数据集版本" prop="dataSetVersion" v-if="dataSetVersion"
            style="display: inline-block;">
            <el-select v-model="ruleForm.dataSetVersion" placeholder="请选择数据集版本"
              v-loadmore='loadDataSetVersion'>
              <el-option v-for="item in dataSetVersionOption" :key="item.datasetId+item.version"
                :label="item.version" :value='item.version'>
              </el-option>
            </el-select>
          </el-form-item>
        </div>

        <el-form-item label="资源规格" :label-width="formLabelWidth" prop="specification">
          <el-select v-model="ruleForm.specification" placeholder="请资源规格" style="width:35%">
            <el-option 
              v-for="(item, index) in resourceList" 
              :key="index" 
              :label="item.label" 
              :value="item.value"
            >
            </el-option>                
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="submit('ruleForm')">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { createNotebook, getMyAlgorithmList, getAlgorithmVersionList, getPublicAlgorithmList, getPresetAlgorithmList } from "@/api/modelDev"
import { getMyDatasetList, getPublicDatasetList, getPresetDatasetList, getVersionList } from "@/api/datasetManager"
import { getMyImage, getPublicImage, getPreImage } from '@/api/imageManager'
import { getResourceList } from "@/api/trainingManager"
import { getErrorMsg } from '@/error/index'
export default {
  name: "notebookCreation",
  data() {
    return {
      ruleForm: {
        name: '',
        desc: '',
        algorithmSource: 'my',
        algorithmId: '',
        algorithmVersion: '',
        imageSource: '',
        imageItem: "",
        dataSetSource: '',
        dataSetId: '',
        dataSetVersion: '',
      },
      rules: {
        name:[
          {
            required: true,
            message: "请输入NoteBook名称",
            trigger: "blur"
          },
        ],
        algorithmSource: [
          { required: true, message: '请选择算法类型', trigger: 'change' }
        ],
        algorithmId: [
          { required: true, message: '请选择算法名称', trigger: 'change' }
        ],
        algorithmVersion: [{ required: true, message: '请选择算法版本', trigger: 'change' }],
        imageSource: [
          { required: true, message: '请选择镜像类型', trigger: 'change' }
        ],
        imageItem: [
          { required: true, message: '请选择镜像名称', trigger: 'change' }
        ],
        // dataSetSource: [
        //   { required: true, message: '请选择数据集类型', trigger: 'change' }
        // ],
        // dataSetId: [
        //   { required: true, message: '请选择数据集名称', trigger: 'change' }
        // ],
        // dataSetVersion: [
        //   { required: true, message: '请选择数据集版本', trigger: 'change' }
        // ],
        specification:[
          {
            required: true,
            message: "请选择资源规格",
            trigger: "blur"
          }
        ]
      },
      CreateFormVisible: true,
      formLabelWidth: "120px",
      pageIndex: 1,
      pageSize: 20,
      // 算法三级框
      algorithmName: false,
      algorithmVersion: false,
      algorithmNameOption: [],
      algorithmVersionOption: [],
      algorithmNameCount: 1,
      algorithmVersionCount: 1,
      algorithmNameTotal: undefined,
      algorithmVersionTotal: undefined,
      // 镜像二级框
      imageName: false,
      imageNameOption: [],
      imageCount: 1,
      imageNameTotal: undefined,
      // 数据集三级框
      dataSetName: false,
      dataSetVersion: false,
      dataSetNameOption: [],
      dataSetVersionOption: [],
      dataSetNameCount: 1,
      dataSetVersionCount: 1,
      dataSetNameTotal: undefined,
      dataSetVersionTotal: undefined,
      resourceOptions: [],
      data: {},
      resourceList: [],
    }
  },
  created(){
    this.getResource()
    this.getAlgorithmNameList()
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
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    handleDialogClose() {
      this.$emit("close", false);
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
    getResource(){
      getResourceList().then(response => {
        if(response.success) {
          response.data.mapResourceSpecIdList.debug.resourceSpecs.forEach(
            item => {
              this.resourceList.push({label: item.name + ' ' + item.price + '机时/h', value: item.id})
            }
          )
          // this.resourceList = response.data.mapResourceSpecIdList.debug.resourceSpecs
        } else {
          this.$message({
            message: this.getErrorMsg(response.error.subcode),
            type: 'warning'
          });
        }
      })
    },
    submit(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.showUpload = true
          const param = {
            name: this.ruleForm.name,
            desc: this.ruleForm.desc,
            imageId: this.ruleForm.imageItem.id,
            imageVersion: this.ruleForm.imageItem.imageVersion,
            resourceSpecId: this.ruleForm.specification,
            algorithmId: this.ruleForm.algorithmId || '',
            algorithmVersion: this.ruleForm.algorithmVersion || '',
            datasetId: this.ruleForm.dataSetId || '',
            datasetVersion: this.ruleForm.dataSetVersion || '',
          }
          createNotebook(param).then(response => {
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
      });
    },
    // 算法三级对话框实现
    changealgorithmSource() {
      this.algorithmName = true
      this.algorithmNameCount = 1
      this.algorithmNameOption = [],
      this.ruleForm.algorithmId = '',
      this.ruleForm.algorithmVersion = ''
      this.algorithmChange = true
      this.getAlgorithmNameList()
    },
    changeAlgorithmName() {
      this.algorithmVersion = true
      this.algorithmVersionCount = 1
      this.algorithmVersionOption = [],
      this.ruleForm.algorithmVersion = ''
      this.getAlgorithmVersionList()
    },
    getAlgorithmNameList() {
      getMyAlgorithmList({ pageIndex: this.algorithmNameCount, pageSize: 10 }).then(response => {
        if (response.data.algorithms.length !== 0) { 
          this.algorithmNameOption = this.algorithmNameOption.concat(response.data.algorithms) 
          this.algorithmNameTotal = response.data.totalSize
        }
      })
    },
    getAlgorithmVersionList() {
      getAlgorithmVersionList({ pageIndex: this.algorithmVersionCount, pageSize: 10, algorithmId: this.ruleForm.algorithmId, fileStatus: 3 }).then(response => {
        if (response.success) {
          this.algorithmVersionOption = this.algorithmVersionOption.concat(response.data.algorithms)
          this.algorithmVersionTotal = response.data.totalSize
        }
      })
    },
    loadAlgorithmName() {
      this.algorithmNameCount = this.algorithmNameCount + 1
      if (this.algorithmNameOption.length < this.algorithmNameTotal) {
        this.getAlgorithmNameList()
      }
    },
    loadAlgorithmVersion() {
      this.algorithmVersionCount = this.algorithmVersionCount + 1
      if (this.algorithmVersionOption.length < this.algorithmVersionTotal) {
        this.getAlgorithmNameList()
      }
    },
    // 镜像二级对话框实现
    changeimageSource() {
      this.imageName = true
      this.imageCount = 1
      this.imageNameOption = [],
      this.ruleForm.imageId = ''
      this.imageChange = true
      this.getImageNameList()
    },
    getImageNameList() {
      if (this.ruleForm.imageSource === 'my') {
        getMyImage({ pageIndex: this.imageCount, pageSize: 10, imageStatus: 3, imageType: 1 }).then(response => {
          if (response.data.images.length !== 0) {
            let data = response.data.images;
            let tableData = [];
            this.imageNameTotal = response.data.totalSize
            data.forEach(item => {
              tableData.push({ ...item.image, isShared: item.isShared })
            })
            this.imageNameOption = this.imageNameOption.concat(tableData)
          }
        })
      }
      if (this.ruleForm.imageSource === 'pre') {
        getPreImage({ pageIndex: this.imageCount, pageSize: 10, imageStatus: 3, imageType: 1 }).then(response => {
          if (response.data.images.length !== 0) { 
            this.imageNameOption = this.imageNameOption.concat(response.data.images) 
            this.imageNameTotal = response.data.totalSize
          }
        })
      }
      if (this.ruleForm.imageSource === 'common') {
        getPublicImage({ pageIndex: this.imageCount, pageSize: 10, imageStatus: 3, imageType: 1 }).then(response => {
          if (response.data.images.length !== 0) { 
            this.imageNameOption = this.imageNameOption.concat(response.data.images) 
            this.imageNameTotal = response.data.totalSize
         }
        })
      }

    },
    loadImageName() {
      this.imageCount = this.imageCount + 1
      this.getImageNameList()
    },
    // 数据集三级对话框
    changedataSetSource() {
      this.dataSetName = true
      this.dataSetNameCount = 1
      this.dataSetNameOption = [],
      this.ruleForm.dataSetId = '',
      this.ruleForm.dataSetVersion = ''
      this.dataSetChange = true
      this.getDataSetNameList()
    },
    changeDataSetName() {
      this.dataSetVersion = true
      this.dataSetVersionCount = 1
      this.dataSetVersionOption = [],
      this.ruleForm.dataSetVersion = '',
      this.getDataSetVersionList()
    },
    getDataSetNameList() {
      if (this.ruleForm.dataSetSource === 'my') {
        getMyDatasetList({ pageIndex: this.dataSetNameCount, pageSize: 10 }).then(response => {
          if (response.data.datasets !== null) { 
            this.dataSetNameOption = this.dataSetNameOption.concat(response.data.datasets) 
            this.dataSetNameTotal = response.data.totalSize
         }
        })
      }
      if (this.ruleForm.dataSetSource === 'pre') {
        getPresetDatasetList({ pageIndex: this.dataSetNameCount, pageSize: 10 }).then(response => {
          if (response.data.datasets !== null) { 
            this.dataSetNameOption = this.dataSetNameOption.concat(response.data.datasets) 
            this.dataSetNameTotal = response.data.totalSize
          }
        })
      }
      if (this.ruleForm.dataSetSource === 'common') {
        getPublicDatasetList({ pageIndex: this.dataSetNameCount, pageSize: 10 }).then(response => {
          if (response.data.datasets !== null) { 
            this.dataSetNameOption = this.dataSetNameOption.concat(response.data.datasets) 
            this.dataSetNameTotal = response.data.totalSize
          }
        })
      }
    },
    getDataSetVersionList() {
      let data = {}
      data.datasetId = this.ruleForm.dataSetId
      data.pageIndex = this.dataSetVersionCount
      data.pageSize = 10
      data.status = 3
      getVersionList(data).then(response => {
        if (response.data.versions !== null) { 
          this.dataSetVersionOption = this.dataSetVersionOption.concat(response.data.versions) 
          this.dataSetVersionTotal = response.data.totalSize
        }
      })
    },
    loadDataSetName() {
      this.dataSetNameCount = this.dataSetNameCount + 1
      if (this.dataSetNameOption.length < this.dataSetNameTotal) {
        this.getDataSetNameList()
      }
    },
    loadDataSetVersion() {
      this.dataSetVersionCount = this.dataSetVersionCount + 1
      if (this.dataSetVersionOption.length < this.dataSetVersionTotal) {
        this.getDataSetVersionList()
      }
    }
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