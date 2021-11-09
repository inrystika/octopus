<template>
    <div>
        <el-dialog title="重新上传" :visible.sync="dialogFormVisible" width="650px" :before-close="handleDialogClose"
            :close-on-click-modal="false" :show-close="close">
            <el-form ref="ruleForm" :model="ruleForm" label-width="100px" class="demo-ruleForm">
                <el-form-item label="模型版本" :label-width="formLabelWidth" prop="modelDescript">
                    <el-input v-model="ruleForm.version" autocomplete="off" :disabled="true" />
                </el-form-item>
                <el-form-item label="模型描述" :label-width="formLabelWidth" prop="modelDescript">
                    <el-input v-model="ruleForm.descript" autocomplete="off" :disabled="true" />
                </el-form-item>
                <el-form-item v-if="showUpload" label="模型上传" :label-width="formLabelWidth">
                    <upload :upload-data="uploadData" @confirm="confirm" @cancel="cancel" @upload="isCloseX" />
                </el-form-item>
            </el-form>

        </el-dialog>
    </div>
</template>

<script>
    import upload from '@/components/upload/index.vue'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "CreateDialog",
        components: {
            upload
        },

        props: {
            isList: {
                type: Boolean
            },
            row: { type: Object, default: () => { } }
        },
        data() {
            return {
                ruleForm: {},
                dialogFormVisible: true,
                formLabelWidth: '120px',
                showUpload: true,
                uploadData: {type:'',data:{}},
                id: undefined,
                close: undefined

            }
        },
        created() {
            this.ruleForm.version = this.row.version
            this.ruleForm.descript = this.row.descript
            this.id = this.row.modelId
            this.uploadData.type = "modelManager"
            this.uploadData.data.modelId=this.row.modelId
            this.uploadData.data.version=this.row.version
            

        },
        beforeDestroy() {

        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            confirm(val) { this.$emit('confirm', val) },
            cancel(val) { this.$emit('cancel', val) },
            handleDialogClose() {
                this.$emit('close', false)
            },
            confirm(val) { this.$emit('confirm', val) },
            cancel(val) { this.$emit('cancel', val) },
            isCloseX(val) {
                this.close = val
            }

        }
    }
</script>
<style lang="scss" scoped>
    .el-dialog--center .el-dialog__body {
        text-align: center;
    }
</style>