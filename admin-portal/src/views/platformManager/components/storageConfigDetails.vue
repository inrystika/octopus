<template>
  <div>
    <el-dialog
      title="配置详情"
      width="35%"
      :visible.sync="CreateFormVisible"
      :before-close="handleDialogClose" 
      :close-on-click-modal="false"
      :append-to-body="true"
    >
      <el-form       
        label-width="100px" 
        class="demo-ruleForm"
      >
        <div v-for="(item, index) in propList" :key="index">
          <el-form-item :label="item.key + ':'">
            {{item.value}}
          </el-form-item>
        </div>
      </el-form>
    </el-dialog>
  </div>
</template>
<script>
export default {
  name: "storageConfigDetails",
  props: {
    storageConfigDetail: {
      type: Array,
      default: []
    },
  },
  data() {
    return {
      CreateFormVisible: true,
      propList: []
    }
  },
  created() {
    this.propList = this.storageConfigDetail
  },
  methods: {
    handleDialogClose() {
      this.$emit('close', false)
    }
  }
}
</script>
<style lang="scss" scoped>

</style>