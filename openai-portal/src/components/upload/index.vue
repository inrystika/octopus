<template>
  <div>
    <el-upload v-if="showUpload" class="upload-demo" action="#" :on-change="upload" :file-list="fileList"
      :http-request="httpRequest" multiple :accept="accept" :before-upload="beforeUpload"
      :disabled="show||progress>0&&progress<100">
      <el-button size="small" type="primary" :disabled="show||progress>0&&progress<100">点击上传
      </el-button>
      <div class="tipText">{{ tipText }}</div>
    </el-upload>
    <el-button v-if="!showUpload" :loading="loadingShow" size="small" type="primary">上传中</el-button>
    <el-progress :text-inside="true" :stroke-width="18" :percentage="progress" class="progress"
      v-if="progress>0&&progress<100" />
    <!-- <div v-if="show" slot="footer" class="dialog-footer">
      <el-button @click="cancel">取 消</el-button>
      <el-button type="primary" @click="confirm">确 定</el-button>
    </div> -->
  </div>
</template>
<script>
  import { uploadImage, finishUpload, uploadMiniIO } from '@/api/imageManager.js'
  import { uploadMyDataset, myDatasetFinishUpload, uploadNewVersion, newVersionFinishUpload } from "@/api/datasetManager.js"
  import { uploadMyAlgorithm, myAlgorithmFinishUpload } from "@/api/modelDev.js";
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
        progress: undefined,
        timer: undefined
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
          this.loadingShow = false
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
          this.show = false
          if (fileForm === 'zip' || fileForm === 'tar') {
            uploadImage({ id: this.uploadData.data.id, fileName: this.fileList[0].name, domain: this.GLOBAL.DOMAIN }).then(response => {
              if (response.success) {
                store.commit('user/SET_PROGRESSID', this.uploadData.data.id)
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw,
                  id: this.uploadData.data.id
                }
                uploadMiniIO(param).then(response => {
                  this.$nextTick(() => {
                    this.loadingShow = false
                    this.show = true
                    this.showUpload = true
                    this.confirm()
                  })

                })
              }
              else {
                this.$message({
                  message: this.getErrorMsg(response.error.subcode),
                  type: 'warning'
                });
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
        } else if (this.uploadData.type === "myDatasetCreation") {
          this.loadingShow = true
          this.showUpload = false
          const param = {
            id: this.uploadData.id,
            fileName: this.fileList[0].name,
            version: this.uploadData.version,
            domain: this.GLOBAL.DOMAIN
          }
          if (fileForm === 'zip') {
            uploadMyDataset(param).then(response => {
              if (response.success) {
                store.commit('user/SET_PROGRESSID', this.uploadData.id+this.uploadData.version)
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw,
                  id: this.uploadData.id+this.uploadData.version,
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
        } else if (this.uploadData.type === 'myAlgorithmCreation') {
          this.loadingShow = true
          this.showUpload = false
          const param = {
            algorithmId: this.uploadData.AlgorithmId,
            FileName: this.fileList[0].name,
            version: this.uploadData.Version,
            domain: this.GLOBAL.DOMAIN
          }
          if (fileForm === 'zip') {
            uploadMyAlgorithm(param).then(response => {
              if (response.success) {
                store.commit('user/SET_PROGRESSID', this.uploadData.AlgorithmId + this.uploadData.Version)
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw,
                  id: this.uploadData.AlgorithmId + this.uploadData.Version,
                }
                minIO(param).then(response => {
                  this.loadingShow = false
                  this.show = true
                  this.showUpload = true,
                    this.confirm()
                })
              } else {
                this.$message({
                  message: this.getErrorMsg(response.error.subcode),
                  type: 'warning'
                });
                // this.loadingShow = false
                // this.show = true
                // this.showUpload = true
                // this.fileList = []
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
        } else if (this.uploadData.type === "newDatasetVersionCreation") {
          this.loadingShow = true
          this.showUpload = false
          const param = {
            id: this.uploadData.id,
            fileName: this.fileList[0].name,
            version: this.uploadData.version,
            domain: this.GLOBAL.DOMAIN
          }
          if (fileForm === 'zip') {
            uploadNewVersion(param).then(response => {
              store.commit('user/SET_PROGRESSID', this.uploadData.id+this.uploadData.version)
              if (response.success) {
                const param = {
                  uploadUrl: response.data.uploadUrl,
                  file: this.fileList[0].raw,
                  id: this.uploadData.id+this.uploadData.version,
                }
                minIO(param).then(response => {
                  this.loadingShow = false
                  this.show = true
                  this.showUpload = true,
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
          finishUpload(this.uploadData.data.id).then(
            response => {
              if (response.success) {
                this.$message({
                  message: '上传成功',
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

            },


          )
        } else if (this.uploadData.type === "myDatasetCreation") {
          const payload = {
            id: this.uploadData.id,
            version: this.uploadData.version,
            fileName: this.fileList[0].name
          }
          myDatasetFinishUpload(payload).then(response => {
            if (response.success) {
              store.commit('user/SET_PROGRESSID', this.uploadData.id)
              this.$message.success("上传数据集成功");
              sessionStorage.setItem(JSON.stringify(store.state.user.progressId), 0),
                this.$emit('confirm', false)
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
        } else if (this.uploadData.type === 'myAlgorithmCreation') {
          const payload = {
            algorithmId: this.uploadData.AlgorithmId,
            version: this.uploadData.Version,
            fileName: this.fileList[0].name
          }
          myAlgorithmFinishUpload(payload).then(response => {
            if (response.success) {
              this.$message.success("上传我的算法成功");
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          }, this.$emit('confirm', false))
        } else if (this.uploadData.type === 'newDatasetVersionCreation') {
          const payload = {
            id: this.uploadData.id,
            version: this.uploadData.version,
            fileName: this.fileList[0].name
          }
          newVersionFinishUpload(payload).then(response => {
            if (response.success) {
              this.$message.success("上传数据集新版本成功");
              sessionStorage.setItem(JSON.stringify(store.state.user.progressId), 0),
                this.$emit('confirm', false)
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
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
          if (store.state.user.progressId == this.uploadData.id) {
            return true
          }
          if (store.state.user.progressId == this.uploadData.AlgorithmId + this.uploadData.Version) {
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
</style>