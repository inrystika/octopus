<template>
  <div>
    <el-dialog title="保存环境" :visible.sync="dialogFormVisible" width="30%" :before-close="handleDialogClose"
      :close-on-click-modal="false">
      <el-form :model="form">
        <el-form-item label="子任务" :label-width="formLabelWidth">
          <el-select v-model="form.region" placeholder="请选择子任务">
            <el-option label="区域一" value="shanghai"></el-option>
            <el-option label="区域二" value="beijing"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="镜像名称:" :label-width="formLabelWidth">
          <el-input v-model="form.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="镜像标签:" :label-width="formLabelWidth">
          <el-input v-model="form.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="增量描述:" :label-width="formLabelWidth">
          <el-input type="textarea" v-model="form.desc" maxlength="100"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="confirm">保 存</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
  export default {
    data() {
      return {
        dialogFormVisible: true,
        form: {
          name: '',
          region: '',
          date1: '',
          date2: '',
          delivery: false,
          type: [],
          resource: '',
          desc: ''
        },
        formLabelWidth: '120px'
      };
    },
    methods: {
      cancel() {
        this.$emit("cancel", false);
      },
      handleDialogClose() {
        this.$emit('close', false)
      },
      confirm(val) {
        this.$emit("confirm", false);
      },
    }
  };
</script>