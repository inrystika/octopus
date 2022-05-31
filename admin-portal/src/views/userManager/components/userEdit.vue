<template>
  <div>
    <el-dialog title="编辑" :visible.sync="createFormVisible" :before-close="handleDialogClose"
      :close-on-click-modal="false" width="35%">
      <el-form :model="ruleForm" ref="ruleForm" label-width="100px">
        <el-form-item label="资源池"  prop="pool">
            <el-select v-model="ruleForm.pool" placeholder="请绑定资源池" multiple>
                <el-option
                    v-for="item in resourcePoolList"
                    :key="item.id"
                    :label="item.id"
                    :value="item.id">
                </el-option>
            </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="update('ruleForm')">确 定</el-button>
      </div>

    </el-dialog>
  </div>
</template>
<script>
  import { editUser } from '@/api/userManager.js'
  export default {
    name: "userEdit",
    props: {
      row: {
        type: Object,
        default: () => { }
      },
      userResourcePoolList: {
        type: Array,
        default: () => []
      }
    },
    data() {
      return {
        createFormVisible: true,
        userInfo: {},
        resourcePoolList: [],
        ruleForm: {
          pool: []
        }
      }
    },
    created() {
      this.resourcePoolList = this.userResourcePoolList
      this.userInfo = this.row
      this.ruleForm.pool = this.row.resourcePools
    },
    methods: {
      handleDialogClose() {
        this.$emit('close', false)
      },
      cancel() {
        this.$emit('cancel', false)
      },
      update(formName) {
        const param = {
            id: this.userInfo.id,
            fullName: this.userInfo.fullName,
            gender: this.userInfo.gender,
            password: undefined,
            phone: this.userInfo.phone,
            resourcePools: this.ruleForm.pool
        }
        editUser(param).then(response => {
            if (response.success) {
                this.$message({
                    message: '资源池绑定成功',
                    type: 'success'
                });
                this.$emit('close', false)
            } else {
                this.$message({
                    message: response.error.message,
                    type: 'error'
                });
            }
        })
      },
    }
  }
</script>
<style lang="scss" scoped>
.el-select {
  width: 60%;
}
</style>