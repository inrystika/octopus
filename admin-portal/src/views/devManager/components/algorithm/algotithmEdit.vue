<template>
  <div>
    <el-dialog title="编辑" :visible.sync="dialogFormVisible" width="30%" :before-close="handleDialogClose"
      :close-on-click-modal="false" v-if="show">
      <el-form :model="form">
        <el-form-item label="模型类别:" :label-width="formLabelWidth">
          <el-select v-model="form.applyId" placeholder="请选择">
            <el-option v-for="item in typeOptions" :key="item.id" :label="item.lableDesc" :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="框架类型:" :label-width="formLabelWidth">
          <el-select v-model="form.frameworkId" placeholder="请选择">
            <el-option v-for="item in useOptions" :key="item.id" :label="item.lableDesc" :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="算法描述:" :label-width="formLabelWidth">
          <el-input v-model="form.algorithmDescript" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="submit">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  import { editAlgorithm, algorithmType, frameType } from "@/api/modelDev";
  export default {
    name: "algotithmEdit",
    props: {
      row: {
        type: Object,
        default: {}
      }
    },
    created() {
      this.form.algorithmId = this.row.algorithmId
      this.form.applyId = this.row.applyId
      this.form.frameworkId = this.row.frameworkId
      this.form.desc = this.row.desc
      this.algorithmType()
      this.frameType()
    },
    data() {
      return {
        form: { algorithmId: '', frameworkId: '', applyId: '', desc: '' },
        dialogFormVisible: true,
        formLabelWidth: '120px',
        typeOptions: [],
        useOptions: [],
        show: false,
        apply: false,
        framework: false

      };
    },
    watch: {
      apply() {
        if (this.apply && this.framework) {
          this.show = true
        }
      },
      framework() {
        if (this.apply && this.framework) {
          this.show = true
        }
      }
    },
    methods: {
      cancel() {
        this.$emit("cancel", false);
      },
      handleDialogClose() {
        this.$emit('close', false)
      },
      confirm(val) {
        this.$emit("confirm", val);
      },
      // 获取算法类型
      algorithmType() {
        algorithmType({ pageIndex: 1, pageSize: 20 }).then(response => {
          if (response.success) {
            this.typeOptions = response.data.lables
            this.apply = true
          } else {
            // this.showUpload = false
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      },
      // 获取算法框架
      frameType() {
        frameType({ pageIndex: 1, pageSize: 20 }).then(response => {
          if (response.success) {
            this.useOptions = response.data.lables
            this.framework = true
          } else {
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      },
      submit() {
        editAlgorithm(this.form).then(response => {
          if (response.success) {
            this.$message({
              message: '编辑成功',
              type: 'success'
            });
            this.$emit('confirm', false)
          } else {

            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      }
    }
  };
</script>