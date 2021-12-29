<template>
  <div>
    <el-dialog title="编辑" :visible.sync="dialogFormVisible" width="30%" :before-close="handleDialogClose"
      :close-on-click-modal="false">
      <el-form :model="form">
        <el-form-item label="数据类型:" :label-width="formLabelWidth">
          <el-select v-model="form.typeId" placeholder="请选择">
            <el-option v-for="item in typeOptions" :key="item.id" :label="item.lableDesc" :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="标注类型:" :label-width="formLabelWidth">
          <el-select v-model="form.applyId" placeholder="请选择">
            <el-option v-for="item in useOptions" :key="item.id" :label="item.lableDesc" :value="item.id">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="数据集描述:" :label-width="formLabelWidth">
          <el-input v-model="form.desc" autocomplete="off" width="100px"></el-input>
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
  import { editDataSet, datasetType, datasetUse } from "@/api/datasetManager.js"
  import { getErrorMsg } from '@/error/index'
  export default {
    name: "dataSetEdit",
    props: {
      data: {
        type: Object,
        default: {}
      }
    },
    created() {
      this.form.datasetId = this.data.id
      this.form.typeId = this.data.typeId
      this.form.applyId = this.data.applyId
      this.form.desc = this.data.desc
      this.datasetType()
      this.datasetUse()
    },
    data() {
      return {
        form: { datasetId: '', typeId: '', applyId: '', desc: '' },
        dialogFormVisible: true,
        formLabelWidth: '120px',
        typeOptions: [],
        useOptions: []
      };
    },
    methods: {
      getErrorMsg(code) {
        return getErrorMsg(code)
      },
      cancel() {
        this.$emit("cancel", false);
      },
      handleDialogClose() {
        this.$emit('close', false)
      },
      confirm(val) {
        this.$emit("confirm", val);
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
            this.$message({
              message: this.getErrorMsg(response.error.subcode),
              type: 'warning'
            });
          }
        })
      },
      submit() {
        editDataSet(this.form).then(response => {
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