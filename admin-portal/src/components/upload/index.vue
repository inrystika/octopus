<template>
  <div>
    <el-upload v-if="showUpload" class="upload-demo" action="#" :on-change="upload" :file-list="fileList"
      :http-request="httpRequest" multiple :accept="accept" :disabled="show||progress>=0&&progress<=100">
      <el-button size="small" type="primary" :disabled="show||progress>=0&&progress<=100" >点击上传
      </el-button>
      <div class="tipText">{{ tipText }}</div>
    </el-upload>
    <el-button v-if="!showUpload" :loading="loadingShow" size="small" type="primary">上传中</el-button>
    <el-tooltip class="item" effect="dark" :content="message" placement="top-start" v-if="!showUpload">
      <i class="el-icon-warning-outline"></i>
    </el-tooltip>
    <el-progress :text-inside="true" :stroke-width="18" :percentage="progress-1" class="progress"
      v-if="progress>0&&progress<=100" />
  </div>
</template>
<script>
  import { uploadPreImage, finishUpload } from '@/api/imageManager.js'
  import { uploadPreDataset, preDatasetFinishUpload, uploadNewVersion, newVersionFinishUpload } from "@/api/dataManager.js"
  import { uploadPreAlgorithm, preAlgorithmFinishUpload } from "@/api/modelDev.js";
  import { uploadModel, modelFinishUpload } from '@/api/modelManager.js'
  import { minIO } from '@/utils/minIO'
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
        progress: undefined,
        timer: undefined,
        message: '上传过程中，本上传页面关闭，不影响上传，但是关闭或者刷新浏览器，上传会被停止'
      }
    },
    computed: {
      ...mapGetters([
        'progressId',
      ])
    },
    created() {
      this.timer = setInterval(() => {
        if (this.showProgress()) {
          if (parseInt(sessionStorage.getItem(JSON.stringify(store.state.user.progressId)))) {
            this.progress = parseInt(sessionStorage.getItem(JSON.stringify(store.state.user.progressId)))
            this.$emit('upload', true)
          }
        }
      }, 1000)
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
          this.loadingShow = false
        }
        if (0 < a < 100) {
          this.loadingShow = true
        }
      }
    },
    methods: {
      beforeUpload(file, fileList) {

        // sessionStorage.setItem(JSON.stringify(store.state.user.progressId), 0);
      },
      upload(file, fileList) {
        // if (this.uploadData.type = "镜像模块") {
        if (file) {
          this.fileList = [file]
          // sessionStorage.setItem(JSON.stringify(store.state.user.progressId), 0);
        }
        // }
      },
      httpRequest() {
        const fileName = this.fileList[0].name
        const fileForm = fileName.slice(fileName.lastIndexOf(".") + 1).toLowerCase() // 获取上传文件格式后缀
        if (this.uploadData.type === "imageManager") {
          this.loadingShow = true
          this.showUpload = false
          this.show = false
          if (fileForm === 'zip' || fileForm === 'tar') {
            this.$emit('upload', false)
            uploadPreImage({ id: this.uploadData.data.id, fileName: this.fileList[0].name, domain: this.GLOBAL.DOMAIN }).then(response => {
              store.commit('user/SET_PROGRESSID', this.uploadData.data.id)
              const param = {
                uploadUrl: response.data.uploadUrl,
                file: this.fileList[0].raw,
                id: this.uploadData.data.id
              }
              minIO(param).then(response => {
                this.$nextTick(() => {
                  this.loadingShow = false
                  this.show = true
                  this.showUpload = true
                  this.confirm()
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
          this.show = false
          if (fileForm === 'zip') {
            this.$emit('upload', false)
            uploadModel({ modelId: this.uploadData.data.modelId, version: this.uploadData.data.version, fileName: this.fileList[0].name, domain: this.GLOBAL.DOMAIN }).then(response => {
              if (response.success) {
                store.commit('user/SET_PROGRESSID', this.uploadData.data.modelId + this.uploadData.data.version)
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw,
                  id: this.uploadData.data.modelId + this.uploadData.data.version
                }
                minIO(param).then(response => {
                  if (response.success) {
                    this.show = true
                    this.loadingShow = false
                    this.showUpload = true
                    this.confirm()
                  }
                })
              }
              else {
                this.$message({
                  message: this.getErrorMsg(response.error.subcode),
                  type: 'warning'
                });
                this.loadingShow = false
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
        } else if (this.uploadData.type === "preDatasetCreation") {
          this.loadingShow = true
          this.showUpload = false
          this.show = false
          const param = {
            id: this.uploadData.id,
            fileName: this.fileList[0].name,
            version: this.uploadData.version,
            domain: this.GLOBAL.DOMAIN
          }
          if (fileForm === 'zip') {
            this.$emit('upload', false)
            uploadPreDataset(param).then(response => {
              if (response.success) {
                // let uploadUrl = response.data.uploadUrl.replace("octopus-dev-minio:9000","192.168.202.73")
                store.commit('user/SET_PROGRESSID', this.uploadData.id + this.uploadData.version)
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw,
                  id: this.uploadData.id + this.uploadData.version
                }
                minIO(param).then(response => {
                  this.loadingShow = false
                  this.show = true
                  this.showUpload = true
                  this.confirm()
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
          this.show = false
          const param = {
            datasetId: this.uploadData.datasetId,
            fileName: this.fileList[0].name,
            version: this.uploadData.version,
            domain: this.GLOBAL.DOMAIN
          }
          if (fileForm === 'zip') {
            this.$emit('upload', false)
            uploadNewVersion(param).then(response => {
              if (response.success) {
                store.commit('user/SET_PROGRESSID', this.uploadData.datasetId + this.uploadData.version)
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw,
                  id: this.uploadData.datasetId + this.uploadData.version
                }
                minIO(param).then(response => {
                  this.loadingShow = false
                  this.show = true
                  this.showUpload = true
                  this.confirm()
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
          this.showUpload = false
          this.show = false
          const param = {
            algorithmId: this.uploadData.algorithmId,
            FileName: this.fileList[0].name,
            version: this.uploadData.version,
            domain: this.GLOBAL.DOMAIN
          }
          if (fileForm === 'zip') {
            this.$emit('upload', false)
            uploadPreAlgorithm(param).then(response => {
              if (response.success) {
                store.commit('user/SET_PROGRESSID', this.uploadData.algorithmId + this.uploadData.version)
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw,
                  id: this.uploadData.algorithmId + this.uploadData.version
                }
                minIO(param).then(response => {
                  this.loadingShow = false
                  this.show = true
                  this.showUpload = true
                  this.confirm()
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
          this.showUpload = false
          this.show = false
          const param = {
            algorithmId: this.uploadData.algorithmId,
            FileName: this.fileList[0].name,
            version: this.uploadData.version,
            domain: this.GLOBAL.DOMAIN
          }
          if (fileForm === 'zip') {
            this.$emit('upload', false)
            uploadPreAlgorithm(param).then(response => {
              store.commit('user/SET_PROGRESSID', this.uploadData.algorithmId + this.uploadData.version)
              if (response.success) {
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw,
                  id: this.uploadData.algorithmId + this.uploadData.version
                }
                minIO(param).then(response => {
                  this.loadingShow = false
                  this.show = true
                  this.showUpload = true
                  this.confirm()
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
                message: '上传成功',
                type: 'success'
              });
              sessionStorage.setItem(JSON.stringify(this.uploadData.data.id), 0),
                this.$emit('confirm', false)
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
              sessionStorage.setItem(JSON.stringify(this.uploadData.data.id), 0)
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
                sessionStorage.setItem(JSON.stringify(this.uploadData.data.modelId + this.uploadData.data.version), 0),
                  this.$emit('confirm', false)
              } else {
                this.$message({
                  message: this.getErrorMsg(response.error.subcode),
                  type: 'warning'
                });
                sessionStorage.setItem(JSON.stringify(this.uploadData.data.modelId + this.uploadData.data.version), 0)
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
              sessionStorage.setItem(JSON.stringify(this.uploadData.id + this.uploadData.version), 0),
                this.$message.success("上传预置数据集成功");
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
              sessionStorage.setItem(JSON.stringify(this.uploadData.id + this.uploadData.version), 0)
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
              sessionStorage.setItem(JSON.stringify(this.uploadData.datasetId + this.uploadData.version), 0),
                this.$message.success("上传预置数据集版本成功");
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
              sessionStorage.setItem(JSON.stringify(this.uploadData.datasetId + this.uploadData.version), 0)
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
              sessionStorage.setItem(JSON.stringify(this.uploadData.algorithmId + this.uploadData.version), 0),
                this.$message.success("上传预置算法成功");
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
              sessionStorage.setItem(JSON.stringify(this.uploadData.algorithmId + this.uploadData.version), 0)
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
              sessionStorage.setItem(JSON.stringify(this.uploadData.algorithmId + this.uploadData.version), 0),
                this.$message.success("上传预置算法新版本成功");
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
              sessionStorage.setItem(JSON.stringify(this.uploadData.algorithmId + this.uploadData.version), 0)
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
      },
      // 显示进度条
      showProgress() {
        if (store.state.user.progressId) {
          if (store.state.user.progressId == this.uploadData.data.id) {
            return true
          }
          if (store.state.user.progressId == this.uploadData.datasetId + this.uploadData.version) {
            return true
          }
          if (store.state.user.progressId == this.uploadData.id + this.uploadData.version) {
            return true
          }
          if (store.state.user.progressId == this.uploadData.algorithmId + this.uploadData.version) {
            return true
          }

          if (store.state.user.progressId == this.uploadData.data.modelId + this.uploadData.data.version) {
            return true
          }
          else { return false }
        }
        else { return false }

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

  .item {
    margin-left: 5px;
    font-size: 16px;
    color: #409EFF;
  }
</style>
