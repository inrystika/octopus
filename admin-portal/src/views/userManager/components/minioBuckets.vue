<template>
  <div>
    <el-dialog 
      title="可读写minio桶" 
      width="35%" 
      :visible.sync="CreateFormVisible"           
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
    <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="60px">
      <el-form-item label="名称" prop="buckets">
        <div v-for="(item, index) in ruleForm.buckets" :key="index">
          <el-input v-model="ruleForm.buckets[index]" style="width: 60%;margin-bottom:10px" />
          <i class="el-icon-delete" @click="deleteItem(index)"></i>
        </div>
        <el-button type="primary" @click="addItem">增加</el-button>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button type="primary" @click="submit" v-preventReClick>提交</el-button>
      <el-button @click="cancel">取消</el-button>
    </div>
    </el-dialog>
  </div>
</template>

<script>
import { createMinIOBucket } from "@/api/userManager.js"
export default {
  name: "minioBuckets",
  props: {
    row: {
      type: Object,
      default: () => { }
    }
  },
  data () {
    var checkName = (rule, value, callback) => {
      const regName = /^[a-z][0-9a-z]{3,3000}$/;
      for(let i = 0; i < value.length; i++) {
        if (value[i].length <= 3 || !regName.test(value[i])) {
          callback(new Error("小写字母开头,长度至少4个字符,只允许由小写字母和数字组成"));
        }
      }
      return callback();
    };
    return {
      CreateFormVisible: true,
      ruleForm: {
        buckets: [],
      },
      rules: {
        buckets: [ 
          {required: true, message: '请输入minio桶的名称', trigger: 'blur'},
          { validator: checkName, trigger: "blur" }
        ]
      }
    }
  },
  created () {
    if(this.row.buckets) {
      this.ruleForm.buckets = this.row.buckets
    }
  },
  methods: {
    handleDialogClose() {
      this.$emit('close', false)
    },
    addItem() {
      this.ruleForm.buckets.push("")
    },
    deleteItem(index) {
      this.ruleForm.buckets.splice(index, 1)
    },
    cancel() {
      this.$emit('cancel', false)
    },
    submit() {
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          const params = {
            buckets: this.ruleForm.buckets,
          }
          const userId = this.row.id
          createMinIOBucket(userId,params).then((response) => {
            if (response.success) {
              this.$message.success("创建成功")
              this.ruleForm.minioPassword = ''
              this.isShow = true
              this.$emit('confirm', false)
            } else {
              this.$message({
                message: this.getErrorMsg(
                  response.error.subcode
                ),
                type: "warning"
              })
            }
          })
        } else {
          console.log('error')
        }
      })
    }
  }
}
</script>
