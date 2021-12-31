<template>
  <div>
    <el-dialog title="保存环境" :visible.sync="dialogFormVisible" width="30%" :before-close="handleDialogClose"
      :close-on-click-modal="false">
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm">
        <el-form-item label="原镜像名称:" :label-width="formLabelWidth">
          <el-input v-model="ruleForm.name" autocomplete="off" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="原镜像标签:" :label-width="formLabelWidth">
          <el-input v-model="ruleForm.version" autocomplete="off" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="描述:" :label-width="formLabelWidth">
          <el-input type="textarea" v-model="ruleForm.desc" maxlength="100" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="子任务" :label-width="formLabelWidth" prop="taskName">
          <el-select v-model="ruleForm.taskName" placeholder="请选择子任务">
            <el-option v-for="item in options" :key="item.value" :label="item.name" :value="item.name">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="镜像名称:" :label-width="formLabelWidth" prop="imageName">
          <el-input v-model="ruleForm.imageName" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="镜像标签:" :label-width="formLabelWidth" prop="imageVersion">
          <el-input v-model="ruleForm.imageVersion" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="增量描述:" :label-width="formLabelWidth" prop="LayerDescription">
          <el-input type="textarea" v-model="ruleForm.LayerDescription" maxlength="100" show-word-limit></el-input>
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
  import { saveNoteBook } from "@/api/modelDev";
  export default {
    props: {
      row: {
        type: Object,
        default: () => { }
      }
    },
    created() {
      this.ruleForm.name = this.row.imageName
      this.ruleForm.version = this.row.imageVersion
      this.ruleForm.desc = this.row.desc
      this.options = this.row.tasks
      this.id = this.row.id
    },
    data() {
       var checkName = (rule, value, callback) => {
        const regName = /^[a-zA-Z][\w|-]*$/;
        if (regName.test(value)) {
          return callback();
        }
        callback(new Error("请输入合法的镜像名称:首字母为大小写字母，其他大小写字母数字或者-"));
      };
      var checkLabel = (rule, value, callback) => {
        const regLabel = /^[a-zA-Z][\w|\-|\.]+$/;
        if (regLabel.test(value)) {
          return callback();
        }
        callback(new Error("请输入合法的标签名称:首字母为英文,其他为英文数字.或者-"));
      };
      return {
        dialogFormVisible: true,
        form: { imageName: '', imageVersion: '', desc: '' },
        ruleForm: {
          name: '',
          version: '',
          desc: '',
          taskName: '',
          imageName: '',
          imageVersion: '',
          LayerDescription: ''
        },
        id: '',
        options: [],
        rules: {
          taskName: [
            { required: true, message: '请选择子任务', trigger: 'change' }
          ],
          imageName: [
            { required: true, message: '请输入镜像名称', trigger: 'blur' },
            { validator: checkName, trigger: "blur" }
          ],
          imageVersion: [
            { required: true, message: '请输入镜像标签', trigger: 'blur' },
            { validator: checkLabel, trigger: "blur" }
          ],
          LayerDescription: [
            { required: true, message: '请填写增量描述', trigger: 'blur' }
          ]
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
        this.$refs['ruleForm'].validate((valid) => {
          if (valid) {
            console.log(this.ruleForm.LayerDescription)
            saveNoteBook({ id: this.id, taskName: this.ruleForm.taskName, imageName: this.ruleForm.imageName, imageVersion: this.ruleForm.imageVersion, layerDescription: this.ruleForm.LayerDescription }).then(response => {
              if (response.success) {
                this.$message({
                  message: '保存成功',
                  type: 'success'
                })
                this.$emit("confirm", false);
              } else {
                this.$message({
                  message: this.getErrorMsg(response.error.subcode),
                  type: 'warning'
                })
              }
            })

          } else {
            console.log('error submit!!');
            return false;
          }
        });

      },
    }
  };
</script>