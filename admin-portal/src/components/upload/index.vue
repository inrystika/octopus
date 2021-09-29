<template>
  <div>
    <el-upload v-if="showUpload" class="upload-demo" action="#" :on-change="upload" :file-list="fileList"
      :http-request="httpRequest" multiple :accept="accept">
      <el-button size="small" type="primary" :disabled="show||progress>0&&progress<100">点击上传
      </el-button>
      <div class="tipText">{{ tipText }}</div>
    </el-upload>
    <el-button v-if="!showUpload" :loading="loadingShow" size="small" type="primary">上传中</el-button>
    <el-progress :text-inside="true" :stroke-width="18" :percentage="progress" class="progress"
      v-if="progress>0&&progress<100" />
    <div v-if="show" slot="footer" class="dialog-footer">
      <el-button @click="cancel">取 消</el-button>
      <el-button type="primary" @click="confirm">确 定</el-button>
    </div>
  </div>
</template>
<script>
  import { uploadPreImage, finishUpload, uploadMiniIO } from '@/api/imageManager.js'
  import { uploadPreDataset, preDatasetFinishUpload, uploadNewVersion, newVersionFinishUpload } from "@/api/dataManager.js"
  import { uploadPreAlgorithm, preAlgorithmFinishUpload } from "@/api/modelDev.js";
  import { uploadModel, modelFinishUpload } from '@/api/modelManager.js'
  import { minIO } from '@/utils/minIO'
  import { getErrorMsg } from '@/error/index'
  import { mapGetters } from 'vuex'
  import store from '@/store'
  export default {
    props: {
      uploadData: {
        type: Object,
        default: () => { }
      }

    },
    data() {
      return {
        fileList: [],
        url: undefined,
        show: false,
        loadingShow: false,
        showUpload: true,
        accept: "application/zip",
        tipText: '上传文件格式为 zip',
        progress:0
      }
    },
    computed: {
      ...mapGetters([
        'progressId',
      ])
    },
    created() {
      this.timer = setInterval(() => {
        if (store.state.user.progressId && store.state.user.progressId == this.uploadData.data.id) {
          if (parseInt(sessionStorage.getItem(JSON.stringify(store.state.user.progressId)))) {
            this.progress = parseInt(sessionStorage.getItem(JSON.stringify(store.state.user.progressId)))
            console.log(this.progress)
          }
        }
      }, 100)

      if (this.uploadData.type === "imageManager") {
        this.accept = "application/zip,.tar"
        this.tipText = '上传文件格式为 zip 或 tar'
      }
    },
    destory() {
      clearInterval(this.timer)
    },
    watch: {
      progress(a, b) {
        if (a == 100) {       
          this.show = true
          this.loadingShow=false
          console.log(this.loadingShow)
        }
        if (0 < a < 100) {
          this.loadingShow = true
        }
      }
    },
    methods: {
      getErrorMsg(code) {
        return getErrorMsg(code)
      },
      beforeUpload() {
        sessionStorage.setItem(JSON.stringify(store.state.user.progressId), 0);
      },
      upload(file, fileList) {
        // if (this.uploadData.type = "镜像模块") {
        if (file) {
          this.fileList = [file]
          sessionStorage.setItem(JSON.stringify(store.state.user.progressId), 0);
        }
        // }
      },
      httpRequest() {
        const fileName = this.fileList[0].name
        const fileForm = fileName.slice(fileName.lastIndexOf(".") + 1).toLowerCase() // 获取上传文件格式后缀
        if (this.uploadData.type === "imageManager") {
          this.loadingShow = true
          this.showUpload = false
          if (fileForm === 'zip' || fileForm === 'tar') {
            uploadPreImage({ id: this.uploadData.data.id, fileName: this.fileList[0].name, domain: this.GLOBAL.DOMAIN }).then(response => {
              const param = {
                uploadUrl: response.data.uploadUrl,
                file: this.fileList[0].raw,
                id:this.uploadData.data.id
              }
              uploadMiniIO(param).then(response => {
                  this.$nextTick(() => {
                    this.loadingShow = false
                    this.show = true
                    this.showUpload = true
                  })

                })
            })
          } else {
            this.loadingShow = false
            this.showUpload = true
            this.fileList = []
            this.$message({
              message: '上传文件格式不正确',
              type: 'warning'
            });
          }
        }
        if (this.uploadData.type === "modelManager") {
          this.loadingShow = true
          this.showUpload = false
          if (fileForm === 'zip') {
            uploadModel({ modelId: this.uploadData.data.modelId, version: this.uploadData.data.version, fileName: this.fileList[0].name, domain: this.GLOBAL.DOMAIN }).then(response => {
              const param = {
                uploadUrl: response.data.uploadUrl,
                file: this.fileList[0].raw
              }
              uploadMiniIO(param).then(response => {
                if (response.success) {
                  this.show = true
                  this.loadingShow = false
                  this.showUpload = true
                }
              })
            })
          } else {
            this.loadingShow = false
            this.showUpload = true
            this.fileList = []
            this.$message({
              message: '上传文件格式不正确',
              type: 'warning'
            });
          }
        } else if (this.uploadData.type === "preDatasetCreation") {
          this.loadingShow = true
          this.showUpload = false
          const param = {
            id: this.uploadData.id,
            fileName: this.fileList[0].name,
            version: this.uploadData.version,
            domain: this.GLOBAL.DOMAIN
          }
          if (fileForm === 'zip') {
            uploadPreDataset(param).then(response => {
              if (response.success) {
                // let uploadUrl = response.data.uploadUrl.replace("octopus-dev-minio:9000","192.168.202.73")
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw
                }
                minIO(param).then(response => {
                  this.loadingShow = false
                  this.show = true
                  this.showUpload = true
                })
              } else {
                this.$message({
                  message: this.getErrorMsg(response.error.subcode),
                  type: 'warning'
                });
                this.loadingShow = false
                this.show = true
                this.showUpload = true
                this.fileList = []
              }
            })
          } else {
            this.loadingShow = false
            this.showUpload = true
            this.fileList = []
            this.$message({
              message: '上传文件格式不正确',
              type: 'warning'
            });
          }
        } else if (this.uploadData.type === "newPreDatasetVersion") {
          this.loadingShow = true
          this.showUpload = false
          const param = {
            datasetId: this.uploadData.datasetId,
            fileName: this.fileList[0].name,
            version: this.uploadData.version,
            domain: this.GLOBAL.DOMAIN
          }
          if (fileForm === 'zip') {
            uploadNewVersion(param).then(response => {
              if (response.success) {
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw
                }
                minIO(param).then(response => {
                  this.loadingShow = false
                  this.show = true
                  this.showUpload = true
                })
              } else {
                this.$message({
                  message: this.getErrorMsg(response.error.subcode),
                  type: 'warning'
                });
                this.loadingShow = false
                this.show = true
                this.showUpload = true
                this.fileList = []
              }
            })
          } else {
            this.loadingShow = false
            this.showUpload = true
            this.fileList = []
            this.$message({
              message: '上传文件格式不正确',
              type: 'warning'
            });
          }
        } else if (this.uploadData.type === 'newPreAlgorithm') {
          this.loadingShow = true
          const param = {
            algorithmId: this.uploadData.algorithmId,
            FileName: this.fileList[0].name,
            version: this.uploadData.version,
            domain: this.GLOBAL.DOMAIN
          }
          if (fileForm === 'zip') {
            uploadPreAlgorithm(param).then(response => {
              if (response.success) {
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw
                }
                minIO(param).then(response => {
                  this.loadingShow = false
                  this.show = true
                  this.showUpload = true
                })
              } else {
                this.$message({
                  message: this.getErrorMsg(response.error.subcode),
                  type: 'warning'
                });
                this.loadingShow = false
                this.show = true
                this.showUpload = true
                this.fileList = []
              }
            })
          } else {
            this.loadingShow = false
            this.showUpload = true
            this.fileList = []
            this.$message({
              message: '上传文件格式不正确',
              type: 'warning'
            });
          }
        } else if (this.uploadData.type === 'newPreAlgorithmVersion') {
          this.loadingShow = true
          const param = {
            algorithmId: this.uploadData.algorithmId,
            FileName: this.fileList[0].name,
            version: this.uploadData.version,
            domain: this.GLOBAL.DOMAIN
          }
          if (fileForm === 'zip') {
            uploadPreAlgorithm(param).then(response => {
              if (response.success) {
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw
                }
                minIO(param).then(response => {
                  this.loadingShow = false
                  this.show = true
                  this.showUpload = true
                })
              } else {
                this.$message({
                  message: this.getErrorMsg(response.error.subcode),
                  type: 'warning'
                });
                this.loadingShow = false
                this.show = true
                this.showUpload = true
                this.fileList = []
              }
            })
          } else {
            this.loadingShow = false
            this.showUpload = true
            this.fileList = []
            this.$message({
              message: '上传文件格式不正确',
              type: 'warning'
            });
          }
        }
      },
      confirm() {
        if (this.uploadData.type === "imageManager") {
          finishUpload({ id: this.uploadData.data.id }).then(response => {
            if (response.success) {
              this.$message({
                message: '创建成功',
                type: 'success'
              });
              sessionStorage.setItem(JSON.stringify(store.state.user.progressId), 0),
              this.$emit('confirm', false)
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
        } else if (this.uploadData.type === "modelManager") {
          modelFinishUpload({ fileName: this.fileList[0].name, modelId: this.uploadData.data.modelId, version: this.uploadData.data.version }).then(
            response => {
              if (response.success) {
                this.$message({
                  message: '创建成功',
                  type: 'success'
                });
                this.$emit('confirm', false)
              } else {
                this.$message({
                  message: this.getErrorMsg(response.error.subcode),
                  type: 'warning'
                });
              }
            }

          )
        } else if (this.uploadData.type === "preDatasetCreation") {
          const payload = {
            id: this.uploadData.id,
            version: this.uploadData.version,
            fileName: this.fileList[0].name || ''
          }
          preDatasetFinishUpload(payload).then(response => {
            if (response.success) {
              this.$message.success("上传预置数据集成功");
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          }, this.$emit('confirm', false))
        } else if (this.uploadData.type === 'newPreDatasetVersion') {
          const payload = {
            datasetId: this.uploadData.datasetId,
            version: this.uploadData.version,
            fileName: this.fileList[0].name
          }
          newVersionFinishUpload(payload).then(response => {
            if (response.success) {
              this.$message.success("上传预置数据集版本成功");
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          }, this.$emit('confirm', false))
        } else if (this.uploadData.type === 'newPreAlgorithm') {
          const payload = {
            algorithmId: this.uploadData.algorithmId,
            version: this.uploadData.version,
            fileName: this.fileList[0].name
          }
          preAlgorithmFinishUpload(payload).then(response => {
            if (response.success) {
              this.$message.success("上传预置算法成功");
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          }, this.$emit('confirm', false))
        } else if (this.uploadData.type === 'newPreAlgorithmVersion') {
          const payload = {
            algorithmId: this.uploadData.algorithmId,
            version: this.uploadData.version,
            fileName: this.fileList[0].name
          }
          preAlgorithmFinishUpload(payload).then(response => {
            if (response.success) {
              this.$message.success("上传预置算法新版本成功");
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          }, this.$emit('confirm', false))
        }
      },
      cancel() {
        this.$confirm('此操作将被取消，是否继续?', '提示', {
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
      }

    }
  }
</script>
<style lang="scss" scoped>
  .dialog-footer {
    text-align: right;
  }

  .tipText {
    float: right;
    margin-left: 10px;
    font-size: 12px
  }

  .progress {
    margin: 5px 0px 10px 0px;
  }

  .dialog-footer {
    margin-top: 10px;
  }
</style>