<template>
  <div>
    <el-dialog title="复制算法" width="35%" :visible.sync="CreateListVisible" :before-close="handleDialogClose"
      :close-on-click-modal="false">
      <el-form ref="ruleForm" v-loading.fullscreen.lock="loading" :model="ruleForm" :rules="rules" label-width="100px">
        <el-form-item label="算法名称：" :label-width="formLabelWidth" prop="name">
          <el-input v-model="ruleForm.name" placeholder="请输入算法名称" />
        </el-form-item>
        <el-form-item label="模型名称：" :label-width="formLabelWidth" prop="modelName">
          <el-input v-model="ruleForm.modelName"></el-input>
        </el-form-item>
        <el-form-item label="算法描述：" :label-width="formLabelWidth" prop="desc">
          <el-input v-model="ruleForm.desc" :autosize="{ minRows: 2, maxRows: 4}" placeholder="请输入算法描述" maxlength="300"
            show-word-limit />
        </el-form-item>
        <el-form-item label="复制版本：" :label-width="formLabelWidth" prop="version">
          <el-select v-model="ruleForm.version" value-key="algorithmVersion" @visible-change="getAlgorithmSource">
            <el-option v-for="item in algorithmList" :key="item.algorithmVersion" :label="item.algorithmVersion"
              :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="说明：" :label-width="formLabelWidth">
          <!-- <ul> -->
          <li style="list-style: none">会复制基础版本代码到新版本中，后续可在新版本中做修改。</li>
          <!-- </ul> -->
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="confirm('ruleForm')">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  import { getAlgorithmVersionList, getPubAlgorithmVersionList, copyAlgorithm } from "@/api/modelDev.js";
  export default {
    name: "AlgorithmCopy",
    props: {
      row: {
        type: Object,
        default: () => { }
      },
      algorithmTabType: { type: Number, default: undefined }
    },
    data() {
      return {
        ruleForm: {
          name: "",
          modelName: "",
          version: "",
          desc: ""
        },
        rules: {
          modelName: [
            {
              required: true,
              message: "请输入模型名称",
              trigger: "blur"
            }
          ],
          name: [
            {
              required: true,
              message: "请输入算法名称",
              trigger: "blur"
            }
          ],
          version: [
            {
              required: true,
              message: "请选择复制版本",
              trigger: "change"
            }
          ]
        },
        pageIndex: 1,
        pageSize: 20,
        algorithmList: [],
        CreateListVisible: true,
        loading: false,
        formLabelWidth: "120px"
      }
    },
    methods: {
      handleDialogClose() {
        this.$emit("close", false);
      },
      getAlgorithmSource() {
        const param = {
          pageIndex: this.pageIndex,
          pageSize: this.pageSize,
          algorithmId: this.row.algorithmId
        }
        if (this.algorithmTabType === 2) {
          getPubAlgorithmVersionList(param).then(response => {
            if (response.success) {
              this.algorithmList = response.data.algorithms
            } else {
              this.$message({
                message: this.getErrorMsg(response.error.subcode),
                type: 'warning'
              });
            }
          })
        } else {
          getAlgorithmVersionList(param).then(response => {
            if (response.success) {
              const newArr = []
              response.data.algorithms.filter(function (item, index) {
                newArr.push(item.algorithmDetail)
              })
              this.algorithmList = newArr
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
      confirm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.loading = true
            const param = {
              version: this.ruleForm.version.algorithmVersion,
              algorithmId: this.ruleForm.version.algorithmId,
              modelName: this.ruleForm.modelName,
              algorithmDescript: this.ruleForm.desc,
              newAlgorithmName: this.ruleForm.name
            }
            copyAlgorithm(param).then(response => {
              if (response.success) {
                this.$message.success("复制成功，请在’我的算法‘中查看");
                this.loading = false
                this.$emit('confirm', false)
              } else {
                this.$message({
                  message: this.getErrorMsg(response.error.subcode),
                  type: 'warning'
                });
                this.loading = false
              }
            })
          } else {
            this.loading = false
            return false;
          }
        });
      }
    }
  }
</script>