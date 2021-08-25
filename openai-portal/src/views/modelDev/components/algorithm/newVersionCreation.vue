<template>
  <div>
    <el-dialog
      title="创建新版本"
      width="35%"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose"
      :close-on-click-modal="false"
    >
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" v-loading.fullscreen.lock="loading">
        <el-form-item label="算法名称：" :label-width="formLabelWidth" prop="name">
          <el-input v-model="ruleForm.name" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="算法描述：" :label-width="formLabelWidth" prop="desc">
          <el-input
            v-model="ruleForm.desc"
            :autosize="{ minRows: 2, maxRows: 4}"
            placeholder="请输入算法描述"
            maxlength="300"
            show-word-limit
          ></el-input>
        </el-form-item>
        <el-form-item label="基础版本：" :label-width="formLabelWidth" prop="version">
          <el-select value-key="algorithmVersion" @visible-change="getAlgorithmSource" v-model="ruleForm.version">
            <el-option 
              v-for="item in algorithmList" 
              :key="item.algorithmVersion" 
              :label="item.algorithmVersion" 
              :value="item"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="说明：" :label-width="formLabelWidth">
          <!-- <ul> -->
            <li style="list-style: none">会在基础版本上创建新版本，后续可在新版本中做修改。</li>
          <!-- </ul> -->
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="submit('ruleForm')">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getAlgorithmVersionList, createNewAlgorithmVersion } from "@/api/modelDev.js";
import { getErrorMsg } from '@/error/index'
export default {
  name: "newVersionCreation",
  props: {
    row:{
      type: Object,
      default: {}
    },
    newVersionName:"",
    versionList: []
  },
  data() {
    return {
      ruleForm: {
        version: "",
        desc: "",
      },
      rules: {
        version:[
          {
            required: true,
            message: "请选择基础版本",
            trigger: "blur"
          },
        ]
      },
      pageIndex: 1,
      pageSize: 20,
      algorithmList: [],
      CreateFormVisible: true,
      loading: false,
      formLabelWidth: "120px"
    }
  },
  created(){
    this.ruleForm.name = this.row.algorithmName
  },
  methods: {
    getErrorMsg(code) {
      return getErrorMsg(code)
    },
    handleDialogClose() {
      this.$emit("close", false);
    },
    getAlgorithmSource(){
      const param = {
        pageIndex: this.pageIndex,
        pageSize: this.pageSize,
        algorithmId: this.row.algorithmId
      }
      getAlgorithmVersionList(param).then(response => {
        if(response.success) {
          let newArr = []
          response.data.algorithms.filter(function(item,index) {
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
    },
    cancel() {
      this.$confirm('此操作将被取消，是否继续?','提示',{
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
    submit(formName) {
      this.loading = true
      this.$refs[formName].validate((valid) => {
        if (valid) {
          const param = {
            // spaceId: this.row.spaceId,
            // userId: this.row.userId,
            algorithmId: this.row.algorithmId,
            oriVersion: this.ruleForm.version.algorithmVersion,
            algorithmDescript: this.ruleForm.desc
          }
          createNewAlgorithmVersion(param).then(response => {
            if(response.success) {            
              this.$message.success("创建成功");
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
    },
  }
}
</script>