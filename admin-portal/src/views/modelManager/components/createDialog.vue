<template>
    <div>
        <el-dialog
            :title="flag?'创建预置模型':'创建模型列表'"
            :visible.sync="dialogFormVisible"
            width="25%"
            :before-close="handleDialogClose"
            :close-on-click-modal="false"
        >
            <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="100px" class="demo-ruleForm">
                <el-form-item label="算法名称" :label-width="formLabelWidth" prop="algorithmId">
                    <el-select
                        v-model="ruleForm.algorithmName"
                        v-loadmore="loadAlgorithmName"
                        placeholder="请选择"
                        :disabled="!flag"
                        @change="changeAlgorithmName"
                    >
                        <el-option
                            v-for="item in algorithmNameOption"
                            :key="item.algorithmId"
                            :label="item.algorithmName"
                            :value="item.algorithmId"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item v-if="algorithmVersion||!flag" label="算法版本" :label-width="formLabelWidth" prop="algorithmVersion">
                    <el-select v-model="ruleForm.algorithmVersion" placeholder="请选择" :disabled="!flag">
                        <el-option
                            v-for="item in algorithmVersionOption"
                            :key="item.algorithmId"
                            :label="item.AlgorithmVersion"
                            :value="item.algorithmVersion"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="模型描述" :label-width="formLabelWidth" prop="modelDescript">
                    <el-input v-model="ruleForm.modelDescript" autocomplete="off" :disabled="showUpload" />
                </el-form-item>
                <el-form-item v-if="showUpload" label="模型上传" :label-width="formLabelWidth">
                    <upload :upload-data="uploadData" @confirm="confirm" @cancel="cancel" />
                </el-form-item>
            </el-form>
            <div v-if="createSuccess" slot="footer" class="dialog-footer">
                <el-button @click="handleDialogClose">取 消</el-button>
                <el-button type="primary" @click="create">下一步</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import { getPresetAlgorithmList, getAlgorithmVersionList } from '@/api/modelDev.js'
    import { addPreModel, addPreList } from '@/api/modelManager.js'
    import upload from '@/components/upload/index.vue'
    import { getErrorMsg } from '@/error/index'
    export default {
        name: "CreateDialog",
        components: {
            upload
        },
        directives: {
            loadmore: {
                inserted: function(el, binding) {
                    const SELECTWRAP_DOM = el.querySelector('.el-select-dropdown .el-select-dropdown__wrap');
                    SELECTWRAP_DOM.addEventListener('scroll', function() {
                        const CONDITION = this.scrollHeight - this.scrollTop <= this.clientHeight;
                        if (CONDITION) {
                            binding.value();
                        }
                    })
                }
            }
        },
        props: {
            isList: {
                type: Boolean
            },
            row: { type: Object, default: () => {} }
        },
        data() {
            return {
                ruleForm: {
                    algorithmName: "",
                    algorithmVersion: "",

                    modelDescript: ""

                },
                rules: {
                    algorithmName: [
                        { required: true, message: '请选择算法名称', trigger: 'change' }
                    ],
                    algorithmVersion: [
                        { required: true, message: '请选择算法版本', trigger: 'change' }
                    ],
                    modelDescript: [
                        { required: true, message: '请输入模型描述', trigger: 'blur' }

                    ]

                },
                algorithmNameOption: [],
                algorithmVersionOption: [],
                dialogFormVisible: true,
                formLabelWidth: '120px',
                flag: undefined,
                algorithmVersion: false,
                algorithmCount: 1,
                showUpload: false,
                uploadData: {},
                id: undefined,
                createSuccess: true
            }
        },
        created() {
            // 新增模型
            if (!this.isList) {
                this.flag = true; this.getPresetAlgorithmList(); this.createSuccess = true
            } else {
                // 新增模型版本
                this.flag = false
            }
            this.ruleForm.algorithmName = this.row.algorithmName
            this.ruleForm.algorithmVersion = this.row.algorithmVersion
            this.ruleForm.modelName = this.row.modelName
            this.ruleForm.modelDescript = this.row.modelDescript
            this.id = this.row.modelId
        },
        beforeDestroy() {

        },
        methods: {
            // 错误码
            getErrorMsg(code) {
                return getErrorMsg(code)
            },
            changeAlgorithmName() {
                this.algorithmVersion = true
                this.algorithmCount = 1
                this.algorithmVersionOption = []
                this.getAlgorithmVersionList()
            },
            getPresetAlgorithmList() {
                getPresetAlgorithmList({ pageIndex: this.algorithmCount, pageSize: 20 }).then(response => {
                    if (response.success) {
                        this.algorithmNameOption = this.algorithmNameOption.concat(response.data.algorithms)
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            getAlgorithmVersionList() {
                getAlgorithmVersionList({ pageIndex: this.algorithmCount, pageSize: 20, algorithmId: this.ruleForm.algorithmName }).then(response => {
                    if (response.success) {
                        this.algorithmVersionOption = response.data.algorithms
                    } else {
                        this.$message({
                            message: this.getErrorMsg(response.error.subcode),
                            type: 'warning'
                        });
                    }
                })
            },
            loadAlgorithmName() {
                this.algorithmCount = this.algorithmCount + 1
                this.getPresetAlgorithmList()
            },
            create() {
                this.$refs['ruleForm'].validate((valid) => {
                    if (true) {
                        const data = JSON.parse(JSON.stringify(this.ruleForm));
                        data.algorithmId = this.ruleForm.algorithmName
                        delete data.algorithmName
                        if (this.flag) {
                            addPreModel(data).then(response => {
                                if (response.success) {
                                    this.uploadData.data = response.data
                                    this.uploadData.type = "modelManager"
                                    this.showUpload = true
                                    this.createSuccess = false
                                } else {
                                    this.$message({
                                        message: this.getErrorMsg(response.error.subcode),
                                        type: 'warning'
                                    });
                                }
                            })
                        } else {
                            addPreList({ modelId: this.id, descript: data.modelDescript }).then(response => {
                                if (response.success) {
                                    this.uploadData.data = response.data
                                    this.uploadData.type = "modelManager"
                                    this.showUpload = true
                                    this.createSuccess = false
                                } else {
                                    this.$message({
                                        message: this.getErrorMsg(response.error.subcode),
                                        type: 'warning'
                                    });
                                }
                            })
                        }
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            confirm(val) { this.$emit('confirm', val) },
            cancel(val) { this.$emit('cancel', val) },
            handleDialogClose() {
                this.$emit('close', false)
            }

        }
    }
</script>
<style lang="scss" scoped>
    .el-dialog--center .el-dialog__body {
        text-align: center;
    }
</style>