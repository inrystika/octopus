<template>
    <div>
        <el-dialog :title="flag?'创建镜像':'编辑镜像'" width="650px" :visible.sync="CreateFormVisible"
            :before-close="handleDialogClose" :close-on-click-modal="false" :show-close="close">
            <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px" class="demo-ruleForm">
                <el-form-item label="镜像名称" :label-width="formLabelWidth" placeholder="请输入镜像名称" prop="imageName">
                    <el-input v-model="ruleForm.imageName" :disabled="!flag||showUpload" />
                </el-form-item>
                <el-form-item label="镜像标签" :label-width="formLabelWidth" placeholder="请输入镜像版本号" prop="imageVersion">
                    <el-input v-model="ruleForm.imageVersion" :disabled="!flag||showUpload" />
                </el-form-item>
                <el-form-item label="镜像描述" :label-width="formLabelWidth">
                    <el-input v-model="ruleForm.imageDesc" type="textarea" :disabled="!flag||showUpload" />
                </el-form-item>
                <el-form-item label="镜像来源" :label-width="formLabelWidth" prop="sourceType">
                    <el-select v-model="ruleForm.sourceType" placeholder="请选择上传类型" :disabled="!flag">
                        <el-option label="文件上传" :value="1" />
                        <el-option label="远程镜像" :value="2" />
                    </el-select>
                    <upload v-if="showUpload && ruleForm.sourceType===1" :upload-data="uploadData" @confirm="confirm"
                        @cancel="cancel" @upload="isCloseX" />
                </el-form-item>
                <el-form-item v-if="ruleForm.sourceType===2" label="远程镜像地址" :label-width="formLabelWidth"
                    placeholder="请输入镜像名称" prop="imageAddr">
                    <el-input v-model="ruleForm.imageAddr" placeholder="请输入远程镜像地址" :disabled="!flag" />
                </el-form-item>
            </el-form>
            <div v-if="ruleForm.sourceType===2" slot="footer" class="dialog-footer">
                <el-button @click="handleDialogClose">取 消</el-button>
                <el-button type="primary" @click="submitAdd('ruleForm')" v-preventReClick>确 定</el-button>
            </div>
            <div v-if="ruleForm.sourceType===1&&!showUpload" slot="footer" class="dialog-footer">
                <el-button @click="cancel">取 消</el-button>
                <el-button type="primary" @click="submitUpload('ruleForm')">下一步</el-button>
            </div>
        </el-dialog>
    </div>
</template>
<script>
    import { createPreImage, editPreImage } from '@/api/imageManager.js'
    import upload from '@/components/upload/index.vue'
    export default {
        name: "DialogCreateForm",
        components: {
            upload
        },
        props: {
            row: {
                type: Object,
                default: () => { }
            },
            flag: {
                type: Boolean,
                default: true
            }
        },
        data() {
            var checkName = (rule, value, callback) => {
                const regName = /^[a-z][a-z0-9\-]*[a-z0-9]$/;
                if (regName.test(value)) {
                    return callback();
                }
                callback(new Error("镜像名称由小写字母、数字或者-组成，小写字母开头，数字或小写字母结尾"));
            };
            var checkLabel = (rule, value, callback) => {
                const regLabel = /^[a-zA-Z][\w|\-|\.]+$/;
                if (regLabel.test(value)) {
                    return callback();
                }
                callback(new Error("请输入合法的标签名称:首字母为英文,其他为英文数字.或者-"));
            };
            return {
                ruleForm: { imageAddr: '', imageDesc: '', imageName: '', imageVersion: '', sourceType: 2, },
                // 镜像id
                id: undefined,
                showUpload: false,
                // 上传完成获得参数
                uploadData: { data: {}, type: undefined },
                CreateFormVisible: true,
                rules: {
                    imageName: [
                        { required: true, message: '请输入镜像名称', trigger: 'blur' },
                        { validator: checkName, trigger: "blur" }

                    ],
                    imageVersion: [
                        { required: true, message: '请输入镜像版本号', trigger: 'blur' },
                        { validator: checkLabel, trigger: "blur" }

                    ],
                    sourceType: [
                        { required: true, message: '请选择镜像上传类型', trigger: 'change' }
                    ],
                    imageAddr: [
                        { required: true, message: '请输入远程镜像地址', trigger: 'blur' }

                    ]

                },
                formLabelWidth: '120px',
                close: true
            }
        },
        created() {
            // 新建镜像
            // 编辑镜像
            if (!this.flag) {
                this.id = this.row.id
                this.ruleForm = { imageType: this.row.imageType, imageDesc: this.row.imageDesc, imageName: this.row.imageName, imageVersion: this.row.imageVersion, imageAddr: this.row.imageAddr, sourceType: this.row.sourceType, imageStatus: this.row.imageStatus }
                if (this.ruleForm.imageStatus === 1 || this.ruleForm.imageStatus === 4) {
                    this.uploadData.data.id = this.id
                    this.uploadData.type = "imageManager"
                    this.showUpload = true
                }
                if (this.ruleForm.sourceType === 1) {
                    this.uploadData.data.id = this.id
                    this.uploadData.type = "imageManager"
                    this.showUpload = true
                }
            }
        },
        beforeDestroy() {
            this.ruleForm = {
                imageName: "",
                imageVersion: "",
                imageDesc: "",
                imageType: "",
                sourceType: 2,
                imageAddr: ""
            }
        },
        methods: {
            submitUpload() {
                if (this.ruleForm.sourceType === 1) {
                    delete this.rules.imageAddr
                    this.$refs['ruleForm'].validate((valid) => {
                        if (valid) {
                            createPreImage(this.ruleForm).then(response => {
                                if (response.success) {
                                    this.uploadData.data.id = response.data.imageId
                                    this.uploadData.type = "imageManager"
                                    this.showUpload = true
                                } else {
                                    this.$message({
                                        message: this.getErrorMsg(response.error.subcode),
                                        type: 'warning'
                                    });
                                }
                            })
                        } else {
                            console.log('error submit!!');
                            return false;
                        }
                    });
                } else { this.rules.imageAddr = [{ required: true, message: '请输入远程镜像地址', trigger: 'blur' }] }
            },
            createPreImage(data) {
                if (this.ruleForm.sourceType === 2) {
                    createPreImage(data).then(response => {
                        if (response.success) {
                            this.$message({
                                message: '创建镜像成功',
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
            },
            editPreImage(data) {
                editPreImage(data).then(response => {
                    if (response.success) {
                        this.$message({
                            message: '编辑镜像成功',
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
            },
            submitAdd(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        if (this.flag) {
                            this.createPreImage(this.ruleForm)
                        } else {
                            const data = { ...this.ruleForm, id: this.id }
                            this.editPreImage(data)
                        }
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
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
</style>